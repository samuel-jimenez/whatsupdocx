package ctypes

import (
	"strings"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common"
	"github.com/samuel-jimenez/whatsupdocx/dml"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/mc"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// A Run is part of a paragraph that has its own style. It could be (CT_R)
// w_CT_R =
type Run struct {
	// Attributes
	// attribute w:rsidRPr { w_ST_LongHexNumber }?,
	RsidRPr *stypes.LongHexNum `xml:"w:rsidRPr,attr,omitempty"` // Revision Identifier for Run Properties
	// attribute w:rsidDel { w_ST_LongHexNumber }?,
	RsidDel *stypes.LongHexNum `xml:"w:rsidDel,attr,omitempty"` // Revision Identifier for Run Deletion
	// attribute w:rsidR { w_ST_LongHexNumber }?,
	RsidR *stypes.LongHexNum `xml:"w:rsidR,attr,omitempty"` // Revision Identifier for Run

	// Sequence:

	//1. Run Properties
	// w_EG_RPr?,
	// w_EG_RPr = element rPr { w_CT_RPr }
	Property *RunProperty `xml:"w:rPr,omitempty"`

	// 2. Choice - Run Inner content
	// w_EG_RunInnerContent*
	Children []RunChild `xml:",group,any,omitempty"`
}

// w_EG_RunInnerContent =
type RunChild struct {

	//specifies that a break shall be placed at the current location in the run content
	// element br { w_CT_Br }
	Break *Break `xml:"w:br,omitempty"`

	//specifies that this run contains literal text which shall be displayed in the document
	// | element t { w_CT_Text }
	Text *Text `xml:"w:t,omitempty"`

	// | element contentPart { w_CT_Rel }

	//specifies that this run contains literal text which shall be displayed in the document
	// | element delText { w_CT_Text }
	DelText *Text `xml:"w:delText,omitempty"`

	//Field Code
	// | element instrText { w_CT_Text }
	InstrText *Text `xml:"w:instrText,omitempty"`

	//Deleted Field Code
	// | element delInstrText { w_CT_Text }
	DelInstrText *Text `xml:"w:delInstrText,omitempty"`

	//Non Breaking Hyphen Character
	// | element noBreakHyphen { w_CT_Empty }
	NoBreakHyphen *common.Empty `xml:"w:noBreakHyphen,omitempty"`

	//Non Breaking Hyphen Character
	// | element softHyphen { w_CT_Empty }?
	SoftHyphen *common.Empty `xml:"w:softHyphen,omitempty"`

	//Date Block - Short Day Format
	// | element dayShort { w_CT_Empty }?
	DayShort *common.Empty `xml:"w:dayShort,omitempty"`

	//Date Block - Short Month Format
	// | element monthShort { w_CT_Empty }?
	MonthShort *common.Empty `xml:"w:monthShort,omitempty"`

	//Date Block - Short Year Format
	// | element yearShort { w_CT_Empty }?
	YearShort *common.Empty `xml:"w:yearShort,omitempty"`

	//Date Block - Long Day Format
	// | element dayLong { w_CT_Empty }?
	DayLong *common.Empty `xml:"w:dayLong,omitempty"`

	//Date Block - Long Month Format
	// | element monthLong { w_CT_Empty }?
	MonthLong *common.Empty `xml:"w:monthLong,omitempty"`

	//Date Block - Long Year Format
	// | element yearLong { w_CT_Empty }?
	YearLong *common.Empty `xml:"w:yearLong,omitempty"`

	//Comment Information Block
	// | element annotationRef { w_CT_Empty }?
	AnnotationRef *common.Empty `xml:"w:annotationRef,omitempty"`

	//Footnote Reference Mark
	// | element footnoteRef { w_CT_Empty }?
	FootnoteRef *common.Empty `xml:"w:footnoteRef,omitempty"`

	//Endnote Reference Mark
	// | element endnoteRef { w_CT_Empty }?
	EndnoteRef *common.Empty `xml:"w:endnoteRef,omitempty"`

	//Footnote/Endnote Separator Mark
	// | element separator { w_CT_Empty }?
	Separator *common.Empty `xml:"w:separator,omitempty"`

	//Continuation Separator Mark
	// | element continuationSeparator { w_CT_Empty }?
	ContSeparator *common.Empty `xml:"w:continuationSeparator,omitempty"`

	//Symbol Character
	// | element sym { w_CT_Sym }?
	Sym *Sym `xml:"w:sym,omitempty"`

	//Page Number Block
	// | element pgNum { w_CT_Empty }?
	PgNumBlock *common.Empty `xml:"w:pgNum,omitempty"`

	//Carriage Return
	// | element cr { w_CT_Empty }?
	CarrRtn *common.Empty `xml:"w:cr,omitempty"`

	//Tab Character
	// | element tab { w_CT_Empty }?
	Tab *common.Empty `xml:"w:tab,omitempty"`

	//TODO:
	// | element object { w_CT_Object }
	// 	w:object    Inline Embedded Object
	// | element fldChar { w_CT_FldChar }
	// w:fldChar    Complex Field Character
	// | element ruby { w_CT_Ruby }
	// w:ruby    Phonetic Guide
	// | element footnoteReference { w_CT_FtnEdnRef }
	// w:footnoteReference    Footnote Reference
	// | element endnoteReference { w_CT_FtnEdnRef }
	// w:endnoteReference    Endnote Reference

	//Comment Content Reference Mark
	// | element commentReference { w_CT_Markup }
	CmntRef *Markup `xml:"w:commentReference,omitempty"`

	//DrawingML Object
	// | element drawing { w_CT_Drawing }
	Drawing *dml.Drawing `xml:"w:drawing,omitempty"`

	//Absolute Position Tab Character
	// | element ptab { w_CT_PTab }?
	PTab *PTab `xml:"w:ptab,omitempty"`

	//Position of Last Calculated Page Break
	// | element lastRenderedPageBreak { w_CT_Empty }?
	LastRenPgBrk *common.Empty `xml:"w:lastRenderedPageBreak,omitempty"`
}

func NewRun() *Run {
	return &Run{}
}

// Sym represents a symbol character in a document.
// w_CT_Sym
// w_CT_Sym =
type Sym struct {
	// attribute w:font { s_ST_String }?,
	Font *string `xml:"w:font,attr,omitempty"`
	// attribute w:char { w_ST_ShortHexNumber }?
	Char *string `xml:"w:char,attr,omitempty"`
}

func NewSym(font, char string) *Sym {
	return &Sym{
		Font: &font,
		Char: &char,
	}
}

// String representation
func (el Run) String() string {
	var builder strings.Builder
	for _, child := range el.Children {
		switch {
		case child.Text != nil:
			t := child.Text
			builder.WriteString(t.Text)
		}
	}
	return builder.String()
}

// Appends a new text to the run.
// Example:
//
//	run.AddText("Hello, World!")
//
// Parameters:
//   - text: A string representing the text to be added to the Run.
func (run *Run) AddText(text string) {
	t := TextFromString(text)
	run.Children = append(run.Children, RunChild{
		Text: t,
	})
}

// Clears the run's children.
func (run *Run) Clear() {
	run.Children = nil
}

// getProp returns the run properties. If not initialized, it creates and returns a new instance.
func (r *Run) getProp() *RunProperty {
	if r.Property == nil {
		r.Property = &RunProperty{}
	}
	return r.Property
}

// Sets the color of the Run.
//
// Example:
//
//	modifiedRun := run.Color("FF0000")
//
// Parameters:
//   - colorCode: A string representing the color code (e.g., "FF0000" for red).
//
// Returns:
//   - *Run: The modified Run instance with the updated color.
func (r *Run) Color(colorCode string) *Run {
	r.getProp().Color = NewColor(colorCode)
	return r
}

// Sets the size of the Run.

// This method takes an integer parameter representing the desired font size.
// It updates the size property of the Run instance with the specified size,
// Example:

// 	modifiedRun := run.Size(12)

// Parameters:
//   - size: An integer representing the font size.

// Returns:
//   - *Run: The modified Run instance with the updated size.
func (r *Run) Size(size uint64) *Run {
	r.getProp().Size = NewFontSize(size * 2)
	return r
}

// Font sets the font for the run.
func (r *Run) Font(font string) *Run {
	if r.getProp().Fonts == nil {
		r.getProp().Fonts = &RunFonts{}
	}

	r.getProp().Fonts.Ascii = font
	r.getProp().Fonts.HAnsi = font
	return r
}

// Shading sets the shading properties (type, color, fill) for the run
func (r *Run) Shading(shdType stypes.Shading, color, fill string) *Run {
	r.getProp().Shading = NewShading().SetShadingType(shdType).SetColor(color).SetFill(fill)
	return r
}

// AddHighlight sets the highlight color for the run.
func (r *Run) Highlight(color string) *Run {
	r.getProp().Highlight = NewCTString(color)
	return r
}

// AddBold enables bold formatting for the run.
func (r *Run) Bold(value bool) *Run {
	r.getProp().Bold = OnOffFromBool(value)
	return r
}

// Italic enables or disables italic formatting for the run.
func (r *Run) Italic(value bool) *Run {
	r.getProp().Italic = OnOffFromBool(value)
	return r
}

// Specifies that the contents of this run shall be displayed with a single horizontal line through the center of the line.
func (r *Run) Strike(value bool) *Run {
	r.getProp().Strike = OnOffFromBool(value)
	return r
}

// Specifies that the contents of this run shall be displayed with two horizontal lines through each character displayed on the line
func (r *Run) DoubleStrike(value bool) *Run {
	r.getProp().DoubleStrike = OnOffFromBool(value)
	return r
}

// Display All Characters As Capital Letters
// Any lowercase characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) Caps(value bool) *Run {
	r.getProp().Caps = OnOffFromBool(value)
	return r
}

