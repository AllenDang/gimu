package nk

/*
#include "nuklear.h"
*/
import "C"

import "unsafe"

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

func (atlas *FontAtlas) DefaultFont() *Font {
	return (*Font)(atlas.default_font)
}

func (f Font) Handle() *UserFont {
	return NewUserFontRef(unsafe.Pointer(&f.handle))
}
