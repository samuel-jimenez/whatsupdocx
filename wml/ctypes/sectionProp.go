package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Document Final Section Properties : w:sectPr
// CT_SectPr
// w_CT_SectPr =
// w_AG_SectPrAttributes,
// w_EG_HdrFtrReferences*,
// w_EG_SectPrContents?,
// element sectPrChange { w_CT_SectPrChange }?

// w_EG_SectPrContents =
// element footnotePr { w_CT_FtnProps }?,
// element endnotePr { w_CT_EdnProps }?,
// element type { w_CT_SectType }?,
// element pgSz { w_CT_PageSz }?,
// element pgMar { w_CT_PageMar }?,
// element paperSrc { w_CT_PaperSource }?,
// element pgBorders { w_CT_PageBorders }?,
// element lnNumType { w_CT_LineNumber }?,
// element pgNumType { w_CT_PageNumber }?,
// element cols { w_CT_Columns }?,
// element formProt { w_CT_OnOff }?,
// element vAlign { w_CT_VerticalJc }?,
// element noEndnote { w_CT_OnOff }?,
// element titlePg { w_CT_OnOff }?,
// element textDirection { w_CT_TextDirection }?,
// element bidi { w_CT_OnOff }?,
// element rtlGutter { w_CT_OnOff }?,
// element docGrid { w_CT_DocGrid }?,
// element printerSettings { w_CT_Rel }?

// w_CT_SectPr =
type SectionProp struct {
	// TODO
	// w_AG_SectPrAttributes,
	// 		// w_AG_SectPrAttributes =
	// 		// attribute w:rsidRPr { w_ST_LongHexNumber }?,
	// 		// attribute w:rsidDel { w_ST_LongHexNumber }?,
	// 		// attribute w:rsidR { w_ST_LongHexNumber }?,
	// 		// attribute w:rsidSect { w_ST_LongHexNumber }?
	//

	// w_EG_HdrFtrReferences*,
	// // w_EG_HdrFtrReferences =
	// // element headerReference { w_CT_HdrFtrRef }?
	// // | element footerReference { w_CT_HdrFtrRef }?
	HeaderReference *HeaderFooterReference `xml:"w:headerReference,omitempty"`
	FooterReference *HeaderFooterReference `xml:"w:footerReference,omitempty"`

	// w_EG_SectPrContents?,
	// w_EG_SectPrContents =
	// TODO
	// element footnotePr { w_CT_FtnProps }?,
	// element endnotePr { w_CT_EdnProps }?,
	// element type { w_CT_SectType }?,
	Type *GenSingleStrVal[stypes.SectionMark] `xml:"w:type,omitempty"`
	// element pgSz { w_CT_PageSz }?,
	PageSize *PageSize `xml:"w:pgSz,omitempty"`
	// element pgMar { w_CT_PageMar }?,
	PageMargin *PageMargin `xml:"w:pgMar,omitempty"`
	// element paperSrc { w_CT_PaperSource }?,
	// element pgBorders { w_CT_PageBorders }?,
	PageBorders *PageBorders `xml:"w:pgBorders,omitempty"`
	// element lnNumType { w_CT_LineNumber }?,
	// element pgNumType { w_CT_PageNumber }?,
	PageNum *PageNumbering `xml:"w:pgNumType,omitempty"`
	// element cols { w_CT_Columns }?,
	// element formProt { w_CT_OnOff }?,
	FormProt *GenSingleStrVal[stypes.OnOff] `xml:"w:formProt,omitempty"`
	// element vAlign { w_CT_VerticalJc }?,
	// element noEndnote { w_CT_OnOff }?,
	// element titlePg { w_CT_OnOff }?,
	TitlePg *GenSingleStrVal[stypes.OnOff] `xml:"w:titlePg,omitempty"`
	// element textDirection { w_CT_TextDirection }?,
	TextDir *GenSingleStrVal[stypes.TextDirection] `xml:"w:textDirection,omitempty"`
	// element bidi { w_CT_OnOff }?,
	// element rtlGutter { w_CT_OnOff }?,
	// element docGrid { w_CT_DocGrid }?,
	DocGrid *DocGrid `xml:"w:docGrid,omitempty"`
	// element printerSettings { w_CT_Rel }?

	// element sectPrChange { w_CT_SectPrChange }?
	// w_CT_SectPrChange =
	// w_CT_TrackChange,
	// element sectPr { w_CT_SectPrBase }?
	// w_CT_SectPrBase = w_AG_SectPrAttributes, w_EG_SectPrContents?
}

func NewSectionProper() *SectionProp {
	return &SectionProp{}
}
