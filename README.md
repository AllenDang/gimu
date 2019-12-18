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

## Installation

Just go get it and everythings ready to work.

```
go get -u github.com/AllenDang/gimu
```

## Getting start

Let's create a simple demo.

```
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

Save, run.


## License

All the code except when stated otherwise is licensed under the [MIT license](https://xlab.mit-license.org).
Nuklear (ANSI C version) is in public domain, authored from 2015-2016 by Micha Mettke.

