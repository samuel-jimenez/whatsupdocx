package dmlct

import (
	"github.com/samuel-jimenez/xml"
)

// LineProperties specifies the properties of a line.
// a_CT_LineProperties =
type LineProperties struct {
	// attribute w { a_ST_LineWidth }?,
	LineWidth *string `xml:"w,attr,omitempty"`

	//TODO
	// attribute cap { a_ST_LineCap }?,
	// attribute cmpd { a_ST_CompoundLine }?,
	// attribute algn { a_ST_PenAlignment }?,

	// a_EG_LineFillProperties?,
	LineFillProperties *LineFillProperties

	//TODO
	// a_EG_LineDashProperties?,
	// a_EG_LineJoinProperties?,
	// element headEnd { a_CT_LineEndProperties }?,
	// element tailEnd { a_CT_LineEndProperties }?,
	// element extLst { a_CT_OfficeArtExtensionList }?
}

func (props LineProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {

	if props.LineWidth != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w"}, Value: string(*props.LineWidth)})
	}

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}
	if props.LineFillProperties != nil {
		propsElement := xml.StartElement{}
		if err = props.LineFillProperties.MarshalXML(e, propsElement); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (props *LineProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "w":
			props.LineWidth = &attr.Value
		}
	}

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "noFill", "solidFill":
				props.LineFillProperties = &LineFillProperties{}
				props.LineFillProperties.UnmarshalXML(d, elem)
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}
