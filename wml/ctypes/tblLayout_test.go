package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapTableLayoutXML(el *TableLayout) *WrapperXML {
	return wrapXML(struct {
		*TableLayout
		XMLName struct{} `xml:"w:tblLayout"`
	}{TableLayout: el})
}

func TestTableLayout_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    *TableLayout
		expected string
	}{{
		input:    &TableLayout{LayoutType: internal.ToPtr(stypes.TableLayoutFixed)},
		expected: `<w:tblLayout w:type="fixed"></w:tblLayout>`,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTableLayoutXML(tt.input))
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

func TestLayout_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   TableLayout
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Overlap Value `never`",
			inputXML: `<w:tblLayout w:type="fixed"></w:tblLayout>`,
			expected: TableLayout{LayoutType: internal.ToPtr(stypes.TableLayoutFixed)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TableLayout
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			err = internal.ComparePtr("Type", tt.expected.LayoutType, result.LayoutType)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
