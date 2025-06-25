package ctypes

// TrackChange represents the complex type for track change
type TrackChange struct {
	ID     int     `xml:"w:id,attr"`
	Author string  `xml:"w:author,attr"`
	Date   *string `xml:"w:date,attr,omitempty"`
}
