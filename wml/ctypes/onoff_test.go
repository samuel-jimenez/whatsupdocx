package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type OnOffXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element OnOff    `xml:"w:rStyle"`
}

func wrapOnOffXML(el OnOff) *OnOffXML {
	return &OnOffXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapOnOffOutput(output string) string {
	return `<OnOffXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</OnOffXML>`
}

func TestOnOff_MarshalXML(t *testing.T) {

	tests := []struct {
		name     string
		input    OnOff
		expected string
	}{
		{
			name:     "With value",
			input:    OnOff{Val: internal.ToPtr(stypes.OnOffFalse)},
			expected: `<w:rStyle w:val="false"></w:rStyle>`,
		},
		{
			name:     "Empty value",
			input:    OnOff{Val: nil},
			expected: `<w:rStyle></w:rStyle>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapOnOffXML(tt.input))
			expected := wrapOnOffOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestOnOff_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected OnOff
	}{
		{
			name:     "With value",
			inputXML: `<w:rStyle w:val="true"></w:rStyle>`,
			expected: OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
		},
		{
			name:     "Empty value",
			inputXML: `<w:rStyle></w:rStyle>`,
			expected: OnOff{Val: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result OnOff

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if err = internal.ComparePtr("Val", tt.expected.Val, result.Val); err != nil {
				t.Error(err)
			}
		})
	}
}
