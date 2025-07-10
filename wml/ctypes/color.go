package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Color represents the color of a text or element.
type Color struct {
	//Run Content Color
	Val string `xml:"w:val,attr"`

	//Run Content Theme Color
	ThemeColor *stypes.ThemeColor `xml:"w:themeColor,attr,omitempty"`

	//Run Content Theme Color Tint
	ThemeTint *string `xml:"w:themeTint,attr,omitempty"`

	//Run Content Theme Color Shade
	ThemeShade *string `xml:"w:themeShade,attr,omitempty"`
}

// NewColor creates a new Color instance with the specified color value.
func NewColor(value string) *Color {
	return &Color{Val: value}
}
