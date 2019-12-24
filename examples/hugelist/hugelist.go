package main

import (
	"fmt"
	"runtime"

	"github.com/AllenDang/gimu"
	"github.com/AllenDang/gimu/nk"
)

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
