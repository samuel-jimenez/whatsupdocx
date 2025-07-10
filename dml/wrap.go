package dml

import (
	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
)

// TODO
// wp_EG_WrapType =
// element wrapNone { wp_CT_WrapNone }
// | element wrapSquare { wp_CT_WrapSquare }
// | element wrapTight { wp_CT_WrapTight }
// | element wrapThrough { wp_CT_WrapThrough }
// | element wrapTopAndBottom { wp_CT_WrapTopBottom }

// wp_EG_WrapType =
type WrapType struct {

	// element wrapNone { wp_CT_WrapNone }
	WrapNone *WrapNone `xml:"wp:wrapNone,omitempty"`

	// | element wrapSquare { wp_CT_WrapSquare }
	WrapSquare *WrapSquare `xml:"wp:wrapSquare,omitempty"`

	// | element wrapTight { wp_CT_WrapTight }

	// | element wrapThrough { wp_CT_WrapThrough }
	WrapThrough *WrapThrough `xml:"wp:wrapThrough,omitempty"`

	// | element wrapTopAndBottom { wp_CT_WrapTopBottom }
	WrapTopBtm *WrapTopBtm `xml:"wp:wrapTopAndBottom,omitempty"`
}

// wp_CT_WrapPath =
// attribute edited { xsd:boolean }?,
// element start { a_CT_Point2D },
// element lineTo { a_CT_Point2D }+

// wp_CT_WrapPath =
type WrapPolygon struct {
	// attribute edited { xsd:boolean }?,
	Edited *bool `xml:"edited,attr,omitempty"`
	// element start { a_CT_Point2D },
	Start dmlct.Point2D `xml:"wp:start"`
	// element lineTo { a_CT_Point2D }+
	LineTo []dmlct.Point2D `xml:"wp:lineTo"`
}

// wp_CT_WrapNone = empty
type WrapNone struct {
	XMLName xml.Name `xml:"wrapNone"`
}

func (w WrapNone) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:wrapNone"
	return e.EncodeElement("", start)
}

// // Dummy implementation to ensure only these types are allowed in the wrap type
// func (w WrapNone) getWrapName()    {}
// func (w WrapSquare) getWrapName()  {}
// func (w WrapThrough) getWrapName() {}
// func (w WrapTopBtm) getWrapName()  {}

// wp_CT_WrapSquare =
// attribute wrapText { wp_ST_WrapText },
// attribute distT { wp_ST_WrapDistance }?,
// attribute distB { wp_ST_WrapDistance }?,
// attribute distL { wp_ST_WrapDistance }?,
// attribute distR { wp_ST_WrapDistance }?,
// element effectExtent { wp_CT_EffectExtent }?

// wp_CT_WrapSquare =
type WrapSquare struct {

	//Text Wrapping Location
	// attribute wrapText { wp_ST_WrapText },
	WrapText dmlst.WrapText `xml:"wrapText,attr"`

	//Distance From Text (Top)
	// attribute distT { wp_ST_WrapDistance }?,
	DistT *uint `xml:"distT,attr,omitempty"`

	//Distance From Text on Bottom Edge
	// attribute distB { wp_ST_WrapDistance }?,
	DistB *uint `xml:"distB,attr,omitempty"`

	//Distance From Text on Left Edge
	// attribute distL { wp_ST_WrapDistance }?,
	DistL *uint `xml:"distL,attr,omitempty"`

	//Distance From Text on Right Edge
	// attribute distR { wp_ST_WrapDistance }?,
	DistR *uint `xml:"distR,attr,omitempty"`

	// element effectExtent { wp_CT_EffectExtent }?
	EffectExtent *EffectExtent `xml:"wp:effectExtent,omitempty"`
}

// wp_CT_WrapTight =
// attribute wrapText { wp_ST_WrapText },
// attribute distL { wp_ST_WrapDistance }?,
// attribute distR { wp_ST_WrapDistance }?,
// element wrapPolygon { wp_CT_WrapPath }

// Tight Wrapping
// wp_CT_WrapTight =
type WrapTight struct {

	// Text Wrapping Location
	// attribute wrapText { wp_ST_WrapText },
	WrapText dmlst.WrapText `xml:"wrapText,attr"`

	// Distance From Text on Left Edge
	// attribute distL { wp_ST_WrapDistance }?,
	DistL *uint `xml:"distL,attr,omitempty"`

	// Distance From Text on Right Edge
	// attribute distR { wp_ST_WrapDistance }?,
	DistR *uint `xml:"distR,attr,omitempty"`

	//Tight Wrapping Extents Polygon
	// element wrapPolygon { wp_CT_WrapPath }
	WrapPolygon WrapPolygon `xml:"wp:wrapPolygon"`
}

// wp_CT_WrapThrough =
// attribute wrapText { wp_ST_WrapText },
// attribute distL { wp_ST_WrapDistance }?,
// attribute distR { wp_ST_WrapDistance }?,
// element wrapPolygon { wp_CT_WrapPath }

// Through Wrapping
// wp_CT_WrapThrough =
type WrapThrough struct {

	// Text Wrapping Location
	// attribute wrapText { wp_ST_WrapText },
	WrapText dmlst.WrapText `xml:"wrapText,attr"`

	// Distance From Text on Left Edge
	// attribute distL { wp_ST_WrapDistance }?,
	DistL *uint `xml:"distL,attr,omitempty"`

	// Distance From Text on Right Edge
	// attribute distR { wp_ST_WrapDistance }?,
	DistR *uint `xml:"distR,attr,omitempty"`

	//Tight Wrapping Extents Polygon
	// element wrapPolygon { wp_CT_WrapPath }
	WrapPolygon WrapPolygon `xml:"wp:wrapPolygon"`
}

// wp_CT_WrapTopBottom =
// attribute distT { wp_ST_WrapDistance }?,
// attribute distB { wp_ST_WrapDistance }?,
// element effectExtent { wp_CT_EffectExtent }?

// Top and Bottom Wrapping
// wp_CT_WrapTopBottom =
type WrapTopBtm struct {

	//Distance From Text (Top)
	// attribute distT { wp_ST_WrapDistance }?,
	DistT *uint `xml:"distT,attr,omitempty"`

	//Distance From Text on Bottom Edge
	// attribute distB { wp_ST_WrapDistance }?,
	DistB *uint `xml:"distB,attr,omitempty"`

	//Wrapping Boundaries
	// element effectExtent { wp_CT_EffectExtent }?
	EffectExtent *EffectExtent `xml:"wp:effectExtent,omitempty"`
}
