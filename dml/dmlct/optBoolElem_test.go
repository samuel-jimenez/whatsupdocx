package dmlct

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
)

type WrapperXML struct {
	XMLName struct{}   `xml:"testwrapper"`
	Attr    []xml.Attr `xml:",any,attr,omitempty"`
	Element any
}

func wrapXML(el any) *WrapperXML {
	return &WrapperXML{
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

func wrapOptBoolElemXML(el *OptBoolElem) *WrapperXML {
	return wrapXML(struct {
		*OptBoolElem
		XMLName struct{} `xml:"w:b"`
	}{OptBoolElem: el})
}

func TestOptBoolElem_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    *OptBoolElem
		expected string
	}{
		{
			name:     "Valid true",
			input:    &OptBoolElem{Val: dmlst.NewOptBool(true)},
			expected: `<w:b w:val="true"></w:b>`,
		},
		{
			name:     "Valid false",
			input:    &OptBoolElem{Val: dmlst.NewOptBool(false)},
			expected: `<w:b w:val="false"></w:b>`,
		},
		{
			name:     "Invalid",
			input:    &OptBoolElem{Val: dmlst.OptBool{Valid: false}},
			expected: `<w:b></w:b>`,
		},
	}

	for _, tt := range tests {
		object := wrapOptBoolElemXML(tt.input)
		expected := wrapXMLOutput(tt.expected)
		t.Run(tt.name, func(t *testing.T) {
			t.Run("MarshalXML", func(t *testing.T) {
				output, err := xml.Marshal(object)
				if err != nil {
					t.Fatalf("Error marshaling to XML: %v", err)
				}
				if got := string(output); got != expected {
					t.Errorf("XML mismatch\nExpected:\n%v\nActual:\n%v", expected, got)
				}
			})
			t.Run("UnMarshalXML", func(t *testing.T) {
				object := tt.input
				expected = tt.expected
				vt := reflect.TypeOf(object)
				dest := reflect.New(vt.Elem()).Interface()
				err := xml.Unmarshal([]byte(expected), dest)
				if err != nil {
					t.Fatalf("Error unmarshaling from XML: %v", err)
				}
				if got, want := dest, object; !reflect.DeepEqual(got, want) {
					t.Errorf("XML mismatch unmarshal(%s):\nExpected:\n%v\nActual:\n%v", tt.expected, want, got)
				}

			})
		})
	}
}
