package docx

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Specifies the background information for this document
//
// This background shall be displayed on all pages of the document, behind all other document content.
type Background struct {
	Color      *string            `xml:"w:color,attr,omitempty"`
	ThemeColor *stypes.ThemeColor `xml:"w:themeColor,attr,omitempty"`
	ThemeTint  *string            `xml:"w:themeTint,attr,omitempty"`
	ThemeShade *string            `xml:"w:themeShade,attr,omitempty"`
}

func NewBackground() *Background {
	return &Background{}
}
