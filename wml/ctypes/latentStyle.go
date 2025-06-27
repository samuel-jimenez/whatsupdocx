package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Latent Style Information
type LatentStyle struct {
	//Default Style Locking Setting
	DefLockedState *stypes.OnOff `xml:"w:defLockedState,attr,omitempty"`

	//Default User Interface Priority Setting
	DefUIPriority *int `xml:"w:defUIPriority,attr,omitempty"`

	//Default Semi-Hidden Setting
	DefSemiHidden *stypes.OnOff `xml:"w:defSemiHidden,attr,omitempty"`

	//Default Hidden Until Used Setting
	DefUnhideWhenUsed *stypes.OnOff `xml:"w:defUnhideWhenUsed,attr,omitempty"`

	//Default Primary Style Setting
	DefQFormat *stypes.OnOff `xml:"w:defQFormat,attr,omitempty"`

	//Latent Style Count
	Count *int `xml:"w:count,attr,omitempty"`

	LsdExceptions []LsdException `xml:"w:lsdException"`
}
