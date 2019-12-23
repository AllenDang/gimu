package nk

/*
#include "nuklear.h"
*/
import "C"

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
