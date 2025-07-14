package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapTableLayoutXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*TableLayout
		XMLName struct{} `xml:"w:tblLayout"`
	}{TableLayout: el.(*TableLayout)})
}

func TestTableLayout(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapTableLayoutXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Input:       &TableLayout{LayoutType: internal.ToPtr(stypes.TableLayoutFixed)},
			ExpectedXML: `<w:tblLayout w:type="fixed"></w:tblLayout>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