// Specifies that all small letter characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) SmallCaps(value bool) *Run {
	r.getProp().Caps = OnOffFromBool(value)
	return r
}

// Outline enables or disables outline formatting for the run.
func (r *Run) Outline(value bool) *Run {
	r.getProp().Outline = OnOffFromBool(value)
	return r
}

// Shadow enables or disables shadow formatting for the run.
func (r *Run) Shadow(value bool) *Run {
	r.getProp().Shadow = OnOffFromBool(value)
	return r
}

// Emboss enables or disables embossing formatting for the run.
func (r *Run) Emboss(value bool) *Run {
	r.getProp().Emboss = OnOffFromBool(value)
	return r
}

// Imprint enables or disables imprint formatting for the run.
func (r *Run) Imprint(value bool) *Run {
	r.getProp().Imprint = OnOffFromBool(value)
	return r
}

// Do Not Check Spelling or Grammar
func (r *Run) NoGrammer(value bool) *Run {
	r.getProp().NoGrammar = OnOffFromBool(value)
	return r
}

// Use Document Grid Settings For Inter-Character Spacing
func (r *Run) SnapToGrid(value bool) *Run {
	r.getProp().SnapToGrid = OnOffFromBool(value)
	return r
}

// Hidden Text
func (r *Run) HideText(value bool) *Run {
	r.getProp().Vanish = OnOffFromBool(value)
	return r
}

