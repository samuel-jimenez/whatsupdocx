package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapCellBordersXML(el *CellBorders) *WrapperXML {
	return wrapXML(struct {
		*CellBorders
		XMLName struct{} `xml:"w:tcBorders"`
	}{CellBorders: el})
}

func TestCellBorders_MarshalXML(t *testing.T) {

	colorRed := "red"
	themeColorAccent1 := stypes.ThemeColor("accent1")
	themeTint := "80"
	themeShade := "20"
	space := "4"
	shadow := stypes.OnOff("true")
	frame := stypes.OnOff("false")

	tests := []struct {
		name     string
		input    *CellBorders
		expected string
	}{{
		name: "With all attributes",
		input: &CellBorders{
			Top:     &Border{Val: stypes.BorderStyleSingle, Color: &colorRed, ThemeColor: &themeColorAccent1, ThemeTint: &themeTint, ThemeShade: &themeShade, Space: &space, Shadow: &shadow, Frame: &frame},
			Left:    &Border{Val: stypes.BorderStyleDouble},
			Bottom:  &Border{Val: stypes.BorderStyleDashed},
			Right:   &Border{Val: stypes.BorderStyleDotted},
			InsideH: &Border{Val: stypes.BorderStyleSingle},
			InsideV: &Border{Val: stypes.BorderStyleDouble},
			TL2BR:   &Border{Val: stypes.BorderStyleThick},
			TR2BL:   &Border{Val: stypes.BorderStyleThick},
		},
		expected: `<w:tcBorders>` +
			`<w:top w:val="single" w:color="red" w:themeColor="accent1" w:themeTint="80" w:themeShade="20" w:space="4" w:shadow="true" w:frame="false"></w:top>` +
			`<w:left w:val="double"></w:left>` +
			`<w:bottom w:val="dashed"></w:bottom>` +
			`<w:right w:val="dotted"></w:right>` +
			`<w:insideH w:val="single"></w:insideH>` +
			`<w:insideV w:val="double"></w:insideV>` +
			`<w:tl2br w:val="thick"></w:tl2br>` +
			`<w:tr2bl w:val="thick"></w:tr2bl>` +
			`</w:tcBorders>`,
	},
		{
			name:     "Default",
			input:    DefaultCellBorders(),
			expected: `<w:tcBorders></w:tcBorders>`,
		},
	}

	for _, tt := range tests {
		object := wrapCellBordersXML(tt.input)
		expected := wrapXMLOutput(tt.expected)
		t.Run(tt.name, func(t *testing.T) {
			t.Run("MarshalXML", func(t *testing.T) {
				output, err := xml.Marshal(object)
				if err != nil {
					t.Fatalf("Error marshaling to XML: %v", err)
				}
				if got := string(output); got != expected {
					t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
				}
			}) //TODO UnmarshalXML_TOO
		})
	}
}

func TestCellBorders_UnmarshalXML(t *testing.T) {
	xmlData := `
	<w:tcBorders xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
		<w:top w:val="single" w:color="red" w:themeColor="accent1" w:themeTint="80" w:themeShade="20" w:space="4" w:shadow="true" w:frame="false"></w:top>
		<w:left w:val="double"></w:left>
		<w:bottom w:val="dashed"></w:bottom>
		<w:right w:val="dotted"></w:right>
		<w:insideH w:val="single"></w:insideH>
		<w:insideV w:val="double"></w:insideV>
		<w:tl2br w:val="thick"></w:tl2br>
		<w:tr2bl w:val="thick"></w:tr2bl>
	</w:tcBorders>`

	expectedBorders := &CellBorders{
		Top:     &Border{Val: stypes.BorderStyleSingle, Color: internal.ToPtr("red"), ThemeColor: internal.ToPtr(stypes.ThemeColor("accent1")), ThemeTint: internal.ToPtr("80"), ThemeShade: internal.ToPtr("20"), Space: internal.ToPtr("4"), Shadow: internal.ToPtr(stypes.OnOff("true")), Frame: internal.ToPtr(stypes.OnOff("false"))},
		Left:    &Border{Val: stypes.BorderStyleDouble},
		Bottom:  &Border{Val: stypes.BorderStyleDashed},
		Right:   &Border{Val: stypes.BorderStyleDotted},
		InsideH: &Border{Val: stypes.BorderStyleSingle},
		InsideV: &Border{Val: stypes.BorderStyleDouble},
		TL2BR:   &Border{Val: stypes.BorderStyleThick},
		TR2BL:   &Border{Val: stypes.BorderStyleThick},
	}

	var unmarshaledBorders CellBorders

	err := xml.Unmarshal([]byte(xmlData), &unmarshaledBorders)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to CellBorders: %v", err)
	}

	// Test each field individually due to pointer comparisons
	if !reflect.DeepEqual(unmarshaledBorders.Top, expectedBorders.Top) {
		t.Errorf("Top border mismatch. Expected %#v, got %#v", expectedBorders.Top, unmarshaledBorders.Top)
	}
	if !reflect.DeepEqual(unmarshaledBorders.Left, expectedBorders.Left) {
		t.Errorf("Left border mismatch. Expected %#v, got %#v", expectedBorders.Left, unmarshaledBorders.Left)
	}
	if !reflect.DeepEqual(unmarshaledBorders.Bottom, expectedBorders.Bottom) {
		t.Errorf("Bottom border mismatch. Expected %#v, got %#v", expectedBorders.Bottom, unmarshaledBorders.Bottom)
	}
	if !reflect.DeepEqual(unmarshaledBorders.Right, expectedBorders.Right) {
		t.Errorf("Right border mismatch. Expected %#v, got %#v", expectedBorders.Right, unmarshaledBorders.Right)
	}
	if !reflect.DeepEqual(unmarshaledBorders.InsideH, expectedBorders.InsideH) {
		t.Errorf("InsideH border mismatch. Expected %#v, got %#v", expectedBorders.InsideH, unmarshaledBorders.InsideH)
	}
	if !reflect.DeepEqual(unmarshaledBorders.InsideV, expectedBorders.InsideV) {
		t.Errorf("InsideV border mismatch. Expected %#v, got %#v", expectedBorders.InsideV, unmarshaledBorders.InsideV)
	}
	if !reflect.DeepEqual(unmarshaledBorders.TL2BR, expectedBorders.TL2BR) {
		t.Errorf("TL2BR border mismatch. Expected %#v, got %#v", expectedBorders.TL2BR, unmarshaledBorders.TL2BR)
	}
	if !reflect.DeepEqual(unmarshaledBorders.TR2BL, expectedBorders.TR2BL) {
		t.Errorf("TR2BL border mismatch. Expected %#v, got %#v", expectedBorders.TR2BL, unmarshaledBorders.TR2BL)
	}
}
