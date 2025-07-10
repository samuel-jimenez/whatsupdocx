package dml

import (
	"strconv"

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

type WrapType struct {
	// 6.1 .wrapNone
	WrapNone *WrapNone `xml:"wp:wrapNone,omitempty"`

	// 6.2. wrapSquare
	WrapSquare *WrapSquare `xml:"wp:wrapSquare,omitempty"`

	// 6.3. wrapThrough
	WrapThrough *WrapThrough `xml:"wp:wrapThrough,omitempty"`

	// 6.4. wrapTopAndBottom
	WrapTopBtm *WrapTopBtm `xml:"wp:wrapTopAndBottom,omitempty"`
}

// TODO fix this

// wp_ST_WrapText = "bothSides" | "left" | "right" | "largest"
// wp_CT_WrapPath =
// attribute edited { xsd:boolean }?,
// element start { a_CT_Point2D },
// element lineTo { a_CT_Point2D }+
// wp_CT_WrapNone = empty
// wp_CT_WrapSquare =
// ISO/IEC 29500-1:2016(E)
// ©ISO/IEC 2016 – All rights reserved
// attribute wrapText { wp_ST_WrapText },
// attribute distT { wp_ST_WrapDistance }?,
// attribute distB { wp_ST_WrapDistance }?,
// attribute distL { wp_ST_WrapDistance }?,
// attribute distR { wp_ST_WrapDistance }?,
// element effectExtent { wp_CT_EffectExtent }?
// wp_CT_WrapTight =
// attribute wrapText { wp_ST_WrapText },
// attribute distL { wp_ST_WrapDistance }?,
// attribute distR { wp_ST_WrapDistance }?,
// element wrapPolygon { wp_CT_WrapPath }
// wp_CT_WrapThrough =
// attribute wrapText { wp_ST_WrapText },
// attribute distL { wp_ST_WrapDistance }?,
// attribute distR { wp_ST_WrapDistance }?,
// element wrapPolygon { wp_CT_WrapPath }
// wp_CT_WrapTopBottom =
// attribute distT { wp_ST_WrapDistance }?,
// attribute distB { wp_ST_WrapDistance }?,
// element effectExtent { wp_CT_EffectExtent }?

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

type WrapSquare struct {
	XMLName xml.Name `xml:"wrapSquare"`

	//Text Wrapping Location
	WrapText dmlst.WrapText `xml:"wrapText,attr"`

	//Distance From Text (Top)
	DistT *uint `xml:"distT,attr,omitempty"`

	//Distance From Text on Bottom Edge
	DistB *uint `xml:"distB,attr,omitempty"`

	//Distance From Text on Left Edge
	DistL *uint `xml:"distL,attr,omitempty"`

	//Distance From Text on Right Edge
	DistR *uint `xml:"distR,attr,omitempty"`

	EffectExtent *EffectExtent `xml:"effectExtent,omitempty"`
}

func (ws WrapSquare) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:wrapSquare"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "wrapText"}, Value: string(ws.WrapText)},
	}

	if ws.DistT != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"}, Value: strconv.FormatUint(uint64(*ws.DistT), 10)})
	}
	if ws.DistB != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"}, Value: strconv.FormatUint(uint64(*ws.DistB), 10)})
	}
	if ws.DistL != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: strconv.FormatUint(uint64(*ws.DistL), 10)})
	}
	if ws.DistR != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: strconv.FormatUint(uint64(*ws.DistR), 10)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if ws.EffectExtent != nil {
		if err := ws.EffectExtent.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type WrapPolygon struct {
	Start  dmlct.Point2D   `xml:"start"`
	LineTo []dmlct.Point2D `xml:"lineTo"`
	Edited *bool           `xml:"edited,attr,omitempty"`
}

func (wp WrapPolygon) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:wrapPolygon"

	start.Attr = []xml.Attr{}

	if wp.Edited != nil {
		if *wp.Edited {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "edited"}, Value: "true"})
		} else {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "edited"}, Value: "false"})
		}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if err := wp.Start.MarshalXML(e, xml.StartElement{
		Name: xml.Name{Local: "wp:start"},
	}); err != nil {
		return err
	}

	for _, lineTo := range wp.LineTo {
		if err := lineTo.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "wp:lineTo"},
		}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// Tight Wrapping
type WrapTight struct {
	XMLName xml.Name `xml:"wrapTight"`

	//Tight Wrapping Extents Polygon
	WrapPolygon WrapPolygon `xml:"wrapPolygon"`

	// Text Wrapping Location
	WrapText dmlst.WrapText `xml:"wrapText,attr"`

	// Distance From Text on Left Edge
	DistL *uint `xml:"distL,attr,omitempty"`

	// Distance From Text on Right Edge
	DistR *uint `xml:"distR,attr,omitempty"`
}

func (w WrapTight) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:wrapTight"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "wrapText"}, Value: string(w.WrapText)},
	}

	if w.DistL != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: strconv.FormatUint(uint64(*w.DistL), 10)})
	}
	if w.DistR != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: strconv.FormatUint(uint64(*w.DistR), 10)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if err := w.WrapPolygon.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// Through Wrapping
type WrapThrough struct {
	XMLName xml.Name `xml:"wrapThrough"`

	//Tight Wrapping Extents Polygon
	WrapPolygon WrapPolygon `xml:"wrapPolygon"`

	// Text Wrapping Location
	WrapText dmlst.WrapText `xml:"wrapText,attr"`

	// Distance From Text on Left Edge
	DistL *uint `xml:"distL,attr,omitempty"`

	// Distance From Text on Right Edge
	DistR *uint `xml:"distR,attr,omitempty"`
}

func (w WrapThrough) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:wrapThrough"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "wrapText"}, Value: string(w.WrapText)},
	}

	if w.DistL != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: strconv.FormatUint(uint64(*w.DistL), 10)})
	}
	if w.DistR != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: strconv.FormatUint(uint64(*w.DistR), 10)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if err := w.WrapPolygon.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// Top and Bottom Wrapping
type WrapTopBtm struct {
	XMLName xml.Name `xml:"wrapTopAndBottom"`

	//Distance From Text (Top)
	DistT *uint `xml:"distT,attr,omitempty"`

	//Distance From Text on Bottom Edge
	DistB *uint `xml:"distB,attr,omitempty"`

	//Wrapping Boundaries
	EffectExtent *EffectExtent `xml:"effectExtent,omitempty"`
}

func (w WrapTopBtm) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:wrapTopAndBottom"

	start.Attr = []xml.Attr{}

	if w.DistT != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"}, Value: strconv.FormatUint(uint64(*w.DistT), 10)})
	}
	if w.DistB != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"}, Value: strconv.FormatUint(uint64(*w.DistB), 10)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if w.EffectExtent != nil {
		if err := w.EffectExtent.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
