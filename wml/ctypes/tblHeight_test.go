package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type TableRowHeightXML struct {
	Attr    xml.Attr       `xml:",any,attr,omitempty"`
	Element TableRowHeight `xml:"w:val"`
}

func wrapTableRowHeightXML(el TableRowHeight) *TableRowHeightXML {
	return &TableRowHeightXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapTableRowHeightOutput(output string) string {
	return `<TableRowHeightXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</TableRowHeightXML>`
}

func TestTableRowHeight_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    TableRowHeight
		expected string
	}{
		{
			name:     "Test with Val only",
			input:    TableRowHeight{Val: internal.ToPtr(500)},
			expected: `<w:val w:val="500"></w:val>`,
		},
		{
			name:     "Test with HRule only",
			input:    TableRowHeight{HRule: internal.ToPtr(stypes.HeightRuleAtLeast)},
			expected: `<w:val w:hRule="atLeast"></w:val>`,
		},
		{
			name:     "Test with Val and HRule",
			input:    TableRowHeight{Val: internal.ToPtr(1000), HRule: internal.ToPtr(stypes.HeightRuleExact)},
			expected: `<w:val w:val="1000" w:hRule="exact"></w:val>`,
		},
		{
			name:     "Test with nil values",
			input:    TableRowHeight{},
			expected: `<w:val></w:val>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTableRowHeightXML(tt.input))
			expected := wrapTableRowHeightOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestTableRowHeight_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   TableRowHeight
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Val attribute",
			inputXML: `<w:val w:val="750"></w:val>`,
			expected: TableRowHeight{Val: internal.ToPtr(750)},
		},
		{
			name:     "Test with HRule attribute",
			inputXML: `<w:val w:hRule="auto"></w:val>`,
			expected: TableRowHeight{HRule: internal.ToPtr(stypes.HeightRuleAuto)},
		},
		{
			name:     "Test with Val and HRule attributes",
			inputXML: `<w:val w:val="500" w:hRule="atLeast"></w:val>`,
			expected: TableRowHeight{Val: internal.ToPtr(500), HRule: internal.ToPtr(stypes.HeightRuleAtLeast)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TableRowHeight
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected TableRowHeight %+v but got %+v", tt.expected, result)
			}
		})
	}
}
