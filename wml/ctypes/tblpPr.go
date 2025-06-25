package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// CT_TblPPr
// w_CT_TblPPr =
type FloatPos struct {
	// attribute w:leftFromText { s_ST_TwipsMeasure }?,
	LeftFromText *uint64 `xml:"w:leftFromText,attr,omitempty"`
	// attribute w:rightFromText { s_ST_TwipsMeasure }?,
	RightFromText *uint64 `xml:"w:rightFromText,attr,omitempty"`
	// attribute w:topFromText { s_ST_TwipsMeasure }?,
	TopFromText *uint64 `xml:"w:topFromText,attr,omitempty"`
	// attribute w:bottomFromText { s_ST_TwipsMeasure }?,
	BottomFromText *uint64 `xml:"w:bottomFromText,attr,omitempty"`

	//Frame Vertical Positioning Base
	// attribute w:vertAnchor { w_ST_VAnchor }?,
	VAnchor *stypes.Anchor `xml:"w:vertAnchor,attr,omitempty"`

	//Frame Horizontal Positioning Base
	// attribute w:horzAnchor { w_ST_HAnchor }?,
	HAnchor *stypes.Anchor `xml:"w:horzAnchor,attr,omitempty"`

	//Relative Horizontal Alignment From Anchor
	// attribute w:tblpXSpec { s_ST_XAlign }?,
	XAlign *stypes.XAlign `xml:"w:tblpXSpec,attr,omitempty"`

	//Relative Vertical Alignment from Anchor
	// attribute w:tblpYSpec { s_ST_YAlign }?,
	YAlign *stypes.YAlign `xml:"w:tblpYSpec,attr,omitempty"`

	//Absolute Horizontal Distance From Anchor
	// attribute w:tblpX { w_ST_SignedTwipsMeasure }?,
	AbsXDist *int `xml:"w:tblpX,attr,omitempty"`

	// Absolute Vertical Distance From Anchor
	// attribute w:tblpY { w_ST_SignedTwipsMeasure }?
	AbsYDist *int `xml:"w:tblpY,attr,omitempty"`
}
