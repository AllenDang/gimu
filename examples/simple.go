package main

import (
	"fmt"
	"image"
	"image/color"
	"runtime"

	"github.com/AllenDang/gimu"
)

var (
	textedit = gimu.NewTextEdit()
)

func updatefn(w *gimu.Window) {
	width, height := w.MasterWindow().GetSize()
	bounds := image.Rect(0, 0, width, height)

	w.Window("Simple Demo", bounds, gimu.WindowNoScrollbar, func(w *gimu.Window) {
		w.Row(25).Dynamic(1)
		w.Label("Hello world!", "LC")
		w.Label("Hello world!", "CC")
		w.Label("Hello world!", "RC")
		w.LabelColored("Colored label", color.RGBA{200, 100, 100, 255}, "LC")
		if w.Button("Click Me") {
			fmt.Println("Button has been clicked")
		}

		w.Row(25).Static(0, 100)
		textedit.Edit(w, gimu.EditField, gimu.EditFilterBinary)
		if w.Button("Print") {
			fmt.Println(textedit.GetString())
		}
	})
}

func main() {
	runtime.LockOSThread()

	wnd := gimu.NewMasterWindow("Simple Demo", 400, 200, gimu.MasterWindowFlagDefault)

	wnd.Main(updatefn)
}
