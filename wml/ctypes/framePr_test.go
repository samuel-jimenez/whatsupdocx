package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapFramePropXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*FrameProp
		XMLName struct{} `xml:"w:framePr"`
	}{FrameProp: el.(*FrameProp)})
}

func TestFrameProp(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapFramePropXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With all attributes",
			Input: &FrameProp{
				Width:      internal.ToPtr(int64(500)),
				Height:     internal.ToPtr(int64((300))),
				DropCap:    internal.ToPtr(stypes.DropCapMargin),
				Lines:      internal.ToPtr(3),
				VSpace:     internal.ToPtr(int64((50))),
				HSpace:     internal.ToPtr(int64((20))),
				Wrap:       internal.ToPtr(stypes.WrapAround),
				HAnchor:    internal.ToPtr(stypes.AnchorMargin),
				VAnchor:    internal.ToPtr(stypes.AnchorPage),
				AbsHPos:    internal.ToPtr(100),
				AbsVPos:    internal.ToPtr(200),
				XAlign:     internal.ToPtr(stypes.XAlignLeft),
				YAlign:     internal.ToPtr(stypes.YAlignCenter),
				HRule:      internal.ToPtr(stypes.HeightRuleExact),
				AnchorLock: internal.ToPtr(stypes.OnOffTrue),
			},
			ExpectedXML: `<w:framePr w:w="500" w:h="300" w:dropCap="margin" w:lines="3" w:hSpace="20" w:vSpace="50" ` +
				`w:wrap="around" w:hAnchor="margin" w:vAnchor="page" w:x="100" w:y="200" w:xAlign="left" w:yAlign="center" ` +
				`w:hRule="exact" w:anchorLock="true"></w:framePr>`,
		},
		{
			Name: "Without optional attributes",
			Input: &FrameProp{
				Width:  internal.ToPtr(int64(500)),
				Height: internal.ToPtr(int64(300)),
			},
			ExpectedXML: `<w:framePr w:w="500" w:h="300"></w:framePr>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
