package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// DocGrid represents the document grid settings.
// w_CT_DocGrid =
type DocGrid struct {
	// attribute w:type { w_ST_DocGrid }?,
	Type stypes.DocGridType `xml:"type,attr,omitempty"`
	// attribute w:linePitch { w_ST_DecimalNumber }?,
	LinePitch *int `xml:"linePitch,attr,omitempty"`
	// attribute w:charSpace { w_ST_DecimalNumber }?
	CharSpace *int `xml:"charSpace,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the DocGrid type.
// It encodes the DocGrid to its corresponding XML representation.
func (docGrid DocGrid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:docGrid"
	if docGrid.Type != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(docGrid.Type)})
	}
	if docGrid.LinePitch != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:linePitch"}, Value: strconv.Itoa(*docGrid.LinePitch)})
	}
	if docGrid.CharSpace != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:charSpace"}, Value: strconv.Itoa(*docGrid.CharSpace)})
	}
	return e.EncodeElement("", start)
}
