package ctypes

// CT_TblBorders
// w_CT_TblBorders =
type TableBorders struct {
	// 1. Table Top Border
	// element top { w_CT_Border }?,
	Top *Border `xml:"w:top,omitempty"`

	// 2. Table Left Border
	// element start { w_CT_Border }?,
	Left *Border `xml:"w:left,omitempty"`

	// 3. Table Bottom Border
	// element bottom { w_CT_Border }?,
	Bottom *Border `xml:"w:bottom,omitempty"`

	// 4. Table Right Border
	// element end { w_CT_Border }?,
	Right *Border `xml:"w:right,omitempty"`

	// 5. Table Inside Horizontal Edges Border
	// element insideH { w_CT_Border }?,
	InsideH *Border `xml:"w:insideH,omitempty"`
	// 6. Table Inside Vertical Edges Border
	// element insideV { w_CT_Border }?
	InsideV *Border `xml:"w:insideV,omitempty"`
}

func DefaultTableBorders() *TableBorders {
	return &TableBorders{}
}
