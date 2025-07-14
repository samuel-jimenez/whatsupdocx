package dmlpic

import (
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/shapes"
)

// a_CT_BlipFillProperties =
// attribute dpi { xsd:unsignedInt }?,
// attribute rotWithShape { xsd:boolean }?,
// element blip { a_CT_Blip }?,
// element srcRect { a_CT_RelativeRect }?,
// a_EG_FillModeProperties?

type BlipFill struct {

	//Attributes:
	DPI          *uint32 `xml:"dpi,attr,omitempty"`          //DPI Setting
	RotWithShape *bool   `xml:"rotWithShape,attr,omitempty"` //Rotate With Shape

	// 1. Blip
	Blip *Blip `xml:"a:blip,omitempty"`

	//2.Source Rectangle
	SrcRect *dmlct.RelativeRect `xml:"a:srcRect,omitempty"`

	// 3. Choice of a:EG_FillModeProperties
	FillModeProps *FillModeProps `xml:",group,any"`
}

// NewBlipFill creates a new BlipFill with the given relationship ID (rID)
// The rID is used to reference the image in the presentation.
func NewBlipFill(rID string) BlipFill {
	return BlipFill{
		Blip: &Blip{
			EmbedID: rID,
		},
	}
}

type FillModeProps struct {
	Stretch *shapes.Stretch `xml:"stretch,omitempty"`
	Tile    *shapes.Tile    `xml:"tile,omitempty"`
}
