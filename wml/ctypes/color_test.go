package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"
)

func wrapColorXML(el *Color) *WrapperXML {
	return wrapXML(struct {
		*Color
		XMLName struct{} `xml:"w:color"`
	}{Color: el})
}

func TestColor(t *testing.T) {
	testColor := NewColor("FF0000")
	tests := []struct {
		name     string
		input    *Color
		expected string
	}{{
		name:     "",
		input:    testColor,
		expected: `<w:color w:val="FF0000"></w:color>`,
	}}

	for _, tt := range tests {
		object := wrapColorXML(tt.input)
		expected := wrapXMLOutput(tt.expected)
		t.Run(tt.name, func(t *testing.T) {
			t.Run("MarshalXML", func(t *testing.T) {
				output, err := xml.Marshal(object)
				if err != nil {
					t.Fatalf("Error marshaling to XML: %v", err)
				}
				if got := string(output); got != expected {
					t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
				}
			}) //TODO UnmarshalXML_TOO
		})

	}

	//TODO
	xmlData, err := xml.Marshal(testColor)
	if err != nil {
		t.Fatalf("Error marshaling Color to XML: %v", err)
	}

	var unmarshaledColor Color
	err = xml.Unmarshal(xmlData, &unmarshaledColor)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Color: %v", err)
	}

	if testColor.Val != unmarshaledColor.Val {
		t.Errorf("Expected color value \n%s\n, got \n%s\n", testColor.Val, unmarshaledColor.Val)
	}

}
