package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
)

type SymXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element Sym      `xml:"w:sym"`
}

func wrapSymXML(el Sym) *SymXML {
	return &SymXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapSymOutput(output string) string {
	return `<SymXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</SymXML>`
}

func TestSym_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Sym
		expected string
	}{
		{
			name:     "Test with Font only",
			input:    Sym{Font: internal.ToPtr("Arial")},
			expected: `<w:sym w:font="Arial"></w:sym>`,
		},
		{
			name:     "Test with Char only",
			input:    Sym{Char: internal.ToPtr("F0")},
			expected: `<w:sym w:char="F0"></w:sym>`,
		},
		{
			name:     "Test with Font and Char",
			input:    Sym{Font: internal.ToPtr("Times New Roman"), Char: internal.ToPtr("03")},
			expected: `<w:sym w:font="Times New Roman" w:char="03"></w:sym>`,
		},
		{
			name:     "Test with nil values",
			input:    Sym{},
			expected: `<w:sym></w:sym>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapSymXML(tt.input))
			expected := wrapSymOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestSym_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   Sym
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Font attribute",
			inputXML: `<w:sym w:font="Verdana"></w:sym>`,
			expected: Sym{Font: internal.ToPtr("Verdana")},
		},
		{
			name:     "Test with Char attribute",
			inputXML: `<w:sym w:char="0E"></w:sym>`,
			expected: Sym{Char: internal.ToPtr("0E")},
		},
		{
			name:     "Test with Font and Char attributes",
			inputXML: `<w:sym w:font="Arial" w:char="F2"></w:sym>`,
			expected: Sym{Font: internal.ToPtr("Arial"), Char: internal.ToPtr("F2")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Sym
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected Sym %+v but got %+v", tt.expected, result)
			}
		})
	}
}
