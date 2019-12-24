package gimu

import (
	"github.com/AllenDang/gimu/nk"
)

type Row struct {
	win    *Window
	ctx    *nk.Context
	height int
}

func (r *Row) Dynamic(col int) {
	nk.NkLayoutRowDynamic(r.ctx, float32(r.height), int32(col))
}

func (r *Row) Static(width ...int) {
	nk.NkLayoutRowTemplateBegin(r.ctx, float32(r.height))

	for _, w := range width {
		if w == 0 {
			nk.NkLayoutRowTemplatePushDynamic(r.ctx)
		} else {
			nk.NkLayoutRowTemplatePushStatic(r.ctx, float32(w))
		}
	}

	nk.NkLayoutRowTemplateEnd(r.ctx)
}

func (r *Row) Ratio(ratio ...float32) {
	nk.NkLayoutRowBegin(r.ctx, nk.LayoutDynamic, float32(r.height), int32(len(ratio)))

	for _, rt := range ratio {
		nk.NkLayoutRowPush(r.ctx, rt)
	}

	nk.NkLayoutRowEnd(r.ctx)

}

func (r *Row) Space(spaceType nk.LayoutFormat, builder BuilderFunc) {
	nk.NkLayoutSpaceBegin(r.ctx, spaceType, float32(r.height), 2147483647)

	builder(r.win)

	nk.NkLayoutSpaceEnd(r.ctx)
}
