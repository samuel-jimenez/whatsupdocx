package dmlct

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapOptBoolElemXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*OptBoolElem
		XMLName struct{} `xml:"w:b"`
	}{OptBoolElem: el.(*OptBoolElem)})
}

func TestOptBoolElem(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapOptBoolElemXML
	xmlTester.WrapXMLOutput = wrapXMLOutput
	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "Valid true",
			Input:       &OptBoolElem{Val: dmlst.NewOptBool(true)},
			ExpectedXML: `<w:b w:val="true"></w:b>`,
		},
		{
			Name:        "Valid false",
			Input:       &OptBoolElem{Val: dmlst.NewOptBool(false)},
			ExpectedXML: `<w:b w:val="false"></w:b>`,
		},
		{
			Name:        "Invalid",
			Input:       &OptBoolElem{Val: dmlst.OptBool{Valid: false}},
			ExpectedXML: `<w:b></w:b>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
