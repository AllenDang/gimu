package main

import (
	"fmt"
	"image"
	"image/color"
	"runtime"

	"github.com/AllenDang/gimu"
)

var (
	textedit      = gimu.NewTextEdit()
	selected      int32
	comboLabel    string
	num1          uint = 11
	num2          uint = 33
	propertyInt   int32
	propertyFloat float32
	checked       bool
	option        int = 1
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

		w.Label("Combobox", "LC")

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

		w.Label("Properties", "LC")
		w.PropertyInt("Age", 1, &propertyInt, 100, 10, 1)
		w.PropertyFloat("Height", 1, &propertyFloat, 10, 0.2, 1)

		w.Label("Checkbox", "LC")
		w.Row(25).Static(0, 100)
		w.Checkbox("Check me", &checked)
		w.Label(fmt.Sprintf("%v", checked), "LC")

		w.Row(25).Dynamic(1)
		w.Label("Radio", "LC")
		w.Row(25).Dynamic(3)
		if op1 := w.Radio("Option 1", option == 1); op1 {
			option = 1
		}
		if op2 := w.Radio("Option 2", option == 2); op2 {
			option = 2
		}
		if op3 := w.Radio("Option 3", option == 3); op3 {
			option = 3
		}

		w.Row(25).Static(0, 100)
		textedit.Edit(w, gimu.EditField, gimu.EditFilterDefault)
		if w.Button("Print") {
			fmt.Println(textedit.GetString())
		}
	})
}

func main() {
	runtime.LockOSThread()

	wnd := gimu.NewMasterWindow("Simple Demo", 400, 500, gimu.MasterWindowFlagDefault)

	wnd.Main(updatefn)
}
