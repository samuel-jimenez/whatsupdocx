package ctypes

import (
	"github.com/samuel-jimenez/xml"

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
	// w_AG_SectPrAttributes =
	// attribute w:rsidRPr { w_ST_LongHexNumber }?,
	// attribute w:rsidDel { w_ST_LongHexNumber }?,
	// attribute w:rsidR { w_ST_LongHexNumber }?,
	// attribute w:rsidSect { w_ST_LongHexNumber }?

	// w_EG_HdrFtrReferences*,
	// w_EG_HdrFtrReferences =
	// element headerReference { w_CT_HdrFtrRef }?
	// | element footerReference { w_CT_HdrFtrRef }?
	HeaderReference *HeaderFooterReference `xml:"headerReference,omitempty"`
	FooterReference *HeaderFooterReference `xml:"footerReference,omitempty"`

	// w_EG_SectPrContents?,
	// w_EG_SectPrContents =
	// TODO
	// element footnotePr { w_CT_FtnProps }?,
	// element endnotePr { w_CT_EdnProps }?,
	// element type { w_CT_SectType }?,
	Type *GenSingleStrVal[stypes.SectionMark] `xml:"type,omitempty"`
	// element pgSz { w_CT_PageSz }?,
	PageSize *PageSize `xml:"pgSz,omitempty"`
	// element pgMar { w_CT_PageMar }?,
	PageMargin *PageMargin `xml:"pgMar,omitempty"`
	// element paperSrc { w_CT_PaperSource }?,
	// element pgBorders { w_CT_PageBorders }?,
	PageBorders *PageBorders `xml:"pgBorders,omitempty"`
	// element lnNumType { w_CT_LineNumber }?,
	// element pgNumType { w_CT_PageNumber }?,
	PageNum *PageNumbering `xml:"pgNumType,omitempty"`
	// element cols { w_CT_Columns }?,
	// element formProt { w_CT_OnOff }?,
	FormProt *GenSingleStrVal[stypes.OnOff] `xml:"formProt,omitempty"`
	// element vAlign { w_CT_VerticalJc }?,
	// element noEndnote { w_CT_OnOff }?,
	// element titlePg { w_CT_OnOff }?,
	TitlePg *GenSingleStrVal[stypes.OnOff] `xml:"titlePg,omitempty"`
	// element textDirection { w_CT_TextDirection }?,
	TextDir *GenSingleStrVal[stypes.TextDirection] `xml:"textDirection,omitempty"`
	// element bidi { w_CT_OnOff }?,
	// element rtlGutter { w_CT_OnOff }?,
	// element docGrid { w_CT_DocGrid }?,
	DocGrid *DocGrid `xml:"docGrid,omitempty"`
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

func (s SectionProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:sectPr"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// w_EG_HdrFtrReferences*,
	// w_EG_HdrFtrReferences =
	// element headerReference { w_CT_HdrFtrRef }?
	// | element footerReference { w_CT_HdrFtrRef }?
	if s.HeaderReference != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:headerReference"}}
		if err = e.EncodeElement(s.HeaderReference, propsElement); err != nil {
			return err
		}
	}

	if s.FooterReference != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:footerReference"}}
		if err = e.EncodeElement(s.FooterReference, propsElement); err != nil {
			return err
		}
	}

	// element footnotePr { w_CT_FtnProps }?,
	// element endnotePr { w_CT_EdnProps }?,
	// element type { w_CT_SectType }?,
	if s.Type != nil {
		if err := s.Type.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:type"},
		}); err != nil {
			return err
		}
	}

	// element pgSz { w_CT_PageSz }?,
	if s.PageSize != nil {
		if err := s.PageSize.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	// element pgMar { w_CT_PageMar }?,
	if s.PageMargin != nil {
		if err = s.PageMargin.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	// element paperSrc { w_CT_PaperSource }?,
	// element pgBorders { w_CT_PageBorders }?,
	if s.PageBorders != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:pgBorders"}}
		if err = s.PageBorders.MarshalXML(e, propsElement); err != nil {
			return err
		}
	}

	// element lnNumType { w_CT_LineNumber }?,
	// element pgNumType { w_CT_PageNumber }?,
	if s.PageNum != nil {
		if err = s.PageNum.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	// element cols { w_CT_Columns }?,
	// element formProt { w_CT_OnOff }?,
	if s.FormProt != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:formProt"}}
		if err = s.FormProt.MarshalXML(e, propsElement); err != nil {
			return err
		}
	}

	// element vAlign { w_CT_VerticalJc }?,
	// element noEndnote { w_CT_OnOff }?,
	// element titlePg { w_CT_OnOff }?,
	if s.TitlePg != nil {
		if err = s.TitlePg.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:titlePg"},
		}); err != nil {
			return err
		}
	}

	// element textDirection { w_CT_TextDirection }?,
	if s.TextDir != nil {
		if s.TextDir.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:textDirection"},
		}); err != nil {
			return err
		}
	}

	// element bidi { w_CT_OnOff }?,
	// element rtlGutter { w_CT_OnOff }?,
	// element docGrid { w_CT_DocGrid }?,
	if s.DocGrid != nil {
		if s.DocGrid.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	// element printerSettings { w_CT_Rel }?

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
