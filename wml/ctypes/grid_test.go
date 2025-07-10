package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
)

func wrapGridXML(el Grid) *WrapperXML {
	return wrapXML(struct {
		Grid
		XMLName struct{} `xml:"w:tblGrid"`
	}{Grid: el})
}

func TestGrid_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Grid
		expected string
	}{
		{
			name: "With Columns and GridChange",
			input: Grid{
				Col: []Column{
					{Width: internal.ToPtr(uint64(500))},
					{Width: internal.ToPtr(uint64(750))},
				},
				GridChange: &GridChange{ID: 1},
			},
			expected: `<w:tblGrid><w:gridCol w:w="500"></w:gridCol><w:gridCol w:w="750"></w:gridCol><w:tblGridChange w:id="1"></w:tblGridChange></w:tblGrid>`,
		},
		{
			name: "With Columns, without GridChange",
			input: Grid{
				Col: []Column{
					{Width: internal.ToPtr(uint64(300))},
					{Width: internal.ToPtr(uint64(600))},
				},
			},
			expected: `<w:tblGrid><w:gridCol w:w="300"></w:gridCol><w:gridCol w:w="600"></w:gridCol></w:tblGrid>`,
		},
		{
			name:     "Empty Grid",
			input:    Grid{},
			expected: `<w:tblGrid></w:tblGrid>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapGridXML(tt.input))
			expected := wrapXMLOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestGrid_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Grid
	}{
		{
			name:     "With Columns and GridChange",
			inputXML: `<w:tblGrid><w:gridCol w:w="500"></w:gridCol><w:gridCol w:w="750"></w:gridCol><w:tblGridChange w:id="1"></w:tblGridChange></w:tblGrid>`,
			expected: Grid{
				Col: []Column{
					{Width: internal.ToPtr(uint64(500))},
					{Width: internal.ToPtr(uint64(750))},
				},
				GridChange: &GridChange{ID: 1},
			},
		},
		{
			name:     "With Columns, without GridChange",
			inputXML: `<w:tblGrid><w:gridCol w:w="300"></w:gridCol><w:gridCol w:w="600"></w:gridCol></w:tblGrid>`,
			expected: Grid{
				Col: []Column{
					{Width: internal.ToPtr(uint64(300))},
					{Width: internal.ToPtr(uint64(600))},
				},
			},
		},
		{
			name:     "Empty Grid",
			inputXML: `<w:tblGrid></w:tblGrid>`,
			expected: Grid{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Grid

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unmarshaled Grid struct does not match expected:\nExpected: %+v\nActual:   %+v", tt.expected, result)
			}
		})
	}
}
