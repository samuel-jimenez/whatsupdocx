package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapGridXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Grid
		XMLName struct{} `xml:"w:tblGrid"`
	}{Grid: el.(*Grid)})
}

func TestGrid(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapGridXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With Columns and GridChange",
			Input: &Grid{
				Col: []Column{
					{Width: internal.ToPtr(uint64(500))},
					{Width: internal.ToPtr(uint64(750))},
				},
				GridChange: &GridChange{ID: 1},
			},
			ExpectedXML: `<w:tblGrid><w:gridCol w:w="500"></w:gridCol><w:gridCol w:w="750"></w:gridCol><w:tblGridChange w:id="1"></w:tblGridChange></w:tblGrid>`,
		},
		{
			Name: "With Columns, without GridChange",
			Input: &Grid{
				Col: []Column{
					{Width: internal.ToPtr(uint64(300))},
					{Width: internal.ToPtr(uint64(600))},
				},
			},
			ExpectedXML: `<w:tblGrid><w:gridCol w:w="300"></w:gridCol><w:gridCol w:w="600"></w:gridCol></w:tblGrid>`,
		},
		{
			Name:        "Empty Grid",
			Input:       &Grid{},
			ExpectedXML: `<w:tblGrid></w:tblGrid>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
