package ctypes

import "github.com/samuel-jimenez/xml"

type PropException struct {
}

func (p PropException) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// TODO:
	return nil
}
