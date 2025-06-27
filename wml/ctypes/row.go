package ctypes

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
