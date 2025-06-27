package ctypes

import (
	"log"
	"strings"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
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

type Hyperlink struct {
	XMLName  xml.Name         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID       string           `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	Run      *Run             `xml:"w:r,omitempty"`
	Children []ParagraphChild `xml:",group,any,omitempty"`
}

func (p Paragraph) String() string {
	var builder strings.Builder
	for _, cElem := range p.Children {
		if cElem.Run != nil {
			for _, child := range cElem.Run.Children {
				switch {
				case child.Text != nil:
					t := child.Text
					builder.WriteString(t.Text)
				}
			}
		}
	}
	return builder.String()
}

// func (p Paragraph) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
// 	log.Println("Paragraph MarshalXML")
//
// 	if p.RsidRPr != nil {
// 		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"}, Value: string(*p.RsidRPr)})
// 	}
//
// 	if p.RsidR != nil {
// 		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"}, Value: string(*p.RsidR)})
// 	}
// 	if p.RsidDel != nil {
// 		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"}, Value: string(*p.RsidDel)})
// 	}
// 	if p.RsidP != nil {
// 		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidP"}, Value: string(*p.RsidP)})
// 	}
// 	if p.RsidRDefault != nil {
// 		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRDefault"}, Value: string(*p.RsidRDefault)})
// 	}
//
// 	if err = e.EncodeToken(start); err != nil {
// 		return err
// 	}
//
// 	if p.Property != nil {
// 		if err = e.EncodeElement(p.Property, xml.StartElement{
// 			// if err = p.Property.MarshalXML(e, xml.StartElement{
// 			Name: xml.Name{Local: "w:pPr"},
// 		}); err != nil {
// 			return err
// 		}
// 	}
//
// 	for _, cElem := range p.Children {
// 		if cElem.Run != nil {
// 			propsElement := xml.StartElement{Name: xml.Name{Local: "w:r"}}
// 			if err := e.EncodeElement(cElem.Run, propsElement); err != nil {
// 				// if err := cElem.Run.MarshalXML(e, propsElement); err != nil {
// 				return err
// 			}
// 		}
//
// 		if cElem.Link != nil {
// 			propsElement := xml.StartElement{Name: xml.Name{Local: "w:hyperlink"}}
// 			if err := e.EncodeElement(cElem.Link, propsElement); err != nil {
// 				return err
// 			}
// 		}
// 	}
//
// 	// Closing </w:p> element
// 	return e.EncodeToken(start.End())
// }

func (p *Paragraph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	// Decode attributes
	log.Println("Paragraph UnmarshalXML")

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "rsidRPr":
			p.RsidRPr = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidR":
			p.RsidR = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidDel":
			p.RsidDel = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidP":
			p.RsidP = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidRDefault":
			p.RsidRDefault = internal.ToPtr(stypes.LongHexNum(attr.Value))
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
			case "r":
				r := NewRun()
				if err = d.DecodeElement(r, &elem); err != nil {
					return err
				}

				p.Children = append(p.Children, ParagraphChild{Run: r})
			case "pPr":
				p.Property = &ParagraphProp{}
				if err = d.DecodeElement(p.Property, &elem); err != nil {
					return err
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

func (p *Paragraph) AddText(text string) *Run {
	t := TextFromString(text)

	runChildren := []RunChild{}
	runChildren = append(runChildren, RunChild{
		Text: t,
	})
	run := &Run{
		Children: runChildren,
	}

	p.Children = append(p.Children, ParagraphChild{Run: run})

	return run
}

func AddParagraph(text string) *Paragraph {
	p := Paragraph{}
	p.AddText(text)
	return &p
}
