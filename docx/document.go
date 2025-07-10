package docx

import (
	"strconv"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/ctypes"
	"github.com/samuel-jimenez/xml"
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
	// Body *Body `xml:"w:body,omitempty"`
	Body *ctypes.Body `xml:"w:body,omitempty"`

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

// addLinkRelation adds a hyperlink relationship to the document's relationships collection.
//
// Parameters:
//   - link: A string representing the target URL or location of the hyperlink.
//
// Returns:
//   - string: The ID ("rId" + relation ID) of the added relationship.
//
// This function generates a new relationship ID, creates a Relationship object with the specified link as the target,
// and appends it to the document's relationships collection (DocRels.Relationships). It returns the generated ID of the relationship.
func (doc *Document) addLinkRelation(link string) string {

	rID := doc.IncRelationID()

	rel := &Relationship{
		ID:         "rId" + strconv.Itoa(rID),
		TargetMode: "External",
		Type:       constants.SourceRelationshipHyperLink,
		Target:     link,
	}

	doc.DocRels.Relationships = append(doc.DocRels.Relationships, rel)

	return "rId" + strconv.Itoa(rID)
}

// addRelation adds a generic relationship to the document's relationships collection.
//
// Parameters:
//   - relType: A string representing the type of relationship (e.g., constants.SourceRelationshipImage).
//   - fileName: A string representing the target file name or location related to the relationship.
//
// Returns:
//   - string: The ID ("rId" + relation ID) of the added relationship.
//
// This function generates a new relationship ID, creates a Relationship object with the specified type and target,
// and appends it to the document's relationships collection (DocRels.Relationships). It returns the generated ID of the relationship.
func (doc *Document) addRelation(relType string, fileName string) string {
	rID := doc.IncRelationID()
	rel := &Relationship{
		ID:     "rId" + strconv.Itoa(rID),
		Type:   relType,
		Target: fileName,
	}

	doc.DocRels.Relationships = append(doc.DocRels.Relationships, rel)

	return "rId" + strconv.Itoa(rID)
}
