package ctypes

// Grid Column Definition
// w_CT_TblGridCol = attribute w:w { s_ST_TwipsMeasure }?
type Column struct {
	Width *uint64 `xml:"w:w,attr,omitempty"` //Grid Column Width
}
