package common

import "github.com/samuel-jimenez/xml"

type Empty struct {
}

func (s Empty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement("", start)
}
