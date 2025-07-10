package docx

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/ctypes"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/common/units"
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

// AddPageBreak adds a page break to the document by inserting a paragraph containing only a page break.
//
// Returns:
//   - *Paragraph: A pointer to the newly created Paragraph object containing the page break.
//
// Example:
//
//	document := whatsupdocx.NewDocument()
//	para := document.AddPageBreak()
func (rd *RootDoc) AddPageBreak() *ctypes.Paragraph {
	p := rd.AddEmptyParagraph()
	p.AddRun().AddBreak(internal.ToPtr(stypes.BreakTypePage))

	return p
}

// Return a heading paragraph newly added to the end of the document.
// The heading paragraph will contain text and have its paragraph style determined by level.
// If level is 0, the style is set to Title.
// The style is set to Heading {level}.
// if level is outside the range 0-9, error will be returned
func (rd *RootDoc) AddHeading(text string, level uint) (*ctypes.Paragraph, error) {
	if level < 0 || level > 9 {
		return nil, errors.New("Heading level not supported")
	}

	p := ctypes.NewParagraph()

	p.Property = ctypes.DefaultParaProperty()

	style := "Title"
	if level != 0 {
		style = fmt.Sprintf("Heading%d", level)
	}

	p.Property.Style = ctypes.NewParagraphStyle(style)

	bodyElem := ctypes.BlockLevel{
		Paragraph: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	p.AddText(text)
	return p, nil
}

// AddEmptyParagraph adds a new empty paragraph to the document.
// It returns the created Paragraph instance.
//
// Returns:
//   - p: The created Paragraph instance.
func (rd *RootDoc) AddEmptyParagraph() *ctypes.Paragraph {
	p := ctypes.NewParagraph()
	bodyElem := ctypes.BlockLevel{
		Paragraph: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	return p
}

// AddParagraph adds a new paragraph with the specified text to the document.
// It returns the created Paragraph instance.
//
// Parameters:
//   - text: The text to be added to the paragraph.
//
// Returns:
//   - p: The created Paragraph instance.
func (rd *RootDoc) AddParagraph(text string) *ctypes.Paragraph {
	p := ctypes.NewParagraph()
	p.AddText(text)
	bodyElem := ctypes.BlockLevel{
		Paragraph: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	return p
}

// AddTable adds a new table to the root document.
//
// It creates and initializes a new table, appends it to the root document's body, and returns a pointer to the created table.
// The table is initially empty, with no rows or cells. To add content to the table, use the provided methods on the returned
// table instance.
//
// Example usage:
//   document := whatsupdocx.NewDocument()
//   table := document.AddTable()
//   table.Style("LightList-Accent2")
//
//   // Add rows and cells to the table
//   row := table.AddRow()
//   cell := row.AddCell()
//   cell.AddParagraph("Hello, World!")
//
// Parameters:
//   - rd: A pointer to the RootDoc instance.
//
// Returns:
//   - *elements.Table: A pointer to the newly added table.

func (rd *RootDoc) AddTable() *ctypes.Table {
	tbl := ctypes.DefaultTable()

	rd.Document.Body.Children = append(rd.Document.Body.Children, ctypes.BlockLevel{
		Table: tbl,
	})

	return tbl
}

// AddPicture adds a new image to the document.
//
// Example usage:
//
//	// Add a picture to the document
//	_, err = document.AddPicture("gopher.png", units.Inch(2.9), units.Inch(2.9))
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//   - path: The path of the image file to be added.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - *PicMeta: Metadata about the added picture, including the Paragraph instance and Inline element.
//   - error: An error, if any occurred during the process.
func (root *RootDoc) AddPicture(path string, width units.Inch, height units.Inch) (*PicMeta, error) {

	p := ctypes.NewParagraph()

	bodyElem := ctypes.BlockLevel{
		Paragraph: p,
	}
	root.Document.Body.Children = append(root.Document.Body.Children, bodyElem)

	return root.AddPicturetoParagraph(p, path, width, height)

}

// AddPicturetoParagraph adds a new image to the document in specified paragraph.
//
// Example usage:
//
//	// Add a picture to the document
//	_, err = document.AddPicturetoParagraph(para, "gopher.png", units.Inch(2.9), units.Inch(2.9))
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//   - para: The paragraph to which the image file will be added.
//   - path: The path of the image file to be added.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - *PicMeta: Metadata about the added picture, including the Paragraph instance and Inline element.
//   - error: An error, if any occurred during the process.
func (root *RootDoc) AddPicturetoParagraph(para *ctypes.Paragraph, path string, width units.Inch, height units.Inch) (*PicMeta, error) {

	imgBytes, err := internal.FileToByte(path)
	if err != nil {
		return nil, err
	}

	imgExt := filepath.Ext(path)
	root.ImageCount += 1
	fileName := fmt.Sprintf("image%d%s", root.ImageCount, imgExt)
	fileIdxPath := fmt.Sprintf("%s%s", constants.MediaPath, fileName)

	imgExtStripDot := strings.TrimPrefix(imgExt, ".")
	imgMIME, err := MIMEFromExt(imgExtStripDot)
	if err != nil {
		return nil, err
	}

	err = root.ContentType.AddExtension(imgExtStripDot, imgMIME)
	if err != nil {
		return nil, err
	}

	overridePart := fmt.Sprintf("/%s%s", constants.MediaPath, fileName)
	err = root.ContentType.AddOverride(overridePart, imgMIME)
	if err != nil {
		return nil, err
	}

	root.FileMap.Store(fileIdxPath, imgBytes)

	relName := fmt.Sprintf("media/%s", fileName)

	rID := root.Document.addRelation(constants.SourceRelationshipImage, relName)

	inline := para.AddDrawing(rID, root.ImageCount, width, height)

	return &PicMeta{
		Paragraph: para,
		Inline:    inline,
	}, nil

}

// GetStyle retrieves the style information applied to the Paragraph.
//
// Returns:
//   - *Style: The style information of the Paragraph.
//   - error: An error if the style information is not found.
func (root *RootDoc) GetStyle(para *ctypes.Paragraph) (*ctypes.Style, error) {
	if para.Property == nil || para.Property.Style == nil {
		return nil, errors.New("No property for the style")
	}

	style := root.GetStyleByID(para.Property.Style.Val, stypes.StyleTypeParagraph)
	if style == nil {
		return nil, errors.New("No style found for the paragraph")
	}

	return style, nil
}

func (root *RootDoc) AddLink(p *ctypes.Paragraph, text string, link string) *ctypes.Hyperlink {
	rId := root.Document.addLinkRelation(link)
	return p.AddLink(rId, text)
}

// GetStyleByID retrieves a style from the document styles collection based on the given style ID and type.
//
// Parameters:
//   - styleID: A string representing the ID of the style to retrieve.
//   - styleType: An stypes.StyleType indicating the type of style (e.g., paragraph, character, table).
//
// Returns:
//   - *ctypes.Style: A pointer to the style matching the provided ID and type, if found; otherwise, nil.
//
// This method searches through the document's style list to find a style with the specified ID and type.
// If no matching style is found or if the document styles collection is nil, it returns nil.
func (rd *RootDoc) GetStyleByID(styleID string, styleType stypes.StyleType) *ctypes.Style {
	if rd.DocStyles == nil {
		return nil
	}

	for _, style := range rd.DocStyles.StyleList {
		if style.ID == nil || style.Type == nil {
			continue
		}

		if *style.ID == styleID && *style.Type == styleType {
			return &style
		}
	}
	return nil
}
