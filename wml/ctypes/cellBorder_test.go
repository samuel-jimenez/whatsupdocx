package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapCellBordersXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*CellBorders
		XMLName struct{} `xml:"w:tcBorders"`
	}{CellBorders: el.(*CellBorders)})
}

func TestCellBorders(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapCellBordersXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	colorRed := "red"
	themeColorAccent1 := stypes.ThemeColor("accent1")
	themeTint := "80"
	themeShade := "20"
	space := "4"
	shadow := stypes.OnOff("true")
	frame := stypes.OnOff("false")
	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With all attributes",
			Input: &CellBorders{
				Top:     &Border{Val: stypes.BorderStyleSingle, Color: &colorRed, ThemeColor: &themeColorAccent1, ThemeTint: &themeTint, ThemeShade: &themeShade, Space: &space, Shadow: &shadow, Frame: &frame},
				Left:    &Border{Val: stypes.BorderStyleDouble},
				Bottom:  &Border{Val: stypes.BorderStyleDashed},
				Right:   &Border{Val: stypes.BorderStyleDotted},
				InsideH: &Border{Val: stypes.BorderStyleSingle},
				InsideV: &Border{Val: stypes.BorderStyleDouble},
				TL2BR:   &Border{Val: stypes.BorderStyleThick},
				TR2BL:   &Border{Val: stypes.BorderStyleThick},
			},
			ExpectedXML: `<w:tcBorders>` +
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
			Name:        "Default",
			Input:       DefaultCellBorders(),
			ExpectedXML: `<w:tcBorders></w:tcBorders>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
