package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapExpaCompXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*ExpaComp
		XMLName struct{} `xml:"w:sz"`
	}{ExpaComp: el.(*ExpaComp)})
}

func TestExpaComp(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapExpaCompXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With value",
			Input:       &ExpaComp{Val: TextScalePtr(24)},
			ExpectedXML: `<w:sz w:val="24"></w:sz>`,
		},
		{
			Name:        "Without value",
			Input:       &ExpaComp{},
			ExpectedXML: `<w:sz></w:sz>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func TextScalePtr(value uint16) *stypes.TextScale {
	ts := stypes.TextScale(value)
	return &ts
}
