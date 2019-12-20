package gimu

import (
	"github.com/AllenDang/gimu/nk"
)

type row struct {
	ctx    *nk.Context
	height int
}

func (r *row) Dynamic(col int) {
	nk.NkLayoutRowDynamic(r.ctx, float32(r.height), int32(col))
}

func (r *row) Static(width ...int) {
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

func (r *row) Ratio(ratio ...float32) {
	nk.NkLayoutRowBegin(r.ctx, nk.LayoutDynamic, float32(r.height), int32(len(ratio)))

	for _, rt := range ratio {
		nk.NkLayoutRowPush(r.ctx, rt)
	}

	nk.NkLayoutRowEnd(r.ctx)

}
