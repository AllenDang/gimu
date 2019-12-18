package gimu

import "github.com/AllenDang/nuklear/nk"

type BuilderFunc func(w *Window)

type MasterWindow struct {
	wnd *nk.MasterWindow
}

type MasterWindowFlag int

const (
	MasterWindowFlagDefault  MasterWindowFlag = MasterWindowFlag(nk.MasterWindowFlagDefault)
	MasterWindowFlagNoResize MasterWindowFlag = MasterWindowFlag(nk.MasterWindowFlagNoResize)
)

func NewMasterWindow(title string, width, height int, flag MasterWindowFlag) *MasterWindow {
	wnd := nk.NewMasterWindow(title, width, height, nk.MasterWindowFlag(flag))
	wnd.LoadDefaultFont()
	wnd.GetContext().SetStyle(nk.THEME_DARK)
	return &MasterWindow{wnd: wnd}
}

func (mw *MasterWindow) Main(updatefn BuilderFunc) {
	window := &Window{
		ctx: mw.wnd.GetContext(),
		mw:  mw,
	}
	fn := func(w *nk.MasterWindow) {
		updatefn(window)
	}
	mw.wnd.Run(fn)
}

func (mw *MasterWindow) GetSize() (width, height int) {
	return mw.wnd.GetSize()
}
