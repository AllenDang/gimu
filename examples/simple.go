package main

import (
	"fmt"
	"image"
	"image/color"
	"runtime"

	"github.com/AllenDang/gimu"
)

var (
	textedit   = gimu.NewTextEdit()
	selected   int32
	comboLabel string
	num1       uint = 11
	num2       uint = 33
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

		selected = w.ComboSimple([]string{"Item1", "Item2", "Item3"}, selected, 25, 0, 200)

		comboLabel = fmt.Sprintf("%d", num1+num2)
		w.ComboLabel(comboLabel, 0, 100, func(w *gimu.Window) {
			w.Row(25).Dynamic(1)
			w.Label("Drag progress bar to see the changes", "LC")

			w.Row(25).Static(0, 30)
			w.Progress(&num1, 100, true)
			w.Label(fmt.Sprintf("%d", num1), "CC")

			w.Progress(&num2, 100, true)
			w.Label(fmt.Sprintf("%d", num2), "CC")
		})

		w.Row(25).Static(0, 100)
		textedit.Edit(w, gimu.EditField, gimu.EditFilterBinary)
		if w.Button("Print") {
			fmt.Println(textedit.GetString())
		}
	})
}

func main() {
	runtime.LockOSThread()

	wnd := gimu.NewMasterWindow("Simple Demo", 400, 400, gimu.MasterWindowFlagDefault)

	wnd.Main(updatefn)
}
