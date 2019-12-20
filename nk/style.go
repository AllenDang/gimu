package nk

/*
#include "nuklear.h"
*/
import "C"

import "bytes"

func (t *StyleText) Color() *Color {
	return (*Color)(&t.color)
}

func (s *Style) Text() *StyleText {
	return (*StyleText)(&s.text)
}

func (s *Style) Window() *StyleWindow {
	return (*StyleWindow)(&s.window)
}

func (s *Style) Combo() *StyleCombo {
	return (*StyleCombo)(&s.combo)
}

func (w *StyleWindow) Background() *Color {
	return (*Color)(&w.background)
}

func (w *StyleWindow) Spacing() *Vec2 {
	return (*Vec2)(&w.spacing)
}

func (w *StyleWindow) Padding() *Vec2 {
	return (*Vec2)(&w.padding)
}

func (w *StyleWindow) GroupPadding() *Vec2 {
	return (*Vec2)(&w.group_padding)
}

func SetSpacing(ctx *Context, v Vec2) {
	*ctx.Style().Window().Spacing() = v
}

func SetPadding(ctx *Context, v Vec2) {
	*ctx.Style().Window().Padding() = v
}

func SetGroupPadding(ctx *Context, v Vec2) {
	*ctx.Style().Window().GroupPadding() = v
}

func SetTextColor(ctx *Context, color Color) {
	*ctx.Style().Text().Color() = color
}

func SetBackgroundColor(ctx *Context, color Color) {
	ctx.Style().Window().fixed_background = C.struct_nk_style_item(NkStyleItemColor(color))
}

func (t *TextEdit) GetGoString() string {
	nkstr := t.GetString()
	b := C.GoBytes(*nkstr.GetBuffer().GetMemory().GetPtr(), C.int(*nkstr.GetBuffer().GetSize()))
	r := bytes.Runes(b)[:*nkstr.GetLen()]
	return string(r)
}
