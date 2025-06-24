package ctypes

import (
	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// PageBorders represents the page borders of a Word document.
// w_CT_PageBorders =
type PageBorders struct {
	// attribute w:zOrder { w_ST_PageBorderZOrder }?,
	// ## default value: front
	PageBorderZOrder *stypes.PageBorderZOrder `xml:"zOrder,attr,omitempty"`

	// attribute w:display { w_ST_PageBorderDisplay }?,
	PageBorderDisplay *stypes.PageBorderDisplay `xml:"display,attr,omitempty"`

	// attribute w:offsetFrom { w_ST_PageBorderOffset }?,
	// ## default value: text
	PageBorderOffset *stypes.PageBorderOffset `xml:"offsetFrom,attr,omitempty"`

	// element top { w_CT_TopPageBorder }?,
	Top *TopPageBorder `xml:"top,omitempty"`

	// element left { w_CT_PageBorder }?,
	Left *PageBorder `xml:"left,omitempty"`

	// element bottom { w_CT_BottomPageBorder }?,
	Bottom *BottomPageBorder `xml:"bottom,omitempty"`

	// element right { w_CT_PageBorder }?
	Right *PageBorder `xml:"right,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the PageMargin type.
// It encodes the PageMargin to its corresponding XML representation.
func (s PageBorders) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Attr = []xml.Attr{}

	if s.PageBorderZOrder != nil {
		propsAttr := xml.Attr{Name: xml.Name{Local: "w:zOrder"}, Value: string(*s.PageBorderZOrder)}
		start.Attr = append(start.Attr, propsAttr)

	}
	if s.PageBorderDisplay != nil {
		propsAttr := xml.Attr{Name: xml.Name{Local: "w:display"}, Value: string(*s.PageBorderDisplay)}
		start.Attr = append(start.Attr, propsAttr)

	}
	if s.PageBorderOffset != nil {
		propsAttr := xml.Attr{Name: xml.Name{Local: "w:offsetFrom"}, Value: string(*s.PageBorderOffset)}
		start.Attr = append(start.Attr, propsAttr)

	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if s.Top != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:top"}}
		if err := e.EncodeElement(s.Top, propsElement); err != nil {
			return err
		}
	}

	if s.Left != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:left"}}
		if err := e.EncodeElement(s.Left, propsElement); err != nil {
			return err
		}
	}
	if s.Bottom != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:bottom"}}
		if err := e.EncodeElement(s.Bottom, propsElement); err != nil {
			return err
		}
	}
	if s.Right != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:right"}}
		if err := e.EncodeElement(s.Right, propsElement); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// PageBorder
// CT_PageBorder
// w_CT_PageBorder = w_CT_Border, r_id?
type PageBorder struct {
	Border
	// r_id = attribute r:id { r_ST_RelationshipId }
	ID *string `xml:"id,attr,omitempty"` //Relationship to Part
}

// BottomPageBorder
// CT_BottomPageBorder
// w_CT_BottomPageBorder = w_CT_PageBorder, r_bottomLeft?, r_bottomRight?
type BottomPageBorder struct {
	PageBorder
	// r_bottomLeft = attribute r:bottomLeft { r_ST_RelationshipId }
	BottomLeft *string `xml:"bottomLeft,attr,omitempty"`
	// r_bottomRight = attribute r:bottomRight { r_ST_RelationshipId }
	BottomRight *string `xml:"bottomRight,attr,omitempty"`
}

// TopPageBorder
// CT_TopPageBorder
// w_CT_TopPageBorder = w_CT_PageBorder, r_topLeft?, r_topRight?
type TopPageBorder struct {
	PageBorder
	// 	r_topLeft = attribute r:topLeft { r_ST_RelationshipId }
	TopLeft *string `xml:"topLeft,attr,omitempty"`
	// r_topRight = attribute r:topRight { r_ST_RelationshipId }
	TopRight *string `xml:"r_topRight,attr,omitempty"`
}
