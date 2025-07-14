package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapCellMarginsXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*CellMargins
		XMLName struct{} `xml:"w:tblCellMar"`
	}{CellMargins: el.(*CellMargins)})
}

func TestCellMargins(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapCellMarginsXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	defaultCellMargins := DefaultCellMargins()
	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Input: &CellMargins{
				Top:    NewTableWidth(0, stypes.TableWidthDxa),
				Left:   NewTableWidth(55, stypes.TableWidthDxa),
				Bottom: NewTableWidth(0, stypes.TableWidthDxa),
				Right:  NewTableWidth(55, stypes.TableWidthDxa),
			},
			ExpectedXML: `<w:tblCellMar><w:top w:w="0" w:type="dxa"></w:top><w:left w:w="55" w:type="dxa"></w:left><w:bottom w:w="0" w:type="dxa"></w:bottom><w:right w:w="55" w:type="dxa"></w:right></w:tblCellMar>`,
		},
		{
			Input:       &defaultCellMargins,
			ExpectedXML: `<w:tblCellMar></w:tblCellMar>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
