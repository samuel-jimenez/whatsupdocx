package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapBorderXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Border
		XMLName struct{} `xml:"w:bdr"`
	}{Border: el.(*Border)})
}

func TestBorder(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapBorderXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With all attributes",
			Input: &Border{
				Val:        stypes.BorderStyleSingle,
				Color:      StringPtr("FF0000"),
				ThemeColor: themeColorPointer(stypes.ThemeColorAccent1),
				ThemeTint:  StringPtr("500"),
				ThemeShade: StringPtr("200"),
				Space:      StringPtr("0"),
				Shadow:     OnOffPtr(stypes.OnOffTrue),
				Size:       internal.ToPtr(19),
				Frame:      OnOffPtr(stypes.OnOffTrue),
			},
			ExpectedXML: `<w:bdr w:val="single" w:sz="19" w:space="0" w:color="FF0000" w:themeColor="accent1" w:themeTint="500" w:themeShade="200" w:shadow="true" w:frame="true"></w:bdr>`,
		},
		{
			Name: "Without optional attributes",
			Input: &Border{
				Val: stypes.BorderStyleDouble,
			},
			ExpectedXML: `<w:bdr w:val="double"></w:bdr>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func StringPtr(s string) *string {
	return &s
}

func OnOffPtr(o stypes.OnOff) *stypes.OnOff {
	return &o
}

func themeColorPointer(t stypes.ThemeColor) *stypes.ThemeColor {
	return &t
}
