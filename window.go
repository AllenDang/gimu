package gimu

import (
	"image"
	"image/color"

	"github.com/AllenDang/gimu/nk"
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
	if nk.NkBegin(w.ctx, title, toNkRect(bounds), nk.Flags(flags)) > 0 {
		builder(w)
		nk.NkEnd(w.ctx)
	}
}

func (w *Window) Row(height float32) *row {
	return &row{
		ctx:    w.ctx,
		height: height,
	}
}

func (w *Window) Spacing(cols int32) {
	nk.NkSpacing(w.ctx, cols)
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

func (w *Window) PropertyInt(label string, min int32, val *int32, max int32, step int32, incPerPixel float32) {
	nk.NkPropertyInt(w.ctx, label, min, val, max, step, incPerPixel)
}

func (w *Window) PropertyFloat(label string, min float32, val *float32, max float32, step float32, incPerPixel float32) {
	nk.NkPropertyFloat(w.ctx, label, min, val, max, step, incPerPixel)
}

func (w *Window) Checkbox(label string, active *bool) {
	i := toInt32(*active)
	nk.NkCheckboxLabel(w.ctx, label, &i)
	*active = i > 0
}

func (w *Window) Radio(label string, active bool) bool {
	return nk.NkOptionLabel(w.ctx, label, toInt32(active)) > 0
}

func (w *Window) SelectableLabel(label string, align string, selected *bool) {
	i := toInt32(*selected)
	nk.NkSelectableLabel(w.ctx, label, toNkFlag(align), &i)
	*selected = i > 0
}

type PopupType int32

const (
	PopupStatic  = iota
	PopupDynamic = 1
)

func (w *Window) Popup(title string, popupType PopupType, flag WindowFlag, bounds image.Rectangle, builder BuilderFunc) bool {
	result := nk.NkPopupBegin(w.ctx, nk.PopupType(popupType), title, nk.Flags(flag), toNkRect(bounds))
	if result > 0 {
		builder(w)
		nk.NkPopupEnd(w.ctx)
	}

	return result > 0
}

func (w *Window) ClosePopup() {
	nk.NkPopupClose(w.ctx)
}

func (w *Window) Group(title string, flag WindowFlag, builder BuilderFunc) bool {
	result := nk.NkGroupBegin(w.ctx, title, nk.Flags(flag))
	if result > 0 {
		builder(w)
		nk.NkGroupEnd(w.ctx)
	}

	return result > 0
}

func (w *Window) Image(texture *Texture) {
	nk.NkImage(w.ctx, texture.image)
}

func (w *Window) Menubar(builder BuilderFunc) {
	nk.NkMenubarBegin(w.ctx)
	builder(w)
	nk.NkMenubarEnd(w.ctx)
}

func (w *Window) Menu(label string, align string, width, height int, builder BuilderFunc) {
	if nk.NkMenuBeginLabel(w.ctx, label, toNkFlag(align), nk.NkVec2(float32(width), float32(height))) > 0 {
		builder(w)
		nk.NkMenuEnd(w.ctx)
	}
}

func (w *Window) MenuItemLabel(label, align string) bool {
	return nk.NkMenuItemLabel(w.ctx, label, toNkFlag(align)) > 0
}

func (w *Window) Tooltip(label string) {
	bounds := nk.NkWidgetBounds(w.ctx)
	if nk.NkInputIsMouseHoveringRect(w.ctx.Input(), bounds) > 0 {
		nk.NkTooltip(w.ctx, label)
	}
}

func (w *Window) GetCanvas() *Canvas {
	c := nk.NkWindowGetCanvas(w.ctx)
	return &Canvas{buffer: c}
}

func (w *Window) GetStyle() *nk.Style {
	return w.ctx.GetStyle()
}

func (w *Window) Contextual(flag WindowFlag, width, height int, builder BuilderFunc) {
	bounds := nk.NkWidgetBounds(w.ctx)
	if nk.NkContextualBegin(w.ctx, nk.Flags(flag), nk.NkVec2(float32(width), float32(height)), bounds) > 0 {
		builder(w)
		nk.NkContextualEnd(w.ctx)
	}
}

func (w *Window) ContextualLabel(label, align string) bool {
	return nk.NkContextualItemLabel(w.ctx, label, toNkFlag(align)) > 0
}

func (w *Window) SliderInt(min int32, val *int32, max int32, step int32) {
	nk.NkSliderInt(w.ctx, min, val, max, step)
}

func (w *Window) SliderFloat(min float32, val *float32, max float32, step float32) {
	nk.NkSliderFloat(w.ctx, min, val, max, step)
}

func (w *Window) Tree(treeType nk.TreeType, title string, initialState nk.CollapseStates, hash string, seed int32, builder BuilderFunc) {
	if nk.NkTreePushHashed(w.ctx, treeType, title, initialState, hash, int32(len(hash)), seed) > 0 {
		builder(w)
		nk.NkTreePop(w.ctx)
	}
}
