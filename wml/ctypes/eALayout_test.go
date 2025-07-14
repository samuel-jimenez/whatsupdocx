package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapEALayoutXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*EALayout
		XMLName struct{} `xml:"w:eastAsianLayout"`
	}{EALayout: el.(*EALayout)})
}

func TestEALayout(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapEALayoutXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "All attributes set",
			Input: &EALayout{
				ID:           internal.ToPtr(1),
				Combine:      internal.ToPtr(stypes.OnOffOn),
				CombineBrkts: internal.ToPtr(stypes.CombineBracketsRound),
				Vert:         internal.ToPtr(stypes.OnOffOff),
				VertCompress: internal.ToPtr(stypes.OnOffOn),
			},
			ExpectedXML: `<w:eastAsianLayout w:id="1" w:combine="on" w:combineBrackets="round" w:vert="off" w:vertCompress="on"></w:eastAsianLayout>`,
		},
		{
			Name: "Only ID set",
			Input: &EALayout{
				ID: internal.ToPtr(2),
			},
			ExpectedXML: `<w:eastAsianLayout w:id="2"></w:eastAsianLayout>`,
		},
		{
			Name: "Only Combine set",
			Input: &EALayout{
				Combine: internal.ToPtr(stypes.OnOffOn),
			},
			ExpectedXML: `<w:eastAsianLayout w:combine="on"></w:eastAsianLayout>`,
		},
		{
			Name: "Only CombineBrkts set",
			Input: &EALayout{
				CombineBrkts: internal.ToPtr(stypes.CombineBracketsSquare),
			},
			ExpectedXML: `<w:eastAsianLayout w:combineBrackets="square"></w:eastAsianLayout>`,
		},
		{
			Name: "Only Vert set",
			Input: &EALayout{
				Vert: internal.ToPtr(stypes.OnOffOff),
			},
			ExpectedXML: `<w:eastAsianLayout w:vert="off"></w:eastAsianLayout>`,
		},
		{
			Name: "Only VertCompress set",
			Input: &EALayout{
				VertCompress: internal.ToPtr(stypes.OnOffOn),
			},
			ExpectedXML: `<w:eastAsianLayout w:vertCompress="on"></w:eastAsianLayout>`,
		},
		{
			Name:        "No attributes set",
			Input:       &EALayout{},
			ExpectedXML: `<w:eastAsianLayout></w:eastAsianLayout>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
