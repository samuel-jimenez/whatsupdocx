package ctypes

import (
	"github.com/samuel-jimenez/xml"
)

type Row struct {
	// 1. Table-Level Property Exceptions
	PropException *PropException

	// 2.Table Row Properties
	Property *RowProperty `xml:"w:trPr,omitempty"`

	// 3.1 Choice
	Contents []TRCellContent
}

func DefaultRow() *Row {
	return &Row{
		Property: &RowProperty{},
	}
}

// TODO  Implement Marshal and Unmarshal properly for all fields

func (r Row) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tr"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	//1.Table-Level Property Exceptions
	if r.PropException != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:tblPrEx"}}
		if err := e.EncodeElement(r.PropException, propsElement); err != nil {
			// if err := r.PropException.MarshalXML(e, propsElement); err != nil {
			return err
		}
	}

	//2. Table Properties
	if r.Property != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:trPr"}}
		if err := e.EncodeElement(r.Property, propsElement); err != nil {
			// if err := r.Property.MarshalXML(e, propsElement); err != nil {
			return err
		}
	}

	// 3.1 Choice
	for _, cont := range r.Contents {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:tc"}}
		if err := e.EncodeElement(cont, propsElement); err != nil {
			// if err := cont.MarshalXML(e, propsElement); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (r *Row) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "trPr":
				prop := RowProperty{}
				if err = d.DecodeElement(&prop, &elem); err != nil {
					return err
				}

				r.Property = &prop
			case "tblPrEx":
				propEx := PropException{}
				if err = d.DecodeElement(&propEx, &elem); err != nil {
					return err
				}

				r.PropException = &propEx
			case "tc":
				cell := Cell{}
				if err = d.DecodeElement(&cell, &elem); err != nil {
					return err
				}

				r.Contents = append(r.Contents, TRCellContent{
					Cell: &cell,
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

//TODO??TODO

type TRCellContent struct {
	Cell *Cell `xml:"w:tc,omitempty"`
}

func (c TRCellContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.Cell != nil {
		return c.Cell.MarshalXML(e, xml.StartElement{})
	}
	return nil
}

type RowContent struct {
	Row *Row `xml:"w:tr,omitempty"`
}

// TODO RowContent
func (r RowContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Row != nil {
		return r.Row.MarshalXML(e, xml.StartElement{})
	}
	return nil
}
