# gimu

Cross-platform GUI for go based on nuklear.

Package nk provides Go bindings for nuklear.h — a small ANSI C gui library. See [github.com/Immediate-Mode-UI/Nuklear](https://github.com/Immediate-Mode-UI/Nuklear).

All the binding code has automatically been generated with rules defined in [nk.yml](/nk.yml).

This package provides a go-style idiomatic wrapper for nuklear.

## Screenshots

<img src="https://github.com/AllenDang/gimu/blob/master/examples/screenshots.png" alt="Simple demo screen shots" width="800"/>

## Overview

Supported platforms are:

* Windows 32-bit
* Windows 64-bit
* OS X
* Linux

The desktop support is achieved using [GLFW](https://github.com/go-gl/glfw) and there are backends written in Go for OpenGL 3.2.

### Installation

Just go get it and everythings ready to work.

```
go get -u github.com/AllenDang/gimu
```

### Getting start

Let's create a simple demo.

```go
package main

import (
	"fmt"
	"image"
	"runtime"

	"github.com/AllenDang/gimu"
)

func builder(w *gimu.Window) {
	// Create a new window inside master window
	width, height := w.MasterWindow().GetSize()
	bounds := image.Rect(0, 0, width, height)

	w.Window("Simple Demo", bounds, gimu.WindowNoScrollbar, func(w *gimu.Window) {
		// Define the row with 25px height, and contains one widget for each row.
		w.Row(25).Dynamic(1)
		// Let's create a label first, note the second parameter "LC" means the text alignment is left-center.
		w.Label("I'm a label", "LC")
		// Create a button.
		clicked := w.Button("Click Me")
		if clicked {
			fmt.Println("The button is clicked")
		}
	})
}

func main() {
	runtime.LockOSThread()

	// Create master window
	wnd := gimu.NewMasterWindow("Simple Demo", 200, 100, gimu.MasterWindowFlagDefault)

	wnd.Main(builder)
}
```

Save and run.

### Deploy

gimu provides a tool to pack compiled executable for several platform to enable app icon and etc.

```
go get -u github.com/AllenDang/gimu/cmd/gmdeploy
```

Run gmdeploy in your project folder.

```
gmdeploy -icon AppIcon.icns .
```

Then you can find bundled executable in [PROJECTDIR]/build/[OS]/

Note:

Currently only MacOS is supported. Windows and linux is WIP.

### Layout system

Layouting in general describes placing widget inside a window with position and size. While in this particular implementation there are two different APIs for layouting

All layouting methods in this library are based around the concept of a row.

A row has a height the window content grows by and a number of columns and each layouting method specifies how each widget is placed inside the row.

After a row has been allocated by calling a layouting functions and then filled with widgets will advance an internal pointer over the allocated row. 

To actually define a layout you just call the appropriate layouting function and each subsequent widget call will place the widget as specified. Important here is that if you define more widgets then columns defined inside the layout functions it will allocate the next row without you having to make another layouting call.

#### Static layout

Define a row with 25px height with two widgets.

```go
w.Row(25).Static(50, 50)
```

Use the magic number 0 to define a widget will auto expand if there is enough space.

```go
w.Row(25).Static(0, 50)
w.Label("I'm a auto growth label", "LC")
w.Button("I'm a button with fixed width")
```

#### Dynamic layout

It provides each widgets with same horizontal space inside the row and dynamically grows if the owning window grows in width. 

Define a row with two widgets each of them will have same width.

```go
w.Row(25).Dynamic(2)
```

#### Flexible Layout

Finally the most flexible API directly allows you to place widgets inside the window. The space layout API is an immediate mode API which does not support row auto repeat and directly sets position and size of a widget. Position and size hereby can be either specified as ratio of allocated space or allocated space local position and pixel size. Since this API is quite powerful there are a number of utility functions to get the available space and convert between local allocated space and screen space.

```go
w.Row(500).Space(nk.Static, func(w *gimu.Window) {
  w.Push(nk.NkRect(0, 0, 150, 150))
    w.Group("Group Left", gimu.WindowBorder|gimu.WindowTitle, func(w *gimu.Window) {
  })

  w.Push(nk.NkRect(160, 0, 150, 240))
    w.Group("Group Top", gimu.WindowBorder|gimu.WindowTitle, func(w *gimu.Window) {
  })

  w.Push(nk.NkRect(160, 250, 150, 250))
  w.Group("Group Bottom", gimu.WindowBorder|gimu.WindowTitle, func(w *gimu.Window) {
  })

  w.Push(nk.NkRect(320, 0, 150, 150))
  w.Group("Group Right Top", gimu.WindowBorder|gimu.WindowTitle, func(w *gimu.Window) {
  })

  w.Push(nk.NkRect(320, 160, 150, 150))
  w.Group("Group Right Center", gimu.WindowBorder|gimu.WindowTitle, func(w *gimu.Window) {
  })

  w.Push(nk.NkRect(320, 320, 150, 150))
  w.Group("Group Right Center", gimu.WindowBorder|gimu.WindowTitle, func(w *gimu.Window) {
  })
})
```

## Widgets usage

Most of the widget's usage are very straight forward.

### Common widgets

#### Label

The second parameter of label indicates the text alignment.

```go
w.Label("Label caption", "LC")
```

"LC" means horizontally left and vertically center.

"LT" means horizontally left and vertically top.

The alignment char layout is listed below, you could use any combinations of those.

   T

L-C-R

   B

#### Selectable Label

Label can be toggled by mouse click.

```go
var selected bool
w.SelectableLabel("Selectable 1", "LC", &selected1)
```

#### Button

Button function will return a bool to indicate whether it was clicked.

```go
clicked := w.Button("Click Me")
if clicked {
  // Do something here
}
```

#### Progressbar

Progress could be readonly or modifiable.

```go
progress := 0
// Modifiable
w.Progress(&progress, 100, true)
// Readonly
w.Progress(&progress, 100, false)
```

To read current progress or update progress bar, just set the progress variable.

#### Slider

Slider behaves like progress bar but step control.

```go
var slider int
w.SliderInt(0, &slider, 100, 1)
```

#### Property widgets

It contains a label and a adjustable control to modify int or float variable.

``` go
var propertyInt int
var propertyFloat float32
w.PropertyInt("Age", 1, &propertyInt, 100, 10, 1)
w.PropertyFloat("Height", 1, &propertyFloat, 10, 0.2, 1)
```

#### Checkbox

```go
var checked bool
w.Checkbox("Check me", &checked)
```

#### Radio

```go
option := 1
if op1 := w.Radio("Option 1", option == 1); op1 {
  option = 1
}
if op2 := w.Radio("Option 2", option == 2); op2 {
  option = 2
}
if op3 := w.Radio("Option 3", option == 3); op3 {
  option = 3
}
```

#### Textedit

Textedit is special because it will retain the input string, so you will have to explicitly create it and call the Edit() function in BuilderFunc.

```go
textedit := gimu.NewTextEdit()

func builder(w *gimu.Window) {
  textedit.Edit(w, gimu.EditField, gimu.EditFilterDefault)
}
```

#### ListView

ListView is designed to display very huge amount of data and only render visible items.

```go
var (
	listview *nk.ListView
	listitem []interface{}
)

func builder(w *gimu.Window) {
	width, height := w.MasterWindow().GetSize()
	bounds := nk.NkRect(0, 0, float32(width), float32(height))
	w.Window("", bounds, 0, func(w *gimu.Window) {
		w.Row(int(height - 18)).Dynamic(1)
		w.ListView(listview, "huge list", nk.WindowBorder, 25,
			listitem,
			func(r *gimu.Row) {
				r.Dynamic(1)
			},
			func(w *gimu.Window, i int, item interface{}) {
				if s, ok := item.(string); ok {
					w.Label(s, "LC")
				}
			})
	})
}

func main() {
	// Init the listview widget
	listview = &nk.ListView{}

	// Create list items
	listitem = make([]interface{}, 12345)
	for i := range listitem {
		listitem[i] = fmt.Sprintf("Label item %d", i)
	}

	runtime.LockOSThread()

	wnd := gimu.NewMasterWindow("Huge list", 800, 600, gimu.MasterWindowFlagNoResize)
	wnd.Main(builder)
}
```

### Popups

#### Tooltip

**Note: Tooltip has to be placed above the widget which wants a tooltip when mouse hovering.**

```go
w.Tooltip("This is a tooltip")
w.Button("Hover me to see tooltip")
```

#### Popup Window

```go
func msgbox(w *gimu.Window) {
  opened := w.Popup(
    "Message", 
    gimu.PopupStatic, 
    gimu.WindowTitle|gimu.WindowNoScrollbar|gimu.WindowClosable, 
    image.Rect(30, 10, 300, 100), 
    func(w *gimu.Window) {
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
```

### Menu

#### Window Menu

**Note: window menu bar has to be the first widget in the builder method.**

```go
// Menu
w.Menubar(func(w *gimu.Window) {
  w.Row(25).Static(60, 60)
  // Menu 1
  w.Menu("Menu1", "CC", 200, 100, func(w *gimu.Window) {
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
```

#### Contextual Menu

**Note: Contextual menu has to be placed above the widget which wants a tooltip when right click.**

You could put any kind of widgets inside the contextual menu.

```go
w.Contextual(0, 100, 300, func(w *gimu.Window) {
  w.Row(25).Dynamic(1)
  w.ContextualLabel("Context menu 1", "LC")
  w.ContextualLabel("Context menu 1", "LC")
  w.SliderInt(0, &slider, 100, 1)
})
w.Button("Right click me")
```

## License

All the code except when stated otherwise is licensed under the [MIT license](https://xlab.mit-license.org).
Nuklear (ANSI C version) is in public domain, authored from 2015-2016 by Micha Mettke.

