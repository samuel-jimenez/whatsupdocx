package dml

import (
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
)

type Anchor struct {
	/// Specifies that this object shall be positioned using the positioning information in the
	/// simplePos child element (§20.4.2.13). This positioning, when specified, positions the
	/// object on the page by placing its top left point at the x-y coordinates specified by that
	/// element.
	/// Reference: http://officeopenxml.com/drwPicFloating-position.php
	//Page Positioning
	SimplePosAttr *int `xml:"simplePos,attr,omitempty"`

	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	//Distance From Text on Top Edge
	DistT uint `xml:"distT,attr,omitempty"`
	//Distance From Text on Bottom Edge
	DistB uint `xml:"distB,attr,omitempty"`
	//Distance From Text on Left Edge
	DistL uint `xml:"distL,attr,omitempty"`
	//Distance From Text on Right Edge
	DistR uint `xml:"distR,attr,omitempty"`

	//Relative Z-Ordering Position
	RelativeHeight int `xml:"relativeHeight,attr"`

	//Layout In Table Cell
	LayoutInCell int `xml:"layoutInCell,attr"`

	//Display Behind Document Text
	BehindDoc int `xml:"behindDoc,attr"`

	//Lock Anchor
	Locked int `xml:"locked,attr"`

	//Allow Objects to Overlap
	AllowOverlap int `xml:"allowOverlap,attr"`

	Hidden *int `xml:"hidden,attr,omitempty"`

	// Child elements:
	// 1. Simple Positioning Coordinates
	SimplePos dmlct.Point2D `xml:"wp:simplePos"`

	// 2. Horizontal Positioning
	PositionH PoistionH `xml:"wp:positionH"`

	// 3. Vertical Positioning
	PositionV PoistionV `xml:"wp:positionV"`

	// 4. Inline Drawing Object Extents
	Extent dmlct.PSize2D `xml:"wp:extent"`

	// 5. EffectExtent
	EffectExtent *EffectExtent `xml:"wp:effectExtent,omitempty"`

	// 6. Wrapping
	// 6.1 .wrapNone
	WrapNone *WrapNone `xml:"wp:wrapNone,omitempty"`

	// 6.2. wrapSquare
	WrapSquare *WrapSquare `xml:"wp:wrapSquare,omitempty"`

	// 6.3. wrapThrough
	WrapThrough *WrapThrough `xml:"wp:wrapThrough,omitempty"`

	// 6.4. wrapTopAndBottom
	WrapTopBtm *WrapTopBtm `xml:"wp:wrapTopAndBottom,omitempty"`

	// 7. Drawing Object Non-Visual Properties
	DocProp DocProp `xml:"wp:docPr"`

	// 8. Common DrawingML Non-Visual Properties
	CNvGraphicFramePr *NonVisualGraphicFrameProp `xml:"wp:cNvGraphicFramePr,omitempty"`

	// 9. Graphic Object
	Graphic Graphic `xml:"wp:graphic"`
}

func NewAnchor() *Anchor {
	return &Anchor{}
}
