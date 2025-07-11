package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapCellXML(el Cell) *WrapperXML {
	return wrapXML(struct {
		Cell
		XMLName struct{} `xml:"w:tc"`
	}{Cell: el})
}

func TestCell_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Cell
		expected string
	}{
		{
			name:     "Empty Cell",
			input:    Cell{},
			expected: `<w:tc></w:tc>`,
		},
		{
			name: "Cell with Property and Paragraph Content",
			input: Cell{
				Property: &CellProperty{
					NoWrap: &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				},
				Contents: []BlockLevel{
					{
						Paragraph: AddParagraph("Test paragraph content"),
					},
				},
			},
			expected: `<w:tc><w:tcPr><w:noWrap w:val="true"></w:noWrap></w:tcPr>` +
				`<w:p><w:r><w:t>Test paragraph content</w:t></w:r></w:p></w:tc>`,
		},
		{
			name: "Cell with Table Content",
			input: Cell{
				Contents: []BlockLevel{
					{
						Table: &Table{},
					},
				},
			},
			expected: `<w:tc><w:tbl><w:tblPr></w:tblPr><w:tblGrid></w:tblGrid></w:tbl></w:tc>`,
		},
	}

	for _, tt := range tests {
		object := wrapCellXML(tt.input)
		expected := wrapXMLOutput(tt.expected)
		t.Run(tt.name, func(t *testing.T) {
			t.Run("MarshalXML", func(t *testing.T) {
				output, err := xml.Marshal(object)
				if err != nil {
					t.Fatalf("Error marshaling to XML: %v", err)
				}
				if got := string(output); got != expected {
					t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
				}
			}) //TODO UnmarshalXML_TOO
		})
	}
}

func TestCellUnmarshalXML(t *testing.T) {
	xmlData := `
<w:tc xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
	<w:tcPr>
		<w:cnfStyle val="001000000000" />
		<w:tcW w:w="5000" w:type="dxa" />
	</w:tcPr>
	<w:p>
		<w:r>
			<w:t>Hello, World!</w:t>
		</w:r>
	</w:p>
	<w:tbl>
		<w:tr>
			<w:tc>
				<w:p>
					<w:r>
						<w:t>Nested Table Cell</w:t>
					</w:r>
				</w:p>
			</w:tc>
		</w:tr>
	</w:tbl>
</w:tc>`

	var cell Cell
	err := xml.Unmarshal([]byte(xmlData), &cell)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Check CellProperty
	if cell.Property == nil {
		t.Fatal("cell.Property should not be nil")
	}
	if cell.Property.CnfStyle == nil {
		t.Fatal("cell.Property.CnfStyle should not be nil")
	}
	if cell.Property.Width == nil {
		t.Fatal("cell.Property.Width should not be nil")
	}

	// Check Contents
	if len(cell.Contents) != 2 {
		t.Fatalf("Expected 2 contents, got %d", len(cell.Contents))
	}
	for i, content := range cell.Contents {
		if content.Paragraph != nil {
			// No need to check deep levels of Paragraph, just ensure it's not nil
			if content.Paragraph == nil {
				t.Fatalf("cell.Contents[%d].Paragraph should not be nil", i)
			}
		} else if content.Table != nil {
			// No need to check deep levels of Table, just ensure it's not nil
			if content.Table == nil {
				t.Fatalf("cell.Contents[%d].Table should not be nil", i)
			}
		} else {
			t.Errorf("Unexpected BlockLevel: %+v", content)
		}
	}
}
