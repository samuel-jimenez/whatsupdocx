package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
)

func wrapColumnXML(el Column) *WrapperXML {
	return wrapXML(struct {
		Column
		XMLName struct{} `xml:"w:gridCol"`
	}{Column: el})
}

func TestColumn_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Column
		expected string
	}{
		{
			name:     "With Width",
			input:    Column{Width: internal.ToPtr(uint64(500))},
			expected: `<w:gridCol w:w="500"></w:gridCol>`,
		},
		{
			name:     "Without Width",
			input:    Column{},
			expected: `<w:gridCol></w:gridCol>`,
		},
	}

	for _, tt := range tests {
		object := wrapColumnXML(tt.input)
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

func TestColumn_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Column
	}{
		{
			name:     "With Width",
			inputXML: `<w:gridCol w:w="750"></w:gridCol>`,
			expected: Column{Width: internal.ToPtr(uint64(750))},
		},
		{
			name:     "Without Width",
			inputXML: `<w:gridCol></w:gridCol>`,
			expected: Column{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Column

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Width == nil && tt.expected.Width == nil {
				// Both are nil, which is fine
			} else if result.Width == nil || tt.expected.Width == nil || *result.Width != *tt.expected.Width {
				t.Errorf("Expected Width %v but got %v", tt.expected.Width, result.Width)
			}
		})
	}
}
