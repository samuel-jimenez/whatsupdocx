package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// PageBorders represents the page borders of a Word document.
// w_CT_PageBorders =
type PageBorders struct {
	// attribute w:zOrder { w_ST_PageBorderZOrder }?,
	// ## default value: front
	PageBorderZOrder *stypes.PageBorderZOrder `xml:"w:zOrder,attr,omitempty"`

	// attribute w:display { w_ST_PageBorderDisplay }?,
	PageBorderDisplay *stypes.PageBorderDisplay `xml:"w:display,attr,omitempty"`

	// attribute w:offsetFrom { w_ST_PageBorderOffset }?,
	// ## default value: text
	PageBorderOffset *stypes.PageBorderOffset `xml:"w:offsetFrom,attr,omitempty"`

	// element top { w_CT_TopPageBorder }?,
	Top *TopPageBorder `xml:"w:top,omitempty"`

	// element left { w_CT_PageBorder }?,
	Left *PageBorder `xml:"w:left,omitempty"`

	// element bottom { w_CT_BottomPageBorder }?,
	Bottom *BottomPageBorder `xml:"w:bottom,omitempty"`

	// element right { w_CT_PageBorder }?
	Right *PageBorder `xml:"w:right,omitempty"`
}

// PageBorder
// CT_PageBorder
// w_CT_PageBorder = w_CT_Border, r_id?
type PageBorder struct {
	Border
	// r_id = attribute r:id { r_ST_RelationshipId }
	ID *string `xml:"r:id,attr,omitempty"` //Relationship to Part
}

// BottomPageBorder
// CT_BottomPageBorder
// w_CT_BottomPageBorder = w_CT_PageBorder, r_bottomLeft?, r_bottomRight?
type BottomPageBorder struct {
	PageBorder
	// r_bottomLeft = attribute r:bottomLeft { r_ST_RelationshipId }
	BottomLeft *string `xml:"r:bottomLeft,attr,omitempty"`
	// r_bottomRight = attribute r:bottomRight { r_ST_RelationshipId }
	BottomRight *string `xml:"r:bottomRight,attr,omitempty"`
}

// TopPageBorder
// CT_TopPageBorder
// w_CT_TopPageBorder = w_CT_PageBorder, r_topLeft?, r_topRight?
type TopPageBorder struct {
	PageBorder
	// 	r_topLeft = attribute r:topLeft { r_ST_RelationshipId }
	TopLeft *string `xml:"r:topLeft,attr,omitempty"`
	// r_topRight = attribute r:topRight { r_ST_RelationshipId }
	TopRight *string `xml:"r:topRight,attr,omitempty"`
}
