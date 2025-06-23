package dml

import (
	"encoding/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlpic"
	"github.com/samuel-jimenez/whatsupdocx/wps"
)

// a_CT_GraphicalObject = element graphicData { a_CT_GraphicalObjectData }
// a_graphic = element graphic { a_CT_GraphicalObject }
type Graphic struct {
	Data *GraphicData `xml:"graphicData,omitempty"`
}

func NewGraphic(data *GraphicData) *Graphic {
	return &Graphic{Data: data}
}

func DefaultGraphic() *Graphic {
	return &Graphic{}
}

// a_CT_GraphicalObjectData =

type GraphicData struct {
	// attribute uri { xsd:token },
	URI string `xml:"uri,attr,omitempty"`

	// a_CT_GraphicalObjectData_any*
	// a_CT_GraphicalObjectData_any =
	// element * - (o:* | v:* | w10:* | x:*) {
	// anyAttribute*,
	// mixed { anyElement* }
	// }
	Pic   *dmlpic.Pic `xml:"pic,omitempty"`
	Shape *wps.Shape  `xml:"wsp,omitempty"`
}

func NewPicGraphic(pic *dmlpic.Pic) *Graphic {
	return &Graphic{
		Data: &GraphicData{
			URI: constants.DrawingMLPicNS,
			Pic: pic,
		},
	}
}

func NewShapeGraphic(shape *wps.Shape) *Graphic {
	return &Graphic{
		Data: &GraphicData{
			URI:   constants.DrawingMLPicNS,
			Shape: shape,
		},
	}
}

func (g Graphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:graphic"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if g.Data != nil {
		if err = g.Data.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (gd GraphicData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:graphicData"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "uri"}, Value: gd.URI},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if gd.Pic != nil {
		if err := e.EncodeElement(gd.Pic, xml.StartElement{Name: xml.Name{Local: "pic:pic"}}); err != nil {
			return err
		}
	}
	if gd.Shape != nil {
		if err := e.EncodeElement(gd.Shape, xml.StartElement{Name: xml.Name{Local: "wps:wsp"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
