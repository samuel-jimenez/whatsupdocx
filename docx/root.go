package docx

import (
	"sync"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/ctypes"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
)

// RootDoc represents the root document of an Office Open XML (OOXML) document.
// It contains information about the document path, file map, the document structure,
// and relationships with other parts of the document.
type RootDoc struct {
	Path        string        // Path represents the path of the document.
	FileMap     sync.Map      // FileMap is a synchronized map for managing files related to the document.
	RootRels    Relationships // RootRels represents relationships at the root level.
	ContentType ContentTypes
	Document    *Document      `xml:"w:document"` // Document is the main document structure.
	DocStyles   *ctypes.Styles `xml:"w:styles"`   // Document styles

	rID        int // rId is used to generate unique relationship IDs.
	ImageCount uint
}

// NewRootDoc creates a new instance of the RootDoc structure.
func NewRootDoc() *RootDoc {
	return &RootDoc{}
}

// LoadDocXml decodes the provided XML data and returns a Document instance.
// It is used to load the main document structure from the document file.
//
// Parameters:
//   - fileName: The name of the document file.
//   - fileBytes: The XML data representing the main document structure.
//
// Returns:
//   - doc: The Document instance containing the decoded main document structure.
//   - err: An error, if any occurred during the decoding process.
func LoadDocXml(rd *RootDoc, fileName string, fileBytes []byte) (*Document, error) {
	// constants.defaultDocNamespace
	doc := Document{
		Root: rd,
	}
	err := xml.Unmarshal(fileBytes, &doc)
	if err != nil {
		return nil, err
	}

	doc.relativePath = fileName
	if len(doc.Attr) == 0 {
		doc.Attr = constants.DefaultNamespacesDoc
	}

	return &doc, nil
}

// Load styles.xml into Styles struct
func LoadStyles(fileName string, fileBytes []byte) (*ctypes.Styles, error) {
	styles := ctypes.Styles{}
	err := xml.Unmarshal(fileBytes, &styles)
	if err != nil {
		return nil, err
	}

	styles.RelativePath = fileName
	if len(styles.Attr) == 0 {
		styles.Attr = constants.DefaultNamespacesStyle
	}

	return &styles, nil
}
