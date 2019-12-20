package nk

/*
#include "nuklear.h"
*/
import "C"
import (
	"bytes"
	"image"
	"unsafe"

	"github.com/go-gl/gl/v3.2-core/gl"
)

var (
	clipboardPlugin ClipboardPlugin
)

type ClipboardPlugin interface {
	GetText() string
	SetText(content string)
}

//export igClipboardPaste
func igClipboardPaste(user C.nk_handle, edit *TextEdit) {
	if clipboardPlugin != nil {
		content := clipboardPlugin.GetText()
		NkTexteditPaste(edit, content, int32(len(content)))
	}
}

//export igClipboardCopy
func igClipboardCopy(user C.nk_handle, text *C.char, len C.int) {
	if clipboardPlugin != nil {
		clipboardPlugin.SetText(C.GoStringN(text, len))
	}
}

func (ctx *Context) SetClipboard(board ClipboardPlugin) {
	clipboardPlugin = board
	NkRegisterClipboard(ctx)
}

func (ctx *Context) Input() *Input {
	return (*Input)(&ctx.input)
}

func (ctx *Context) Current() *Window {
	return (*Window)(ctx.current)
}

func (ctx *Context) Style() *Style {
	return (*Style)(&ctx.style)
}

func (ctx *Context) Memory() *Buffer {
	return (*Buffer)(&ctx.memory)
}

func (ctx *Context) Clip() *Clipboard {
	return (*Clipboard)(&ctx.clip)
}

func (ctx *Context) LastWidgetState() Flags {
	return (Flags)(ctx.last_widget_state)
}

func (ctx *Context) DeltaTimeSeconds() float32 {
	return (float32)(ctx.delta_time_seconds)
}

func (ctx *Context) ButtonBehavior() ButtonBehavior {
	return (ButtonBehavior)(ctx.button_behavior)
}

func (ctx *Context) Stacks() *ConfigurationStacks {
	return (*ConfigurationStacks)(&ctx.stacks)
}

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

func NkFontAtlasAddFromBytes(atlas *FontAtlas, data []byte, height float32, config *FontConfig) *Font {
	dataPtr := unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)
	return NkFontAtlasAddFromMemory(atlas, dataPtr, Size(len(data)), height, config)
}

func (fc *FontConfig) SetPixelSnap(b bool) {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	fc.pixel_snap = (C.uchar)(i)
}

func (fc *FontConfig) SetOversample(v, h int) {
	fc.oversample_v = (C.uchar)(v)
	fc.oversample_h = (C.uchar)(h)
}

func (fc *FontConfig) SetRange(r *Rune) {
	fc._range = (*C.nk_rune)(r)
}

func (fc *FontConfig) SetRangeGoRune(r []rune) {
	fc._range = (*C.nk_rune)(unsafe.Pointer(&r[0]))
}

func (h Handle) ID() int {
	return int(*(*int64)(unsafe.Pointer(&h)))
}

func (h Handle) Ptr() uintptr {
	return *(*uintptr)(unsafe.Pointer(&h))
}

