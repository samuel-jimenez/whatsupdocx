package ctypes

import (
	"strconv"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// This element specifies the set of table-wide properties applied to the current table. These properties affect the appearance of all rows and cells within the parent table, but may be overridden by individual table-level exception, row, and cell level properties as defined by each TableProp.
// w_CT_TblPr =
// w_CT_TblPrBase,
type TableProp struct {
	// 	w_CT_TblPrBase =
	// element tblStyle { w_CT_String }?,

	// 1. Referenced Table Style
	// TableStyle represents the style of a table in a document.
	// This is applicable when creating a new document. When using this style in a new document, you need to ensure
	// that the specified style ID exists in your document's style base or is manually created through the library.
	//
	// Some examples of predefined style IDs that can be used for table styles:
	//
	//   - "LightShading"
	//   - "LightShading-Accent1"
	//   - "LightShading-Accent2"
	//   - "LightShading-Accent3"
	//   - "LightShading-Accent4"
	//   - "LightShading-Accent5"
	//   - "LightShading-Accent6"
	//   - "LightList"
	//   - "LightList-Accent1"..."LightList-Accent6"
	//   - "LightGrid"
	//   - "LightGrid-Accent1"..."LightGrid-Accent6"
	//   - "MediumShading"
	//   - "MediumShading-Accent1"..."MediumShading-Accent6"
	//   - "MediumShading2"
	//   - "MediumShading2-Accent1"..."MediumShading2-Accent6"
	//   - "MediumList1"
	//   - "MediumList1-Accent1"..."MediumList1-Accent6"
	//   - "MediumList2"
	//   - "MediumList2-Accent1"..."MediumList2-Accent6"
	//   - "TableGrid"
	//   - "MediumGrid1"
	//   - "MediumGrid1-Accent1"..."MediumGrid1-Accent6"
	//   - "MediumGrid2"
	//   - "MediumGrid2-Accent1"..."MediumGrid2-Accent6"
	//   - "MediumGrid3"
	//   - "MediumGrid3-Accent1"..."MediumGrid3-Accent6"
	//   - "DarkList"
	//   - "DarkList-Accent1"..."DarkList-Accent6"
	//   - "ColorfulShading"
	//   - "ColorfulShading-Accent1"..."ColorfulShading-Accent6"
	//   - "ColorfulList"
	//   - "ColorfulList-Accent1"..."ColorfulList-Accent6"
	//   - "ColorfulGrid"
	//   - "ColorfulGrid-Accent1"..."ColorfulGrid-Accent6"
	Style *CTString `xml:"tblStyle,omitempty"`

	// 2. Floating Table Positioning
	// element tblpPr { w_CT_TblPPr }?,
	FloatPos *FloatPos `xml:"tblpPr,omitempty"`

	// 3.Floating Table Allows Other Tables to Overlap
	// element tblOverlap { w_CT_TblOverlap }?,
	Overlap *GenSingleStrVal[stypes.TblOverlap] `xml:"tblOverlap,omitempty"`

	// 4. Visually Right to Left Table
	// element bidiVisual { w_CT_OnOff }?,
	BidiVisual *OnOff `xml:"bidiVisual,omitempty"`

	// 5. Number of Rows in Row Band
	// element tblStyleRowBandSize { w_CT_DecimalNumber }?,
	RowCountInRowBand *DecimalNum `xml:"tblStyleRowBandSize,omitempty"`

	// 6. Number of Columns in Column Band
	// element tblStyleColBandSize { w_CT_DecimalNumber }?,
	RowCountInColBand *DecimalNum `xml:"tblStyleColBandSize,omitempty"`

	// 7. Preferred Table Width
	// element tblW { w_CT_TblWidth }?,
	Width *TableWidth `xml:"tblW,omitempty"`

	// 8.Table Alignment
	// element jc { w_CT_JcTable }?,
	Justification *GenSingleStrVal[stypes.Justification] `xml:"jc,omitempty"`

	// 9.Table Cell Spacing Default
	// element tblCellSpacing { w_CT_TblWidth }?,
	CellSpacing *TableWidth `xml:"blCellSpacing,omitempty"`

	// 10. Table Indent from Leading Margin
	// element tblInd { w_CT_TblWidth }?,
	Indent *TableWidth `xml:"tblInd,omitempty"`

	// 11. Table Indent from Leading Margin
	// element tblBorders { w_CT_TblBorders }?,
	Borders *TableBorders `xml:"tblBorders,omitempty"`

	// 12. Table Shading
	// element shd { w_CT_Shd }?,
	Shading *Shading `xml:"shd,omitempty"`

	// 13. Table Layout
	// element tblLayout { w_CT_TblLayoutType }?,
	Layout *TableLayout `xml:"tblLayout,omitempty"`

	// 14. Table Cell Margin Defaults
	// element tblCellMar { w_CT_TblCellMar }?,
	CellMargin *CellMargins `xml:"tblCellMar,omitempty"`

	// 15. Table Style Conditional Formatting Settings
	// element tblLook { w_CT_TblLook }?,
	TableLook *CTString `xml:"tblLook,omitempty"`

	//TODO
	// element tblCaption { w_CT_String }?,
	// element tblDescription { w_CT_String }?

	//16. Revision Information for Table Properties
	// element tblPrChange { w_CT_TblPrChange }?
	PrChange *TblPrChange `xml:"tblPrChange,omitempty"`
}

func DefaultTableProp() *TableProp {
	return &TableProp{}
}

func (t TableProp) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tblPr"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. tblStyle
	if t.Style != nil {
		if err = t.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblStyle"},
		}); err != nil {
			return err
		}
	}

	// 2. tblpPr
	if t.FloatPos != nil {
		if err = t.FloatPos.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblpPr"},
		}); err != nil {
			return err
		}
	}

	// 3. tblOverlap
	if t.Overlap != nil {
		if err = t.Overlap.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblOverlap"},
		}); err != nil {
			return err
		}
	}

	// 4. tblOverlap
	if t.BidiVisual != nil {
		if err = t.BidiVisual.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bidiVisual"},
		}); err != nil {
			return err
		}
	}

	// 5. tblStyleRowBandSize
	if t.RowCountInRowBand != nil {
		if err = t.RowCountInRowBand.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblStyleRowBandSize"},
		}); err != nil {
			return err
		}
	}

	// 6. tblStyleColBandSize
	if t.RowCountInColBand != nil {
		if err = t.RowCountInColBand.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblStyleColBandSize"},
		}); err != nil {
			return err
		}
	}

	// 7. tblStyleColBandSize
	if t.Width != nil {
		if err = t.Width.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblW"},
		}); err != nil {
			return err
		}
	}

	// 8. jc
	if t.Justification != nil {
		if err = t.Justification.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:jc"},
		}); err != nil {
			return err
		}
	}

	// 9. blCellSpacing
	if t.CellSpacing != nil {
		if err = t.CellSpacing.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:blCellSpacing"},
		}); err != nil {
			return err
		}
	}

	// 10. tblInd
	if t.Indent != nil {
		if err = t.Indent.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblInd"},
		}); err != nil {
			return err
		}
	}

	// 11. tblBorders
	if t.Borders != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:tblBorders"}}
		if err := e.EncodeElement(t.Borders, propsElement); err != nil {
			// if err := t.Borders.MarshalXML(e, propsElement); err != nil {
			return err
		}
	}

	// 12. shd
	if t.Shading != nil {
		if err = t.Shading.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:shd"},
		}); err != nil {
			return err
		}
	}

	// 13. tblLayout
	if t.Layout != nil {
		if err = t.Layout.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblLayout"},
		}); err != nil {
			return err
		}
	}

	// 14. CellMargin
	if t.CellMargin != nil {
		if err = t.CellMargin.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblCellMar"},
		}); err != nil {
			return err
		}
	}

	// 15. TableLook
	if t.TableLook != nil {
		if err = t.TableLook.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblLook"},
		}); err != nil {
			return err
		}
	}

	// 16. tblPrChange
	if t.PrChange != nil {
		if err = t.PrChange.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblPrChange"},
		}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type TblPrChange struct {
	ID     int       `xml:"id,attr"`
	Author string    `xml:"author,attr"`
	Date   *string   `xml:"date,attr,omitempty"`
	Prop   TableProp `xml:"tblPr"`
}

func (t TblPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblPrChange"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(t.ID)},
		{Name: xml.Name{Local: "w:author"}, Value: t.Author},
	}

	if t.Date != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"}, Value: *t.Date})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if err := t.Prop.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
