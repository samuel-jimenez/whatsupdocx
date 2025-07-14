package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapCellPropertyXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*CellProperty
		XMLName struct{} `xml:"w:tcPr"`
	}{CellProperty: el.(*CellProperty)})
}

func TestCellProperty(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapCellPropertyXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "Empty CellProperty",
			Input:       &CellProperty{},
			ExpectedXML: `<w:tcPr></w:tcPr>`,
		},
		{
			Name: "With Width",
			Input: &CellProperty{
				Width: NewTableWidth(100, stypes.TableWidthDxa),
			},
			ExpectedXML: `<w:tcPr><w:tcW w:w="100" w:type="dxa"></w:tcW></w:tcPr>`,
		},
		{
			Name: "With Shading",
			Input: &CellProperty{
				Shading: &Shading{Val: "clear"},
			},
			ExpectedXML: `<w:tcPr><w:shd w:val="clear"></w:shd></w:tcPr>`,
		},
		{
			Name: "With All Fields",
			Input: &CellProperty{
				CnfStyle:      &CTString{Val: "TestCnfStyle"},
				Width:         NewTableWidth(100, stypes.TableWidthDxa),
				GridSpan:      &DecimalNum{Val: 2},
				HMerge:        NewGenOptStrVal(stypes.MergeCellContinue),
				VMerge:        NewGenOptStrVal(stypes.MergeCellRestart),
				Borders:       &CellBorders{},
				Shading:       &Shading{Val: "clear"},
				NoWrap:        &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				Margins:       &CellMargins{},
				TextDirection: NewGenSingleStrVal(stypes.TextDirectionBtLr),
				FitText:       &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				VAlign:        NewGenSingleStrVal(stypes.VerticalJcCenter),
				HideMark:      &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				PrChange:      &TCPrChange{ID: 1, Author: "Author", Date: nil},
			},
			ExpectedXML: `<w:tcPr>` +
				`<w:cnfStyle w:val="TestCnfStyle"></w:cnfStyle>` +
				`<w:tcW w:w="100" w:type="dxa"></w:tcW>` +
				`<w:gridSpan w:val="2"></w:gridSpan>` +
				`<w:hMerge w:val="continue"></w:hMerge>` +
				`<w:vMerge w:val="restart"></w:vMerge>` +
				`<w:tcBorders></w:tcBorders>` +
				`<w:shd w:val="clear"></w:shd>` +
				`<w:noWrap w:val="true"></w:noWrap>` +
				`<w:tcMar></w:tcMar>` +
				`<w:textDirection w:val="btLr"></w:textDirection>` +
				`<w:tcFitText w:val="true"></w:tcFitText>` +
				`<w:vAlign w:val="center"></w:vAlign>` +
				`<w:hideMark w:val="true"></w:hideMark>` +
				`<w:tcPrChange w:id="1" w:author="Author"><w:tcPr></w:tcPr></w:tcPrChange>` +
				`</w:tcPr>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
