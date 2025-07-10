package ctypes

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"
	"github.com/stretchr/testify/assert"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func TestParagraph_Style(t *testing.T) {
	f := func(styleValue string, expectedStyleValue string) {
		t.Helper()

		p := &Paragraph{}

		p.Style(styleValue)

		assert.NotNil(t, p.Property)
		assert.NotNil(t, p.Property.Style)
		assert.Equal(t, p.Property.Style.Val, expectedStyleValue)
	}

	f("Heading1", "Heading1")
	f("Normal", "Normal")
}
func TestParagraph_Justification(t *testing.T) {
	f := func(justificationValue, expectedJustificationValue stypes.Justification) {
		t.Helper()

		p := &Paragraph{}

		p.Justification(justificationValue)

		assert.NotNil(t, p.Property, "Expected ct.Property to be non-nil")
		assert.NotNil(t, p.Property.Justification, "Expected ct.Property.Justification to be non-nil")
		assert.Equal(t, p.Property.Justification.Val, expectedJustificationValue, "Paragraph.Justification() value mismatch")
	}

	f(stypes.JustificationCenter, stypes.JustificationCenter)
	f(stypes.JustificationLeft, stypes.JustificationLeft)
	f(stypes.JustificationRight, stypes.JustificationRight)
	f(stypes.JustificationBoth, stypes.JustificationBoth)
}

func TestParagraph_Numbering(t *testing.T) {
	f := func(id int, level int, expectedNumID int, expectedILvl int) {
		t.Helper()

		p := &Paragraph{}

		p.Numbering(id, level)

		assert.NotNil(t, p.Property, "Expected ct.Property to be non-nil")
		assert.NotNil(t, p.Property.NumProp, "Expected ct.Property.NumProp to be non-nil")
		assert.Equal(t, expectedNumID, p.Property.NumProp.NumID.Val, "Paragraph.Numbering() NumID value mismatch")
		assert.Equal(t, expectedILvl, p.Property.NumProp.ILvl.Val, "Paragraph.Numbering() ILvl value mismatch")
	}

	f(1, 0, 1, 0)
	f(2, 1, 2, 1)
	f(3, 2, 3, 2)
	f(4, 3, 4, 3)
}

func TestParagraph_Indent(t *testing.T) {
	f := func(indentValue, expectedIndentValue Indent) {
		t.Helper()

		p := &Paragraph{}

		p.Indent(&indentValue)

		assert.NotNil(t, p.Property, "Expected ct.Property to be non-nil")
		assert.NotNil(t, p.Property.Indent, "Expected ct.Property.Indent to be non-nil")
		assert.Equal(t, p.Property.Indent, &expectedIndentValue, "Paragraph.Indent() value mismatch")
	}

	var size6 int = 6
	var size360 int = 360
	var sizeu420 uint64 = 420

	indentLeft := Indent{Left: &size360, Hanging: &sizeu420}
	indentRight := Indent{Right: &size360, Hanging: &sizeu420}
	indentFirst := Indent{FirstLine: &sizeu420}
	indentLeftChars := Indent{LeftChars: &size6}
	indentRightChars := Indent{RightChars: &size6}
	indentFirstChars := Indent{FirstLineChars: &size6}

	f(indentLeft, indentLeft)
	f(indentRight, indentRight)
	f(indentFirst, indentFirst)
	f(indentLeftChars, indentLeftChars)
	f(indentRightChars, indentRightChars)
	f(indentFirstChars, indentFirstChars)
}

func TestParagraph_AddText(t *testing.T) {
	f := func(text string, expectedText string) {
		t.Helper()

		p := &Paragraph{
			Children: []ParagraphChild{},
		}

		run := p.AddText(text)

		assert.NotNil(t, run, "Expected AddText() to return a non-nil Run")
		assert.Equal(t, len(p.Children), 1, "Expected one Run to be added to Paragraph")

		assert.NotNil(t, p.Children[0].Run, "Expected ct.Children[0].Run to be non-nil")
		assert.GreaterOrEqual(t, len(p.Children[0].Run.Children), 1, "Expected at least one Text element in Run")
		assert.NotNil(t, p.Children[0].Run.Children[0].Text, "Expected Text element in Run to be non-nil")
		assert.Equal(t, p.Children[0].Run.Children[0].Text.Text, expectedText, "Paragraph.AddText() Text value mismatch")
	}

	f("Hello, World!", "Hello, World!")
	f("Another test", "Another test")
	f("A third text", "A third text")
	f("Sample text", "Sample text")
}

func TestParagraph_AddRun(t *testing.T) {
	p := &Paragraph{
		Children: []ParagraphChild{},
	}

	run := p.AddRun()

	assert.NotNil(t, run, "Expected AddRun() to return a non-nil Run")

	assert.Equal(t, 1, len(p.Children), "Expected one Run to be added to Paragraph")
	assert.NotNil(t, p.Children[0].Run, "Expected ct.Children[0].Run to be non-nil")
	assert.Equal(t, run, p.Children[0].Run, "Expected the Run returned by AddRun() to match the Run added to Paragraph")

	assert.Empty(t, run.Children, "Expected new Run to have no initial Children")

	assert.Equal(t, 0, len(p.Children[0].Run.Children), "Expected the new Run to have no initial Children")
}

func TestParagraphXML(t *testing.T) {
	// Create a sample paragraph
	p := Paragraph{}
	// p.Style("Heading1")
	// p.Numbering(1, 0)
	p.AddText("This is a sample paragraph.")

	// Marshal the paragraph to XML
	var buf bytes.Buffer

	encoder := xml.NewEncoder(&buf)
	start := xml.StartElement{Name: xml.Name{Local: "fake"}}
	if err := encoder.EncodeElement(p, start); err != nil {
		t.Errorf("Error during MarshalXML: %v", err)
	}

	err := encoder.Flush()
	if err != nil {
		t.Errorf("Error flushing encoder: %v", err)
	}

	// Unmarshal the XML back to a paragraph
	var paraUnmarshaled Paragraph
	ns := map[string]string{
		"w": constants.WMLNamespace,
	}
	decoder := xml.NewDecoder(&buf)
	decoder.DefaultSpace = constants.WMLNamespace
	decoder.Entity = ns

	if err := decoder.Decode(&paraUnmarshaled); err != nil {
		t.Errorf("Error unmarshaling XML to paragraph: %v", err)
		return
	}

	if !reflect.DeepEqual(p, paraUnmarshaled) {
		t.Errorf("Original and unmarshaled paragraphs are not equal.")
	}
}
