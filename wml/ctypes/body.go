package ctypes

import (
	"github.com/samuel-jimenez/xml"
)

// This element specifies the contents of the body of the document – the main document editing surface.
// w_CT_Body =
type Body struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`

	// w_EG_BlockLevelElts*,
	Children []BlockLevel `xml:",group,any,omitempty"`
	// element sectPr { w_CT_SectPr }?
	SectPr *SectionProp `xml:"w:sectPr,omitempty"`
}

// Use this function to initialize a new Body before adding content to it.
func NewBody() *Body {
	return &Body{}
}
