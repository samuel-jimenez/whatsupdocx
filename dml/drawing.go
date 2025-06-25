package dml

type DrawingPositionType string

const (
	DrawingPositionAnchor DrawingPositionType = "wp:anchor"
	DrawingPositionInline DrawingPositionType = "wp:inline"
)

type Drawing struct {
	Inline []Inline  `xml:"wp:inline,omitempty"`
	Anchor []*Anchor `xml:"wp:anchor,omitempty"`
}
