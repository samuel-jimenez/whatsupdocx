package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type TableWidthXML struct {
	Attr    xml.Attr   `xml:",any,attr,omitempty"`
	Element TableWidth `xml:"w:tblW"`
}

func wrapTableWidthXML(el TableWidth) *TableWidthXML {
	return &TableWidthXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapTableWidthOutput(output string) string {
	return `<TableWidthXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</TableWidthXML>`
}

func TestTableWidth_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    TableWidth
		expected string
	}{
		{
			name:     "Test with Width only",
			input:    TableWidth{Width: internal.ToPtr(500)},
			expected: `<w:tblW w:w="500"></w:tblW>`,
		},
		{
			name:     "Test with Type only",
			input:    TableWidth{WidthType: internal.ToPtr(stypes.TableWidthDxa)},
			expected: `<w:tblW w:type="dxa"></w:tblW>`,
		},
		{
			name:     "Test with Width and Type",
			input:    TableWidth{Width: internal.ToPtr(1000), WidthType: internal.ToPtr(stypes.TableWidthAuto)},
			expected: `<w:tblW w:w="1000" w:type="auto"></w:tblW>`,
		},
		{
			name:     "Test with nil values",
			input:    TableWidth{},
			expected: `<w:tblW></w:tblW>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTableWidthXML(tt.input))
			expected := wrapTableWidthOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestTableWidth_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   TableWidth
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Width attribute",
			inputXML: `<w:tblW w:w="750"></w:tblW>`,
			expected: TableWidth{Width: internal.ToPtr(750)},
		},
		{
			name:     "Test with Type attribute",
			inputXML: `<w:tblW w:type="dxa"></w:tblW>`,
			expected: TableWidth{WidthType: internal.ToPtr(stypes.TableWidthDxa)},
		},
		{
			name:     "Test with Width and Type attributes",
			inputXML: `<w:tblW w:w="500" w:type="pct"></w:tblW>`,
			expected: TableWidth{Width: internal.ToPtr(500), WidthType: internal.ToPtr(stypes.TableWidthPct)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TableWidth
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected TableWidth %+v but got %+v", tt.expected, result)
			}
		})
	}
}
