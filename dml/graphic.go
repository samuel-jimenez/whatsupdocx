package dml

import (
	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlpic"
	"github.com/samuel-jimenez/whatsupdocx/wps"
)

// http://www.datypic.com/sc/ooxml/e-a_graphic.html
// a_CT_GraphicalObject = element graphicData { a_CT_GraphicalObjectData }
// a_graphic = element graphic { a_CT_GraphicalObject }
type Graphic struct {
	DrawingMLMainNS string       `xml:"xmlns:a,attr,omitempty"`
	Data            *GraphicData `xml:"a:graphicData,omitempty"`
}

func NewGraphic(data *GraphicData) *Graphic {
	return &Graphic{DrawingMLMainNS: constants.DrawingMLMainNS, Data: data}
}

func DefaultGraphic() *Graphic {
	return &Graphic{DrawingMLMainNS: constants.DrawingMLMainNS}
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
	Pic   *dmlpic.Pic `xml:"pic:pic,omitempty"`
	Shape *wps.Shape  `xml:"wps:wsp,omitempty"`
}

func NewPicGraphic(pic *dmlpic.Pic) *Graphic {
	return &Graphic{
		DrawingMLMainNS: constants.DrawingMLMainNS,
		Data: &GraphicData{
			URI: constants.DrawingMLPicNS,
			Pic: pic,
		},
	}
}

func NewShapeGraphic(shape *wps.Shape) *Graphic {
	return &Graphic{
		DrawingMLMainNS: constants.DrawingMLMainNS,
		Data: &GraphicData{
			URI:   constants.DrawingMLPicNS,
			Shape: shape,
		},
	}
}

func (gd GraphicData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

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
