package ctypes

import (
	"fmt"
	"strings"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/common/units"
	"github.com/samuel-jimenez/whatsupdocx/dml"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlpic"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// CT_P
// w_CT_P =
type Paragraph struct {
	id string

	// Attributes
	// attribute w:rsidRPr { w_ST_LongHexNumber }?,
	RsidRPr *stypes.LongHexNum `xml:"w:rsidRPr,attr,omitempty"` // Revision Identifier for Paragraph Glyph Formatting
	// attribute w:rsidR { w_ST_LongHexNumber }?,
	RsidR *stypes.LongHexNum `xml:"w:rsidR,attr,omitempty"` // Revision Identifier for Paragraph
	// attribute w:rsidDel { w_ST_LongHexNumber }?,
	RsidDel *stypes.LongHexNum `xml:"w:rsidDel,attr,omitempty"` // Revision Identifier for Paragraph Deletion
	// attribute w:rsidP { w_ST_LongHexNumber }?,
	RsidP *stypes.LongHexNum `xml:"w:rsidP,attr,omitempty"` // Revision Identifier for Paragraph Properties
	// attribute w:rsidRDefault { w_ST_LongHexNumber }?,
	RsidRDefault *stypes.LongHexNum `xml:"w:rsidRDefault,attr,omitempty"` // Default Revision Identifier for Runs

	// 1. Paragraph Properties
	// element pPr { w_CT_PPr }?,
	Property *ParagraphProp `xml:"w:pPr,omitempty"`

	// 2. Choices (Slice of Child elements)
	// w_EG_PContent*
	Children []ParagraphChild `xml:",group,any,omitempty"`
}

type ParagraphChild struct {
	Link *Hyperlink `xml:"w:hyperlink,omitempty"` // w:hyperlink
	Run  *Run       `xml:"w:r,omitempty"`
}

func (p Paragraph) String() string {
	var builder strings.Builder
	for _, cElem := range p.Children {
		if run := cElem.Run; run != nil {
			builder.WriteString(run.String())
		}
	}
	return builder.String()
}

// paraOption defines a type for functions that can configure a Paragraph.
type paraOption func(*Paragraph)

// NewParagraph creates and initializes a new Paragraph instance with given options.
func NewParagraph(opts ...paraOption) *Paragraph {
	p := &Paragraph{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// paraWithText is an option for adding text to a Paragraph.
func paraWithText(text string) paraOption {
	return func(p *Paragraph) {
		p.AddText(text)
	}
}

func (p *Paragraph) AddText(text string) *Run {

	run := p.AddRun()
	run.AddText(text)

	return run
}

func AddParagraph(text string) *Paragraph {
	p := Paragraph{}
	p.AddText(text)
	return &p
}

func (p *Paragraph) AddRun() *Run {

	run := &Run{}

	p.Children = append(p.Children, ParagraphChild{Run: run})

	return run
}

func (p *Paragraph) ensureProp() {
	if p.Property == nil {
		p.Property = DefaultParaProperty()
	}
}

/*
@param before: Spacing Above Paragraph in twips
@param after: Spacing Below Paragraph in twips
*/
func (p *Paragraph) Spacing(before uint64, after uint64) {
	p.ensureProp()
	p.Property.Spacing = NewParagraphSpacing(before, after)
}

// Style sets the paragraph style.
//
// Parameters:
//   - value: A string representing the style value. It can be any valid style defined in the WordprocessingML specification.
//
// Example:
//
//	p1 := document.AddParagraph("Example para")
//	paragraph.Style("List Number")
func (p *Paragraph) Style(value string) {
	p.ensureProp()
	p.Property.Style = NewParagraphStyle(value)
}

// Justification sets the paragraph justification type.
//
// Parameters:
//   - value: A value of type stypes.Justification representing the justification type.
//     It can be one of the Justification type values defined in the stypes package.
//
// Example:
//
//	p1 := document.AddParagraph("Example justified para")
//	p1.Justification(stypes.JustificationCenter) // Center justification
func (p *Paragraph) Justification(value stypes.Justification) {
	p.ensureProp()

	p.Property.Justification = NewGenSingleStrVal(value)
}

// Numbering sets the paragraph numbering properties.
//
// This function assigns a numbering definition ID and a level to the paragraph,
// which affects how numbering is displayed in the document.
//
// Parameters:
//   - id: An integer representing the numbering definition ID.
//   - level: An integer representing the level within the numbering definition.
//
// Example:
//
//	p1 := document.AddParagraph("Example numbered para")
//	p1.Numbering(1, 0)
//
// In this example, the paragraph p1 is assigned the numbering properties
// defined by numbering definition ID 1 and level 0.
func (p *Paragraph) Numbering(id int, level int) {

	p.ensureProp()

	if p.Property.NumProp == nil {
		p.Property.NumProp = &NumProp{}
	}

	p.Property.NumProp.NumID = NewDecimalNum(id)
	p.Property.NumProp.ILvl = NewDecimalNum(level)
}

// Indent sets the paragraph indentation properties.
//
// This function assigns an indent definition to the paragraph,
// which affects how exactly the paragraph is going to be indented.
//
// Parameters:
//   - indentProp: A Indent instance pointer representing exact indentation
//     measurements to use.
//
// Example:
//
//	var size360 int = 360
//	var sizeu420 uint64 = 420
//	indent360 := Indent{Left: &size360, Hanging: &sizeu420}
//
//	p1 := document.AddParagraph("Example indented para")
//	p1.Indent(&indent360)
func (p *Paragraph) Indent(indentProp *Indent) {

	p.ensureProp()

	p.Property.Indent = indentProp
}

// AddDrawing adds a new drawing (image) to the Paragraph.
//
// Parameters:
//   - rID: The relationship ID of the image in the document.
//   - imgCount: The count of images in the document.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - *dml.Inline: The created Inline instance representing the added drawing.
func (p *Paragraph) AddDrawing(rID string, imgCount uint, width units.Inch, height units.Inch) *dml.Inline {
	eWidth := width.ToEmu()
	eHeight := height.ToEmu()

	inline := dml.NewInline(
		*dmlct.NewPostvSz2D(eWidth, eHeight),
		dml.DocProp{
			ID:   uint64(imgCount),
			Name: fmt.Sprintf("Image%d", imgCount),
		},
		*dml.NewPicGraphic(dmlpic.NewPic(rID, imgCount, eWidth, eHeight)),
	)

	runChildren := []RunChild{}
	drawing := &dml.Drawing{}

	drawing.Inline = append(drawing.Inline, inline)

	runChildren = append(runChildren, RunChild{
		Drawing: drawing,
	})

	run := &Run{
		Children: runChildren,
	}

	p.Children = append(p.Children, ParagraphChild{Run: run})

	return &inline
}

func (para *Paragraph) AddLink(rId, text string) *Hyperlink {

	runChildren := []RunChild{}
	runChildren = append(runChildren, RunChild{
		Text: TextFromString(text),
	})
	run := &Run{
		Children: runChildren,
		Property: &RunProperty{
			Style: &CTString{
				Val: constants.HyperLinkStyle,
			},
		},
	}

	hyperLink := &Hyperlink{
		ID:  rId,
		Run: run,
	}

	para.Children = append(para.Children, ParagraphChild{Link: hyperLink})

	return hyperLink
}
