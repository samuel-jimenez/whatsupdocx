package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// PageNumbering represents the page numbering format in a Word document.
type PageNumbering struct {
	Format stypes.NumFmt `xml:"w:fmt,attr,omitempty"`
}
