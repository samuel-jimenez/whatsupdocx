package dmlpic

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapBlipXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Blip
		XMLName struct{} `xml:"a:blip"`
	}{Blip: el.(*Blip)})
}

func TestBlip(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapBlipXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "All Attributes",
			Input: &Blip{
				EmbedID: "rId0",
			},
			ExpectedXML: `<a:blip r:embed="rId0"></a:blip>`,
		},
		{
			Name:        "Minimal Attributes",
			Input:       &Blip{},
			ExpectedXML: `<a:blip></a:blip>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
