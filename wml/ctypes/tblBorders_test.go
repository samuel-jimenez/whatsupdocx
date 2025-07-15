package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type TableBordersXML struct {
	Attr    xml.Attr      `xml:",any,attr,omitempty"`
	Element *TableBorders `xml:"w:tblBorders"`
}

func wrapTableBordersXML(el *TableBorders) *TableBordersXML {
	return &TableBordersXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapTableBordersOutput(output string) string {
	return `<TableBordersXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</TableBordersXML>`
}

func TestTableBorders_MarshalXML(t *testing.T) {
	colorRed := "red"
	themeColorAccent1 := stypes.ThemeColor("accent1")
	themeTint := "80"
	themeShade := "20"
	space := "4"
	size := 150
	shadow := stypes.OnOff("true")
	frame := stypes.OnOff("false")

	tests := []struct {
		name     string
		input    *TableBorders
		expected string
	}{{
		input: &TableBorders{
			Top:     &Border{Val: stypes.BorderStyleSingle, Color: &colorRed, ThemeColor: &themeColorAccent1, ThemeTint: &themeTint, ThemeShade: &themeShade, Size: &size, Space: &space, Shadow: &shadow, Frame: &frame},
			Left:    &Border{Val: stypes.BorderStyleDouble},
			Bottom:  &Border{Val: stypes.BorderStyleDashed},
			Right:   &Border{Val: stypes.BorderStyleDotted},
			InsideH: &Border{Val: stypes.BorderStyleSingle},
			InsideV: &Border{Val: stypes.BorderStyleDouble},
		},
		expected: `<w:tblBorders>` +
			`<w:top w:val="single" w:sz="150" w:space="4" w:color="red" w:themeColor="accent1" w:themeTint="80" w:themeShade="20" w:shadow="true" w:frame="false"></w:top>` +
			`<w:left w:val="double"></w:left>` +
			`<w:bottom w:val="dashed"></w:bottom>` +
			`<w:right w:val="dotted"></w:right>` +
			`<w:insideH w:val="single"></w:insideH>` +
			`<w:insideV w:val="double"></w:insideV>` +
			`</w:tblBorders>`,
	}, {
		name:     "Default",
		input:    DefaultTableBorders(),
		expected: `<w:tblBorders></w:tblBorders>`,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTableBordersXML(tt.input))
			expected := wrapTableBordersOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestTableBorders_UnmarshalXML_Valid(t *testing.T) {
	xmlData := `
	<w:tblBorders xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
		<w:top w:val="single"></w:top>
		<w:left w:val="double"></w:left>
		<w:bottom w:val="dashed"></w:bottom>
		<w:right w:val="dotted"></w:right>
		<w:insideH w:val="single"></w:insideH>
		<w:insideV w:val="double"></w:insideV>
	</w:tblBorders>`

	expectedBorders := &TableBorders{
		Top:     &Border{Val: stypes.BorderStyleSingle},
		Left:    &Border{Val: stypes.BorderStyleDouble},
		Bottom:  &Border{Val: stypes.BorderStyleDashed},
		Right:   &Border{Val: stypes.BorderStyleDotted},
		InsideH: &Border{Val: stypes.BorderStyleSingle},
		InsideV: &Border{Val: stypes.BorderStyleDouble},
	}

	var unmarshaledBorders TableBorders

	err := xml.Unmarshal([]byte(xmlData), &unmarshaledBorders)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to TableBorders: %v", err)
	}

	if !reflect.DeepEqual(&unmarshaledBorders, expectedBorders) {
		t.Errorf("Expected %#v, got %#v", expectedBorders, unmarshaledBorders)
	}
}
