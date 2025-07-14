package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapSectionPropXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*SectionProp
		XMLName struct{} `xml:"w:sectPr"`
	}{SectionProp: el.(*SectionProp)})
}

func TestSectionProp(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapSectionPropXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{

		{
			Name: "All attributes",
			Input: &SectionProp{
				HeaderReference: &HeaderFooterReference{Type: "default", ID: "rId1"},
				FooterReference: &HeaderFooterReference{Type: "default", ID: "rId2"},
				PageSize: &PageSize{
					Width:  uint64Ptr(12240),
					Height: uint64Ptr(15840),
				},
				Type:       NewGenSingleStrVal(stypes.SectionMarkNextPage),
				PageMargin: &PageMargin{Top: intPtr(1440), Bottom: intPtr(1440), Left: intPtr(1440), Right: intPtr(1440)},
				PageNum:    &PageNumbering{Format: stypes.NumFmtDecimal},
				FormProt:   NewGenSingleStrVal(stypes.OnOffTrue),
				TitlePg:    NewGenSingleStrVal(stypes.OnOffTrue),
				TextDir:    NewGenSingleStrVal(stypes.TextDirectionLrTb),
				DocGrid:    &DocGrid{Type: "default", LinePitch: intPtr(360)},
			},
			ExpectedXML: `<w:sectPr><w:headerReference r:id="rId1" w:type="default"></w:headerReference><w:footerReference r:id="rId2" w:type="default"></w:footerReference><w:type w:val="nextPage"></w:type><w:pgSz w:w="12240" w:h="15840"></w:pgSz><w:pgMar w:top="1440" w:right="1440" w:bottom="1440" w:left="1440"></w:pgMar><w:pgNumType w:fmt="decimal"></w:pgNumType><w:formProt w:val="true"></w:formProt><w:titlePg w:val="true"></w:titlePg><w:textDirection w:val="lrTb"></w:textDirection><w:docGrid w:type="default" w:linePitch="360"></w:docGrid></w:sectPr>`,
		},
		{
			Name:        "No attributes",
			Input:       &SectionProp{},
			ExpectedXML: `<w:sectPr></w:sectPr>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
