package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapFitTextXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*FitText
		XMLName struct{} `xml:"w:fitText"`
	}{FitText: el.(*FitText)})
}

func TestFitText(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapFitTextXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With ID",
			Input:       &FitText{Val: 123, ID: IntPtr(456)},
			ExpectedXML: `<w:fitText w:val="123" w:id="456"></w:fitText>`,
		},
		{
			Name:        "Without ID",
			Input:       &FitText{Val: 789},
			ExpectedXML: `<w:fitText w:val="789"></w:fitText>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func IntPtr(i int) *int {
	return &i
}
