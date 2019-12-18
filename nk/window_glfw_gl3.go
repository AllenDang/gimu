package nk

import (
	"time"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/xlab/closer"
)

// Master window flag
type MasterWindowFlag int

const (
	// MasterWindowFlagNoResize - Create an not resizable window
	MasterWindowFlagDefault MasterWindowFlag = iota
	MasterWindowFlagNoResize
)

func (this MasterWindowFlag) HasFlag(flag MasterWindowFlag) bool {
	return this|flag == this
}

// Master window
type MasterWindow struct {
	win              *glfw.Window
	ctx              *Context
	maxVertexBuffer  int
	maxElementBuffer int
	bgColor          Color
}

func NewMasterWindow(title string, width, height int, flags MasterWindowFlag) *MasterWindow {
	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	if flags.HasFlag(MasterWindowFlagNoResize) {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		closer.Fatalln(err)
	}
	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		closer.Fatalln("opengl: init failed:", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))

	ctx := NkPlatformInit(win, PlatformInstallCallbacks)

	return &MasterWindow{
		win:              win,
		ctx:              ctx,
		maxVertexBuffer:  512 * 1024,
		maxElementBuffer: 128 * 1024,
		bgColor:          NkRgba(28, 48, 62, 255),
	}
}

func (w *MasterWindow) GetSize() (width, height int) {
	return w.win.GetSize()
}

func (w *MasterWindow) SetBgColor(color Color) {
	w.bgColor = color
}

func (w *MasterWindow) SetTitle(title string) {
	w.win.SetTitle(title)
}

func (w *MasterWindow) GetContext() *Context {
	return w.ctx
}

func (w *MasterWindow) LoadDefaultFont() {
	atlas := NewFontAtlas()
	NkFontStashBegin(&atlas)
	NkFontStashEnd()
}

func (w *MasterWindow) LoadFontFromFile(filePath string, size float32, config *FontConfig) {
	atlas := NewFontAtlas()
	NkFontStashBegin(&atlas)

	f := NkFontAtlasAddFromFile(atlas, filePath, size, config)

	NkFontStashEnd()

	if f == nil {
		closer.Fatalln("Failed to load font")
	}

	NkStyleSetFont(w.GetContext(), f.Handle())
}

func (w *MasterWindow) internalUpdateFn(updateFunc func(win *MasterWindow)) {
	NkPlatformNewFrame()

	updateFunc(w)

	// Render
	bg := make([]float32, 4)
	NkColorFv(bg, w.bgColor)
	width, height := w.GetSize()
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(bg[0], bg[1], bg[2], bg[3])
	NkPlatformRender(AntiAliasingOn, w.maxVertexBuffer, w.maxElementBuffer)
	w.win.SwapBuffers()
}

func (w *MasterWindow) Run(updateFunc func(win *MasterWindow)) {
	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	closer.Bind(func() {
		close(exitC)
		<-doneC
	})

	fpsTicker := time.NewTicker(time.Second / 30)
	for {
		select {
		case <-exitC:
			NkPlatformShutdown()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if w.win.ShouldClose() {
				close(exitC)
				continue
			}
			glfw.PollEvents()
			w.internalUpdateFn(updateFunc)
		}
	}
}
