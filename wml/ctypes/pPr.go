package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Numbering Level Associated Paragraph Properties
// w_CT_PPr =
// w_CT_PPrBase,
// element rPr { w_CT_ParaRPr }?,
// element sectPr { w_CT_SectPr }?,
// element pPrChange { w_CT_PPrChange }?
// w_CT_PPrBase =
type ParagraphProp struct {

	// 1. This element specifies the style ID of the paragraph style which shall be used to format the contents of this paragraph.
	// element pStyle { w_CT_String }?,
	Style *CTString `xml:"w:pStyle,omitempty"`

	// 2. Keep Paragraph With Next Paragraph
	// element keepNext { w_CT_OnOff }?,
	KeepNext *OnOff `xml:"w:keepNext,omitempty"`

	// 3. Keep All Lines On One Page
	// element keepLines { w_CT_OnOff }?,
	KeepLines *OnOff `xml:"w:keepLines,omitempty"`

	// 4. Start Paragraph on Next Page
	// element pageBreakBefore { w_CT_OnOff }?,
	PageBreakBefore *OnOff `xml:"w:pageBreakBefore,omitempty"`

	// 5. Text Frame Properties
	// element framePr { w_CT_FramePr }?,
	FrameProp *FrameProp `xml:"w:framePr,omitempty"`

	// 6. Allow First/Last Line to Display on a Separate Page
	// element widowControl { w_CT_OnOff }?,
	WindowControl *OnOff `xml:"w:widowControl,omitempty"`

	// 7. Numbering Definition Instance Reference
	// element numPr { w_CT_NumPr }?,
	NumProp *NumProp `xml:"w:w:numPr,omitempty"`

	// 8. Suppress Line Numbers for Paragraph
	// element suppressLineNumbers { w_CT_OnOff }?,
	SuppressLineNmbrs *OnOff `xml:"w:suppressLineNumbers,omitempty"`

	// 9. Paragraph Borders
	// element pBdr { w_CT_PBdr }?,
	Border *ParaBorder `xml:"w:w:pBdr,omitempty"`

	// 10. This element specifies the shading applied to the contents of the paragraph.
	// element shd { w_CT_Shd }?,
	Shading *Shading `xml:"w:shd,omitempty"`

	// 11. Set of Custom Tab Stops
	// element tabs { w_CT_Tabs }?,
	Tabs Tabs `xml:"w:tabs,omitempty"`

	// 12. Suppress Hyphenation for Paragraph
	// element suppressAutoHyphens { w_CT_OnOff }?,
	SuppressAutoHyphens *OnOff `xml:"w:suppressAutoHyphens,omitempty"`

	// 13. Use East Asian Typography Rules for First and Last Character per Line
	// element kinsoku { w_CT_OnOff }?,
	Kinsoku *OnOff `xml:"w:kinsoku,omitempty"`

	// 14. Allow Line Breaking At Character Level
	// element wordWrap { w_CT_OnOff }?,
	WordWrap *OnOff `xml:"w:wordWrap,omitempty"`

	// 15. Allow Punctuation to Extent Past Text Extents
	// element overflowPunct { w_CT_OnOff }?,
	OverflowPunct *OnOff `xml:"w:overflowPunct,omitempty"`

	// 16. Compress Punctuation at Start of a Line
	// element topLinePunct { w_CT_OnOff }?,
	TopLinePunct *OnOff `xml:"w:topLinePunct,omitempty"`

	// 17. Automatically Adjust Spacing of Latin and East Asian Text
	// element autoSpaceDE { w_CT_OnOff }?,
	AutoSpaceDE *OnOff `xml:"w:autoSpaceDE,omitempty"`

	// 18. Automatically Adjust Spacing of East Asian Text and Numbers
	// element autoSpaceDN { w_CT_OnOff }?,
	AutoSpaceDN *OnOff `xml:"w:autoSpaceDN,omitempty"`

	// 19. Right to Left Paragraph Layout
	// element bidi { w_CT_OnOff }?,
	Bidi *OnOff `xml:"w:bidi,omitempty"`

	// 20. Automatically Adjust Right Indent When Using Document Grid
	// element adjustRightInd { w_CT_OnOff }?,
	AdjustRightInd *OnOff `xml:"w:adjustRightInd,omitempty"`

	// 21. Use Document Grid Settings for Inter-Line Paragraph Spacing
	// element snapToGrid { w_CT_OnOff }?,
	SnapToGrid *OnOff `xml:"w:snapToGrid,omitempty"`

	// 22. Spacing Between Lines and Above/Below Paragraph
	// element spacing { w_CT_Spacing }?,
	Spacing *Spacing `xml:"w:spacing,omitempty"`

	// 23. Paragraph Indentation
	// element ind { w_CT_Ind }?,
	Indent *Indent `xml:"w:ind,omitempty"`

	// 24. Ignore Spacing Above and Below When Using Identical Styles
	// element contextualSpacing { w_CT_OnOff }?,
	CtxlSpacing *OnOff `xml:"w:contextualSpacing,omitempty"`

	// 25. Use Left/Right Indents as Inside/Outside Indents
	// element mirrorIndents { w_CT_OnOff }?,
	MirrorIndents *OnOff `xml:"w:mirrorIndents,omitempty"`

	// 26. Prevent Text Frames From Overlapping
	// element suppressOverlap { w_CT_OnOff }?,
	SuppressOverlap *OnOff `xml:"w:suppressOverlap,omitempty"`

	// 27. Paragraph Alignment
	// element jc { w_CT_Jc }?,
	Justification *GenSingleStrVal[stypes.Justification] `xml:"w:jc,omitempty"`

	// 28. Paragraph Text Flow Direction
	// element textDirection { w_CT_TextDirection }?,
	TextDirection *GenSingleStrVal[stypes.TextDirection] `xml:"w:textDirection,omitempty"`

	// 29. Vertical Character Alignment on Line
	// element textAlignment { w_CT_TextAlignment }?,
	TextAlignment *GenSingleStrVal[stypes.TextAlign] `xml:"w:textAlignment,omitempty"`

	// 30.Allow Surrounding Paragraphs to Tight Wrap to Text Box Contents
	// element textboxTightWrap { w_CT_TextboxTightWrap }?,
	TextboxTightWrap *GenSingleStrVal[stypes.TextboxTightWrap] `xml:"w:textboxTightWrap,omitempty"`

	// 31. Associated Outline Level
	// element outlineLvl { w_CT_DecimalNumber }?,
	OutlineLvl *DecimalNum `xml:"w:outlineLvl,omitempty"`

	// 32. Associated HTML div ID
	// element divId { w_CT_DecimalNumber }?,
	DivID *DecimalNum `xml:"w:divId,omitempty"`

	// 33. Paragraph Conditional Formatting
	// element cnfStyle { w_CT_Cnf }?
	CnfStyle *CTString `xml:"w:cnfStyle,omitempty"`

	// 34. Run Properties for the Paragraph Mark
	// element rPr { w_CT_ParaRPr }?,
	RunProperty *RunProperty `xml:"w:rPr,omitempty"`

	// 35. Section Properties
	// element sectPr { w_CT_SectPr }?,
	SectPr *SectionProp `xml:"w:sectPr,omitempty"`

	// 36. Revision Information for Paragraph Properties
	// element pPrChange { w_CT_PPrChange }?
	PPrChange *PPrChange `xml:"w:pPrChange,omitempty"`
}

// NewParagraphStyle creates a new ParagraphStyle.
func NewParagraphStyle(val string) *CTString {
	return &CTString{Val: val}
}

// DefaultParagraphStyle creates the default ParagraphStyle with the value "Normal".
func DefaultParagraphStyle() *CTString {
	return &CTString{Val: "Normal"}
}

func DefaultParaProperty() *ParagraphProp {
	return &ParagraphProp{}
}

// <== ParaProp ends here ==>

// Revision Information for Paragraph Properties
type PPrChange struct {
	ID       int            `xml:"id,attr"`
	Author   string         `xml:"author,attr"`
	Date     *string        `xml:"date,attr,omitempty"`
	ParaProp *ParagraphProp `xml:"w:pPr"`
}

// <== PPrChange ends here ==>
