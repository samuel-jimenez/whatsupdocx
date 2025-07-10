package dml

type DrawingPositionType string

const (
	DrawingPositionAnchor DrawingPositionType = "wp:anchor"
	DrawingPositionInline DrawingPositionType = "wp:inline"
)

// w_CT_Drawing = (wp_anchor? | wp_inline?)+
// wp_inline = element inline { wp_CT_Inline }
// wp_anchor = element anchor { wp_CT_Anchor }
type Drawing struct {
	Inline []Inline  `xml:"wp:inline,omitempty"`
	Anchor []*Anchor `xml:"wp:anchor,omitempty"`
}

/*
 * TODO
// w_CT_Drawing = (wp_anchor? | wp_inline?)+
type Drawing struct {
	DrawingChild []*DrawingChild `xml:",group,any,omitempty"`
}

// w_CT_Drawing = (wp_anchor? | wp_inline?)+
// wp_inline = element inline { wp_CT_Inline }
// wp_anchor = element anchor { wp_CT_Anchor }
type DrawingChild struct {
	Inline *Inline `xml:"wp:inline,omitempty"`
	Anchor *Anchor `xml:"wp:anchor,omitempty"`
}
*/
