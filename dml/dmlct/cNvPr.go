package dmlct

import (
	"github.com/samuel-jimenez/xml"
	"strconv"
)

// Non-Visual Drawing Properties
// CT_NonVisualDrawingProps
// a_CT_NonVisualDrawingProps =
type CNvPr struct {
	// attribute id { a_ST_DrawingElementId },
	ID uint `xml:"id,attr,omitempty"`
	// attribute name { xsd:string },
	Name string `xml:"name,attr,omitempty"`

	//Alternative Text for Object - Default value is "".
	// attribute descr { xsd:string }?,
	Description string `xml:"descr,attr,omitempty"`

	// Hidden - Default value is "false".
	// ## default value: false
	// attribute hidden { xsd:boolean }?,
	Hidden *bool `xml:"hidden,attr,omitempty"`

	// attribute title { xsd:string }?,
	// element hlinkClick { a_CT_Hyperlink }?,
	// element hlinkHover { a_CT_Hyperlink }?,
	// element extLst { a_CT_OfficeArtExtensionList }?
	// TODO: implement child elements
	// Sequence [1..1]
	// a:hlinkClick [0..1]    Drawing Element On Click Hyperlink
	// a:hlinkHover [0..1]    Hyperlink for Hover
	// a:extLst [0..1]    Extension List
}

func NewNonVisProp(id uint, name string) *CNvPr {
	return &CNvPr{
		ID:   id,
		Name: name,
	}
}

func (c CNvPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// ! NOTE: Disabling the empty name check for the Picture
	//  since popular docx tools allow them
	// if c.Name == "" {
	// 	return fmt.Errorf("invalid Name for Non-Visual Drawing Properties when marshaling")
	// }

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: strconv.FormatUint(uint64(c.ID), 10)},
		{Name: xml.Name{Local: "name"}, Value: c.Name},
	}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "descr"}, Value: c.Description})

	if c.Hidden != nil {
		if *c.Hidden {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"}, Value: "true"})
		} else {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"}, Value: "false"})
		}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
