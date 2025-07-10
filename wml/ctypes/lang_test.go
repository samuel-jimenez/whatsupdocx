package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"
)

func wrapLangXML(el Lang) *WrapperXML {
	return wrapXML(struct {
		Lang
		XMLName struct{} `xml:"w:lang"`
	}{Lang: el})
}

func TestLang_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Lang
		expected string
	}{
		{
			name:     "All attributes set",
			input:    Lang{Val: strPtr("en-US"), EastAsia: strPtr("ja-JP"), Bidi: strPtr("ar-SA")},
			expected: `<w:lang w:val="en-US" w:eastAsia="ja-JP" w:bidi="ar-SA"></w:lang>`,
		},
		{
			name:     "Only val set",
			input:    Lang{Val: strPtr("en-US")},
			expected: `<w:lang w:val="en-US"></w:lang>`,
		},
		{
			name:     "Only eastAsia set",
			input:    Lang{EastAsia: strPtr("ja-JP")},
			expected: `<w:lang w:eastAsia="ja-JP"></w:lang>`,
		},
		{
			name:     "Only bidi set",
			input:    Lang{Bidi: strPtr("ar-SA")},
			expected: `<w:lang w:bidi="ar-SA"></w:lang>`,
		},
		{
			name:     "No attributes set",
			input:    Lang{},
			expected: `<w:lang></w:lang>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapLangXML(tt.input))
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

func TestLang_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Lang
	}{
		{
			name:     "All attributes set",
			inputXML: `<w:lang w:val="en-US" w:eastAsia="ja-JP" w:bidi="ar-SA"></w:lang>`,
			expected: Lang{Val: strPtr("en-US"), EastAsia: strPtr("ja-JP"), Bidi: strPtr("ar-SA")},
		},
		{
			name:     "Only val set",
			inputXML: `<w:lang w:val="en-US"></w:lang>`,
			expected: Lang{Val: strPtr("en-US")},
		},
		{
			name:     "Only eastAsia set",
			inputXML: `<w:lang w:eastAsia="ja-JP"></w:lang>`,
			expected: Lang{EastAsia: strPtr("ja-JP")},
		},
		{
			name:     "Only bidi set",
			inputXML: `<w:lang w:bidi="ar-SA"></w:lang>`,
			expected: Lang{Bidi: strPtr("ar-SA")},
		},
		{
			name:     "No attributes set",
			inputXML: `<w:lang></w:lang>`,
			expected: Lang{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lang Lang
			err := xml.Unmarshal([]byte(tt.inputXML), &lang)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !langEqual(lang, tt.expected) {
				t.Errorf("Expected %+v but got %+v", tt.expected, lang)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}

func langEqual(a, b Lang) bool {
	if (a.Val == nil && b.Val != nil) || (a.Val != nil && b.Val == nil) {
		return false
	}
	if (a.EastAsia == nil && b.EastAsia != nil) || (a.EastAsia != nil && b.EastAsia == nil) {
		return false
	}
	if (a.Bidi == nil && b.Bidi != nil) || (a.Bidi != nil && b.Bidi == nil) {
		return false
	}
	if a.Val != nil && *a.Val != *b.Val {
		return false
	}
	if a.EastAsia != nil && *a.EastAsia != *b.EastAsia {
		return false
	}
	if a.Bidi != nil && *a.Bidi != *b.Bidi {
		return false
	}
	return true
}
