package ctypes

import (
	"github.com/samuel-jimenez/xml"
	"strconv"
)

// PageMargin represents the page margins of a Word document.
// w_CT_PageMar =
type PageMargin struct {
	// attribute w:top { w_ST_SignedTwipsMeasure },
	Top *int `xml:"top,attr,omitempty"`
	// attribute w:right { s_ST_TwipsMeasure },
	Right *int `xml:"right,attr,omitempty"`
	// attribute w:bottom { w_ST_SignedTwipsMeasure },
	Bottom *int `xml:"bottom,attr,omitempty"`
	// attribute w:left { s_ST_TwipsMeasure },
	Left *int `xml:"left,attr,omitempty"`
	// attribute w:header { s_ST_TwipsMeasure },
	Header *int `xml:"header,attr,omitempty"`
	// attribute w:footer { s_ST_TwipsMeasure },
	Footer *int `xml:"footer,attr,omitempty"`
	// attribute w:gutter { s_ST_TwipsMeasure }
	Gutter *int `xml:"gutter,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the PageMargin type.
// It encodes the PageMargin to its corresponding XML representation.
func (p PageMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgMar"

	start.Attr = []xml.Attr{}

	attrs := []struct {
		local string
		value *int
	}{
		{"w:top", p.Top},
		{"w:right", p.Right},
		{"w:bottom", p.Bottom},
		{"w:left", p.Left},
		{"w:header", p.Header},
		{"w:footer", p.Footer},
		{"w:gutter", p.Gutter},
	}

	for _, attr := range attrs {
		if attr.value == nil {
			continue
		}
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: attr.local}, Value: strconv.Itoa(*attr.value)})
	}

	return e.EncodeElement("", start)
}
