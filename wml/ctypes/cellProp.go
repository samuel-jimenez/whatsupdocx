package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type CellProperty struct {

	// Sequnce:

	// 1. Table Cell Conditional Formatting
	CnfStyle *CTString `xml:"w:cnfStyle,omitempty"`

	// 2. Preferred Table Cell Width
	Width *TableWidth `xml:"w:tcW,omitempty"`

	// 3.Grid Columns Spanned by Current Table Cell
	GridSpan *DecimalNum `xml:"w:gridSpan,omitempty"`

	// 4.Horizontally Merged Cell
	HMerge *GenOptStrVal[stypes.MergeCell] `xml:"w:hMerge,omitempty"`

	// 5.Vertically Merged Cell
	VMerge *GenOptStrVal[stypes.MergeCell] `xml:"w:vMerge,omitempty"`

	// 6.Table Cell Borders
	Borders *CellBorders `xml:"w:tcBorders,omitempty"`

	//7.Table Cell Shading
	Shading *Shading `xml:"w:shd,omitempty"`

	//8.Don't Wrap Cell Content
	NoWrap *OnOff `xml:"w:noWrap,omitempty"`

	//9.Single Table Cell Margins
	Margins *CellMargins `xml:"w:tcMar,omitempty"`

	//10.Table Cell Text Flow Direction
	TextDirection *GenSingleStrVal[stypes.TextDirection] `xml:"w:textDirection,omitempty"`

	//11.Fit Text Within Cell
	FitText *OnOff `xml:"w:tcFitText,omitempty"`

	//12. Table Cell Vertical Alignment
	VAlign *GenSingleStrVal[stypes.VerticalJc] `xml:"w:vAlign,omitempty"`

	//13.Ignore End Of Cell Marker In Row Height Calculation
	HideMark *OnOff `xml:"w:hideMark,omitempty"`

	//14. Choice - ZeroOrOne
	// At max only one of these element should exist

	//Table Cell Insertion
	CellInsertion *TrackChange `xml:"w:cellIns,omitempty"`

	//Table Cell Deletion
	CellDeletion *TrackChange `xml:"w:cellDel,omitempty"`

	//Vertically Merged/Split Table Cells
	CellMerge *CellMerge `xml:"w:cellMerge,omitempty"`

	//15.Revision Information for Table Cell Properties
	PrChange *TCPrChange `xml:"w:tcPrChange,omitempty"`
}

type TCPrChange struct {
	ID     int          `xml:"w:id,attr"`
	Author string       `xml:"w:author,attr"`
	Date   *string      `xml:"w:date,attr,omitempty"`
	Prop   CellProperty `xml:"w:tcPr"`
}