func (atlas *FontAtlas) DefaultFont() *Font {
	return (*Font)(atlas.default_font)
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

func (f Font) Handle() *UserFont {
	return NewUserFontRef(unsafe.Pointer(&f.handle))
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

// Allocated is the total amount of memory allocated.
func (b *Buffer) Allocated() int {
	return (int)(b.allocated)
}

// Size is the current size of the buffer.
func (b *Buffer) Size() int {
	return (int)(b.size)
}

// Type is the memory management type of the buffer.
func (b *Buffer) Type() AllocationType {
	return (AllocationType)(b._type)
}

func (l *ListView) Begin() int {
	return (int)(l.begin)
}

func (l *ListView) End() int {
	return (int)(l.end)
}

func (l *ListView) Count() int {
	return (int)(l.count)
}

func (panel *Panel) Bounds() *Rect {
	return (*Rect)(&panel.bounds)
}

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

type Theme int

const (
	THEME_WHITE Theme = iota
	THEME_DARK
	THEME_RED
	THEME_BLUE
)

func (ctx *Context) SetStyle(theme Theme) {
	table := make([]Color, ColorCount)

	switch theme {
	case THEME_WHITE:
		table[ColorText] = NkRgba(70, 70, 70, 255)
		table[ColorWindow] = NkRgba(175, 175, 175, 255)
		table[ColorHeader] = NkRgba(175, 175, 175, 255)
		table[ColorBorder] = NkRgba(0, 0, 0, 255)
		table[ColorButton] = NkRgba(185, 185, 185, 255)
		table[ColorButtonHover] = NkRgba(170, 170, 170, 255)
		table[ColorButtonActive] = NkRgba(160, 160, 160, 255)
		table[ColorToggle] = NkRgba(150, 150, 150, 255)
		table[ColorToggleHover] = NkRgba(120, 120, 120, 255)
		table[ColorToggleCursor] = NkRgba(175, 175, 175, 255)
		table[ColorSelect] = NkRgba(190, 190, 190, 255)
		table[ColorSelectActive] = NkRgba(175, 175, 175, 255)
		table[ColorSlider] = NkRgba(190, 190, 190, 255)
		table[ColorSliderCursor] = NkRgba(80, 80, 80, 255)
		table[ColorSliderCursorHover] = NkRgba(70, 70, 70, 255)
		table[ColorSliderCursorActive] = NkRgba(60, 60, 60, 255)
		table[ColorProperty] = NkRgba(175, 175, 175, 255)
		table[ColorEdit] = NkRgba(150, 150, 150, 255)
		table[ColorEditCursor] = NkRgba(0, 0, 0, 255)
		table[ColorCombo] = NkRgba(175, 175, 175, 255)
		table[ColorChart] = NkRgba(160, 160, 160, 255)
		table[ColorChartColor] = NkRgba(45, 45, 45, 255)
		table[ColorChartColorHighlight] = NkRgba(255, 0, 0, 255)
		table[ColorScrollbar] = NkRgba(180, 180, 180, 255)
		table[ColorScrollbarCursor] = NkRgba(140, 140, 140, 255)
		table[ColorScrollbarCursorHover] = NkRgba(150, 150, 150, 255)
		table[ColorScrollbarCursorActive] = NkRgba(160, 160, 160, 255)
		table[ColorTabHeader] = NkRgba(180, 180, 180, 255)
		NkStyleFromTable(ctx, table)
	case THEME_DARK:
		table[ColorText] = NkRgba(210, 210, 210, 255)
		table[ColorWindow] = NkRgba(57, 67, 71, 215)
		table[ColorHeader] = NkRgba(51, 51, 56, 220)
		table[ColorBorder] = NkRgba(46, 46, 46, 255)
		table[ColorButton] = NkRgba(48, 83, 111, 255)
		table[ColorButtonHover] = NkRgba(58, 93, 121, 255)
		table[ColorButtonActive] = NkRgba(63, 98, 126, 255)
		table[ColorToggle] = NkRgba(50, 58, 61, 255)
		table[ColorToggleHover] = NkRgba(45, 53, 56, 255)
		table[ColorToggleCursor] = NkRgba(48, 83, 111, 255)
		table[ColorSelect] = NkRgba(57, 67, 61, 255)
		table[ColorSelectActive] = NkRgba(48, 83, 111, 255)
		table[ColorSlider] = NkRgba(50, 58, 61, 255)
		table[ColorSliderCursor] = NkRgba(48, 83, 111, 245)
		table[ColorSliderCursorHover] = NkRgba(53, 88, 116, 255)
		table[ColorSliderCursorActive] = NkRgba(58, 93, 121, 255)
		table[ColorProperty] = NkRgba(50, 58, 61, 255)
		table[ColorEdit] = NkRgba(50, 58, 61, 225)
		table[ColorEditCursor] = NkRgba(210, 210, 210, 255)
		table[ColorCombo] = NkRgba(50, 58, 61, 255)
		table[ColorChart] = NkRgba(50, 58, 61, 255)
		table[ColorChartColor] = NkRgba(48, 83, 111, 255)
		table[ColorChartColorHighlight] = NkRgba(255, 0, 0, 255)
		table[ColorScrollbar] = NkRgba(50, 58, 61, 255)
		table[ColorScrollbarCursor] = NkRgba(48, 83, 111, 255)
		table[ColorScrollbarCursorHover] = NkRgba(53, 88, 116, 255)
		table[ColorScrollbarCursorActive] = NkRgba(58, 93, 121, 255)
		table[ColorTabHeader] = NkRgba(48, 83, 111, 255)
		NkStyleFromTable(ctx, table)
	case THEME_RED:
		table[ColorText] = NkRgba(190, 190, 190, 255)
		table[ColorWindow] = NkRgba(30, 33, 40, 215)
		table[ColorHeader] = NkRgba(181, 45, 69, 220)
		table[ColorBorder] = NkRgba(51, 55, 67, 255)
		table[ColorButton] = NkRgba(181, 45, 69, 255)
		table[ColorButtonHover] = NkRgba(190, 50, 70, 255)
		table[ColorButtonActive] = NkRgba(195, 55, 75, 255)
		table[ColorToggle] = NkRgba(51, 55, 67, 255)
		table[ColorToggleHover] = NkRgba(45, 60, 60, 255)
		table[ColorToggleCursor] = NkRgba(181, 45, 69, 255)
		table[ColorSelect] = NkRgba(51, 55, 67, 255)
		table[ColorSelectActive] = NkRgba(181, 45, 69, 255)
		table[ColorSlider] = NkRgba(51, 55, 67, 255)
		table[ColorSliderCursor] = NkRgba(181, 45, 69, 255)
		table[ColorSliderCursorHover] = NkRgba(186, 50, 74, 255)
		table[ColorSliderCursorActive] = NkRgba(191, 55, 79, 255)
		table[ColorProperty] = NkRgba(51, 55, 67, 255)
		table[ColorEdit] = NkRgba(51, 55, 67, 225)
		table[ColorEditCursor] = NkRgba(190, 190, 190, 255)
		table[ColorCombo] = NkRgba(51, 55, 67, 255)
		table[ColorChart] = NkRgba(51, 55, 67, 255)
		table[ColorChartColor] = NkRgba(170, 40, 60, 255)
		table[ColorChartColorHighlight] = NkRgba(255, 0, 0, 255)
		table[ColorScrollbar] = NkRgba(30, 33, 40, 255)
		table[ColorScrollbarCursor] = NkRgba(64, 84, 95, 255)
		table[ColorScrollbarCursorHover] = NkRgba(70, 90, 100, 255)
		table[ColorScrollbarCursorActive] = NkRgba(75, 95, 105, 255)
		table[ColorTabHeader] = NkRgba(181, 45, 69, 220)
		NkStyleFromTable(ctx, table)
	case THEME_BLUE:
		table[ColorText] = NkRgba(20, 20, 20, 255)
		table[ColorWindow] = NkRgba(202, 212, 214, 215)
		table[ColorHeader] = NkRgba(137, 182, 224, 220)
		table[ColorBorder] = NkRgba(140, 159, 173, 255)
		table[ColorButton] = NkRgba(137, 182, 224, 255)
		table[ColorButtonHover] = NkRgba(142, 187, 229, 255)
		table[ColorButtonActive] = NkRgba(147, 192, 234, 255)
		table[ColorToggle] = NkRgba(177, 210, 210, 255)
		table[ColorToggleHover] = NkRgba(182, 215, 215, 255)
		table[ColorToggleCursor] = NkRgba(137, 182, 224, 255)
		table[ColorSelect] = NkRgba(177, 210, 210, 255)
		table[ColorSelectActive] = NkRgba(137, 182, 224, 255)
		table[ColorSlider] = NkRgba(177, 210, 210, 255)
		table[ColorSliderCursor] = NkRgba(137, 182, 224, 245)
		table[ColorSliderCursorHover] = NkRgba(142, 188, 229, 255)
		table[ColorSliderCursorActive] = NkRgba(147, 193, 234, 255)
		table[ColorProperty] = NkRgba(210, 210, 210, 255)
		table[ColorEdit] = NkRgba(210, 210, 210, 225)
		table[ColorEditCursor] = NkRgba(20, 20, 20, 255)
		table[ColorCombo] = NkRgba(210, 210, 210, 255)
		table[ColorChart] = NkRgba(210, 210, 210, 255)
		table[ColorChartColor] = NkRgba(137, 182, 224, 255)
		table[ColorChartColorHighlight] = NkRgba(255, 0, 0, 255)
		table[ColorScrollbar] = NkRgba(190, 200, 200, 255)
		table[ColorScrollbarCursor] = NkRgba(64, 84, 95, 255)
		table[ColorScrollbarCursorHover] = NkRgba(70, 90, 100, 255)
		table[ColorScrollbarCursorActive] = NkRgba(75, 95, 105, 255)
		table[ColorTabHeader] = NkRgba(156, 193, 220, 255)
		NkStyleFromTable(ctx, table)
	}
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
