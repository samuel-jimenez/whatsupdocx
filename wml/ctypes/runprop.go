package ctypes

import (
	"fmt"

	"github.com/samuel-jimenez/xml"

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
	Style *CTString `xml:"rStyle,omitempty"`

	//2. Run Fonts
	// element rFonts { w_CT_Fonts }?&
	Fonts *RunFonts `xml:"rFonts,omitempty"`

	//3. Bold
	// element b { w_CT_OnOff }?&
	Bold *OnOff `xml:"b,omitempty"`

	//4.Complex Script Bold
	// element bCs { w_CT_OnOff }?&
	BoldCS *OnOff `xml:"bCs,omitempty"`

	// 5.Italics
	// element i { w_CT_OnOff }?&
	Italic *OnOff `xml:"i,omitempty"`

	//6.Complex Script Italics
	// element iCs { w_CT_OnOff }?&
	ItalicCS *OnOff `xml:"iCs,omitempty"`

	//7.Display All Characters As Capital Letters
	// element caps { w_CT_OnOff }?&
	Caps *OnOff `xml:"caps,omitempty"`

	//8.Small Caps
	// element smallCaps { w_CT_OnOff }?&
	SmallCaps *OnOff `xml:"smallCaps,omitempty"`

	//9.Single Strikethrough
	// element strike { w_CT_OnOff }?&
	Strike *OnOff `xml:"strike,omitempty"`

	//10.Double Strikethrough
	// element dstrike { w_CT_OnOff }?&
	DoubleStrike *OnOff `xml:"dstrike,omitempty"`

	//11.Display Character Outline
	// element outline { w_CT_OnOff }?&
	Outline *OnOff `xml:"outline,omitempty"`

	//12.Shadow
	// element shadow { w_CT_OnOff }?&
	Shadow *OnOff `xml:"shadow,omitempty"`

	//13.Embossing
	// element emboss { w_CT_OnOff }?&
	Emboss *OnOff `xml:"emboss,omitempty"`

	//14.Imprinting
	// element imprint { w_CT_OnOff }?&
	Imprint *OnOff `xml:"imprint,omitempty"`

	//15.Do Not Check Spelling or Grammar
	// element noProof { w_CT_OnOff }?&
	NoGrammar *OnOff `xml:"noProof,omitempty"`

	//16.Use Document Grid Settings For Inter-Character Spacing
	// element snapToGrid { w_CT_OnOff }?&
	SnapToGrid *OnOff `xml:"snapToGrid,omitempty"`

	//17.Hidden Text
	// element vanish { w_CT_OnOff }?&
	Vanish *OnOff `xml:"vanish,omitempty"`

	//18.Web Hidden Text
	// element webHidden { w_CT_OnOff }?&
	WebHidden *OnOff `xml:"webHidden,omitempty"`

	//19.Run Content Color
	// element color { w_CT_Color }?&
	Color *Color `xml:"color,omitempty"`

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
	Position *DecimalNum `xml:"position,omitempty"`

	//24.Font Size
	// element sz { w_CT_HpsMeasure }?&
	Size *FontSize `xml:"sz,omitempty"`

	//25.Complex Script Font Size
	// element szCs { w_CT_HpsMeasure }?&
	SizeCs *FontSizeCS `xml:"szCs,omitempty"`

	//26.Text Highlighting
	// element highlight { w_CT_Highlight }?&
	Highlight *CTString `xml:"highlight,omitempty"`

	//27.Underline
	// element u { w_CT_Underline }?&
	Underline *GenSingleStrVal[stypes.Underline] `xml:"u,omitempty"`

	//28.Animated Text Effect
	// element effect { w_CT_TextEffect }?&
	Effect *Effect `xml:"effect,omitempty"`

	//29.Text Border
	// element bdr { w_CT_Border }?&
	Border *Border `xml:"w:bdr,omitempty"`

	//30.Run Shading
	// element shd { w_CT_Shd }?&
	Shading *Shading `xml:"shd,omitempty"`

	//31.Manual Run Width
	// element fitText { w_CT_FitText }?&
	FitText *FitText `xml:"fitText,omitempty"`

	//32.Subscript/Superscript Text
	// element vertAlign { w_CT_VerticalAlignRun }?&
	VertAlign *GenSingleStrVal[stypes.VerticalAlignRun] `xml:"vertAlign,omitempty"`

	//33.Right To Left Text
	// element rtl { w_CT_OnOff }?&
	RightToLeft *OnOff `xml:"rtl,omitempty"`

	//34.Use Complex Script Formatting on Run
	// element cs { w_CT_OnOff }?&
	CSFormat *OnOff `xml:"cs,omitempty"`

	//35.Emphasis Mark
	// element em { w_CT_Em }?&
	Em *GenSingleStrVal[stypes.Em] `xml:"em,omitempty"`

	//36.Languages for Run Content
	// element lang { w_CT_Language }?&
	Lang *Lang `xml:"lang,omitempty"`

	//37.East Asian Typography Settings
	// element eastAsianLayout { w_CT_EastAsianLayout }?&
	EALayout *EALayout `xml:"eastAsianLayout,omitempty"`

	//38.Paragraph Mark Is Always Hidden
	// element specVanish { w_CT_OnOff }?&
	SpecVanish *OnOff `xml:"specVanish,omitempty"`

	//39.Office Open XML Math
	// element oMath { w_CT_OnOff }?
	OMath *OnOff `xml:"oMath,omitempty"`

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

// MarshalXML marshals RunProperty to XML.
func (rp RunProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPr"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Referenced Character Style
	if rp.Style != nil {
		if err = rp.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rStyle"},
		}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	//2.Run Fonts
	if rp.Fonts != nil {
		if err = rp.Fonts.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Fonts: %w", err)
		}
	}

	set1 := []optBoolElems{
		{rp.Bold, "w:b"},                //3.Bold
		{rp.BoldCS, "w:bCs"},            //4.Complex Script Bold
		{rp.Italic, "w:i"},              //5.Italics
		{rp.ItalicCS, "w:iCs"},          //6.Complex Script Italics
		{rp.Caps, "w:caps"},             //7.Display All Characters As Capital Letters
		{rp.SmallCaps, "w:smallCaps"},   //8.Small Caps
		{rp.Strike, "w:strike"},         //9.Single Strikethrough
		{rp.DoubleStrike, "w:dstrike"},  //10.Double Strikethrough
		{rp.Outline, "w:outline"},       //11.Display Character Outline
		{rp.Shadow, "w:shadow"},         //12.Shadow
		{rp.Emboss, "w:emboss"},         //13.Embossing
		{rp.Imprint, "w:imprint"},       //14.Imprinting
		{rp.NoGrammar, "w:noProof"},     //15.Do Not Check Spelling or Grammar
		{rp.SnapToGrid, "w:snapToGrid"}, //16.Use Document Grid Settings For Inter-Character Spacing
		{rp.Vanish, "w:vanish"},         //17.Hidden Text
		{rp.WebHidden, "w:webHidden"},   //18.Web Hidden Text
	}

	for _, entry := range set1 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", entry.XMLName, err)
		}
	}

	//19.Run Content Color
	if rp.Color != nil {
		if err = rp.Color.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("color: %w", err)
		}
	}

	//20. Character Spacing Adjustment
	if rp.Spacing != nil {
		if err = rp.Spacing.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:spacing"},
		}); err != nil {
			return fmt.Errorf("spacing: %w", err)
		}
	}

	//21.Expanded/Compressed Text
	if rp.ExpaComp != nil {
		if err = rp.ExpaComp.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:w"},
		}); err != nil {
			return fmt.Errorf("expand/compression text: %w", err)
		}
	}

	//22.Font Kerning
	if rp.Kern != nil {
		if err = rp.Kern.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:kern"},
		}); err != nil {
			return fmt.Errorf("kern: %w", err)
		}
	}

	//23. Vertically Raised or Lowered Text
	if rp.Position != nil {
		if err = rp.Position.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:position"},
		}); err != nil {
			return fmt.Errorf("position: %w", err)
		}
	}

	//24.Font Size
	if rp.Size != nil {
		if err = rp.Size.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("size: %w", err)
		}
	}

	//25.Complex Script Font Size
	if rp.SizeCs != nil {
		if err = rp.SizeCs.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("size complex script: %w", err)
		}
	}

	//26.Text Highlighting
	if rp.Highlight != nil {
		if err = rp.Highlight.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:highlight"},
		}); err != nil {
			return fmt.Errorf("highlight: %w", err)
		}
	}

	//27.Underline
	if rp.Underline != nil {
		if err = rp.Underline.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:u"},
		}); err != nil {
			return fmt.Errorf("underline: %w", err)
		}
	}

	//28.Animated Text Effect
	if rp.Effect != nil {
		if err = rp.Effect.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:effect"},
		}); err != nil {
			return fmt.Errorf("effect: %w", err)
		}
	}

	//29.Text Border
	if rp.Border != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:bdr"}}
		if err := e.EncodeElement(rp.Border, propsElement); err != nil {
			// if err := rp.Border.MarshalXML(e, propsElement); err != nil {
			return fmt.Errorf("border: %w", err)
		}
	}
	//30.Run Shading
	if rp.Shading != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:shd"}}
		if err := e.EncodeElement(rp.Shading, propsElement); err != nil {
			// if err := rp.Shading.MarshalXML(e, propsElement); err != nil {
			return fmt.Errorf("shading: %w", err)
		}
	}

	//31.Manual Run Width
	if rp.FitText != nil {
		if err = rp.FitText.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("fit text: %w", err)
		}
	}

	//32.Subscript/Superscript Text
	if rp.VertAlign != nil {
		if err = rp.VertAlign.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:vertAlign"},
		}); err != nil {
			return fmt.Errorf("vertical align: %w", err)
		}
	}

	//33.Right To Left Text
	if rp.RightToLeft != nil {
		if err = rp.RightToLeft.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rtl"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "RightToLeft", err)
		}
	}

	//34.Use Complex Script Formatting on Run
	if rp.CSFormat != nil {
		if err = rp.CSFormat.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cs"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "CSFormat", err)
		}
	}

	//35.Emphasis Mark
	if rp.Em != nil {
		if err = rp.Em.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:em"},
		}); err != nil {
			return fmt.Errorf("emphasis mark: %w", err)
		}
	}

	//36.Languages for Run Content
	if rp.Lang != nil {
		if err = rp.Lang.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("languages for Run Content: %w", err)
		}
	}

	//37.East Asian Typography Settings
	if rp.EALayout != nil {
		if err = rp.EALayout.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("East Asian Typography Settings: %w", err)
		}
	}

	//38.Paragraph Mark Is Always Hidden
	if rp.SpecVanish != nil {
		if err = rp.SpecVanish.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:specVanish"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "specVanish", err)
		}
	}

	//39.Office Open XML Math
	if rp.OMath != nil {
		if err = rp.OMath.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:oMath"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "oMath", err)
		}
	}

	return e.EncodeToken(start.End())
}
