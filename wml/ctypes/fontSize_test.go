package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"
)

func wrapFontSizeXML(el FontSize) *WrapperXML {
	return wrapXML(struct {
		FontSize
		XMLName struct{} `xml:"w:sz"`
	}{FontSize: el})
}
func wrapFontSizeCSXML(el FontSize) *WrapperXML {
	return wrapXML(struct {
		FontSize
		XMLName struct{} `xml:"w:szCs"`
	}{FontSize: el})
}

func TestFontSize_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FontSize
		expected string
	}{
		{
			name:     "With value",
			input:    *NewFontSize(24),
			expected: `<w:sz w:val="24"></w:sz>`,
		},
		{
			name:     "Without value",
			input:    FontSize{},
			expected: `<w:sz w:val="0"></w:sz>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapFontSizeXML(tt.input))
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

func TestFontSize_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FontSize
	}{
		{
			name:     "With value",
			inputXML: `<w:sz w:val="24"></w:sz>`,
			expected: FontSize{Value: 24},
		},
		{
			name:     "Without value",
			inputXML: `<w:sz></w:sz>`,
			expected: FontSize{Value: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FontSize

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Value != tt.expected.Value {
				t.Errorf("Expected Value %d but got %d", tt.expected.Value, result.Value)
			}
		})
	}
}

func TestFontSizeCS_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FontSize
		expected string
	}{
		{
			name:     "With value",
			input:    *NewFontSize(24),
			expected: `<w:szCs w:val="24"></w:szCs>`,
		},
		{
			name:     "Without value",
			input:    FontSize{},
			expected: `<w:szCs w:val="0"></w:szCs>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapFontSizeCSXML(tt.input))
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

func TestFontSizeCS_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FontSize
	}{
		{
			name:     "With value",
			inputXML: `<w:szCs w:val="24"></w:szCs>`,
			expected: FontSize{Value: 24},
		},
		{
			name:     "Without value",
			inputXML: `<w:szCs></w:szCs>`,
			expected: FontSize{Value: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FontSize

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Value != tt.expected.Value {
				t.Errorf("Expected Value %d but got %d", tt.expected.Value, result.Value)
			}
		})
	}
}
