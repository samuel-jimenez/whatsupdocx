package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// w_CT_TextEffect = attribute w:val { w_ST_TextEffect }
type Effect struct {
	Val *stypes.TextEffect `xml:"w:val,attr,omitempty"`
}
