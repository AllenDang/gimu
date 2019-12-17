package gimu

import (
	"image/color"

	"github.com/AllenDang/nuklear/nk"
)

func toNkFlag(align string) nk.Flags {
	var a nk.Flags
	switch align {
	case "RT":
		a = nk.TextAlignRight | nk.TextAlignTop
	case "RC":
		a = nk.TextAlignRight | nk.TextAlignMiddle
	case "RB":
		a = nk.TextAlignRight | nk.TextAlignBottom
	case "CT":
		a = nk.TextAlignCentered | nk.TextAlignTop
	case "CC":
		a = nk.TextAlignCentered | nk.TextAlignMiddle
	case "CB":
		a = nk.TextAlignCentered | nk.TextAlignBottom
	case "LT":
		a = nk.TextAlignLeft | nk.TextAlignTop
	case "LB":
		a = nk.TextAlignLeft | nk.TextAlignBottom
	case "LC":
		a = nk.TextAlignLeft | nk.TextAlignMiddle
	default:
		a = nk.TextAlignLeft | nk.TextAlignMiddle
	}

	return a
}

func toNkColor(c color.RGBA) nk.Color {
	nc := nk.NewColor()
	nc.SetRGBA(nk.Byte(c.R), nk.Byte(c.G), nk.Byte(c.B), nk.Byte(c.A))
	return *nc
}
