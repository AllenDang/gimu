package gimu

import "github.com/AllenDang/nuklear/nk"

type row struct {
	ctx    *nk.Context
	height float32
}

func (r *row) Dynamic(col int32) {
	nk.NkLayoutRowDynamic(r.ctx, r.height, col)
}

func (r *row) Static(width ...float32) {
	nk.NkLayoutRowBegin(r.ctx, nk.LayoutStatic, r.height, int32(len(width)))

	for _, w := range width {
		nk.NkLayoutRowPush(r.ctx, w)
	}

	nk.NkLayoutRowEnd(r.ctx)
}

func (r *row) Ratio(ratio ...float32) {
	nk.NkLayoutRowBegin(r.ctx, nk.LayoutDynamic, r.height, int32(len(ratio)))

	for _, rt := range ratio {
		nk.NkLayoutRowPush(r.ctx, rt)
	}

	nk.NkLayoutRowEnd(r.ctx)

}
