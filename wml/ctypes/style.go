package ctypes

import (
	"fmt"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

var defaultStyleNSAttrs = map[string]string{
	"xmlns:w":      "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	"xmlns:mc":     "http://schemas.openxmlformats.org/markup-compatibility/2006",
	"xmlns:w14":    "http://schemas.microsoft.com/office/word/2010/wordml",
	"mc:Ignorable": "w14",
}

// Style Definitions
// w_styles = element styles { w_CT_Styles }
//
// w_ST_StyleType =
// string "paragraph"
// | string "character"
// | string "table"
// | string "numbering"
//
// w_CT_Style =
// attribute w:type { w_ST_StyleType }?,
// attribute w:styleId { s_ST_String }?,
// attribute w:default { s_ST_OnOff }?,
// attribute w:customStyle { s_ST_OnOff }?,
// element name { w_CT_String }?,
// element aliases { w_CT_String }?,
// element basedOn { w_CT_String }?,
// element next { w_CT_String }?,
// element link { w_CT_String }?,
// element autoRedefine { w_CT_OnOff }?,
// element hidden { w_CT_OnOff }?,
// element uiPriority { w_CT_DecimalNumber }?,
// element semiHidden { w_CT_OnOff }?,
// element unhideWhenUsed { w_CT_OnOff }?,
// element qFormat { w_CT_OnOff }?,
// element locked { w_CT_OnOff }?,
// element personal { w_CT_OnOff }?,
// element personalCompose { w_CT_OnOff }?,
// element personalReply { w_CT_OnOff }?,
// element rsid { w_CT_LongHexNumber }?,
// element pPr { w_CT_PPrGeneral }?,
// element rPr { w_CT_RPr }?,
// element tblPr { w_CT_TblPrBase }?,
// element trPr { w_CT_TrPr }?,
// element tcPr { w_CT_TcPr }?,
// element tblStylePr { w_CT_TblStylePr }*
// w_CT_LsdException =
// attribute w:name { s_ST_String },
// attribute w:locked { s_ST_OnOff }?,
// attribute w:uiPriority { w_ST_DecimalNumber }?,
// attribute w:semiHidden { s_ST_OnOff }?,
// attribute w:unhideWhenUsed { s_ST_OnOff }?,
// attribute w:qFormat { s_ST_OnOff }?
// w_CT_LatentStyles =
// attribute w:defLockedState { s_ST_OnOff }?,
// attribute w:defUIPriority { w_ST_DecimalNumber }?,
// attribute w:defSemiHidden { s_ST_OnOff }?,
// attribute w:defUnhideWhenUsed { s_ST_OnOff }?,
// attribute w:defQFormat { s_ST_OnOff }?,
// attribute w:count { w_ST_DecimalNumber }?,
// element lsdException { w_CT_LsdException }*
//
// w_CT_Styles =
// element docDefaults { w_CT_DocDefaults }?,
// element latentStyles { w_CT_LatentStyles }?,
// element style { w_CT_Style }*
type Styles struct {
	XMLName xml.Name `xml:"w:styles"`

	RelativePath string `xml:"-"`
	Attr         []xml.Attr

	// Sequence

	//1. Document Default Paragraph and Run Properties
	DocDefaults *DocDefault `xml:"w:docDefaults,omitempty"`

	//2. Latent Style Information
	LatentStyle *LatentStyle `xml:"w:latentStyles,omitempty"`

	//3. Style Definition
	StyleList []Style `xml:",any"`
}

func (s *Styles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:styles"

	if len(s.Attr) == 0 {
		for key, value := range defaultStyleNSAttrs {
			attr := xml.Attr{Name: xml.Name{Local: key}, Value: value}
			start.Attr = append(start.Attr, attr)
		}
	} else {
		start.Attr = s.Attr
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 1. Document Default Paragraph and Run Properties
	if s.DocDefaults != nil {
		if err := s.DocDefaults.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:docDefaults"}}); err != nil {
			return fmt.Errorf("docDefaults: %w", err)
		}
	}

	// 2. Latent Style Information
	if s.LatentStyle != nil {
		if err := s.LatentStyle.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:latentStyles"}}); err != nil {
			return fmt.Errorf("latentStyle: %w", err)
		}
	}

	//3. Style Definition
	for _, elem := range s.StyleList {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:style"}}
		if err := e.EncodeElement(elem, propsElement); err != nil {
			// if err := elem.MarshalXML(e, propsElement); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}

type Style struct {
	//Sequence:

	//1. Primary Style Name
	Name *CTString `xml:"w:name,omitempty"`

	//2. Alternate Style Names
	Alias *CTString `xml:"w:alias,omitempty"`

	//3. Parent Style ID
	BasedOn *CTString `xml:"w:basedOn,omitempty"`

	//4. Style For Next Paragraph
	Next *CTString `xml:"w:next,omitempty"`

	//5. Linked Style Reference
	Link *CTString `xml:"w:link,omitempty"`

	//6.Automatically Merge User Formatting Into Style Definition
	AutoRedefine *OnOff `xml:"w:autoRedefine,omitempty"`

	//7.Hide Style From User Interface
	Hidden *OnOff `xml:"w:hidden,omitempty"`

	//8.Optional User Interface Sorting Order
	UIPriority *DecimalNum `xml:"w:uiPriority,omitempty"`

	// 9. Hide Style From Main User Interface
	SemiHidden *OnOff `xml:"w:semiHidden,omitempty"`

	// 10. Remove Semi-Hidden Property When Style Is Used
	UnhideWhenUsed *OnOff `xml:"w:unhideWhenUsed,omitempty"`

	// 11. Primary Style
	QFormat *OnOff `xml:"w:qFormat,omitempty"`

	// 12. Style Cannot Be Applied
	Locked *OnOff `xml:"w:locked,omitempty"`

	// 13. E-Mail Message Text Style
	Personal *OnOff `xml:"w:personal,omitempty"`

	// 14. E-Mail Message Composition Style
	PersonalCompose *OnOff `xml:"w:personalCompose,omitempty"`

	// 15. E-Mail Message Reply Style
	PersonalReply *OnOff `xml:"w:personalReply,omitempty"`

	//16. Revision Identifier for Style Definition
	RevID *GenSingleStrVal[stypes.LongHexNum] `xml:"w:rsid,omitempty"`

	//17. Style Paragraph Properties
	ParaProp *ParagraphProp `xml:"w:pPr,omitempty"`

	//18. Run Properties
	RunProp *RunProperty `xml:"w:rPr,omitempty"`

	//19. Style Table Properties
	TableProp *TableProp `xml:"w:tblPr,omitempty"`

	//20. Style Table Row Properties
	TableRowProp *RowProperty `xml:"w:trPr,omitempty"`

	//21. Style Table Cell Properties
	TableCellProp *CellProperty `xml:"w:tcPr,omitempty"`

	//22.Style Conditional Table Formatting Properties
	TableStylePr []TableStyleProp `xml:",any"`

	// Attributes

	//Style Type
	Type *stypes.StyleType `xml:"w:type,attr,omitempty"`

	//Style ID
	ID *string `xml:"w:styleId,attr,omitempty"`

	//Default Style
	Default *stypes.OnOff `xml:"w:default,attr,omitempty"`

	//User-Defined Style
	CustomStyle *stypes.OnOff `xml:"w:customStyle,attr,omitempty"`
}
