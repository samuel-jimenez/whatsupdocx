package dmlpic

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/stretchr/testify/suite"
	// "github.com/davecgh/go-spew/spew"
	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlprops"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
	"github.com/samuel-jimenez/whatsupdocx/dml/shapes"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
)

func wrapPicXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Pic
		XMLName struct{} `xml:"pic:pic"`
	}{Pic: el.(*Pic)})
}

func TestPic(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapPicXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "All Attributes",
			Input: &Pic{
				Attr: []xml.Attr{constants.NameSpaceDrawingMLPic},
				NonVisualPicProp: NonVisualPicProp{
					CNvPr: dmlct.CNvPr{
						ID:          1,
						Name:        "Pic 1",
						Description: "Description",
					},
					CNvPicPr: CNvPicPr{
						PicLocks: &dmlprops.PicLocks{
							NoChangeAspect:     dmlst.NewOptBool(true),
							NoChangeArrowheads: dmlst.NewOptBool(true),
						},
					},
				},
				BlipFill: BlipFill{
					Blip: &Blip{
						EmbedID: "rId1",
					},
					FillModeProps: &FillModeProps{
						Stretch: &shapes.Stretch{
							FillRect: &dmlct.RelativeRect{},
						},
					},
				},
				PicShapeProp: PicShapeProp{
					TransformGroup: &TransformGroup{
						Offset: &Offset{
							X: 0,
							Y: 0,
						},
						Extent: &dmlct.PSize2D{
							Width:  100000,
							Height: 100000,
						},
					},
					PresetGeometry: &PresetGeometry{
						Preset: "rect",
					},
				},
			},
			ExpectedXML: `<pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture"><pic:nvPicPr><pic:cNvPr id="1" name="Pic 1" descr="Description"></pic:cNvPr><pic:cNvPicPr><a:picLocks noChangeAspect="1" noChangeArrowheads="1"></a:picLocks></pic:cNvPicPr></pic:nvPicPr><pic:blipFill><a:blip r:embed="rId1"></a:blip><a:stretch><a:fillRect></a:fillRect></a:stretch></pic:blipFill><pic:spPr><a:xfrm><a:off x="0" y="0"></a:off><a:ext cx="100000" cy="100000"></a:ext></a:xfrm><a:prstGeom prst="rect"></a:prstGeom></pic:spPr></pic:pic>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
