package dmlpic

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/shapes"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapBlipFillXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*BlipFill
		XMLName struct{} `xml:"pic:blipFill"`
	}{BlipFill: el.(*BlipFill)})
}

func TestBlipFill(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapBlipFillXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "All Attributes",
			Input: &BlipFill{
				DPI:          internal.ToPtr(uint32(1)),
				RotWithShape: internal.ToPtr(true),
				Blip: &Blip{
					EmbedID: "rId0",
				},
				FillModeProps: &FillModeProps{
					Stretch: &shapes.Stretch{
						FillRect: &dmlct.RelativeRect{},
					},
				},
			},
			ExpectedXML: `<pic:blipFill dpi="1" rotWithShape="true"><a:blip r:embed="rId0"></a:blip><a:stretch><a:fillRect></a:fillRect></a:stretch></pic:blipFill>`,
		},
		{
			Name:        "Minimal Attributes",
			Input:       &BlipFill{},
			ExpectedXML: `<pic:blipFill></pic:blipFill>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
