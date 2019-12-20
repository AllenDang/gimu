package nk

/*
#include "nuklear.h"
*/
import "C"

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
