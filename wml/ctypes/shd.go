package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Shading represents the shading properties for a run in a WordprocessingML document.
type Shading struct {
	Val            stypes.Shading     `xml:"w:val,attr"`
	Color          *string            `xml:"w:color,attr,omitempty"`
	ThemeColor     *stypes.ThemeColor `xml:"w:themeColor,attr,omitempty"`
	ThemeFill      *stypes.ThemeColor `xml:"w:themeFill,attr,omitempty"`
	ThemeTint      *string            `xml:"w:themeTint,attr,omitempty"`
	ThemeShade     *string            `xml:"w:themeShade,attr,omitempty"`
	Fill           *string            `xml:"w:fill,attr,omitempty"`
	ThemeFillTint  *string            `xml:"w:themeFillTint,attr,omitempty"`
	ThemeFillShade *string            `xml:"w:themeFillShade,attr,omitempty"`
}

// DefaultShading creates a new Shading with default values.
func DefaultShading() *Shading {
	color := "auto"
	fill := "FFFFFF"
	return &Shading{
		Val:   stypes.ShdClear,
		Color: &color,
		Fill:  &fill,
	}
}

// NewShading creates a new Shading.
func NewShading() *Shading {
	return DefaultShading()
}

// Color sets the color for the shading.
func (s *Shading) SetColor(color string) *Shading {
	s.Color = &color
	return s
}

// Fill sets the fill for the shading.
func (s *Shading) SetFill(fill string) *Shading {
	s.Fill = &fill
	return s
}

// ShadingType sets the shading type for the shading.
func (s *Shading) SetShadingType(shdType stypes.Shading) *Shading {
	s.Val = shdType
	return s
}
