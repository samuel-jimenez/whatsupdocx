package ctypes

// FontSize represents the font size of a text or element.
//
//	(Half Point Measurement)
//
// Specifies a positive measurement specified in half-points (1/144 of an inch).
// CT_HpsMeasure
// w_CT_HpsMeasure =
type FontSize struct {
	// attribute w:val { w_ST_HpsMeasure }
	Value uint64 `xml:"w:val,attr"`
}

// NewFontSize creates a new FontSize with the specified font size value.
func NewFontSize(value uint64) *FontSize {
	return &FontSize{Value: value}
}
