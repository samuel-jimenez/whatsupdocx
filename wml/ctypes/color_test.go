package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapColorXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Color
		XMLName struct{} `xml:"w:color"`
	}{Color: el.(*Color)})
}

func TestColor(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapColorXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	testColor := NewColor("FF0000")
	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Input:       testColor,
			ExpectedXML: `<w:color w:val="FF0000"></w:color>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
