package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Border Properties (CT_Border)
// w_CT_Border =
type Border struct {
	// attribute w:val { w_ST_Border },
	Val stypes.BorderStyle `xml:"val,attr"`
	// attribute w:color { w_ST_HexColor }?,
	// ## default value: auto
	Color *string `xml:"color,attr,omitempty"`
	// attribute w:themeColor { w_ST_ThemeColor }?,
	ThemeColor *stypes.ThemeColor `xml:"themeColor,attr,omitempty"`
	// attribute w:themeTint { w_ST_UcharHexNumber }?,
	ThemeTint *string `xml:"themeTint,attr,omitempty"`
	// attribute w:themeShade { w_ST_UcharHexNumber }?,
	ThemeShade *string `xml:"themeShade,attr,omitempty"`
	// attribute w:sz { w_ST_EighthPointMeasure }?,
	Size *int `xml:"sz,attr,omitempty"`
	// attribute w:space { w_ST_PointMeasure }?,
	// ## default value: 0
	Space *string `xml:"space,attr,omitempty"`
	// attribute w:shadow { s_ST_OnOff }?,
	Shadow *stypes.OnOff `xml:"shadow,attr,omitempty"`
	// attribute w:frame { s_ST_OnOff }?
	Frame *stypes.OnOff `xml:"frame,attr,omitempty"`
}

/*
new cell border
@param style: border style
@param color: border color; @see github.com/samuel-jimenez/whatsupdocx/wml/color
@param space: the gap between the border and the cell content
@param size:  border size
*/
func NewCellBorder(style stypes.BorderStyle, color string, space string, size int) *Border {
	if size < 0 {
		size = 0
	}
	return &Border{
		Val:   style,
		Color: &color,
		Space: &space,
		Size:  &size,
	}
}

func (t Border) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(t.Val)})

	if t.Color != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"}, Value: *t.Color})
	}
	if t.ThemeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeColor"}, Value: string(*t.ThemeColor)})
	}
	if t.ThemeTint != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"}, Value: *t.ThemeTint})
	}
	if t.ThemeShade != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"}, Value: *t.ThemeShade})
	}
	if t.Size != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:sz"}, Value: strconv.Itoa(*t.Size)})
	}
	if t.Space != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:space"}, Value: *t.Space})
	}
	if t.Shadow != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:shadow"}, Value: string(*t.Shadow)})
	}
	if t.Frame != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:frame"}, Value: string(*t.Frame)})
	}

	return e.EncodeElement("", start)
}
