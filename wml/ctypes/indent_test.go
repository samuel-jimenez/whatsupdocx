package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/stretchr/testify/suite"
)

func wrapIndentXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Indent
		XMLName struct{} `xml:"w:ind"`
	}{Indent: el.(*Indent)})
}

func TestIndent(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapIndentXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name: "With all attributes",
			Input: &Indent{
				Left:           internal.ToPtr(720),
				LeftChars:      internal.ToPtr(2),
				Right:          internal.ToPtr(360),
				RightChars:     internal.ToPtr(1),
				Hanging:        internal.ToPtr(uint64(360)),
				HangingChars:   internal.ToPtr(1),
				FirstLine:      internal.ToPtr(uint64(720)),
				FirstLineChars: internal.ToPtr(2),
			},
			ExpectedXML: `<w:ind w:left="720" w:leftChars="2" w:right="360" w:rightChars="1" w:hanging="360" w:hangingChars="1" w:firstLine="720" w:firstLineChars="2"></w:ind>`,
		},
		{
			Name:        "Without attributes",
			Input:       &Indent{},
			ExpectedXML: `<w:ind></w:ind>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
