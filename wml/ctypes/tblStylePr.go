package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Style Conditional Table Formatting Properties
type TableStyleProp struct {
	// Sequence:

	//1.Table Style Conditional Formatting Paragraph Properties
	ParaProp *ParagraphProp `xml:"w:pPr,omitempty"`

	//2.Table Style Conditional Formatting Run Properties
	RunProp *RunProperty `xml:"w:rPr,omitempty"`

	//3.Table Style Conditional Formatting Table Properties
	TableProp *TableProp `xml:"w:tblPr,omitempty"`

	//4.Table Style Conditional Formatting Table Row Properties
	RowProp *RowProperty `xml:"w:trPr,omitempty"`

	//5.Table Style Conditional Formatting Table Cell Properties
	CellProp *CellProperty `xml:"w:tcPr,omitempty"`

	//Attributes:

	//Table Style Conditional Formatting Type
	Type stypes.TblStyleOverrideType `xml:"w:type,attr"`
}
