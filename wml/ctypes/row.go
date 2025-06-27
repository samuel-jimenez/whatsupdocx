package ctypes

import (
	"github.com/samuel-jimenez/xml"
)

// w_CT_Row =
type Row struct {
	// attribute w:rsidRPr { w_ST_LongHexNumber }?,
	// attribute w:rsidR { w_ST_LongHexNumber }?,
	// attribute w:rsidDel { w_ST_LongHexNumber }?,
	// attribute w:rsidTr { w_ST_LongHexNumber }?,

	// 1. Table-Level Property Exceptions
	// element tblPrEx { w_CT_TblPrEx }?,
	PropException *PropException `xml:"w:tblPrEx,omitempty"`

	// 2.Table Row Properties
	// element trPr { w_CT_TrPr }?,
	Property *RowProperty `xml:"w:trPr,omitempty"`

	// 3.1 Choice
	// w_EG_ContentCellContent*
	Contents []TRCellContent `xml:",group,any,omitempty"`
	// `xml:"w:tc,omitempty"`
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

// w_CT_SdtCell =
// element sdtPr { w_CT_SdtPr }?,
// element sdtEndPr { w_CT_SdtEndPr }?,
// element sdtContent { w_CT_SdtContentCell }?

// w_EG_ContentCellContent =
// element tc { w_CT_Tc }*
// | element customXml { w_CT_CustomXmlCell }
// | element sdt { w_CT_SdtCell }
// | w_EG_RunLevelElts*
type TRCellContent struct {
	Cell *Cell `xml:"w:tc,omitempty"`
}

func (c TRCellContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.Cell != nil {
		return c.Cell.MarshalXML(e, xml.StartElement{})
	}
	return nil
}

// w_EG_ContentBlockContent =
// element customXml { w_CT_CustomXmlBlock }
// | element sdt { w_CT_SdtBlock }
// | element p { w_CT_P }*
// | element tbl { w_CT_Tbl }*
// | w_EG_RunLevelElts*
// w_CT_SdtContentBlock = w_EG_ContentBlockContent*

//
// w_CT_SdtContentRow = w_EG_ContentRowContent*
//

// w_EG_ContentRowContent
// w_EG_ContentRowContent =
type RowContent struct {
	// element tr { w_CT_Row }*
	Row *Row `xml:"w:tr,omitempty"`
	// | element customXml { w_CT_CustomXmlRow }
	// | element sdt { w_CT_SdtRow }
	// | w_EG_RunLevelElts*
}
