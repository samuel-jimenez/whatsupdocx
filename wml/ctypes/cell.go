package ctypes

// w_CT_Tc =
type Cell struct {
	// attribute w:id { s_ST_String }?,

	// 1.Table Cell Properties
	// element tcPr { w_CT_TcPr }?,
	Property *CellProperty `xml:"w:tcPr,omitempty"`

	// 2.1 Choice: ZeroOrMore
	// Any number of elements can exists within this choice group
	// w_EG_BlockLevelElts+
	Contents []BlockLevel `xml:",group,any,omitempty"`

	//TODO: Remaining choices
}

func DefaultCell() *Cell {
	return &Cell{
		Property: &CellProperty{
			Shading: DefaultShading(),
		},
	}
}
