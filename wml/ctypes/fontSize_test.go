package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapFontSizeXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*FontSize
		XMLName struct{} `xml:"w:sz"`
	}{FontSize: el.(*FontSize)})
}

func wrapFontSizeCSXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*FontSize
		XMLName struct{} `xml:"w:szCs"`
	}{FontSize: el.(*FontSize)})
}

func TestFontSize(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapFontSizeXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With value",
			Input:       NewFontSize(24),
			ExpectedXML: `<w:sz w:val="24"></w:sz>`,
		},
		{
			Name:        "Without value",
			Input:       &FontSize{},
			ExpectedXML: `<w:sz w:val="0"></w:sz>`,
		},
		{
			Name:          "Without value",
			Input:         &FontSize{Value: 0},
			ExpectedXML:   `<w:sz></w:sz>`,
			UnmarshalOnly: true,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func TestFontSizeCS(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapFontSizeCSXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With value",
			Input:       NewFontSize(24),
			ExpectedXML: `<w:szCs w:val="24"></w:szCs>`,
		},
		{
			Name:        "Without value",
			Input:       &FontSize{},
			ExpectedXML: `<w:szCs w:val="0"></w:szCs>`,
		},
		{
			Name:          "Without value",
			Input:         &FontSize{Value: 0},
			ExpectedXML:   `<w:szCs></w:szCs>`,
			UnmarshalOnly: true,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
