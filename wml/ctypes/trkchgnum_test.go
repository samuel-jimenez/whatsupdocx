package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapTrackChangeNumXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*TrackChangeNum
		XMLName struct{} `xml:"TrackChangeNum"`
	}{TrackChangeNum: el.(*TrackChangeNum)})
}

func TestTrackChangeNum(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapTrackChangeNumXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With all attributes",
			Input: &TrackChangeNum{
				ID:       123,
				Author:   "John Doe",
				Date:     internal.ToPtr("2023-06-18T12:34:56Z"),
				Original: internal.ToPtr("42"),
			},
			ExpectedXML: `<TrackChangeNum w:id="123" w:author="John Doe" w:date="2023-06-18T12:34:56Z" w:original="42"></TrackChangeNum>`,
		},
		{
			Name: "Without optional attributes",
			Input: &TrackChangeNum{
				ID:     124,
				Author: "Jane Doe",
			},
			ExpectedXML: `<TrackChangeNum w:id="124" w:author="Jane Doe"></TrackChangeNum>`,
		},
		{
			Name: "With only date attribute",
			Input: &TrackChangeNum{
				ID:     125,
				Author: "Alice",
				Date:   internal.ToPtr("2024-06-18T12:34:56Z"),
			},
			ExpectedXML: `<TrackChangeNum w:id="125" w:author="Alice" w:date="2024-06-18T12:34:56Z"></TrackChangeNum>`,
		},
		{
			Name: "With only original attribute",
			Input: &TrackChangeNum{
				ID:       126,
				Author:   "Bob",
				Original: internal.ToPtr("99"),
			},
			ExpectedXML: `<TrackChangeNum w:id="126" w:author="Bob" w:original="99"></TrackChangeNum>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
