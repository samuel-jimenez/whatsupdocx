package ctypes

import (
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
	Style *CTString `xml:"w:tblStyle,omitempty"`

	// 2. Floating Table Positioning
	// element tblpPr { w_CT_TblPPr }?,
	FloatPos *FloatPos `xml:"w:tblpPr,omitempty"`

	// 3.Floating Table Allows Other Tables to Overlap
	// element tblOverlap { w_CT_TblOverlap }?,
	Overlap *GenSingleStrVal[stypes.TblOverlap] `xml:"w:tblOverlap,omitempty"`

	// 4. Visually Right to Left Table
	// element bidiVisual { w_CT_OnOff }?,
	BidiVisual *OnOff `xml:"w:bidiVisual,omitempty"`

	// 5. Number of Rows in Row Band
	// element tblStyleRowBandSize { w_CT_DecimalNumber }?,
	RowCountInRowBand *DecimalNum `xml:"w:tblStyleRowBandSize,omitempty"`

	// 6. Number of Columns in Column Band
	// element tblStyleColBandSize { w_CT_DecimalNumber }?,
	RowCountInColBand *DecimalNum `xml:"w:tblStyleColBandSize,omitempty"`

	// 7. Preferred Table Width
	// element tblW { w_CT_TblWidth }?,
	Width *TableWidth `xml:"w:tblW,omitempty"`

	// 8.Table Alignment
	// element jc { w_CT_JcTable }?,
	Justification *GenSingleStrVal[stypes.Justification] `xml:"w:jc,omitempty"`

	// 9.Table Cell Spacing Default
	// element tblCellSpacing { w_CT_TblWidth }?,
	CellSpacing *TableWidth `xml:"w:blCellSpacing,omitempty"`

	// 10. Table Indent from Leading Margin
	// element tblInd { w_CT_TblWidth }?,
	Indent *TableWidth `xml:"w:tblInd,omitempty"`

	// 11. Table Indent from Leading Margin
	// element tblBorders { w_CT_TblBorders }?,
	Borders *TableBorders `xml:"w:tblBorders,omitempty"`

	// 12. Table Shading
	// element shd { w_CT_Shd }?,
	Shading *Shading `xml:"w:shd,omitempty"`

	// 13. Table Layout
	// element tblLayout { w_CT_TblLayoutType }?,
	Layout *TableLayout `xml:"w:tblLayout,omitempty"`

	// 14. Table Cell Margin Defaults
	// element tblCellMar { w_CT_TblCellMar }?,
	CellMargin *CellMargins `xml:"w:tblCellMar,omitempty"`

	// 15. Table Style Conditional Formatting Settings
	// element tblLook { w_CT_TblLook }?,
	TableLook *CTString `xml:"w:tblLook,omitempty"`

	//TODO
	// element tblCaption { w_CT_String }?,
	// element tblDescription { w_CT_String }?

	//16. Revision Information for Table Properties
	// element tblPrChange { w_CT_TblPrChange }?
	PrChange *TblPrChange `xml:"w:tblPrChange,omitempty"`
}

func DefaultTableProp() *TableProp {
	return &TableProp{}
}

type TblPrChange struct {
	ID     int       `xml:"w:id,attr"`
	Author string    `xml:"w:author,attr"`
	Date   *string   `xml:"w:date,attr,omitempty"`
	Prop   TableProp `xml:"w:tblPr"`
}
