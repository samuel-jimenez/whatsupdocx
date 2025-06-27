package ctypes

//TODO crossref  w_EG_ContentBlockContent

//TODO
// DocumentChild

// DocumentChild represents a child element within a Word document, which can be a Paragraph or a Table.
// w_EG_BlockLevelElts = w_EG_BlockLevelChunkElts*
// | element altChunk { w_CT_AltChunk }*

// w_CT_AltChunk =
// r_id?,
// element altChunkPr { w_CT_AltChunkPr }?
// w_CT_AltChunkPr = element matchSrc { w_CT_OnOff }?

// w_EG_BlockLevelChunkElts = w_EG_ContentBlockContent*
// w_EG_ContentBlockContent =
// element customXml { w_CT_CustomXmlBlock }
// | element sdt { w_CT_SdtBlock }
// | element p { w_CT_P }*
// | element tbl { w_CT_Tbl }*
// | w_EG_RunLevelElts*

// w_EG_RunLevelElts =
// element proofErr { w_CT_ProofErr }?
// | element permStart { w_CT_PermStart }?
// | element permEnd { w_CT_Perm }?
// | w_EG_RangeMarkupElements*
// | element ins { w_CT_RunTrackChange }?
// | element del { w_CT_RunTrackChange }?
// | element moveFrom { w_CT_RunTrackChange }
// | element moveTo { w_CT_RunTrackChange }
// | w_EG_MathContent*

// Table Cell - ContentBlockContent
// w_EG_BlockLevelChunkElts = w_EG_ContentBlockContent*
// w_EG_BlockLevelElts =
// w_EG_BlockLevelChunkElts*
// | element altChunk { w_CT_AltChunk }*
type BlockLevel struct {
	//Paragraph
	//	- ZeroOrMore: Any number of times Paragraph can repeat within cell
	// | element p { w_CT_P }*
	Paragraph *Paragraph `xml:"w:p,omitempty"`
	//Table
	//	- ZeroOrMore: Any number of times Table can repeat within cell
	// | element tbl { w_CT_Tbl }*
	Table *Table `xml:"w:tbl,omitempty"`
}
