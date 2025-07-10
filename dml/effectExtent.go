package dml

type EffectExtent struct {
	LeftEdge   int64 `xml:"l,attr"`
	TopEdge    int64 `xml:"t,attr"`
	RightEdge  int64 `xml:"r,attr"`
	BottomEdge int64 `xml:"b,attr"`
}

func NewEffectExtent(left, top, right, bottom int64) *EffectExtent {
	return &EffectExtent{
		LeftEdge:   left,
		TopEdge:    top,
		RightEdge:  right,
		BottomEdge: bottom,
	}
}
