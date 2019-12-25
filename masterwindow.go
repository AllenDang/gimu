package gimu

import (
	"image/color"
	"log"
	"time"

	"github.com/AllenDang/gimu/nk"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type BuilderFunc func(w *Window)

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
	ctx              *nk.Context
	maxVertexBuffer  int
	maxElementBuffer int
	bgColor          nk.Color
	defaultFont      *nk.Font
}

func NewMasterWindow(title string, width, height int, flags MasterWindowFlag) *MasterWindow {
	if err := glfw.Init(); err != nil {
		log.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	if flags.HasFlag(MasterWindowFlagNoResize) {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init failed:", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))

	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)

	return &MasterWindow{
		win:              win,
		ctx:              ctx,
		maxVertexBuffer:  512 * 1024,
		maxElementBuffer: 128 * 1024,
		bgColor:          nk.NkRgba(28, 48, 62, 255),
	}
}

func (w *MasterWindow) GetSize() (width, height int) {
	gw, gh := w.win.GetSize()
	return gw, gh
}

func (w *MasterWindow) SetBgColor(color color.RGBA) {
	w.bgColor = toNkColor(color)
}

func (w *MasterWindow) SetTitle(title string) {
	w.win.SetTitle(title)
}

func (w *MasterWindow) GetContext() *nk.Context {
	return w.ctx
}

func (w *MasterWindow) GetDefaultFont() *nk.Font {
	return w.defaultFont
}

func (w *MasterWindow) Main(builder BuilderFunc) {
	// Load default font
	w.defaultFont = LoadDefaultFont()
	w.GetContext().SetStyle(nk.THEME_DARK)

	window := Window{
		ctx: w.ctx,
		mw:  w,
	}

	for !w.win.ShouldClose() {
		glfw.PollEvents()
		nk.NkPlatformNewFrame()

		builder(&window)

		// Render
		bg := make([]float32, 4)
		nk.NkColorFv(bg, w.bgColor)
		width, height := w.GetSize()
		gl.Viewport(0, 0, int32(width), int32(height))
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.ClearColor(bg[0], bg[1], bg[2], bg[3])
		nk.NkPlatformRender(nk.AntiAliasingOn, w.maxVertexBuffer, w.maxElementBuffer)
		w.win.SwapBuffers()

		// 30 FPS
		time.Sleep(time.Second / 30)
	}

	nk.NkPlatformShutdown()
	glfw.Terminate()
}
