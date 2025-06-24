package ctypes

import "github.com/samuel-jimenez/xml"

// Range Markup elements
type RngMarkupElem struct {
}

func (r RngMarkupElem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// TODO:
	return nil
}
