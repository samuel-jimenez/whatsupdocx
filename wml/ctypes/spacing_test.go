package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type SpacingXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element Spacing  `xml:"w:spacing"`
}

func wrapSpacingXML(el Spacing) *SpacingXML {
	return &SpacingXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapSpacingOutput(output string) string {
	return `<SpacingXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</SpacingXML>`
}

func TestSpacing_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Spacing
		expected string
	}{
		{
			name: "All fields set",
			input: Spacing{
				Before:            internal.ToPtr(uint64(120)),
				After:             internal.ToPtr(uint64(240)),
				BeforeLines:       internal.ToPtr(10),
				BeforeAutospacing: internal.ToPtr(stypes.OnOffOn),
				AfterAutospacing:  internal.ToPtr(stypes.OnOffOff),
				Line:              internal.ToPtr(360),
				LineRule:          internal.ToPtr(stypes.LineSpacingRuleExact),
			},
			expected: `<w:spacing w:before="120" w:after="240" w:beforeLines="10" w:beforeAutospacing="on" w:afterAutospacing="off" w:line="360" w:lineRule="exact"></w:spacing>`,
		},
		{
			name: "Some fields set",
			input: Spacing{
				Before:   internal.ToPtr(uint64(120)),
				After:    internal.ToPtr(uint64(240)),
				Line:     internal.ToPtr(360),
				LineRule: internal.ToPtr(stypes.LineSpacingRuleAuto),
			},
			expected: `<w:spacing w:before="120" w:after="240" w:line="360" w:lineRule="auto"></w:spacing>`,
		},
		{
			name:     "No fields set",
			input:    Spacing{},
			expected: `<w:spacing></w:spacing>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapSpacingXML(tt.input))
			expected := wrapSpacingOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}
