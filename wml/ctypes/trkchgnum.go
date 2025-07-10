package ctypes

// TrackChangeNum represents the complex type for track change numbering
type TrackChangeNum struct {
	ID       int     `xml:"w:id,attr"`
	Author   string  `xml:"w:author,attr"`
	Date     *string `xml:"w:date,attr,omitempty"`
	Original *string `xml:"w:original,attr,omitempty"`
}
