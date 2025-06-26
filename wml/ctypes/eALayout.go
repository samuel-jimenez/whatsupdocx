package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// East Asian Typography Settings
type EALayout struct {
	ID           *int                    `xml:"w:id,attr,omitempty"`
	Combine      *stypes.OnOff           `xml:"w:combine,attr,omitempty"`
	CombineBrkts *stypes.CombineBrackets `xml:"w:combineBrackets,attr,omitempty"`
	Vert         *stypes.OnOff           `xml:"w:vert,attr,omitempty"`
	VertCompress *stypes.OnOff           `xml:"w:vertCompress,attr,omitempty"`
}
