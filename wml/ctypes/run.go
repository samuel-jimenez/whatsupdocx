package ctypes

import (
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
	Children []RunChild `xml:",omitempty"`
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

// 	func (group RunChild) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
//
// 	if group.NoFillProperties != nil {
// 		propsElement := xml.StartElement{Name: xml.Name{Local: "a:noFill"}}
// 		if err = e.EncodeElement(group.NoFillProperties, propsElement); err != nil {
// 			return err
// 		}
// 	}
// 	if group.SolidColorFillProperties != nil {
// 		propsElement := xml.StartElement{Name: xml.Name{Local: "a:solidFill"}}
// 		if err = e.EncodeElement(group.SolidColorFillProperties, propsElement); err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
//
// }

func NewRun() *Run {
	return &Run{}
}

func (r Run) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {

	if r.RsidRPr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"}, Value: string(*r.RsidRPr)})
	}
	if r.RsidR != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"}, Value: string(*r.RsidR)})
	}
	if r.RsidDel != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"}, Value: string(*r.RsidDel)})
	}

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Property
	if r.Property != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		if err = e.EncodeElement(r.Property, propsElement); err != nil {
			return err
		}
	}

	// 2. Remaining Child elemens
	if err = r.MarshalChild(e); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
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

// Sym represents a symbol character in a document.
type Sym struct {
	Font *string `xml:"w:font,attr,omitempty"`
	Char *string `xml:"w:char,attr,omitempty"`
}

func NewSym(font, char string) *Sym {
	return &Sym{
		Font: &font,
		Char: &char,
	}
}

func (r *Run) MarshalChild(e *xml.Encoder) error {
	var err error
	for _, child := range r.Children {
		switch {
		case child.Break != nil:
			err = child.Break.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:br"}})
		case child.Text != nil:
			err = child.Text.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:t"}})
		case child.DelText != nil:
			err = child.DelText.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:delText"}})
		case child.InstrText != nil:
			err = child.InstrText.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:instrText"}})
		case child.DelInstrText != nil:
			err = child.DelInstrText.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:delInstrText"}})
		case child.NoBreakHyphen != nil:
			err = child.NoBreakHyphen.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:noBreakHyphen"}})
		case child.SoftHyphen != nil:
			err = child.SoftHyphen.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:softHyphen"}})
		case child.DayShort != nil:
			err = child.DayShort.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:dayShort"}})
		case child.MonthShort != nil:
			err = child.MonthShort.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:monthShort"}})
		case child.YearShort != nil:
			err = child.YearShort.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:yearShort"}})
		case child.DayLong != nil:
			err = child.DayLong.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:dayLong"}})
		case child.MonthLong != nil:
			err = child.MonthLong.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:monthLong"}})
		case child.YearLong != nil:
			err = child.YearLong.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:yearLong"}})
		case child.AnnotationRef != nil:
			err = child.AnnotationRef.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:annotationRef"}})
		case child.FootnoteRef != nil:
			err = child.FootnoteRef.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:footnoteRef"}})
		case child.EndnoteRef != nil:
			err = child.EndnoteRef.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:endnoteRef"}})
		case child.Separator != nil:
			err = child.Separator.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:separator"}})
		case child.ContSeparator != nil:
			err = child.ContSeparator.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:continuationSeparator"}})
		case child.Sym != nil:
			propsElement := xml.StartElement{Name: xml.Name{Local: "w:sym"}}
			err = e.EncodeElement(child.Sym, propsElement)
		case child.PgNumBlock != nil:
			err = child.PgNumBlock.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:pgNum"}})
		case child.CarrRtn != nil:
			err = child.CarrRtn.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:cr"}})
		case child.Tab != nil:
			err = child.Tab.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tab"}})
		case child.Drawing != nil:
			propsElement := xml.StartElement{Name: xml.Name{Local: "w:drawing"}}
			err = e.EncodeElement(child.Drawing, propsElement)
		case child.LastRenPgBrk != nil:
			err = child.LastRenPgBrk.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:lastRenderedPageBreak"}})
		case child.PTab != nil:
			err = child.PTab.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:ptab"}})
		case child.CmntRef != nil:
			err = child.CmntRef.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:commentReference"}})

		}

		if err != nil {
			return err
		}
	}
	return nil
}
