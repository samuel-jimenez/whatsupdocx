package dmlct

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapCNvPrXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*CNvPr
		XMLName struct{} `xml:"pic:cNvPr"`
	}{CNvPr: el.(*CNvPr)})
}

func TestCNvPr(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapCNvPrXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With Description",
			Input: &CNvPr{
				ID:          1,
				Name:        "Drawing1",
				Description: "Description of Drawing1",
			},
			ExpectedXML: `<pic:cNvPr id="1" name="Drawing1" descr="Description of Drawing1"></pic:cNvPr>`,
		},
		{
			Name: "No Description",
			Input: &CNvPr{
				ID:   2,
				Name: "Drawing2",
			},
			ExpectedXML: `<pic:cNvPr id="2" name="Drawing2"></pic:cNvPr>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
