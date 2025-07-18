package dml

import (
	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
)

type GraphicFrameLocks struct {
	//Disallow Aspect Ratio Change
	NoChangeAspect dmlst.OptBool
}

// TODO
func (g GraphicFrameLocks) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:graphicFrameLocks"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
	}

	if g.NoChangeAspect.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeAspect"}, Value: g.NoChangeAspect.ToStringFlag()})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// TODO
func (g *GraphicFrameLocks) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "noChangeAspect" {
			g.NoChangeAspect = dmlst.OptBoolFromStr(a.Value)
		}
	}

	return decoder.Skip()
}
