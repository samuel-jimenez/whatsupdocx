package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapEffectXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Effect
		XMLName struct{} `xml:"w:effect"`
	}{Effect: el.(*Effect)})
}

func TestEffect(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapEffectXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With value",
			Input:       &Effect{Val: TextEffectPtr(stypes.TextEffectBlinkBackground)},
			ExpectedXML: `<w:effect w:val="blinkBackground"></w:effect>`,
		},
		{
			Name:        "Without value",
			Input:       &Effect{},
			ExpectedXML: `<w:effect></w:effect>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func TextEffectPtr(value stypes.TextEffect) *stypes.TextEffect {
	return &value
}
