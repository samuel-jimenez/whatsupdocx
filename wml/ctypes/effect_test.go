package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapEffectXML(el Effect) *WrapperXML {
	return wrapXML(struct {
		Effect
		XMLName struct{} `xml:"w:effect"`
	}{Effect: el})
}

func TestEffect_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Effect
		expected string
	}{
		{
			name:     "With value",
			input:    Effect{Val: TextEffectPtr(stypes.TextEffectBlinkBackground)},
			expected: `<w:effect w:val="blinkBackground"></w:effect>`,
		},
		{
			name:     "Without value",
			input:    Effect{},
			expected: `<w:effect></w:effect>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapEffectXML(tt.input))
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

func TestEffect_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Effect
	}{
		{
			name:     "With value",
			inputXML: `<w:effect w:val="blinkBackground"></w:effect>`,
			expected: Effect{Val: TextEffectPtr(stypes.TextEffectBlinkBackground)},
		},
		{
			name:     "Without value",
			inputXML: `<w:effect></w:effect>`,
			expected: Effect{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Effect

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if tt.expected.Val != nil {
				if result.Val == nil {
					t.Errorf("Expected Value %s but got nil", *tt.expected.Val)
				} else if *tt.expected.Val != *result.Val {
					t.Errorf("Expected Value %s but got %s", *tt.expected.Val, *result.Val)
				}
			} else {
				if result.Val != nil {
					t.Errorf("Expected nil but got %s", *result.Val)
				}
			}

		})
	}
}

func TextEffectPtr(value stypes.TextEffect) *stypes.TextEffect {
	return &value
}
