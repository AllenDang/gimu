package gimu

import "github.com/AllenDang/gimu/nk"

type row struct {
	ctx    *nk.Context
	height float32
}

func (r *row) Dynamic(col int32) {
	nk.NkLayoutRowDynamic(r.ctx, r.height, col)
}

func (r *row) Static(width ...float32) {
	nk.NkLayoutRowTemplateBegin(r.ctx, r.height)

	for _, w := range width {
		if w == 0 {
			nk.NkLayoutRowTemplatePushDynamic(r.ctx)
		} else {
			nk.NkLayoutRowTemplatePushStatic(r.ctx, w)
		}
	}

	nk.NkLayoutRowTemplateEnd(r.ctx)
}

func (r *row) Ratio(ratio ...float32) {
	nk.NkLayoutRowBegin(r.ctx, nk.LayoutDynamic, r.height, int32(len(ratio)))

	for _, rt := range ratio {
		nk.NkLayoutRowPush(r.ctx, rt)
	}

	nk.NkLayoutRowEnd(r.ctx)

}
