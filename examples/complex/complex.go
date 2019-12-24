package main

import (
	"runtime"

	"github.com/AllenDang/gimu"
	"github.com/AllenDang/gimu/nk"
)

func builder(w *gimu.Window) {
	width, height := w.MasterWindow().GetSize()

	w.Window("", nk.NkRect(0, 0, float32(width), float32(height)), nk.WindowNoScrollbar, func(w *gimu.Window) {
		w.Row(500).Space(nk.Static, func(w *gimu.Window) {
			w.Push(nk.NkRect(0, 0, 150, 150))
			w.Group("Group Left", nk.WindowBorder|nk.WindowTitle, func(w *gimu.Window) {
			})

			w.Push(nk.NkRect(160, 0, 150, 240))
			w.Group("Group Top", nk.WindowBorder|nk.WindowTitle, func(w *gimu.Window) {
			})

			w.Push(nk.NkRect(160, 250, 150, 250))
			w.Group("Group Bottom", nk.WindowBorder|nk.WindowTitle, func(w *gimu.Window) {
			})

			w.Push(nk.NkRect(320, 0, 150, 150))
			w.Group("Group Right Top", nk.WindowBorder|nk.WindowTitle, func(w *gimu.Window) {
			})

			w.Push(nk.NkRect(320, 160, 150, 150))
			w.Group("Group Right Center", nk.WindowBorder|nk.WindowTitle, func(w *gimu.Window) {
			})

			w.Push(nk.NkRect(320, 320, 150, 150))
			w.Group("Group Right Center", nk.WindowBorder|nk.WindowTitle, func(w *gimu.Window) {
			})
		})
	})
}

func main() {
	runtime.LockOSThread()

	wnd := gimu.NewMasterWindow("Complex layout", 500, 600, gimu.MasterWindowFlagNoResize)

	wnd.Main(builder)
}
