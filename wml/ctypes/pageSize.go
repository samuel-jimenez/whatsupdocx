package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Page Size : w:pgSz
type PageSize struct {
	Width  *uint64           `xml:"w:w,attr,omitempty"`
	Height *uint64           `xml:"w:h,attr,omitempty"`
	Orient stypes.PageOrient `xml:"w:orient,attr,omitempty"`
	Code   *int              `xml:"w:code,attr,omitempty"`
}
