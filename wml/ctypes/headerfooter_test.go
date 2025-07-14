package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapHeaderReferenceXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*HeaderFooterReference
		XMLName struct{} `xml:"w:headerReference"`
	}{HeaderFooterReference: el.(*HeaderFooterReference)})
}

func wrapFooterReferenceXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*HeaderFooterReference
		XMLName struct{} `xml:"w:footerReference"`
	}{HeaderFooterReference: el.(*HeaderFooterReference)})
}

func TestHeaderReference(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapHeaderReferenceXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With ID and Type",
			Input: &HeaderFooterReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
			ExpectedXML: `<w:headerReference r:id="rId1" w:type="first"></w:headerReference>`,
		},
		{
			Name: "With Type only",
			Input: &HeaderFooterReference{
				Type: stypes.HdrFtrEven,
			},
			ExpectedXML: `<w:headerReference w:type="even"></w:headerReference>`,
		},
		{
			Name: "With ID only",
			Input: &HeaderFooterReference{
				ID: "rId2",
			},
			ExpectedXML: `<w:headerReference r:id="rId2"></w:headerReference>`,
		},
		{
			Name:        "With neither ID nor Type",
			Input:       &HeaderFooterReference{},
			ExpectedXML: `<w:headerReference></w:headerReference>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func TestFooterReference(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapFooterReferenceXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With ID and Type",
			Input: &HeaderFooterReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
			ExpectedXML: `<w:footerReference r:id="rId1" w:type="first"></w:footerReference>`,
		},
		{
			Name: "With Type only",
			Input: &HeaderFooterReference{
				Type: stypes.HdrFtrEven,
			},
			ExpectedXML: `<w:footerReference w:type="even"></w:footerReference>`,
		},
		{
			Name: "With ID only",
			Input: &HeaderFooterReference{
				ID: "rId2",
			},
			ExpectedXML: `<w:footerReference r:id="rId2"></w:footerReference>`,
		},
		{
			Name:        "With neither ID nor Type",
			Input:       &HeaderFooterReference{},
			ExpectedXML: `<w:footerReference></w:footerReference>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
