package dmlpic

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapPicShapePropXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*PicShapeProp
		XMLName struct{} `xml:"pic:spPr"`
	}{PicShapeProp: el.(*PicShapeProp)})
}

func TestPicShapeProp(t *testing.T) {
	bwMode := BlackWhiteModeGray

	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapPicShapePropXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With attributes",
			Input: &PicShapeProp{
				BwMode:         &bwMode,
				TransformGroup: &TransformGroup{},
				PresetGeometry: &PresetGeometry{
					Preset: "rect",
				},
			},
			ExpectedXML: `<pic:spPr bwMode="gray"><a:xfrm></a:xfrm><a:prstGeom prst="rect"></a:prstGeom></pic:spPr>`,
		},
		{
			Name:        "Empty",
			Input:       &PicShapeProp{},
			ExpectedXML: `<pic:spPr></pic:spPr>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
