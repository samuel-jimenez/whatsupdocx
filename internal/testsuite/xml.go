package testsuite

import (
	"reflect"

	"github.com/samuel-jimenez/xml"

	"github.com/stretchr/testify/suite"
)

type WrapperXML struct {
	XMLName struct{}   `xml:"testwrapper"`
	Attr    []xml.Attr `xml:",any,attr,omitempty"`
	Element any
}

type XMLTestData struct {
	Name          string
	Input         any
	ExpectedXML   string
	XMLName       string
	UnmarshalOnly bool
}

type XMLUnMarshalTester struct {
	suite.Suite
	Stats *suite.SuiteInformation
	Tests []XMLTestData
}

func (suite *XMLUnMarshalTester) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	suite.Stats = stats
}

func (suite *XMLUnMarshalTester) TestUnMarshalXML() {

	for _, tt := range suite.Tests {
		suite.Run(tt.Name, func() {
			object := tt.Input
			expectedXML := tt.ExpectedXML
			vt := reflect.TypeOf(object)
			dest := reflect.New(vt.Elem()).Interface()
			err := xml.Unmarshal([]byte(expectedXML), dest)
			if err != nil {
				suite.Fail("Error unmarshaling from XML", err)
			}
			suite.Equal(object, dest)
		})
	}
}

type XMLTester struct {
	XMLUnMarshalTester
	WrapXMLInput  func(any) *WrapperXML
	WrapXMLOutput func(string) string
}

func (suite *XMLTester) TestMarshalXML() {
	for _, tt := range suite.Tests {
		if !tt.UnmarshalOnly {
			suite.Run(tt.Name, func() {
				object := suite.WrapXMLInput(tt.Input)
				expectedXML := suite.WrapXMLOutput(tt.ExpectedXML)
				output, err := xml.Marshal(object)
				if err != nil {
					suite.Fail("Error marshaling to XML", err)
				}
				suite.Equal(expectedXML, string(output))
			})
		}
	}
}

type XMLNamedTester struct {
	XMLUnMarshalTester
	WrapXMLInput  func(any, string) *WrapperXML
	WrapXMLOutput func(string) string
}

func (suite *XMLNamedTester) TestMarshalXML() {
	for _, tt := range suite.Tests {
		suite.Run(tt.Name, func() {
			object := suite.WrapXMLInput(tt.Input, tt.XMLName)
			expectedXML := suite.WrapXMLOutput(tt.ExpectedXML)
			output, err := xml.Marshal(object)
			if err != nil {
				suite.Fail("Error marshaling to XML", err)
			}
			suite.Equal(expectedXML, string(output))
		})
	}
}
