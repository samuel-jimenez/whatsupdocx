package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"
	"github.com/stretchr/testify/suite"

	"github.com/samuel-jimenez/whatsupdocx/common"
	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
)

// !--- XML Testsuite wrappers ---!

func wrapXML(el any) *testsuite.WrapperXML {
	return &testsuite.WrapperXML{
		Attr: []xml.Attr{
			constants.NameSpaceWordprocessingML,
			constants.NameSpaceR,
		},
		Element: el,
	}
}

func wrapXMLOutput(output string) string {
	return `<testwrapper` +
		` xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"` +
		` xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"` +
		`>` + output + `</testwrapper>`
}

// !--- Tests for CTString start here ---!

func wrapCTStringXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*CTString
		XMLName struct{} `xml:"w:rStyle"`
	}{CTString: el.(*CTString)})
}

func TestCTString(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapCTStringXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{

		{
			Name:        "With value",
			Input:       &CTString{Val: "example"},
			ExpectedXML: `<w:rStyle w:val="example"></w:rStyle>`,
		},
		{
			Name:        "Empty value",
			Input:       &CTString{Val: ""},
			ExpectedXML: `<w:rStyle w:val=""></w:rStyle>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

// !--- Tests for CTString end here ---!

// !--- Tests for DecimalNum start here ---!

func wrapDecimalNumXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*DecimalNum
		XMLName struct{} `xml:"w:outlineLvl"`
	}{DecimalNum: el.(*DecimalNum)})
}

func TestDecimalNum(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapDecimalNumXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With positive value",
			Input:       &DecimalNum{Val: 10},
			ExpectedXML: `<w:outlineLvl w:val="10"></w:outlineLvl>`,
		},
		{
			Name:        "With negative  value",
			Input:       &DecimalNum{Val: -1},
			ExpectedXML: `<w:outlineLvl w:val="-1"></w:outlineLvl>`,
		},
		{
			Name:          "With leading zeroes",
			Input:         &DecimalNum{Val: 122},
			ExpectedXML:   `<w:outlineLvl w:val="00122"></w:outlineLvl>`,
			UnmarshalOnly: true,
		},
		{
			Name:          "With sign",
			Input:         &DecimalNum{Val: 3},
			ExpectedXML:   `<w:outlineLvl w:val="+3"></w:outlineLvl>`,
			UnmarshalOnly: true,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

// !--- Tests for DecimalNum end here ---!

// !--- Tests for Uint64 start here ---!

func wrapUint64ElemXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Uint64Elem
		XMLName struct{} `xml:"w:kern"`
	}{Uint64Elem: el.(*Uint64Elem)})
}

func TestUint64Elem(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapUint64ElemXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With value",
			Input:       &Uint64Elem{Val: 10},
			ExpectedXML: `<w:kern w:val="10"></w:kern>`,
		},
		{
			Name:        "Empty value",
			Input:       &Uint64Elem{Val: 18446744073709551615},
			ExpectedXML: `<w:kern w:val="18446744073709551615"></w:kern>`,
		},
		{
			Name:          "With leading zeroes",
			Input:         &Uint64Elem{Val: 122},
			ExpectedXML:   `<w:kern w:val="00122"></w:kern>`,
			UnmarshalOnly: true,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

// !--- Tests for Uint64 end here ---!

// !--- Tests for GenSingleStrVal start here ---!

func wrapGenSingleStrValXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*GenSingleStrVal[string]
		XMLName struct{} `xml:"GenSingleStrVal"`
	}{GenSingleStrVal: el.(*GenSingleStrVal[string])})
}

func TestGenSingleStrVal(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapGenSingleStrValXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{Name: "Test1", Input: NewGenSingleStrVal("Hello"), ExpectedXML: `<GenSingleStrVal w:val="Hello"></GenSingleStrVal>`},
		{Name: "Test2", Input: NewGenSingleStrVal("World"), ExpectedXML: `<GenSingleStrVal w:val="World"></GenSingleStrVal>`},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

// !--- Tests for GenSingleStrVal end here ---!

// !--- Tests for Empty start here ---!

func wrapEmptyXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*common.Empty
		XMLName struct{} `xml:"w:tab"`
	}{Empty: el.(*common.Empty)})
}

func TestEmpty(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapEmptyXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "Empty element",
			Input:       &common.Empty{},
			ExpectedXML: `<w:tab></w:tab>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

// !--- Tests for Empty end here ---!

// !--- Tests for Markup start here ---!

func wrapMarkupXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Markup
		XMLName struct{} `xml:"Markup"`
	}{Markup: el.(*Markup)})
}

func TestMarkup(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapMarkupXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "With ID",
			Input:       &Markup{ID: 42},
			ExpectedXML: `<Markup w:id="42"></Markup>`,
		},
		{
			Name:        "Zero ID",
			Input:       &Markup{ID: 0},
			ExpectedXML: `<Markup w:id="0"></Markup>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

// !--- Tests for Markup end here ---!

// !--- Tests for GenOptStrVal start here ---!

func wrapGenOptStrValXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*GenOptStrVal[string]
		XMLName struct{} `xml:"element"`
	}{GenOptStrVal: el.(*GenOptStrVal[string])})
}

func TestGenOptStrVal(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapGenOptStrValXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "WithValue",
			Input:       &GenOptStrVal[string]{Val: internal.ToPtr("test")},
			ExpectedXML: `<element w:val="test"></element>`,
		},
		{
			Name:        "WithNilValue",
			Input:       &GenOptStrVal[string]{Val: nil},
			ExpectedXML: `<element></element>`,
		},
		{
			Name:        "EmptyValue",
			Input:       &GenOptStrVal[string]{Val: internal.ToPtr("")},
			ExpectedXML: `<element w:val=""></element>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}

// !--- Tests for GenOptStrVal[string] end here ---!
