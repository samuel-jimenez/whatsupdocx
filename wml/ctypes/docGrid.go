package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// DocGrid represents the document grid settings.
// w_CT_DocGrid =
type DocGrid struct {
	// attribute w:type { w_ST_DocGrid }?,
	Type stypes.DocGridType `xml:"w:type,attr,omitempty"`
	// attribute w:linePitch { w_ST_DecimalNumber }?,
	LinePitch *int `xml:"w:linePitch,attr,omitempty"`
	// attribute w:charSpace { w_ST_DecimalNumber }?
	CharSpace *int `xml:"w:charSpace,attr,omitempty"`
}
