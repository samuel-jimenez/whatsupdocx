package dmlct

import (
	"testing"

	"github.com/samuel-jimenez/xml"
	"github.com/stretchr/testify/suite"

	"github.com/samuel-jimenez/whatsupdocx/common/units"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
)

func wrapPSize2DXML(el any, tag string) *testsuite.WrapperXML {
	return wrapXML(struct {
		*PSize2D
		XMLName xml.Name
	}{PSize2D: el.(*PSize2D), XMLName: xml.Name{Local: tag}})
}

func TestPSize2D(t *testing.T) {
	xmlTester := new(testsuite.XMLNamedTester)
	xmlTester.WrapXMLInput = wrapPSize2DXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "w:extent",
			Input:       NewPostvSz2D(units.Emu(100), units.Emu(200)),
			ExpectedXML: `<w:extent cx="100" cy="200"></w:extent>`,
			XMLName:     "w:extent",
		},
		{
			Name:        "a:ext",
			Input:       NewPostvSz2D(units.Emu(150), units.Emu(250)),
			ExpectedXML: `<a:ext cx="150" cy="250"></a:ext>`,
			XMLName:     "a:ext",
		},
		{
			Name: "a:extent",
			Input: &PSize2D{
				Width:  150,
				Height: 250,
			},
			ExpectedXML: `<a:extent cx="150" cy="250"></a:extent>`,
			XMLName:     "a:extent",
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

func TestNewPSize2D(t *testing.T) {
	width := units.Emu(100)
	height := units.Emu(200)
	extent := NewPostvSz2D(width, height)

	if extent.Width != uint64(width) {
		t.Errorf("Width does not match. Expected %d, got %d", width, extent.Width)
	}

	if extent.Height != uint64(height) {
		t.Errorf("Height does not match. Expected %d, got %d", height, extent.Height)
	}
}