// Spacing sets the spacing between characters in the run.
func (r *Run) Spacing(value int) *Run {
	r.getProp().Spacing = NewDecimalNum(value)
	return r
}

// Underline sets the underline style for the run.
func (r *Run) Underline(value stypes.Underline) *Run {
	r.getProp().Underline = NewGenSingleStrVal(value)
	return r
}

// Add a break element of `stypes.BreakType` to this run.
func (r *Run) AddBreak(breakType *stypes.BreakType) {
	// clear := stypes.BreakClearNone
	// switch breakType{
	// case stypes.BreakType:

	// }
	br := Break{}

	if breakType != nil {
		br.BreakType = breakType
	}

	r.Children = append(r.Children, RunChild{
		Break: &br,
	})
}

// Style sets the style of the run.
func (r *Run) Style(value string) *Run {
	r.getProp().Style = NewRunStyle(value)
	return r
}

// VerticalAlign sets the vertical alignment for the run text.
//
// Parameter: A value from the stypes.VerticalAlignRun type indicating the desired vertical alignment. One of:
//
//	VerticalAlignRunBaseline, VerticalAlignRunSuperscript, VerticalAlignRunSubscript
//
// Returns: The modified Run instance with the updated vertical alignment.
func (r *Run) VerticalAlign(value stypes.VerticalAlignRun) *Run {
	r.getProp().VertAlign = NewGenSingleStrVal(value)
	return r
}

func (r *Run) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	// Decode attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "rsidRPr":
			r.RsidRPr = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidR":
			r.RsidR = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidDel":
			r.RsidDel = internal.ToPtr(stypes.LongHexNum(attr.Value))
		}
	}

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "t":
				txt := NewText()
				if err = d.DecodeElement(txt, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, RunChild{Text: txt})
			case "rPr":
				r.Property = &RunProperty{}
				if err = d.DecodeElement(r.Property, &elem); err != nil {
					return err
				}
			case "tab":
				tabElem := &common.Empty{}
				if err = d.DecodeElement(tabElem, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, RunChild{
					Tab: tabElem,
				})
			case "br":
				br := Break{}
				if err = d.DecodeElement(&br, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, RunChild{
					Break: &br,
				})
			case "drawing":
				drawingElem := &dml.Drawing{}
				if err = d.DecodeElement(drawingElem, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, RunChild{
					Drawing: drawingElem,
				})
				// <mc:AlternateContent>
			case "AlternateContent":
				alternateContent := &mc.AlternateContent{}

				if err = d.DecodeElement(alternateContent, &elem); err != nil {
					return err
				}

				if alternateContent.Drawing != nil {

					r.Children = append(r.Children, RunChild{
						Drawing: alternateContent.Drawing,
					})
				}

			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}
