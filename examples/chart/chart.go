package main

import (
	"math/rand"
	"runtime"

	"github.com/AllenDang/gimu"
	"github.com/AllenDang/gimu/nk"
)

var (
	lineValues  []float32
	line2Values []float32
	line3Values []float32
)

func builder(w *gimu.Window) {
	width, height := w.MasterWindow().GetSize()
	w.Window("", nk.NkRect(0, 0, float32(width), float32(height)), nk.WindowNoScrollbar, func(w *gimu.Window) {
		w.Row(25).Dynamic(1)
		w.Label("Simple Charts", "LC")

		w.Row(100).Dynamic(1)
		w.Chart(nk.ChartLines, 0, 100, lineValues)
		w.ChartColored(nk.ChartLines, nk.NkRgb(255, 0, 0), nk.NkRgb(150, 0, 0), 0, 100, line2Values)

		w.Chart(nk.ChartColumn, 0, 100, lineValues)
		w.ChartColored(nk.ChartColumn, nk.NkRgb(255, 0, 0), nk.NkRgb(150, 0, 0), 0, 100, line2Values)

		w.Row(25).Dynamic(1)
		w.Label("Mixed Charts", "LC")

		w.Row(100).Dynamic(1)
		w.ChartMixed([]gimu.ChartSeries{
			gimu.ChartSeries{
				ChartType:   nk.ChartLines,
				Min:         0,
				Max:         100,
				Data:        lineValues,
				Color:       nk.NkRgb(0, 100, 0),
				ActiveColor: nk.NkRgb(0, 200, 0),
			},
			gimu.ChartSeries{
				ChartType:   nk.ChartLines,
				Min:         0,
				Max:         100,
				Data:        line2Values,
				Color:       nk.NkRgb(0, 0, 100),
				ActiveColor: nk.NkRgb(0, 0, 200),
			},
			gimu.ChartSeries{
				ChartType:   nk.ChartColumn,
				Min:         0,
				Max:         10,
				Data:        line3Values,
				Color:       nk.NkRgb(100, 0, 100),
				ActiveColor: nk.NkRgb(100, 0, 200),
			},
		})
	})
}

func main() {
	runtime.LockOSThread()

	lineValues = make([]float32, 100)
	line2Values = make([]float32, 100)
	line3Values = make([]float32, 100)

	rand.Seed(42)
	for i := range lineValues {
		lineValues[i] = float32(rand.Intn(100))
		line2Values[i] = float32(rand.Intn(100))
		line3Values[i] = float32(rand.Intn(10))
	}

	wnd := gimu.NewMasterWindow("Chart Demo", 800, 600, gimu.MasterWindowFlagNoResize)
	wnd.Main(builder)
}
