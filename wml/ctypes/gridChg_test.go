package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapGridChangeXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*GridChange
		XMLName struct{} `xml:"w:tblGridChange"`
	}{GridChange: el.(*GridChange)})
}

func TestGridChange(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapGridChangeXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With ID",
			Input:       &GridChange{ID: 1},
			ExpectedXML: `<w:tblGridChange w:id="1"></w:tblGridChange>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
