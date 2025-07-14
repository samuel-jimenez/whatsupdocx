package dmlct

import (
	"github.com/samuel-jimenez/whatsupdocx/common/units"
)

// Complex Type: CT_PositiveSize2D
type PSize2D struct {
	Width  uint64 `xml:"cx,attr"`
	Height uint64 `xml:"cy,attr"`
}

func NewPostvSz2D(width units.Emu, height units.Emu) *PSize2D {
	return &PSize2D{
		Height: uint64(height),
		Width:  uint64(width),
	}
}
