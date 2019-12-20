package nk

/*
#include "nuklear.h"
*/
import "C"

func (input *Input) Mouse() *Mouse {
	return (*Mouse)(&input.mouse)
}

func (input *Input) Keyboard() *Keyboard {
	return (*Keyboard)(&input.keyboard)
}

func (keyboard *Keyboard) Text() string {
	return C.GoStringN(&keyboard.text[0], keyboard.text_len)
}

func (mouse *Mouse) Grab() bool {
	return mouse.grab == True
}

func (mouse *Mouse) Grabbed() bool {
	return mouse.grabbed == True
}

func (mouse *Mouse) Ungrab() bool {
	return mouse.ungrab == True
}

func (mouse *Mouse) ScrollDelta() Vec2 {
	return (Vec2)(mouse.scroll_delta)
}

func (mouse *Mouse) Pos() (int32, int32) {
	return (int32)(mouse.pos.x), (int32)(mouse.pos.y)
}

func (mouse *Mouse) SetPos(x, y int32) {
	mouse.pos.x = (C.float)(x)
	mouse.pos.y = (C.float)(y)
}

func (mouse *Mouse) Prev() (int32, int32) {
	return (int32)(mouse.prev.x), (int32)(mouse.prev.y)
}

func (mouse *Mouse) Delta() (int32, int32) {
	return (int32)(mouse.delta.x), (int32)(mouse.delta.y)
}
