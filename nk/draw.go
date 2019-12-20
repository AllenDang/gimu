package nk

/*
#include "nuklear.h"
*/
import "C"

import (
	"image"
	"unsafe"

	"github.com/go-gl/gl/v3.2-core/gl"
)

var VertexLayoutEnd = DrawVertexLayoutElement{
	Attribute: VertexAttributeCount,
	Format:    FormatCount,
	Offset:    0,
}

func NkDrawForeach(ctx *Context, b *Buffer, fn func(cmd *DrawCommand)) {
	cmd := Nk_DrawBegin(ctx, b)
	for {
		if cmd == nil {
			break
		}
		fn(cmd)
		cmd = Nk_DrawNext(cmd, b, ctx)
	}
}

func (h Handle) ID() int {
	return int(*(*int64)(unsafe.Pointer(&h)))
}

func (h Handle) Ptr() uintptr {
	return *(*uintptr)(unsafe.Pointer(&h))
}

func (c Color) R() Byte {
	return Byte(c.r)
}

func (c Color) G() Byte {
	return Byte(c.g)
}

func (c Color) B() Byte {
	return Byte(c.b)
}

func (c Color) A() Byte {
	return Byte(c.a)
}

func (c Color) RGBA() (Byte, Byte, Byte, Byte) {
	return Byte(c.r), Byte(c.g), Byte(c.b), Byte(c.a)
}

func (c Color) RGBAi() (int32, int32, int32, int32) {
	return int32(c.r), int32(c.g), int32(c.b), int32(c.a)
}

func (c *Color) SetR(r Byte) {
	c.r = (C.nk_byte)(r)
}

func (c *Color) SetG(g Byte) {
	c.g = (C.nk_byte)(g)
}

func (c *Color) SetB(b Byte) {
	c.b = (C.nk_byte)(b)
}

func (c *Color) SetA(a Byte) {
	c.a = (C.nk_byte)(a)
}

func (c *Color) SetRGBA(r, g, b, a Byte) {
	c.r = (C.nk_byte)(r)
	c.g = (C.nk_byte)(g)
	c.b = (C.nk_byte)(b)
	c.a = (C.nk_byte)(a)
}

func (c *Color) SetRGBAi(r, g, b, a int32) {
	c.r = (C.nk_byte)(r)
	c.g = (C.nk_byte)(g)
	c.b = (C.nk_byte)(b)
	c.a = (C.nk_byte)(a)
}

func (cmd *DrawCommand) ElemCount() int {
	return (int)(cmd.elem_count)
}

func (cmd *DrawCommand) Texture() Handle {
	return (Handle)(cmd.texture)
}

func (cmd *DrawCommand) ClipRect() *Rect {
	return (*Rect)(&cmd.clip_rect)
}

func (r *Rect) X() float32 {
	return (float32)(r.x)
}

func (r *Rect) Y() float32 {
	return (float32)(r.y)
}

func (r *Rect) W() float32 {
	return (float32)(r.w)
}

func (r *Rect) H() float32 {
	return (float32)(r.h)
}

func (v *Vec2) X() float32 {
	return (float32)(v.x)
}

func (v *Vec2) Y() float32 {
	return (float32)(v.y)
}

func (v *Vec2) SetX(x float32) {
	v.x = (C.float)(x)
}

func (v *Vec2) SetY(y float32) {
	v.y = (C.float)(y)
}

func (v *Vec2) Reset() {
	v.x = 0
	v.y = 0
}

func RgbaToNkImage(tex *uint32, rgba *image.RGBA) Image {
	if tex == nil {
		gl.GenTextures(1, tex)
	}
	gl.BindTexture(gl.TEXTURE_2D, *tex)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR_MIPMAP_NEAREST)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(rgba.Bounds().Dx()), int32(rgba.Bounds().Dy()),
		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&rgba.Pix[0]))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	return NkImageId(int32(*tex))
}
