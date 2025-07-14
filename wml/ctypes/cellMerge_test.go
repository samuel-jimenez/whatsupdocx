package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapCellMergeXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*CellMerge
		XMLName struct{} `xml:"CellMerge"`
	}{CellMerge: el.(*CellMerge)})
}

func TestCellMerge(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapCellMergeXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "All Attributes",
			Input: &CellMerge{
				ID:         0,
				Author:     "John Smith",
				Date:       xmlStrPtr("2025-07-15"), // Helper function to get pointer to string
				VMerge:     internal.ToPtr(AnnotationVMergeCont),
				VMergeOrig: internal.ToPtr(AnnotationVMergeRest),
			},
			ExpectedXML: `<CellMerge w:id="0" w:author="John Smith" w:date="2025-07-15" w:vMerge="cont" w:vMergeOrig="rest"></CellMerge>`,
		},
		{
			Name: "vMergeOrig",
			Input: &CellMerge{
				ID:         1,
				Author:     "John Doe",
				Date:       nil,
				VMerge:     internal.ToPtr(AnnotationVMergeCont),
				VMergeOrig: internal.ToPtr(AnnotationVMergeRest),
			},
			ExpectedXML: `<CellMerge w:id="1" w:author="John Doe" w:vMerge="cont" w:vMergeOrig="rest"></CellMerge>`,
		},
		{
			Name: "date",
			Input: &CellMerge{
				ID:         2,
				Author:     "Jane Smith",
				Date:       xmlStrPtr("2024-06-25"), // Helper function to get pointer to string
				VMerge:     internal.ToPtr(AnnotationVMergeRest),
				VMergeOrig: nil,
			},
			ExpectedXML: `<CellMerge w:id="2" w:author="Jane Smith" w:date="2024-06-25" w:vMerge="rest"></CellMerge>`,
		},
		{
			Name: "Required Attributes",
			Input: &CellMerge{
				ID:     3,
				Author: "Tim Robbins",
			},
			ExpectedXML: `<CellMerge w:id="3" w:author="Tim Robbins"></CellMerge>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

// Helper function to return pointer to string
func xmlStrPtr(s string) *string {
	return &s
}
