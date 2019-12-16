package gimu

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #define IM_OFFSETOF(_TYPE,_MEMBER) ((size_t)&(((_TYPE*)0)->_MEMBER))
// #cgo LDFLAGS: ./cimgui/cimgui.so
// #include "cimgui/cimgui.h"
// inline ImTextureID nativeHandleCast(uintptr_t id) {
//   return (ImTextureID)id;
// }
import "C"

func ShowDemoWindow() {
	C.igShowDemoWindow(nil)
}
