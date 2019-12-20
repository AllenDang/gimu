package gimu

import (
	"github.com/AllenDang/gimu/nk"
)

type Input struct {
	input *nk.Input
}

func (i *Input) IsMouseHoveringRect(rect nk.Rect) bool {
	return nk.NkInputIsMouseHoveringRect(i.input, rect) > 0
}

func (i *Input) IsMousePrevHoveringRect(rect nk.Rect) bool {
	return nk.NkInputIsMousePrevHoveringRect(i.input, rect) > 0
}

func (i *Input) IsMouseDown(buttons nk.Buttons) bool {
	return nk.NkInputIsMouseDown(i.input, buttons) > 0
}

func (i *Input) Mouse() *nk.Mouse {
	return i.input.Mouse()
}
