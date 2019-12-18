package gimu

import (
	"image"
	"image/color"

	"github.com/AllenDang/nuklear/nk"
)

type WindowFlag int32

const (
	WindowBorder         = nk.WindowBorder
	WindowMovable        = nk.WindowMovable
	WindowScalable       = nk.WindowScalable
	WindowClosable       = nk.WindowClosable
	WindowMinimizable    = nk.WindowMinimizable
	WindowNoScrollbar    = nk.WindowNoScrollbar
	WindowTitle          = nk.WindowTitle
	WindowScrollAutoHide = nk.WindowScrollAutoHide
	WindowBackground     = nk.WindowBackground
	WindowScaleLeft      = nk.WindowScaleLeft
	WindowNoInput        = nk.WindowNoInput
)

type Window struct {
	ctx *nk.Context
	mw  *MasterWindow
}

func (w *Window) MasterWindow() *MasterWindow {
	return w.mw
}

func (w *Window) Window(title string, bounds image.Rectangle, flags WindowFlag, builder BuilderFunc) {
	rect := nk.NkRect(float32(bounds.Min.X), float32(bounds.Min.Y), float32(bounds.Max.X), float32(bounds.Max.Y))
	if nk.NkBegin(w.ctx, title, rect, nk.Flags(flags)) > 0 {
		builder(w)
	}
	nk.NkEnd(w.ctx)
}

func (w *Window) Row(height float32) *row {
	return &row{
		ctx:    w.ctx,
		height: height,
	}
}

func (w *Window) Label(content string, align string) {
	nk.NkLabel(w.ctx, content, toNkFlag(align))
}

func (w *Window) LabelColored(content string, textColor color.RGBA, align string) {
	nk.NkLabelColored(w.ctx, content, toNkFlag(align), toNkColor(textColor))
}

func (w *Window) Button(content string) bool {
	return nk.NkButtonLabel(w.ctx, content) > 0
}

func (w *Window) Progress(current *uint, max uint, modifiable bool) {
	nk.NkProgress(w.ctx, (*nk.Size)(current), nk.Size(max), toInt32(modifiable))
}

func (w *Window) ComboSimple(labels []string, selected int32, itemHeight int32, dropDownWidth, dropDownHeight float32) int32 {
	if dropDownWidth == 0 {
		dropDownWidth = getDynamicWidth(w.ctx)
	}
	return nk.NkCombo(w.ctx, labels, int32(len(labels)), selected, itemHeight, nk.NkVec2(dropDownWidth, dropDownHeight))
}

func (w *Window) ComboLabel(label string, dropDownWidth, dropDownHeight float32, itemBuilder BuilderFunc) {
	if dropDownWidth == 0 {
		dropDownWidth = getDynamicWidth(w.ctx)
	}

	if nk.NkComboBeginLabel(w.ctx, label, nk.NkVec2(dropDownWidth, dropDownHeight)) > 0 {
		itemBuilder(w)
		nk.NkComboEnd(w.ctx)
	}
}
