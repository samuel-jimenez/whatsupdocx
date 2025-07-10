package ctypes

// Table Grid
// w_CT_TblGridBase = element gridCol { w_CT_TblGridCol }*
// w_CT_TblGrid =
// w_CT_TblGridBase,
// element tblGridChange { w_CT_TblGridChange }?
type Grid struct {
	//1. Grid Column Definition
	// element gridCol { w_CT_TblGridCol }*
	Col []Column `xml:"w:gridCol,omitempty"`

	//2.Revision Information for Table Grid Column Definitions
	// element tblGridChange { w_CT_TblGridChange }?
	GridChange *GridChange `xml:"w:tblGridChange,omitempty"`
}
