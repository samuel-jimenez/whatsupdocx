package dml

import (
	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
)

// This element specifies that the DrawingML object located at this position in the document is an inline object. Within a WordprocessingML document, drawing objects can exist in two states:
//
//• Inline - The drawing object is in line with the text, and affects the line height and layout of its line (like a character glyph of similar size).

type Inline struct {
	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	//
	// NOTE!: As per http://www.datypic.com/sc/ooxml/e-wp_inline.html, Dist* attributes is optional
	// But MS Word requires them to be there

	//Distance From Text on Top Edge
	DistT uint `xml:"distT,attr,omitempty"`

	//Distance From Text on Bottom Edge
	DistB uint `xml:"distB,attr,omitempty"`

	//Distance From Text on Left Edge
	DistL uint `xml:"distL,attr,omitempty"`

	//Distance From Text on Right Edge
	DistR uint `xml:"distR,attr,omitempty"`

	Attr []xml.Attr `xml:",any,attr,omitempty"`

	// Child elements:
	// 1. Drawing Object Size
	Extent dmlct.PSize2D `xml:"wp:extent,omitempty"`

	// 2. Inline Wrapping Extent
	EffectExtent *EffectExtent `xml:"wp:effectExtent,omitempty"`

	// 3. Drawing Object Non-Visual Properties
	DocProp DocProp `xml:"wp:docPr,omitempty"`

	//4.Common DrawingML Non-Visual Properties
	CNvGraphicFramePr *NonVisualGraphicFrameProp `xml:"wp:cNvGraphicFramePr,omitempty"`

	//5.Graphic Object
	Graphic Graphic `xml:"a:graphic,omitempty"`
}

func NewInline(extent dmlct.PSize2D, docProp DocProp, graphic Graphic) Inline {
	return Inline{
		Attr: []xml.Attr{
			//TODO extract
			{Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
			{Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
		},
		Extent:  extent,
		DocProp: docProp,
		Graphic: graphic,
		CNvGraphicFramePr: &NonVisualGraphicFrameProp{
			GraphicFrameLocks: &GraphicFrameLocks{
				NoChangeAspect: dmlst.NewOptBool(true),
			},
		},
	}
}
