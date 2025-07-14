package dmlct

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/common"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapSolidColorFillPropertiesXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*SolidColorFillProperties
	}{SolidColorFillProperties: el.(*SolidColorFillProperties)})
}

func TestSolidColorFillProperties(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapSolidColorFillPropertiesXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{

		{
			Name: "With ColorChoice",
			Input: &SolidColorFillProperties{ColorChoice: &ColorChoice{
				SchemeColor: &SchemeColor{Val: "accent1",
					ColorTransform: []ColorTransform{{LumMod: &Percentage{Val: "50000"}}},
				},
			}},
			ExpectedXML: `<Element><a:schemeClr val="accent1"><a:lumMod val="50000"></a:lumMod></a:schemeClr></Element>`,
		},
		{
			Name: "With SchemeColor",
			Input: &SolidColorFillProperties{ColorChoice: &ColorChoice{
				SchemeColor: &SchemeColor{Val: "accent1"},
			}},
			ExpectedXML: `<Element><a:schemeClr val="accent1"></a:schemeClr></Element>`,
		},
		{
			Name:        "Default",
			Input:       &SolidColorFillProperties{},
			ExpectedXML: `<Element></Element>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func wrapLineFillPropertiesXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*LineFillProperties
	}{LineFillProperties: el.(*LineFillProperties)})
}

func TestLineFillProperties(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapLineFillPropertiesXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "noFill",
			Input: &LineFillProperties{
				NoFillProperties: &common.Empty{},
			},
			ExpectedXML: `<Element><a:noFill></a:noFill></Element>`,
		},
		{
			Name: "solidFill",
			Input: &LineFillProperties{
				SolidColorFillProperties: &SolidColorFillProperties{ColorChoice: &ColorChoice{
					SchemeColor: &SchemeColor{Val: "accent1",
						ColorTransform: []ColorTransform{{LumMod: &Percentage{Val: "50000"}}}},
				}},
			},
			ExpectedXML: `<Element><a:solidFill><a:schemeClr val="accent1"><a:lumMod val="50000"></a:lumMod></a:schemeClr></a:solidFill></Element>`,
		},
		{
			Name:        "Default",
			Input:       &LineFillProperties{},
			ExpectedXML: `<Element></Element>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func wrapFontReferenceXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*FontReference
	}{FontReference: el.(*FontReference)})
}

func TestFontReference(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapFontReferenceXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With ColorChoice",
			Input: &FontReference{
				Id: FontCollectionIndexMajor,
				ColorChoice: &ColorChoice{
					SchemeColor: &SchemeColor{Val: "accent1",
						ColorTransform: []ColorTransform{{LumMod: &Percentage{Val: "50000"}}},
					},
				}},
			ExpectedXML: `<Element idx="major"><a:schemeClr val="accent1"><a:lumMod val="50000"></a:lumMod></a:schemeClr></Element>`,
		},
		{
			Name: "With SchemeColor",
			Input: &FontReference{
				Id: FontCollectionIndexMinor,
				ColorChoice: &ColorChoice{
					SchemeColor: &SchemeColor{Val: "accent1"},
				}},
			ExpectedXML: `<Element idx="minor"><a:schemeClr val="accent1"></a:schemeClr></Element>`,
		},
		{
			Name:        "Default",
			Input:       &FontReference{Id: FontCollectionIndexNone},
			ExpectedXML: `<Element idx="none"></Element>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

/*
	NoFillProperties *common.Empty `xml:"a:noFill,omitempty"`
	// | element solidFill { a_CT_SolidColorFillProperties }
	SolidColorFillProperties *SolidColorFillProperties `xml:"a:solidFill,omitempty"`

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
*/
