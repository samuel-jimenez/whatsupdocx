package dml

import (
	"github.com/samuel-jimenez/xml"
	"errors"

	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
)

type PositionH struct {
	RelativeFrom dmlst.RelFromH `xml:"relativeFrom,attr"`
	PosOffset    int            `xml:"posOffset"`
}

type PositionV struct {
	RelativeFrom dmlst.RelFromV `xml:"relativeFrom,attr"`
	PosOffset    int            `xml:"posOffset"`
}

func (p PositionH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if p.RelativeFrom == "" {
		return errors.New("Invalid RelativeFrom in PoistionH")
	}

	start.Name.Local = "wp:positionH"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeFrom"}, Value: string(p.RelativeFrom)})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	offsetElem := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
	if err = e.EncodeElement(p.PosOffset, offsetElem); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (p PositionV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if p.RelativeFrom == "" {
		return errors.New("Invalid RelativeFrom in PoistionV")
	}

	start.Name.Local = "wp:positionV"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeFrom"}, Value: string(p.RelativeFrom)})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	offsetElem := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
	if err = e.EncodeElement(p.PosOffset, offsetElem); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
