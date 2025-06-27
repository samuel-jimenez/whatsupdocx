package docx

import (
	"log"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/ctypes"
)

// This element specifies the contents of the body of the document – the main document editing surface.
// w_CT_Body =
type Body struct {
	root    *RootDoc `xml:"-"`
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`

	// w_EG_BlockLevelElts*,
	Children []DocumentChild `xml:",group,any,omitempty"`
	// element sectPr { w_CT_SectPr }?
	SectPr *ctypes.SectionProp `xml:"w:sectPr,omitempty"`
}

// DocumentChild represents a child element within a Word document, which can be a Paragraph or a Table.
// w_EG_BlockLevelElts = w_EG_BlockLevelChunkElts*
// | element altChunk { w_CT_AltChunk }*

// w_CT_AltChunk =
// r_id?,
// element altChunkPr { w_CT_AltChunkPr }?
// w_CT_AltChunkPr = element matchSrc { w_CT_OnOff }?

// w_EG_BlockLevelChunkElts = w_EG_ContentBlockContent*
// w_EG_ContentBlockContent =
// element customXml { w_CT_CustomXmlBlock }
// | element sdt { w_CT_SdtBlock }
// | element p { w_CT_P }*
// | element tbl { w_CT_Tbl }*
// | w_EG_RunLevelElts*

// w_EG_RunLevelElts =
// element proofErr { w_CT_ProofErr }?
// | element permStart { w_CT_PermStart }?
// | element permEnd { w_CT_Perm }?
// | w_EG_RangeMarkupElements*
// | element ins { w_CT_RunTrackChange }?
// | element del { w_CT_RunTrackChange }?
// | element moveFrom { w_CT_RunTrackChange }
// | element moveTo { w_CT_RunTrackChange }
// | w_EG_MathContent*

// w_EG_ContentBlockContent =
type DocumentChild struct {

	// element customXml { w_CT_CustomXmlBlock }
	// | element sdt { w_CT_SdtBlock }

	// | element p { w_CT_P }*
	Para *Paragraph `xml:"w:p,omitempty"`
	// | element tbl { w_CT_Tbl }*
	Table *Table `xml:"w:tbl,omitempty"`
	// | w_EG_RunLevelElts*

	// w_EG_RunLevelElts =
	// element proofErr { w_CT_ProofErr }?
	// | element permStart { w_CT_PermStart }?
	// | element permEnd { w_CT_Perm }?
	// | w_EG_RangeMarkupElements*
	// | element ins { w_CT_RunTrackChange }?
	// | element del { w_CT_RunTrackChange }?
	// | element moveFrom { w_CT_RunTrackChange }
	// | element moveTo { w_CT_RunTrackChange }

}

// Use this function to initialize a new Body before adding content to it.
func NewBody(root *RootDoc) *Body {
	return &Body{
		root: root,
	}
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Body type.
// It decodes the XML representation of the Body.
func (body *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "p":
				para := newParagraph(body.root)
				if err := para.unmarshalXML(d, elem); err != nil {
					return err
				}
				log.Println("Body UnmarshalXML", para)
				body.Children = append(body.Children, DocumentChild{Para: para})
			case "tbl":
				tbl := NewTable(body.root)
				if err := tbl.unmarshalXML(d, elem); err != nil {
					return err
				}
				body.Children = append(body.Children, DocumentChild{Table: tbl})
			case "sectPr":
				body.SectPr = ctypes.NewSectionProper()
				if err := d.DecodeElement(body.SectPr, &elem); err != nil {
					return err
				}
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}
