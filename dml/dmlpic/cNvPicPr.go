package dmlpic

import (
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlprops"
)

// Non-Visual Picture Drawing Properties
// CT_NonVisualPictureProperties
// a_CT_NonVisualPictureProperties =
type CNvPicPr struct {
	//Relative Resize Preferred	- Default value is "true"(i.e when attr not specified).
	// ## default value: true
	// attribute preferRelativeResize { xsd:boolean }?,
	PreferRelativeResize *bool `xml:"preferRelativeResize,attr,omitempty"`

	//1. Picture Locks
	// element picLocks { a_CT_PictureLocking }?,
	PicLocks *dmlprops.PicLocks `xml:"a:picLocks,omitempty"`

	//TODO:
	// 2. Extension List
	// element extLst { a_CT_OfficeArtExtensionList }?
}

func NewCNvPicPr() CNvPicPr {
	return CNvPicPr{}
}

// Non-Visual Picture Properties
type NonVisualPicProp struct {
	// 1. Non-Visual Drawing Properties
	CNvPr dmlct.CNvPr `xml:"pic:cNvPr,omitempty"`

	// 2.Non-Visual Picture Drawing Properties
	CNvPicPr CNvPicPr `xml:"pic:cNvPicPr,omitempty"`
}

func NewNVPicProp(cNvPr dmlct.CNvPr, cNvPicPr CNvPicPr) NonVisualPicProp {
	return NonVisualPicProp{
		CNvPr:    cNvPr,
		CNvPicPr: cNvPicPr,
	}
}

func DefaultNVPicProp(id uint, name string) NonVisualPicProp {
	cnvPicPr := NewCNvPicPr()
	cnvPicPr.PicLocks = dmlprops.DefaultPicLocks()
	return NonVisualPicProp{
		CNvPr:    *dmlct.NewNonVisProp(id, name),
		CNvPicPr: cnvPicPr,
	}
}
