package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Table Cell Margin Defaults
type CellMargins struct {
	// 1. Table Cell Top Margin Default
	Top *TableWidth `xml:"w:top,omitempty"`

	// 2. Table Cell Left Margin Default
	Left *TableWidth `xml:"w:left,omitempty"`

	// 3. Table Cell Bottom Margin Default
	Bottom *TableWidth `xml:"w:bottom,omitempty"`

	// 4. Table Cell Right Margin Default
	Right *TableWidth `xml:"w:right,omitempty"`
}

func DefaultCellMargins() CellMargins {
	return CellMargins{}
}

func (tcm CellMargins) Margin(top, left, bottom, right int) CellMargins {
	tcm.Top = NewTableWidth(top, stypes.TableWidthDxa)
	tcm.Left = NewTableWidth(left, stypes.TableWidthDxa)
	tcm.Bottom = NewTableWidth(bottom, stypes.TableWidthDxa)
	tcm.Right = NewTableWidth(right, stypes.TableWidthDxa)
	return tcm
}

func (tcm CellMargins) MarginTop(v int, t stypes.TableWidth) CellMargins {
	tcm.Top = NewTableWidth(v, t)
	return tcm
}

func (tcm CellMargins) MarginRight(v int, t stypes.TableWidth) CellMargins {
	tcm.Right = NewTableWidth(v, t)
	return tcm
}

func (tcm CellMargins) MarginLeft(v int, t stypes.TableWidth) CellMargins {
	tcm.Left = NewTableWidth(v, t)
	return tcm
}

func (tcm CellMargins) MarginBottom(v int, t stypes.TableWidth) CellMargins {
	tcm.Bottom = NewTableWidth(v, t)
	return tcm
}
