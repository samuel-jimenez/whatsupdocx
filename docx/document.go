package docx

import (
	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// w_CT_DocumentBase = element background { w_CT_Background }?
// w_CT_Document =
// w_CT_DocumentBase,
// element body { w_CT_Body }?,
// attribute w:conformance { s_ST_ConformanceClass }?
// w_CT_GlossaryDocument =
// w_CT_DocumentBase,
// element docParts { w_CT_DocParts }?
// w_document = element document { w_CT_Document }

// This element specifies the contents of a main document part in a WordprocessingML document.
// w_CT_Document =
type Document struct {
	XMLName xml.Name   `xml:"w:document"`
	Attr    []xml.Attr `xml:",any,attr,omitempty"`

	// Reference to the RootDoc
	Root *RootDoc `xml:"-"`

	// attribute w:conformance { s_ST_ConformanceClass }?
	//TODO

	// Elements
	// w_CT_DocumentBase,
	// w_CT_DocumentBase = element background { w_CT_Background }?
	Background *Background `xml:"w:background,omitempty"`
	// element body { w_CT_Body }?,
	Body *Body `xml:"w:body,omitempty"`

	DocRels      Relationships `xml:"-"` // DocRels represents relationships specific to the document.
	RID          int           `xml:"-"`
	relativePath string        `xml:"-"`
}

// IncRelationID increments the relation ID of the document and returns the new ID.
// This method is used to generate unique IDs for relationships within the document.
func (doc *Document) IncRelationID() int {
	doc.RID += 1
	return doc.RID
}

// AddPageBreak adds a page break to the document by inserting a paragraph containing only a page break.
//
// Returns:
//   - *Paragraph: A pointer to the newly created Paragraph object containing the page break.
//
// Example:
//
//	document := godocx.NewDocument()
//	para := document.AddPageBreak()
func (rd *RootDoc) AddPageBreak() *Paragraph {
	p := rd.AddEmptyParagraph()
	p.AddRun().AddBreak(internal.ToPtr(stypes.BreakTypePage))

	return p
}
