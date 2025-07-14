package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapLatentStyleXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*LatentStyle
		XMLName struct{} `xml:"w:latentStyles"`
	}{LatentStyle: el.(*LatentStyle)})
}

func TestLatentStyle(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapLatentStyleXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "All attributes set",
			Input: &LatentStyle{
				DefLockedState:    internal.ToPtr(stypes.OnOffOn),
				DefUIPriority:     internal.ToPtr(99),
				DefSemiHidden:     internal.ToPtr(stypes.OnOffOn),
				DefUnhideWhenUsed: internal.ToPtr(stypes.OnOffOn),
				DefQFormat:        internal.ToPtr(stypes.OnOffOn),
				Count:             internal.ToPtr(3),
				LsdExceptions: []LsdException{
					{
						Name:           "Heading1",
						Locked:         internal.ToPtr(stypes.OnOffOn),
						UIPriority:     internal.ToPtr(99),
						SemiHidden:     internal.ToPtr(stypes.OnOffOn),
						UnhideWhenUsed: internal.ToPtr(stypes.OnOffOn),
						QFormat:        internal.ToPtr(stypes.OnOffOn),
					},
					{
						Name: "Heading2",
					},
				},
			},
			ExpectedXML: `<w:latentStyles w:defLockedState="on" w:defUIPriority="99" w:defSemiHidden="on" w:defUnhideWhenUsed="on" w:defQFormat="on" w:count="3"><w:lsdException w:name="Heading1" w:locked="on" w:uiPriority="99" w:semiHidden="on" w:unhideWhenUsed="on" w:qFormat="on"></w:lsdException><w:lsdException w:name="Heading2"></w:lsdException></w:latentStyles>`,
		},
		{
			Name: "Only required attributes set",
			Input: &LatentStyle{
				LsdExceptions: []LsdException{
					{
						Name: "Heading1",
					},
				},
			},
			ExpectedXML: `<w:latentStyles><w:lsdException w:name="Heading1"></w:lsdException></w:latentStyles>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
