package ctypes

// Table Cell Borders
type CellBorders struct {
	// 1. Table Cell Top Border
	Top *Border `xml:"w:top,omitempty"`

	// 2. Table Cell Left Border
	Left *Border `xml:"w:left,omitempty"`

	// 3. Table Cell Bottom Border
	Bottom *Border `xml:"w:bottom,omitempty"`

	// 4. Table Cell Right Border
	Right *Border `xml:"w:right,omitempty"`

	// 5. Table Cell Inside Horizontal Edges Border
	InsideH *Border `xml:"w:insideH,omitempty"`

	// 6. Table Cell Inside Vertical Edges Border
	InsideV *Border `xml:"w:insideV,omitempty"`

	// 7. Table Cell Top Left to Bottom Right Diagonal Border
	TL2BR *Border `xml:"w:tl2br,omitempty"`

	// 8. Table Cell Top Right to Bottom Left Diagonal Border
	TR2BL *Border `xml:"w:tr2bl,omitempty"`
}

func DefaultCellBorders() *CellBorders {
	return &CellBorders{}
}
