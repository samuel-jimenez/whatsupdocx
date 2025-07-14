package ctypes

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapDocGridXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*DocGrid
		XMLName struct{} `xml:"w:docGrid"`
	}{DocGrid: el.(*DocGrid)})
}

func TestDocGrid(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapDocGridXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	linePitch := 240
	charSpace := 120
	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "All Attributes",
			Input: &DocGrid{
				Type:      stypes.DocGridLinesAndChars,
				LinePitch: &linePitch,
				CharSpace: &charSpace,
			},
			ExpectedXML: `<w:docGrid w:type="linesAndChars" w:linePitch="240" w:charSpace="120"></w:docGrid>`,
		}, {
			Name: "Minimal Attributes",
			Input: &DocGrid{
				Type: stypes.DocGridLines,
			},
			ExpectedXML: `<w:docGrid w:type="lines"></w:docGrid>`,
		}, {
			Name:        "No Attributes",
			Input:       &DocGrid{},
			ExpectedXML: `<w:docGrid></w:docGrid>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
