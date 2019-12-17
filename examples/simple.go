package main

import (
	"fmt"
	"image"
	"image/color"
	"runtime"

	"github.com/AllenDang/gimu"
)

func updatefn(w *gimu.Window) {
	width, height := w.MasterWindow().GetSize()
	bounds := image.Rect(0, 0, width, height)

	if w.Begin("Simple Demo", bounds, gimu.WindowNoScrollbar) {
		w.Row(25).Dynamic(1)
		w.Label("Hello world!", "LC")
		w.Label("Hello world!", "CC")
		w.Label("Hello world!", "RC")
		w.LabelColored("Colored label", color.RGBA{200, 100, 100, 255}, "LC")
		if w.Button("Click Me") {
			fmt.Println("Button has been clicked")
		}
	}
	w.End()
}

func main() {
	runtime.LockOSThread()

	wnd := gimu.NewMasterWindow("Simple Demo", 400, 200, gimu.MasterWindowFlagNoResize)

	wnd.Main(updatefn)
}
