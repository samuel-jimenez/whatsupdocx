package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapLangXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Lang
		XMLName struct{} `xml:"w:lang"`
	}{Lang: el.(*Lang)})
}

func TestLang(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapLangXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{

		{
			Name:        "All attributes set",
			Input:       &Lang{Val: strPtr("en-US"), EastAsia: strPtr("ja-JP"), Bidi: strPtr("ar-SA")},
			ExpectedXML: `<w:lang w:val="en-US" w:eastAsia="ja-JP" w:bidi="ar-SA"></w:lang>`,
		},
		{
			Name:        "Only val set",
			Input:       &Lang{Val: strPtr("en-US")},
			ExpectedXML: `<w:lang w:val="en-US"></w:lang>`,
		},
		{
			Name:        "Only eastAsia set",
			Input:       &Lang{EastAsia: strPtr("ja-JP")},
			ExpectedXML: `<w:lang w:eastAsia="ja-JP"></w:lang>`,
		},
		{
			Name:        "Only bidi set",
			Input:       &Lang{Bidi: strPtr("ar-SA")},
			ExpectedXML: `<w:lang w:bidi="ar-SA"></w:lang>`,
		},
		{
			Name:        "No attributes set",
			Input:       &Lang{},
			ExpectedXML: `<w:lang></w:lang>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func strPtr(s string) *string {
	return &s
}
