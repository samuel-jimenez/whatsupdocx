package ctypes

import (
	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type Break struct {
	BreakType *stypes.BreakType  `xml:"w:type,attr,omitempty"`
	Clear     *stypes.BreakClear `xml:"w:clear,attr,omitempty"`
}

// NewBreak creates a new Break element with the given break type.
func NewBreak(breakType stypes.BreakType) *Break {
	return &Break{
		BreakType: &breakType,
	}
}

// MarshalXML implements the xml.Marshaler interface.
func (b Break) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:br"

	if b.BreakType != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(*b.BreakType)})
	}

	if b.Clear != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:clear"}, Value: string(*b.Clear)})
	}

	return e.EncodeElement("", start)
}
