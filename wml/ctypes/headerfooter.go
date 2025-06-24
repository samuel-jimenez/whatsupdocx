package ctypes

import (
	"encoding/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Header Reference
// w_CT_HdrFtrRef =
type HeaderFooterReference struct {
	// w_CT_Rel,
	// w_CT_Rel = r_id
	// r_id = attribute r:id { r_ST_RelationshipId }
	ID string `xml:"id,attr,omitempty"` //Relationship to Part
	// attribute w:type { w_ST_HdrFtr }
	Type stypes.HdrFtrType `xml:"type,attr,omitempty"` //Header or Footer Type

}

func (h HeaderFooterReference) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if h.Type != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(h.Type)})
	}

	if h.ID != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"}, Value: h.ID})
	}

	return e.EncodeElement("", start)
}
