package ctypes

type ParaBorder struct {
	Top     *Border `xml:"w:top,omitempty"`
	Left    *Border `xml:"w:left,omitempty"`
	Right   *Border `xml:"w:right,omitempty"`
	Bottom  *Border `xml:"w:bottom,omitempty"`
	Between *Border `xml:"w:between,omitempty"`
	Bar     *Border `xml:"w:bar,omitempty"`
}
