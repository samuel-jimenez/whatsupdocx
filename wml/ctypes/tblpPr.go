package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// CT_TblPPr
// w_CT_TblPPr =
type FloatPos struct {
	// attribute w:leftFromText { s_ST_TwipsMeasure }?,
	LeftFromText *uint64 `xml:"leftFromText,attr,omitempty"`
	// attribute w:rightFromText { s_ST_TwipsMeasure }?,
	RightFromText *uint64 `xml:"rightFromText,attr,omitempty"`
	// attribute w:topFromText { s_ST_TwipsMeasure }?,
	TopFromText *uint64 `xml:"topFromText,attr,omitempty"`
	// attribute w:bottomFromText { s_ST_TwipsMeasure }?,
	BottomFromText *uint64 `xml:"bottomFromText,attr,omitempty"`

	//Frame Vertical Positioning Base
	// attribute w:vertAnchor { w_ST_VAnchor }?,
	VAnchor *stypes.Anchor `xml:"vertAnchor,attr,omitempty"`

	//Frame Horizontal Positioning Base
	// attribute w:horzAnchor { w_ST_HAnchor }?,
	HAnchor *stypes.Anchor `xml:"horzAnchor,attr,omitempty"`

	//Relative Horizontal Alignment From Anchor
	// attribute w:tblpXSpec { s_ST_XAlign }?,
	XAlign *stypes.XAlign `xml:"tblpXSpec,attr,omitempty"`

	//Relative Vertical Alignment from Anchor
	// attribute w:tblpYSpec { s_ST_YAlign }?,
	YAlign *stypes.YAlign `xml:"tblpYSpec,attr,omitempty"`

	//Absolute Horizontal Distance From Anchor
	// attribute w:tblpX { w_ST_SignedTwipsMeasure }?,
	AbsXDist *int `xml:"tblpX,attr,omitempty"`

	// Absolute Vertical Distance From Anchor
	// attribute w:tblpY { w_ST_SignedTwipsMeasure }?
	AbsYDist *int `xml:"tblpY,attr,omitempty"`
}

func (t FloatPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblpPr"

	if t.LeftFromText != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:leftFromText"}, Value: strconv.FormatUint(*t.LeftFromText, 10)})
	}
	if t.RightFromText != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rightFromText"}, Value: strconv.FormatUint(*t.RightFromText, 10)})
	}
	if t.TopFromText != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:topFromText"}, Value: strconv.FormatUint(*t.TopFromText, 10)})
	}
	if t.BottomFromText != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:bottomFromText"}, Value: strconv.FormatUint(*t.BottomFromText, 10)})
	}
	if t.VAnchor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vertAnchor"}, Value: string(*t.VAnchor)})
	}
	if t.HAnchor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:horzAnchor"}, Value: string(*t.HAnchor)})
	}
	if t.XAlign != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpXSpec"}, Value: string(*t.XAlign)})
	}
	if t.YAlign != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpYSpec"}, Value: string(*t.YAlign)})
	}
	if t.AbsXDist != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpX"}, Value: strconv.Itoa(*t.AbsXDist)})
	}
	if t.AbsYDist != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpY"}, Value: strconv.Itoa(*t.AbsYDist)})
	}

	return e.EncodeElement("", start)
}
