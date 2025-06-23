package mc

import (
	"encoding/xml"

	"github.com/samuel-jimenez/whatsupdocx/dml"
)

// This element describes alternate content.
type AlternateContent struct {
	Drawing *dml.Drawing `xml:"drawing,omitempty"`
}

func (ac *AlternateContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	accepted := false

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "Choice":
				if !accepted {
					for _, attr := range elem.Attr {
						switch attr.Name.Local {
						case "Requires":
							switch attr.Value {
							case "wps":
								accepted = true
							}
						}
					}
				}
			case "drawing":
				drawingElem := &dml.Drawing{}
				if err = d.DecodeElement(drawingElem, &elem); err != nil {
					return err
				}
				ac.Drawing = drawingElem
			case "Fallback":
				if !accepted {
					accepted = true
					// TODO:
					// w:pict    VML Object
				}

			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			switch elem.Name.Local {
			case "Choice":
			case "Fallback":
				continue
			case "AlternateContent":
				break loop
			}
		}
	}
	return nil
}
