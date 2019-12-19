# gimu

Immediate GUI for go based nuklear.

Package nk provides Go bindings for nuklear.h — a small ANSI C gui library. See [github.com/Immediate-Mode-UI/Nuklear](https://github.com/Immediate-Mode-UI/Nuklear).<br />
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


## License

All the code except when stated otherwise is licensed under the [MIT license](https://xlab.mit-license.org).
Nuklear (ANSI C version) is in public domain, authored from 2015-2016 by Micha Mettke.

