package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
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
	selected1     bool
	selected2     bool
	showPopup     bool
	picture       *gimu.Texture
	slider        int32 = 33
)

func msgbox(w *gimu.Window) {
	opened := w.Popup("Message", gimu.PopupStatic, gimu.WindowTitle|gimu.WindowNoScrollbar|gimu.WindowClosable, image.Rect(30, 10, 300, 100), func(w *gimu.Window) {
		w.Row(25).Dynamic(1)
		w.Label("Here is a pop up window", "LC")
		if w.Button("Close") {
			showPopup = false
			w.ClosePopup()
		}
	})
	if !opened {
		showPopup = false
	}
}

func widgets(w *gimu.Window) {
	w.Row(25).Dynamic(1)
	w.Label("Hello world!", "LC")
	w.Label("Hello world!", "CC")
	w.Label("Hello world!", "RC")
	w.LabelColored("Colored label", color.RGBA{200, 100, 100, 255}, "LC")
	if w.Button("Click me to show a popup window") {
		showPopup = true
	}

	if showPopup {
		msgbox(w)
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

	w.Label("Slider", "LC")
	w.SliderInt(0, &slider, 100, 1)

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

	w.Row(25).Dynamic(1)
	w.Label("Selectable label", "LC")
	w.Row(25).Dynamic(2)
	w.SelectableLabel("Selectable 1", "LC", &selected1)
	w.SelectableLabel("Selectable 2", "LC", &selected2)

	w.Row(25).Static(0, 100)
	textedit.Edit(w, gimu.EditField, gimu.EditFilterDefault)
	if w.Button("Print") {
		fmt.Println(textedit.GetString())
	}
}

func updatefn(w *gimu.Window) {
	width, height := w.MasterWindow().GetSize()
	bounds := image.Rect(0, 0, width, height)

	w.Window("Simple Demo", bounds, gimu.WindowNoScrollbar, func(w *gimu.Window) {
		_, h := w.MasterWindow().GetSize()
		w.Row(float32(h - 10)).Dynamic(2)
		w.Group("Group1", gimu.WindowBorder|gimu.WindowTitle, func(w *gimu.Window) {
			widgets(w)
		})
		w.Group("Group2", gimu.WindowBorder|gimu.WindowTitle, func(w *gimu.Window) {
			// Menu
			w.Menubar(func(w *gimu.Window) {
				w.Row(25).Static(60, 60)
				// Menu 1
				w.Menu("Menu1", "CC", 100, 100, func(w *gimu.Window) {
					w.Row(25).Dynamic(1)
					w.MenuItemLabel("Menu item 1", "LC")
					w.MenuItemLabel("Menu item 2", "LC")
					w.Button("Button inside menu")
				})
				// Menu 2
				w.Menu("Menu2", "CC", 100, 100, func(w *gimu.Window) {
					w.Row(25).Dynamic(1)
					w.MenuItemLabel("Menu item 1", "LC")
					w.SliderInt(0, &slider, 100, 1)
					w.MenuItemLabel("Menu item 2", "LC")
				})

			})

			// Image
			w.Row(170).Static(300)
			if picture != nil {
				w.Image(picture)
			}

			// Tooltip
			w.Row(25).Dynamic(1)
			w.Tooltip("This is a tooltip")
			w.Button("Hover me to see tooltip")

			// Contextual menu
			w.Contextual(0, 100, 300, func(w *gimu.Window) {
				w.Row(25).Dynamic(1)
				w.ContextualLabel("Context menu 1", "LC")
				w.ContextualLabel("Context menu 1", "LC")
				w.SliderInt(0, &slider, 100, 1)
			})
			w.Button("Right click me")
		})
	})
}

func main() {
	runtime.LockOSThread()

	// Create master window
	wnd := gimu.NewMasterWindow("Simple Demo", 1000, 700, gimu.MasterWindowFlagDefault)

	// Load png image
	fn, err := os.Open("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	defer fn.Close()

	img, err := png.Decode(fn)
	if err != nil {
		log.Fatal(err)
	}
	if img != nil {
		rgba := gimu.ImgToRgba(img)
		picture = gimu.RgbaToTexture(rgba)
	}

	wnd.Main(updatefn)
}
