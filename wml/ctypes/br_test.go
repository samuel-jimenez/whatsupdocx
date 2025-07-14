package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapBreakXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Break
		XMLName struct{} `xml:"w:br"`
	}{Break: el.(*Break)})
}

func TestBreak(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapBreakXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	breakTypePage := stypes.BreakTypePage
	breakTypeColumn := stypes.BreakTypeColumn
	breakClearAll := stypes.BreakClearAll
	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With all attributes",
			Input: &Break{
				BreakType: &breakTypeColumn,
				Clear:     &breakClearAll,
			},
			ExpectedXML: `<w:br w:type="column" w:clear="all"></w:br>`,
		}, {
			Name:        "BreakTypePage",
			Input:       NewBreak(breakTypePage),
			ExpectedXML: `<w:br w:type="page"></w:br>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
