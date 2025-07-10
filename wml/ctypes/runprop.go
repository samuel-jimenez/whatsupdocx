package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// RunProperty represents the properties of a run of text within a paragraph.
// w_CT_RPr = w_EG_RPrContent?

// w_EG_RPrContent =
// w_EG_RPrBase?,
// element rPrChange { w_CT_RPrChange }?

// w_EG_RPrBase =
// element rStyle { w_CT_String }?&
// element rFonts { w_CT_Fonts }?&
// element b { w_CT_OnOff }?&
// element bCs { w_CT_OnOff }?&
// element i { w_CT_OnOff }?&
// element iCs { w_CT_OnOff }?&
// element caps { w_CT_OnOff }?&
// element smallCaps { w_CT_OnOff }?&
// element strike { w_CT_OnOff }?&
// element dstrike { w_CT_OnOff }?&
// element outline { w_CT_OnOff }?&
// element shadow { w_CT_OnOff }?&
// element emboss { w_CT_OnOff }?&
// element imprint { w_CT_OnOff }?&
// element noProof { w_CT_OnOff }?&
// element snapToGrid { w_CT_OnOff }?&
// element vanish { w_CT_OnOff }?&
// element webHidden { w_CT_OnOff }?&
// element color { w_CT_Color }?&
// element spacing { w_CT_SignedTwipsMeasure }?&
// element w { w_CT_TextScale }?&
// element kern { w_CT_HpsMeasure }?&
// element position { w_CT_SignedHpsMeasure }?&
// element sz { w_CT_HpsMeasure }?&
// element szCs { w_CT_HpsMeasure }?&
// element highlight { w_CT_Highlight }?&
// element u { w_CT_Underline }?&
// element effect { w_CT_TextEffect }?&
// element bdr { w_CT_Border }?&
// element shd { w_CT_Shd }?&
// element fitText { w_CT_FitText }?&
// element vertAlign { w_CT_VerticalAlignRun }?&
// element rtl { w_CT_OnOff }?&
// element cs { w_CT_OnOff }?&
// element em { w_CT_Em }?&
// element lang { w_CT_Language }?&
// element eastAsianLayout { w_CT_EastAsianLayout }?&
// element specVanish { w_CT_OnOff }?&
// element oMath { w_CT_OnOff }?
type RunProperty struct {
	//1. Referenced Character Style
	// element rStyle { w_CT_String }?&
	Style *CTString `xml:"w:rStyle,omitempty"`

	//2. Run Fonts
	// element rFonts { w_CT_Fonts }?&
	Fonts *RunFonts `xml:"w:rFonts,omitempty"`

	//3. Bold
	// element b { w_CT_OnOff }?&
	Bold *OnOff `xml:"w:b,omitempty"`

	//4.Complex Script Bold
	// element bCs { w_CT_OnOff }?&
	BoldCS *OnOff `xml:"w:bCs,omitempty"`

	// 5.Italics
	// element i { w_CT_OnOff }?&
	Italic *OnOff `xml:"w:i,omitempty"`

	//6.Complex Script Italics
	// element iCs { w_CT_OnOff }?&
	ItalicCS *OnOff `xml:"w:iCs,omitempty"`

	//7.Display All Characters As Capital Letters
	// element caps { w_CT_OnOff }?&
	Caps *OnOff `xml:"w:caps,omitempty"`

	//8.Small Caps
	// element smallCaps { w_CT_OnOff }?&
	SmallCaps *OnOff `xml:"w:smallCaps,omitempty"`

	//9.Single Strikethrough
	// element strike { w_CT_OnOff }?&
	Strike *OnOff `xml:"w:strike,omitempty"`

	//10.Double Strikethrough
	// element dstrike { w_CT_OnOff }?&
	DoubleStrike *OnOff `xml:"w:dstrike,omitempty"`

	//11.Display Character Outline
	// element outline { w_CT_OnOff }?&
	Outline *OnOff `xml:"w:outline,omitempty"`

	//12.Shadow
	// element shadow { w_CT_OnOff }?&
	Shadow *OnOff `xml:"w:shadow,omitempty"`

	//13.Embossing
	// element emboss { w_CT_OnOff }?&
	Emboss *OnOff `xml:"w:emboss,omitempty"`

	//14.Imprinting
	// element imprint { w_CT_OnOff }?&
	Imprint *OnOff `xml:"w:imprint,omitempty"`

	//15.Do Not Check Spelling or Grammar
	// element noProof { w_CT_OnOff }?&
	NoGrammar *OnOff `xml:"w:noProof,omitempty"`

	//16.Use Document Grid Settings For Inter-Character Spacing
	// element snapToGrid { w_CT_OnOff }?&
	SnapToGrid *OnOff `xml:"w:snapToGrid,omitempty"`

	//17.Hidden Text
	// element vanish { w_CT_OnOff }?&
	Vanish *OnOff `xml:"w:vanish,omitempty"`

	//18.Web Hidden Text
	// element webHidden { w_CT_OnOff }?&
	WebHidden *OnOff `xml:"w:webHidden,omitempty"`

	//19.Run Content Color
	// element color { w_CT_Color }?&
	Color *Color `xml:"w:color,omitempty"`

	//20. Character Spacing Adjustment
	// element spacing { w_CT_SignedTwipsMeasure }?&
	Spacing *DecimalNum `xml:"w:spacing,omitempty"`

	//21.Expanded/Compressed Text
	// element w { w_CT_TextScale }?&
	ExpaComp *ExpaComp `xml:"w:w,omitempty"`

	//22.Font Kerning
	// element kern { w_CT_HpsMeasure }?&
	Kern *Uint64Elem `xml:"w:kern,omitempty"`

	//23. Vertically Raised or Lowered Text
	// element position { w_CT_SignedHpsMeasure }?&
	Position *DecimalNum `xml:"w:position,omitempty"`

	//24.Font Size
	// element sz { w_CT_HpsMeasure }?&
	Size *FontSize `xml:"w:sz,omitempty"`

	//25.Complex Script Font Size
	// element szCs { w_CT_HpsMeasure }?&
	SizeCs *FontSize `xml:"w:szCs,omitempty"`

	//26.Text Highlighting
	// element highlight { w_CT_Highlight }?&
	Highlight *CTString `xml:"w:highlight,omitempty"`

	//27.Underline
	// element u { w_CT_Underline }?&
	Underline *GenSingleStrVal[stypes.Underline] `xml:"w:u,omitempty"`

	//28.Animated Text Effect
	// element effect { w_CT_TextEffect }?&
	Effect *Effect `xml:"w:effect,omitempty"`

	//29.Text Border
	// element bdr { w_CT_Border }?&
	Border *Border `xml:"w:bdr,omitempty"`

	//30.Run Shading
	// element shd { w_CT_Shd }?&
	Shading *Shading `xml:"w:shd,omitempty"`

	//31.Manual Run Width
	// element fitText { w_CT_FitText }?&
	FitText *FitText `xml:"w:fitText,omitempty"`

	//32.Subscript/Superscript Text
	// element vertAlign { w_CT_VerticalAlignRun }?&
	VertAlign *GenSingleStrVal[stypes.VerticalAlignRun] `xml:"w:vertAlign,omitempty"`

	//33.Right To Left Text
	// element rtl { w_CT_OnOff }?&
	RightToLeft *OnOff `xml:"w:rtl,omitempty"`

	//34.Use Complex Script Formatting on Run
	// element cs { w_CT_OnOff }?&
	CSFormat *OnOff `xml:"w:cs,omitempty"`

	//35.Emphasis Mark
	// element em { w_CT_Em }?&
	Em *GenSingleStrVal[stypes.Em] `xml:"w:em,omitempty"`

	//36.Languages for Run Content
	// element lang { w_CT_Language }?&
	Lang *Lang `xml:"w:lang,omitempty"`

	//37.East Asian Typography Settings
	// element eastAsianLayout { w_CT_EastAsianLayout }?&
	EALayout *EALayout `xml:"w:eastAsianLayout,omitempty"`

	//38.Paragraph Mark Is Always Hidden
	// element specVanish { w_CT_OnOff }?&
	SpecVanish *OnOff `xml:"w:specVanish,omitempty"`

	//39.Office Open XML Math
	// element oMath { w_CT_OnOff }?
	OMath *OnOff `xml:"w:oMath,omitempty"`

	// element rPrChange { w_CT_RPrChange }?
}

// NewRunProperty creates a new RunProperty with default values.
func NewRunProperty() RunProperty {
	return RunProperty{}
}

type optBoolElems struct {
	elem    *OnOff
	XMLName string
}
