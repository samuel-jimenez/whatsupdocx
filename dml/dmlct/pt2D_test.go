package dmlct

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/xml"
	"github.com/stretchr/testify/suite"
)

func TestNewPoint2D(t *testing.T) {
	x := int64(100)
	y := int64(200)
	start := NewPoint2D(x, y)

	if start.XAxis != x {
		t.Errorf("NewPoint2D() failed: Expected XAxis %d, got %d", x, start.XAxis)
	}

	if start.YAxis != y {
		t.Errorf("NewPoint2D() failed: Expected YAxis %d, got %d", y, start.YAxis)
	}
}

func wrapPoint2DXML(el any, tag string) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Point2D
		XMLName xml.Name
	}{Point2D: el.(*Point2D), XMLName: xml.Name{Local: tag}})
}

func TestPoint2D(t *testing.T) {
	xmlTester := new(testsuite.XMLNamedTester)
	xmlTester.WrapXMLInput = wrapPoint2DXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "wp:start",
			Input: &Point2D{
				XAxis: 100,
				YAxis: 200,
			},
			ExpectedXML: `<wp:start x="100" y="200"></wp:start>`,
			XMLName:     "wp:start",
		},
		{
			Name: "negative",
			Input: &Point2D{
				XAxis: -100,
				YAxis: -200,
			},
			ExpectedXML: `<wp:start x="-100" y="-200"></wp:start>`,
			XMLName:     "wp:start",
		},

		{
			Name: "wp:lineTo",
			Input: &Point2D{
				XAxis: 0,
				YAxis: 90,
			},
			ExpectedXML: `<wp:lineTo x="0" y="90"></wp:lineTo>`,
			XMLName:     "wp:lineTo",
		},
		{
			Name:        "origin",
			Input:       internal.ToPtr(NewPoint2D(150, 250)),
			ExpectedXML: `<origin x="150" y="250"></origin>`,
			XMLName:     "origin",
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
