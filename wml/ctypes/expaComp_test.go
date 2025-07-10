package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapExpaCompXML(el ExpaComp) *WrapperXML {
	return wrapXML(struct {
		ExpaComp
		XMLName struct{} `xml:"w:sz"`
	}{ExpaComp: el})
}

func TestExpaComp_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    ExpaComp
		expected string
	}{
		{
			name:     "With value",
			input:    ExpaComp{Val: TextScalePtr(24)},
			expected: `<w:sz w:val="24"></w:sz>`,
		},
		{
			name:     "Without value",
			input:    ExpaComp{},
			expected: `<w:sz></w:sz>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapExpaCompXML(tt.input))
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

func TestExpaComp_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected ExpaComp
	}{
		{
			name:     "With value",
			inputXML: `<w:w w:val="24"></w:w>`,
			expected: ExpaComp{Val: TextScalePtr(24)},
		},
		{
			name:     "Without value",
			inputXML: `<w:w></w:w>`,
			expected: ExpaComp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result ExpaComp

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if tt.expected.Val != nil {
				if result.Val == nil {
					t.Errorf("Expected Value %d but got nil", *tt.expected.Val)
				} else if *tt.expected.Val != *result.Val {
					t.Errorf("Expected Value %d but got %d", *tt.expected.Val, *result.Val)
				}
			} else {
				if result.Val != nil {
					t.Errorf("Expected nil but got %d", *result.Val)
				}
			}

		})
	}
}

func TextScalePtr(value uint16) *stypes.TextScale {
	ts := stypes.TextScale(value)
	return &ts
}
