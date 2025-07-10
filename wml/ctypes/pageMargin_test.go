package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/xml"
)

type PageMarginXML struct {
	Attr    xml.Attr   `xml:",any,attr,omitempty"`
	Element PageMargin `xml:"w:titlePg"`
}

func wrapPageMarginXML(el PageMargin) *PageMarginXML {
	return &PageMarginXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapPageMarginOutput(output string) string {
	return `<PageMarginXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</PageMarginXML>`
}

func TestPageMargin_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PageMargin
		expected string
	}{
		{
			name: "All attributes",
			input: PageMargin{
				Left:   intPtr(1440),
				Right:  intPtr(1440),
				Gutter: intPtr(0),
				Header: intPtr(720),
				Top:    intPtr(1440),
				Footer: intPtr(720),
				Bottom: intPtr(1440),
			},
			expected: `<w:pgMar w:top="1440" w:right="1440" w:bottom="1440" w:left="1440" w:header="720" w:footer="720" w:gutter="0"></w:pgMar>`,
		},
		{
			name: "Some attributes",
			input: PageMargin{
				Left:   intPtr(1440),
				Right:  intPtr(1440),
				Top:    intPtr(1440),
				Bottom: intPtr(1440),
			},
			expected: `<w:pgMar w:top="1440" w:right="1440" w:bottom="1440" w:left="1440"></w:pgMar>`,
		},
		{
			name:     "No attributes",
			input:    PageMargin{},
			expected: `<w:pgMar></w:pgMar>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapPageMarginXML(tt.input))
			expected := wrapPageMarginOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestPageMargin_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected PageMargin
	}{
		{
			name:     "All attributes",
			inputXML: `<w:pgMar w:left="1440" w:right="1440" w:gutter="0" w:header="720" w:top="1440" w:footer="720" w:bottom="1440"></w:pgMar>`,
			expected: PageMargin{
				Left:   intPtr(1440),
				Right:  intPtr(1440),
				Gutter: intPtr(0),
				Header: intPtr(720),
				Top:    intPtr(1440),
				Footer: intPtr(720),
				Bottom: intPtr(1440),
			},
		},
		{
			name:     "Some attributes",
			inputXML: `<w:pgMar w:left="1440" w:right="1440" w:top="1440" w:bottom="1440"></w:pgMar>`,
			expected: PageMargin{
				Left:   intPtr(1440),
				Right:  intPtr(1440),
				Top:    intPtr(1440),
				Bottom: intPtr(1440),
			},
		},
		{
			name:     "No attributes",
			inputXML: `<w:pgMar></w:pgMar>`,
			expected: PageMargin{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result PageMargin

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			comparePageMargins(t, result, tt.expected)
		})
	}
}

func intPtr(i int) *int {
	return &i
}

func comparePageMargins(t *testing.T, got, want PageMargin) {
	if !compareIntPtr(got.Left, want.Left) {
		t.Errorf("Left = %v, want %v", got.Left, want.Left)
	}
	if !compareIntPtr(got.Right, want.Right) {
		t.Errorf("Right = %v, want %v", got.Right, want.Right)
	}
	if !compareIntPtr(got.Gutter, want.Gutter) {
		t.Errorf("Gutter = %v, want %v", got.Gutter, want.Gutter)
	}
	if !compareIntPtr(got.Header, want.Header) {
		t.Errorf("Header = %v, want %v", got.Header, want.Header)
	}
	if !compareIntPtr(got.Top, want.Top) {
		t.Errorf("Top = %v, want %v", got.Top, want.Top)
	}
	if !compareIntPtr(got.Footer, want.Footer) {
		t.Errorf("Footer = %v, want %v", got.Footer, want.Footer)
	}
	if !compareIntPtr(got.Bottom, want.Bottom) {
		t.Errorf("Bottom = %v, want %v", got.Bottom, want.Bottom)
	}
}

func compareIntPtr(got, want *int) bool {
	if got == nil && want == nil {
		return true
	}
	if got == nil || want == nil {
		return false
	}
	return *got == *want
}
