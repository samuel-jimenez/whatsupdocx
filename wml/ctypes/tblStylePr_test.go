package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type TableStylePropXML struct {
	Attr    xml.Attr        `xml:",any,attr,omitempty"`
	Element *TableStyleProp `xml:"w:tblStylePr"`
}

func wrapTableStylePropXML(el *TableStyleProp) *TableStylePropXML {
	return &TableStylePropXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapTableStylePropOutput(output string) string {
	return `<TableStylePropXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</TableStylePropXML>`
}

func TestTableStyleProp_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    *TableStyleProp
		expected string
	}{
		{
			name: "all properties set",
			input: &TableStyleProp{
				ParaProp:  &ParagraphProp{},
				RunProp:   &RunProperty{},
				TableProp: &TableProp{},
				RowProp:   &RowProperty{},
				CellProp:  &CellProperty{},
				Type:      stypes.TblStyleOverrideType("testType"),
			},
			expected: `<w:tblStylePr w:type="testType"><w:pPr></w:pPr><w:rPr></w:rPr><w:tblPr></w:tblPr><w:trPr></w:trPr><w:tcPr></w:tcPr></w:tblStylePr>`,
		},
		{
			name: "some properties nil",
			input: &TableStyleProp{
				ParaProp:  nil,
				RunProp:   &RunProperty{},
				TableProp: nil,
				RowProp:   &RowProperty{},
				CellProp:  nil,
				Type:      stypes.TblStyleOverrideType("testType"),
			},
			expected: `<w:tblStylePr w:type="testType"><w:rPr></w:rPr><w:trPr></w:trPr></w:tblStylePr>`,
		},
		{
			name: "all properties nil",
			input: &TableStyleProp{
				ParaProp:  nil,
				RunProp:   nil,
				TableProp: nil,
				RowProp:   nil,
				CellProp:  nil,
				Type:      stypes.TblStyleOverrideType("testType"),
			},
			expected: `<w:tblStylePr w:type="testType"></w:tblStylePr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTableStylePropXML(tt.input))
			expected := wrapTableStylePropOutput(tt.expected)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestTableStyleProp_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected *TableStyleProp
	}{
		{
			name: "all properties set",
			xmlInput: `<w:tblStylePr w:type="neCell">
							<w:pPr></w:pPr>
							<w:rPr></w:rPr>
							<w:tblPr></w:tblPr>
							<w:trPr></w:trPr>
							<w:tcPr></w:tcPr>
						</w:tblStylePr>`,
			expected: &TableStyleProp{
				ParaProp:  &ParagraphProp{},
				RunProp:   &RunProperty{},
				TableProp: &TableProp{},
				RowProp:   &RowProperty{},
				CellProp:  &CellProperty{},
				Type:      stypes.TblStyleOverrideNeCell,
			},
		},
		{
			name: "some properties nil",
			xmlInput: `<w:tblStylePr w:type="neCell">
							<w:rPr></w:rPr>
							<w:trPr></w:trPr>
						</w:tblStylePr>`,
			expected: &TableStyleProp{
				RunProp: &RunProperty{},
				RowProp: &RowProperty{},
				Type:    stypes.TblStyleOverrideNeCell,
			},
		},
		{
			name:     "all properties nil",
			xmlInput: `<w:tblStylePr w:type="neCell"></w:tblStylePr>`,
			expected: &TableStyleProp{
				Type: stypes.TblStyleOverrideNeCell,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var prop TableStyleProp
			err := xml.Unmarshal([]byte(tt.xmlInput), &prop)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(&prop, tt.expected) {
				t.Errorf("expected %+v, but got %+v", tt.expected, &prop)
			}
		})
	}
}
