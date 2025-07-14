package dmlpic

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlprops"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapNonVisualPicPropXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*NonVisualPicProp
		XMLName struct{} `xml:"pic:nvPicPr"`
	}{NonVisualPicProp: el.(*NonVisualPicProp)})
}

func TestNonVisualPicProp(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapNonVisualPicPropXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Input: &NonVisualPicProp{
				CNvPr: dmlct.CNvPr{
					ID:          1,
					Name:        "Pic1",
					Description: "Description of Pic1",
				},
				CNvPicPr: CNvPicPr{
					PicLocks: &dmlprops.PicLocks{
						NoChangeAspect: dmlst.NewOptBool(true),
					},
				},
			},
			ExpectedXML: `<pic:nvPicPr><pic:cNvPr id="1" name="Pic1" descr="Description of Pic1"></pic:cNvPr><pic:cNvPicPr><a:picLocks noChangeAspect="1"></a:picLocks></pic:cNvPicPr></pic:nvPicPr>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
