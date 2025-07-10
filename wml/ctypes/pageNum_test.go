package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type PageNumberingXML struct {
	Attr    xml.Attr      `xml:",any,attr,omitempty"`
	Element PageNumbering `xml:"w:pgNumType"`
}

func wrapPageNumberingXML(el PageNumbering) *PageNumberingXML {
	return &PageNumberingXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapPageNumberingOutput(output string) string {
	return `<PageNumberingXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</PageNumberingXML>`
}

func TestPageNumbering_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PageNumbering
		expected string
	}{
		{
			name:     "With format",
			input:    PageNumbering{Format: stypes.NumFmtDecimal},
			expected: `<w:pgNumType w:fmt="decimal"></w:pgNumType>`,
		},
		{
			name:     "Without format",
			input:    PageNumbering{},
			expected: `<w:pgNumType></w:pgNumType>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapPageNumberingXML(tt.input))
			expected := wrapPageNumberingOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestPageNumbering_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected PageNumbering
	}{
		{
			name:     "With format",
			inputXML: `<w:pgNumType w:fmt="decimal"></w:pgNumType>`,
			expected: PageNumbering{Format: stypes.NumFmtDecimal},
		},
		{
			name:     "Without format",
			inputXML: `<w:pgNumType></w:pgNumType>`,
			expected: PageNumbering{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result PageNumbering

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			if result.Format != tt.expected.Format {
				t.Errorf("Expected Format %s but got %s", tt.expected.Format, result.Format)
			}
		})
	}
}
