package dmlct

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapColorChoiceXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*ColorChoice
	}{ColorChoice: el.(*ColorChoice)})
}

func TestColorChoice(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapColorChoiceXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With lumMod",
			Input: &ColorChoice{
				SchemeColor: &SchemeColor{Val: "accent1",
					ColorTransform: []ColorTransform{{LumMod: &Percentage{Val: "50000"}}},
				},
			},
			ExpectedXML: `<Element><a:schemeClr val="accent1"><a:lumMod val="50000"></a:lumMod></a:schemeClr></Element>`,
		},
		{
			Name: "With SchemeColor",
			Input: &ColorChoice{
				SchemeColor: &SchemeColor{Val: "accent1"},
			},
			ExpectedXML: `<Element><a:schemeClr val="accent1"></a:schemeClr></Element>`,
		},
		{
			Name:        "Default",
			Input:       &ColorChoice{},
			ExpectedXML: `<Element></Element>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func wrapSchemeColorXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*SchemeColor
		XMLName struct{} `xml:"a:schemeClr"`
	}{SchemeColor: el.(*SchemeColor)})
}

func TestSchemeColor(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapSchemeColorXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With lumMod",
			Input: &SchemeColor{Val: "accent1",
				ColorTransform: []ColorTransform{
					{LumMod: &Percentage{Val: "50000"}}},
			},
			ExpectedXML: `<a:schemeClr val="accent1"><a:lumMod val="50000"></a:lumMod></a:schemeClr>`,
		},
		{
			Name:        "Attributes",
			Input:       &SchemeColor{Val: "accent1"},
			ExpectedXML: `<a:schemeClr val="accent1"></a:schemeClr>`,
		},
		{
			Name:        "Default",
			Input:       &SchemeColor{},
			ExpectedXML: `<a:schemeClr val=""></a:schemeClr>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
