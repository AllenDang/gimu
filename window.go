package gimu

import (
	"fmt"
	"image/color"

	"github.com/AllenDang/gimu/nk"
)

type Window struct {
	ctx *nk.Context
	mw  *MasterWindow
}

func (w *Window) MasterWindow() *MasterWindow {
	return w.mw
}

func (w *Window) Window(title string, bounds nk.Rect, flags nk.Flags, builder BuilderFunc) {
	if nk.NkBegin(w.ctx, title, bounds, flags) > 0 {
		builder(w)
		nk.NkEnd(w.ctx)
	}
}

func (w *Window) Row(height int) *Row {
	return &Row{
		win:    w,
		ctx:    w.ctx,
		height: height,
	}
}

func (w *Window) Push(rect nk.Rect) {
	nk.NkLayoutSpacePush(w.ctx, rect)
}

func (w *Window) Spacing(cols int) {
	nk.NkSpacing(w.ctx, int32(cols))
}

func (w *Window) L(content string) {
	w.Label(content, "LC")
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

func (w *Window) ButtonColor(color nk.Color) bool {
	return nk.NkButtonColor(w.ctx, color) > 0
}

func (w *Window) ButtonImage(img nk.Image) bool {
	return nk.NkButtonImage(w.ctx, img) > 0
}

func (w *Window) ButtonImageLabel(img nk.Image, label, align string) bool {
	return nk.NkButtonImageLabel(w.ctx, img, label, toNkFlag(align)) > 0
}

func (w *Window) ButtonSymbol(symbol nk.SymbolType) bool {
	return nk.NkButtonSymbol(w.ctx, symbol) > 0
}

func (w *Window) ButtonSymbolLabel(symbol nk.SymbolType, label, align string) bool {
	return nk.NkButtonSymbolLabel(w.ctx, symbol, label, toNkFlag(align)) > 0
}

func (w *Window) Progress(current *uint, max uint, modifiable bool) {
	nk.NkProgress(w.ctx, (*nk.Size)(current), nk.Size(max), toInt32(modifiable))
}

func (w *Window) ComboSimple(labels []string, selected int, itemHeight int, dropDownWidth, dropDownHeight float32) int {
	if dropDownWidth == 0 {
		dropDownWidth = getDynamicWidth(w.ctx)
	}
	return int(nk.NkCombo(w.ctx, labels, int32(len(labels)), int32(selected), int32(itemHeight), nk.NkVec2(dropDownWidth, dropDownHeight)))
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

func (w *Window) PropertyInt(label string, min int, val *int32, max int, step int, incPerPixel float32) {
	nk.NkPropertyInt(w.ctx, label, int32(min), val, int32(max), int32(step), incPerPixel)
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

func (w *Window) SelectableSymbolLabel(symbol nk.SymbolType, label, align string, selected *bool) {
	i := make([]int32, 1)
	i[0] = toInt32(*selected)
	nk.NkSelectableSymbolLabel(w.ctx, symbol, label, toNkFlag(align), i)
	*selected = i[0] > 0
}

func (w *Window) Popup(title string, popupType nk.PopupType, flag nk.Flags, bounds nk.Rect, builder BuilderFunc) bool {
	result := nk.NkPopupBegin(w.ctx, popupType, title, flag, bounds)
	if result > 0 {
		builder(w)
		nk.NkPopupEnd(w.ctx)
	}

	return result > 0
}

func (w *Window) ClosePopup() {
	nk.NkPopupClose(w.ctx)
}

func (w *Window) Group(title string, flag nk.Flags, builder BuilderFunc) bool {
	result := nk.NkGroupBegin(w.ctx, title, flag)
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

func (w *Window) Contextual(flag nk.Flags, width, height int, builder BuilderFunc) {
	bounds := nk.NkWidgetBounds(w.ctx)
	if nk.NkContextualBegin(w.ctx, flag, nk.NkVec2(float32(width), float32(height)), bounds) > 0 {
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

func (w *Window) WidgetBounds() nk.Rect {
	return nk.NkWidgetBounds(w.ctx)
}

func (w *Window) GetInput() *Input {
	return &Input{input: w.ctx.Input()}
}

type ListViewItemBuilder func(w *Window, i int, item interface{})

// Should only called in ListView to define the row layout.
type RowLayoutFunc func(r *Row)

func (w *Window) ListView(view *nk.ListView, id string, flags nk.Flags, rowHeight int, items []interface{}, rowLayoutFunc RowLayoutFunc, builder ListViewItemBuilder) {
	if nk.NkListViewBegin(w.ctx, view, id, flags, int32(rowHeight), int32(len(items))) > 0 {
		rowLayoutFunc(w.Row(rowHeight))

		for i := 0; i < view.Count(); i++ {
			builder(w, i, items[view.Begin()+i])
		}
		nk.NkListViewEnd(view)
	}
}

func (w *Window) Chart(chartType nk.ChartType, min, max float32, data []float32) {
	selected := -1

	if nk.NkChartBegin(w.ctx, chartType, int32(len(data)), min, max) > 0 {
		for i, d := range data {
			res := nk.NkChartPush(w.ctx, d)
			if res == nk.ChartHovering {
				selected = i
			}
		}
		nk.NkChartEnd(w.ctx)
	}

	// Show tooltip on mouse hovering
	if selected != -1 {
		nk.NkTooltip(w.ctx, fmt.Sprintf("%.2f", data[selected]))
	}
}

func (w *Window) ChartColored(chartType nk.ChartType, color, activeColor nk.Color, min, max float32, data []float32) {
	selected := -1
	if nk.NkChartBeginColored(w.ctx, chartType, color, activeColor, int32(len(data)), min, max) > 0 {
		for i, d := range data {
			res := nk.NkChartPush(w.ctx, d)
			if res == nk.ChartHovering {
				selected = i
			}
		}
		nk.NkChartEnd(w.ctx)
	}

	// Show tooltip on mouse hovering
	if selected != -1 {
		nk.NkTooltip(w.ctx, fmt.Sprintf("%.2f", data[selected]))
	}
}

type ChartSeries struct {
	ChartType   nk.ChartType
	Min, Max    float32
	Data        []float32
	Color       nk.Color
	ActiveColor nk.Color
}

func (w *Window) ChartMixed(series []ChartSeries) {
	if len(series) > 0 {
		first := series[0]
		if nk.NkChartBeginColored(w.ctx, first.ChartType, first.Color, first.ActiveColor, int32(len(first.Data)), first.Min, first.Max) > 0 {

			for i, s := range series {
				if i > 0 {
					nk.NkChartAddSlotColored(w.ctx, s.ChartType, s.Color, s.ActiveColor, int32(len(s.Data)), s.Min, s.Max)
				}
			}

			for i, s := range series {
				selected := -1
				for di, d := range s.Data {
					res := nk.NkChartPushSlot(w.ctx, d, int32(i))
					if res == nk.ChartHovering {
						selected = di
					}
				}

				if selected != -1 {
					nk.NkTooltip(w.ctx, fmt.Sprintf("%.2f", s.Data[selected]))
				}
			}

			nk.NkChartEnd(w.ctx)
		}
	}
}
