package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapColumnXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Column
		XMLName struct{} `xml:"w:gridCol"`
	}{Column: el.(*Column)})
}

func TestColumn(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapColumnXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With Width",
			Input:       &Column{Width: internal.ToPtr(uint64(500))},
			ExpectedXML: `<w:gridCol w:w="500"></w:gridCol>`,
		},
		{
			Name:        "Without Width",
			Input:       &Column{},
			ExpectedXML: `<w:gridCol></w:gridCol>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
