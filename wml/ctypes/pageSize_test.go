package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type PageSizeXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element PageSize `xml:"w:pgSz"`
}

func wrapPageSizeXML(el PageSize) *PageSizeXML {
	return &PageSizeXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapPageSizeOutput(output string) string {
	return `<PageSizeXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</PageSizeXML>`
}

func TestPageSize_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PageSize
		expected string
	}{
		{
			name: "All attributes",
			input: PageSize{
				Width:  uint64Ptr(12240),
				Height: uint64Ptr(15840),
				Orient: stypes.PageOrientLandscape,
				Code:   intPtr(1),
			},
			expected: `<w:pgSz w:w="12240" w:h="15840" w:orient="landscape" w:code="1"></w:pgSz>`,
		},
		{
			name: "Some attributes",
			input: PageSize{
				Width:  uint64Ptr(12240),
				Height: uint64Ptr(15840),
			},
			expected: `<w:pgSz w:w="12240" w:h="15840"></w:pgSz>`,
		},
		{
			name:     "No attributes",
			input:    PageSize{},
			expected: `<w:pgSz></w:pgSz>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapPageSizeXML(tt.input))
			expected := wrapPageSizeOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestPageSize_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected PageSize
	}{
		{
			name:     "All attributes",
			inputXML: `<w:pgSz w:w="12240" w:h="15840" w:orient="landscape" w:code="1"></w:pgSz>`,
			expected: PageSize{
				Width:  uint64Ptr(12240),
				Height: uint64Ptr(15840),
				Orient: stypes.PageOrientLandscape,
				Code:   intPtr(1),
			},
		},
		{
			name:     "Some attributes",
			inputXML: `<w:pgSz w:w="12240" w:h="15840"></w:pgSz>`,
			expected: PageSize{
				Width:  uint64Ptr(12240),
				Height: uint64Ptr(15840),
			},
		},
		{
			name:     "No attributes",
			inputXML: `<w:pgSz></w:pgSz>`,
			expected: PageSize{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result PageSize

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			comparePageSizes(t, result, tt.expected)
		})
	}
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}

func comparePageSizes(t *testing.T, got, want PageSize) {
	if !compareUint64Ptr(got.Width, want.Width) {
		t.Errorf("Width = %v, want %v", got.Width, want.Width)
	}
	if !compareUint64Ptr(got.Height, want.Height) {
		t.Errorf("Height = %v, want %v", got.Height, want.Height)
	}
	if got.Orient != want.Orient {
		t.Errorf("Orient = %v, want %v", got.Orient, want.Orient)
	}
	if !compareIntPtr(got.Code, want.Code) {
		t.Errorf("Code = %v, want %v", got.Code, want.Code)
	}
}

func compareUint64Ptr(got, want *uint64) bool {
	if got == nil && want == nil {
		return true
	}
	if got == nil || want == nil {
		return false
	}
	return *got == *want
}
