package nk

/*
#include "nuklear.h"
*/
import "C"
import "bytes"

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

func (t *TextEdit) GetGoString() string {
	nkstr := t.GetString()
	b := C.GoBytes(*nkstr.GetBuffer().GetMemory().GetPtr(), C.int(*nkstr.GetBuffer().GetSize()))
	r := bytes.Runes(b)[:*nkstr.GetLen()]
	return string(r)
}
