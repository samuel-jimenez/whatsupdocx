package dmlpic

import (
	"fmt"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/common/units"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/geom"
	"github.com/samuel-jimenez/whatsupdocx/dml/shapes"
)

type Pic struct {
	Attr []xml.Attr `xml:",any,attr,omitempty"`

	// 1. Non-Visual Picture Properties
	NonVisualPicProp NonVisualPicProp `xml:"pic:nvPicPr,omitempty"`

	// 2.Picture Fill
	BlipFill BlipFill `xml:"pic:blipFill,omitempty"`

	// 3.Shape Properties
	PicShapeProp PicShapeProp `xml:"pic:spPr,omitempty"`
}

func NewPic(rID string, imgCount uint, width units.Emu, height units.Emu) *Pic {
	shapeProp := NewPicShapeProp(
		WithTransformGroup(
			WithTFExtent(width, height),
		),
		WithPrstGeom("rect"),
	)

	nvPicProp := DefaultNVPicProp(imgCount, fmt.Sprintf("Image%v", imgCount))

	blipFill := NewBlipFill(rID)

	blipFill.FillModeProps = &FillModeProps{
		Stretch: &shapes.Stretch{
			FillRect: &dmlct.RelativeRect{},
		},
	}

	return &Pic{
		Attr:             []xml.Attr{constants.NameSpaceDrawingMLPic},
		BlipFill:         blipFill,
		NonVisualPicProp: nvPicProp,
		PicShapeProp:     *shapeProp,
	}
}

// a_CT_Transform2D =
type TransformGroup struct {

	// ## default value: 0
	// attribute rot { a_ST_Angle }?,

	// ## default value: false
	// attribute flipH { xsd:boolean }?,

	// ## default value: false
	// attribute flipV { xsd:boolean }?,

	// element off { a_CT_Point2D }?,
	Offset *Offset `xml:"a:off,omitempty"`
	// element ext { a_CT_PositiveSize2D }?
	Extent *dmlct.PSize2D `xml:"a:ext,omitempty"`
}

type TFGroupOption func(*TransformGroup)

func NewTransformGroup(options ...TFGroupOption) *TransformGroup {
	tf := &TransformGroup{}

	for _, opt := range options {
		opt(tf)
	}

	return tf
}

func WithTFExtent(width units.Emu, height units.Emu) TFGroupOption {
	return func(tf *TransformGroup) {
		tf.Extent = dmlct.NewPostvSz2D(width, height)
	}
}

type Offset struct {
	X uint64 `xml:"x,attr"`
	Y uint64 `xml:"y,attr"`
}

type PresetGeometry struct {
	Preset       string             `xml:"prst,attr,omitempty"`
	AdjustValues *geom.AdjustValues `xml:"a:avLst,omitempty"`
}

func NewPresetGeom(preset string) *PresetGeometry {
	return &PresetGeometry{
		Preset: preset,
	}
}
