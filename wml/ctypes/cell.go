package ctypes

import (
	"github.com/samuel-jimenez/xml"
)

// w_CT_Tc =
type Cell struct {
	// attribute w:id { s_ST_String }?,

	// 1.Table Cell Properties
	// element tcPr { w_CT_TcPr }?,
	Property *CellProperty

	// 2.1 Choice: ZeroOrMore
	// Any number of elements can exists within this choice group
	// w_EG_BlockLevelElts+
	Contents []TCBlockContent

	//TODO: Remaining choices
}

func DefaultCell() *Cell {
	return &Cell{
		Property: &CellProperty{
			Shading: DefaultShading(),
		},
	}
}

func (c Cell) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tc"

	if err = e.EncodeToken(start); err != nil {
		return err
	}

	//1.Table Cell Properties
	if c.Property != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:tcPr"}}
		if err := e.EncodeElement(c.Property, propsElement); err != nil {
			// if err := c.Property.MarshalXML(e, propsElement); err != nil {
			return err
		}
	}

	//2.1 Choice
	for _, elem := range c.Contents {
		if err = elem.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (c *Cell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "tcPr":
				prop := CellProperty{}
				if err = d.DecodeElement(&prop, &elem); err != nil {
					return err
				}

				c.Property = &prop
			case "p":
				para := Paragraph{}
				if err = d.DecodeElement(&para, &elem); err != nil {
					return err
				}

				c.Contents = append(c.Contents, TCBlockContent{
					Paragraph: &para,
				})
			case "tbl":
				tbl := Table{}
				if err = d.DecodeElement(&tbl, &elem); err != nil {
					return err
				}

				c.Contents = append(c.Contents, TCBlockContent{
					Table: &tbl,
				})
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}

//TODO crossref  w_EG_ContentBlockContent

// Table Cell - ContentBlockContent
// w_EG_BlockLevelChunkElts = w_EG_ContentBlockContent*
// w_EG_BlockLevelElts =
// w_EG_BlockLevelChunkElts*
// | element altChunk { w_CT_AltChunk }*
type TCBlockContent struct {
	//Paragraph
	//	- ZeroOrMore: Any number of times Paragraph can repeat within cell
	// | element p { w_CT_P }*
	Paragraph *Paragraph `xml:"w:p,omitempty"`
	//Table
	//	- ZeroOrMore: Any number of times Table can repeat within cell
	// | element tbl { w_CT_Tbl }*
	Table *Table `xml:"w:tbl,omitempty"`
}

func (t TCBlockContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.Paragraph != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:p"}}
		return e.EncodeElement(t.Paragraph, propsElement)
	}

	if t.Table != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:tbl"}}
		return e.EncodeElement(t.Paragraph, propsElement)
	}

	return nil
}
