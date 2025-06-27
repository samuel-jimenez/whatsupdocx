package dmlpic

import (
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
)

const (
	BlackWhiteModeClr        = "clr"
	BlackWhiteModeAuto       = "auto"
	BlackWhiteModeGray       = "gray"
	BlackWhiteModeLtGray     = "ltGray"
	BlackWhiteModeInvGray    = "invGray"
	BlackWhiteModeGrayWhite  = "grayWhite"
	BlackWhiteModeBlackGray  = "blackGray"
	BlackWhiteModeBlackWhite = "blackWhite"
	BlackWhiteModeBlack      = "black"
	BlackWhiteModeWhite      = "white"
	BlackWhiteModeHidden     = "hidden"
)

/*

   <wps:cNvCnPr/>
   <wps:spPr>
       <a:xfrm>
           <a:off x="0" y="0"/>
           <a:ext cx="7096125" cy="28575"/>
       </a:xfrm>
       <a:prstGeom prst="line">
           <a:avLst/>
       </a:prstGeom>
       <a:ln w="57150">
           <a:solidFill>
               <a:schemeClr val="accent1">
                   <a:lumMod val="50000"/>
               </a:schemeClr>
           </a:solidFill>
       </a:ln>
*/

// ShapeProperties *spPr // element ([ISO/IEC29500-1:2016] section A.4.1) that specifies the visual shape properties that can be applied to a shape.<121>
// CT_ShapeProperties
// a_CT_ShapeProperties =
type PicShapeProp struct {
	// -- Attributes --
	//Black and White Mode
	// attribute bwMode { a_ST_BlackWhiteMode }?,
	BwMode *string `xml:"bwMode,attr,omitempty"`

	// -- Child Elements --
	//1.2D Transform for Individual Objects
	// element xfrm { a_CT_Transform2D }?,
	TransformGroup *TransformGroup `xml:"a:xfrm,omitempty"`

	// 2. Choice
	//TODO: Modify it as Geometry choice
	// a_EG_Geometry?,
	PresetGeometry *PresetGeometry `xml:"a:prstGeom,omitempty"`

	//TODO
	// a_EG_FillProperties?,
	// FillProperties *dmlct.FillProperties

	// element ln { a_CT_LineProperties }?,
	LineProperties *dmlct.LineProperties `xml:"a:ln,omitempty"`

	//TODO: Remaining sequence of elements
	// a_EG_EffectProperties?,
	// element scene3d { a_CT_Scene3D }?,
	// element sp3d { a_CT_Shape3D }?,
	// element extLst { a_CT_OfficeArtExtensionList }?
}

type PicShapePropOption func(*PicShapeProp)

func WithTransformGroup(options ...TFGroupOption) PicShapePropOption {
	return func(p *PicShapeProp) {
		p.TransformGroup = NewTransformGroup(options...)
	}
}

func WithPrstGeom(preset string) PicShapePropOption {
	return func(p *PicShapeProp) {
		p.PresetGeometry = NewPresetGeom(preset)
	}
}

// func WithLineProperties(preset string) PicShapePropOption {
// 	return func(p *PicShapeProp) {
// 		p.PresetGeometry = NewPresetGeom(preset)
// 	}
// }

func NewPicShapeProp(options ...PicShapePropOption) *PicShapeProp {
	p := &PicShapeProp{}

	for _, opt := range options {
		opt(p)
	}

	return p
}
