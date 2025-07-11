package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"
)

func wrapFitTextXML(el FitText) *WrapperXML {
	return wrapXML(struct {
		FitText
		XMLName struct{} `xml:"w:fitText"`
	}{FitText: el})
}

func TestFitText_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FitText
		expected string
	}{
		{
			name:     "With ID",
			input:    FitText{Val: 123, ID: IntPtr(456)},
			expected: `<w:fitText w:val="123" w:id="456"></w:fitText>`,
		},
		{
			name:     "Without ID",
			input:    FitText{Val: 789},
			expected: `<w:fitText w:val="789"></w:fitText>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapFitTextXML(tt.input))
			expected := wrapXMLOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestFitText_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FitText
	}{
		{
			name:     "With ID",
			inputXML: `<w:fitText w:val="123" w:id="456"></w:fitText>`,
			expected: FitText{Val: 123, ID: IntPtr(456)},
		},
		{
			name:     "Without ID",
			inputXML: `<w:fitText w:val="789"></w:fitText>`,
			expected: FitText{Val: 789},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FitText

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %+v but got %+v", tt.expected, result)
			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}
