package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Table Row Properties
type RowProperty struct {
	//1. Choice - ZeroOrMore

	//Table Row Conditional Formatting
	Cnf *CTString `xml:"w:cnfStyle,omitempty"`

	// Associated HTML div ID
	DivId *DecimalNum `xml:"w:divId,omitempty"`

	//Grid Columns Before First Cell
	GridBefore *DecimalNum `xml:"w:gridBefore,omitempty"`

	//Grid Columns After Last Cell
	GridAfter *DecimalNum `xml:"w:gridAfter,omitempty"`

	//Preferred Width Before Table Row
	WidthBefore *TableWidth `xml:"w:tblWBefore,omitempty"`

	//Preferred Width After Table Row
	WidthAfter *TableWidth `xml:"w:tblWAfter,omitempty"`

	//Table Row Cannot Break Across Pages
	CantSplit *OnOff `xml:"w:cantSplit,omitempty"`

	//Table Row Height
	Height *TableRowHeight `xml:"w:trHeight,omitempty"`

	//Repeat Table Row on Every New Page
	Header *OnOff `xml:"w:tblHeader,omitempty"`

	//Table Row Cell Spacing
	CellSpacing *TableWidth `xml:"w:tblCellSpacing,omitempty"`

	// Table Row Alignment
	JC *GenSingleStrVal[stypes.Justification] `xml:"w:jc,omitempty"`

	//Hidden Table Row Marker
	Hidden *OnOff `xml:"w:hidden,omitempty"`

	//2.Inserted Table Row
	Ins *TrackChange `xml:"w:ins,omitempty"`

	//3. Deleted Table Row
	Del *TrackChange `xml:"w:del,omitempty"`

	//4.Revision Information for Table Row Properties
	Change *TRPrChange `xml:"w:trPrChange,omitempty"`
}

// NewRowProperty creates a new RowProperty instance.
func DefaultRowProperty() *RowProperty {
	return &RowProperty{}
}

type TRPrChange struct {
	ID     int         `xml:"w:id,attr"`
	Author string      `xml:"w:author,attr"`
	Date   *string     `xml:"w:date,attr,omitempty"`
	Prop   RowProperty `xml:"w:tcPr"`
}
