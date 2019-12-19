package gimu

import (
	"image"
	"image/color"

	"github.com/AllenDang/gimu/nk"
)

type Canvas struct {
	buffer *nk.CommandBuffer
}

func (c *Canvas) FillRect(rect image.Rectangle, rounding float32, color color.RGBA) {
	nk.NkFillRect(c.buffer, toNkRect(rect), rounding, toNkColor(color))
}

func (c *Canvas) FillCircle(rect image.Rectangle, color color.RGBA) {
	nk.NkFillCircle(c.buffer, toNkRect(rect), toNkColor(color))
}

func (c *Canvas) FillTriangle(p1, p2, p3 image.Point, color color.RGBA) {
	nk.NkFillTriangle(c.buffer, float32(p1.X), float32(p1.Y), float32(p2.X), float32(p2.Y), float32(p3.X), float32(p3.Y), toNkColor(color))
}

func (c *Canvas) FillPolygon(points []float32, color color.RGBA) {
	nk.NkFillPolygon(c.buffer, points, int32(len(points)), toNkColor(color))
}