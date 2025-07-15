package dml

import (
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
)

// CT_Anchor

// wp_CT_Anchor =
// attribute distT { wp_ST_WrapDistance }?,
// attribute distB { wp_ST_WrapDistance }?,
// attribute distL { wp_ST_WrapDistance }?,
// attribute distR { wp_ST_WrapDistance }?,
// attribute simplePos { xsd:boolean }?,
// attribute relativeHeight { xsd:unsignedInt },
// attribute behindDoc { xsd:boolean },
// attribute locked { xsd:boolean },
// attribute layoutInCell { xsd:boolean },
// attribute hidden { xsd:boolean }?,
// attribute allowOverlap { xsd:boolean },
// element simplePos { a_CT_Point2D },
// element positionH { wp_CT_PosH },
// element positionV { wp_CT_PosV },
// element extent { a_CT_PositiveSize2D },
// element effectExtent { wp_CT_EffectExtent }?,
// wp_EG_WrapType,
// element docPr { a_CT_NonVisualDrawingProps },
// element cNvGraphicFramePr { a_CT_NonVisualGraphicFrameProperties }?,
// a_graphic

// NOTE!: As per http://www.datypic.com/sc/ooxml/e-wp_anchor.html, Dist* attributes is optional
// But we include them to lower diffs with MS Word

// wp_CT_Anchor =
type Anchor struct {

	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	//Distance From Text on Top Edge
	// attribute distT { wp_ST_WrapDistance }?,
	DistT uint `xml:"distT,attr"`
	//Distance From Text on Bottom Edge
	// attribute distB { wp_ST_WrapDistance }?,
	DistB uint `xml:"distB,attr"`
	//Distance From Text on Left Edge
	// attribute distL { wp_ST_WrapDistance }?,
	DistL uint `xml:"distL,attr"`
	//Distance From Text on Right Edge
	// attribute distR { wp_ST_WrapDistance }?,
	DistR uint `xml:"distR,attr"`

	/// Specifies that this object shall be positioned using the positioning information in the
	/// simplePos child element (§20.4.2.13). This positioning, when specified, positions the
	/// object on the page by placing its top left point at the x-y coordinates specified by that
	/// element.
	/// Reference: http://officeopenxml.com/drwPicFloating-position.php
	//Page Positioning
	// attribute simplePos { xsd:boolean }?,
	SimplePosAttr *int `xml:"simplePos,attr,omitempty"`

	//Relative Z-Ordering Position
	// attribute relativeHeight { xsd:unsignedInt },
	RelativeHeight int `xml:"relativeHeight,attr"`

	//Display Behind Document Text
	// attribute behindDoc { xsd:boolean },
	BehindDoc int `xml:"behindDoc,attr"`

	//Lock Anchor
	// attribute locked { xsd:boolean },
	Locked int `xml:"locked,attr"`

	//Layout In Table Cell
	// attribute layoutInCell { xsd:boolean },
	LayoutInCell int `xml:"layoutInCell,attr"`

	// attribute hidden { xsd:boolean }?,
	Hidden *int `xml:"hidden,attr,omitempty"`

	//Allow Objects to Overlap
	// attribute allowOverlap { xsd:boolean },
	AllowOverlap int `xml:"allowOverlap,attr"`

	// Child elements:
	// 1. Simple Positioning Coordinates
	// element simplePos { a_CT_Point2D },
	SimplePos dmlct.Point2D `xml:"wp:simplePos"`

	// 2. Horizontal Positioning
	// element positionH { wp_CT_PosH },
	PositionH PositionH `xml:"wp:positionH"`

	// 3. Vertical Positioning
	// element positionV { wp_CT_PosV },
	PositionV PositionV `xml:"wp:positionV"`

	// 4. Inline Drawing Object Extents
	// element extent { a_CT_PositiveSize2D },
	Extent dmlct.PSize2D `xml:"wp:extent"`

	// 5. EffectExtent
	// element effectExtent { wp_CT_EffectExtent }?,
	EffectExtent *EffectExtent `xml:"wp:effectExtent,omitempty"`

	// 6. Wrapping
	// wp_EG_WrapType,
	WrapType WrapType `xml:",group"`

	// 7. Drawing Object Non-Visual Properties
	// element docPr { a_CT_NonVisualDrawingProps },
	DocProp DocProp `xml:"wp:docPr"`

	// 8. Common DrawingML Non-Visual Properties
	// element cNvGraphicFramePr { a_CT_NonVisualGraphicFrameProperties }?,
	CNvGraphicFramePr *NonVisualGraphicFrameProp `xml:"wp:cNvGraphicFramePr,omitempty"`

	// 9. Graphic Object
	// a_graphic
	Graphic Graphic `xml:"a:graphic"`
}

func NewAnchor() *Anchor {
	return &Anchor{}
}
