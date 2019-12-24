package main

import (
	"runtime"

	"github.com/AllenDang/gimu"
	"github.com/AllenDang/gimu/nk"
)

var (
	leftWidth     int = 200
	rightWidth    int
	splitterWidth int = 5
)

func builder(w *gimu.Window) {
	width, height := w.MasterWindow().GetSize()

	w.Window("", nk.NkRect(0, 0, float32(width), float32(height)), nk.WindowNoScrollbar, func(w *gimu.Window) {
		rightWidth = int(width) - leftWidth - splitterWidth - 25

		w.Row(int(height-10)).Static(leftWidth, splitterWidth, rightWidth)
		w.Group("Left Group", nk.WindowTitle|nk.WindowBorder|nk.WindowNoScrollbar, func(w *gimu.Window) {
			w.Row(25).Dynamic(1)
			w.Label("Content", "LC")
		})

		//Splitter
		bounds := w.WidgetBounds()
		w.Spacing(1)
		input := w.GetInput()
		if (input.IsMouseHoveringRect(bounds) || input.IsMousePrevHoveringRect(bounds)) && input.IsMouseDown(nk.ButtonLeft) {
			x, _ := input.Mouse().Delta()
			leftWidth += int(x)
			rightWidth -= int(x)
		}

		w.Group("Right Group", nk.WindowTitle|nk.WindowBorder|nk.WindowNoScrollbar, func(w *gimu.Window) {
			w.Row(25).Dynamic(1)
			w.Label("Drag the space between two group to resize", "LC")
		})
	})
}

func main() {
	runtime.LockOSThread()

	wnd := gimu.NewMasterWindow("Splitter Demo", 800, 600, gimu.MasterWindowFlagDefault)

	wnd.Main(builder)
}
