package ctypes

import (
	"reflect"
	"strings"
	"testing"

	"github.com/samuel-jimenez/xml"
)

func wrapGridChangeXML(el GridChange) *WrapperXML {
	return wrapXML(struct {
		GridChange
		XMLName struct{} `xml:"w:tblGridChange"`
	}{GridChange: el})
}

func TestGridChange_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    GridChange
		expected string
	}{
		{
			name:     "With ID",
			input:    GridChange{ID: 1},
			expected: `<w:tblGridChange w:id="1"></w:tblGridChange>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapGridChangeXML(tt.input))
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

func TestGridChange_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   GridChange
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "With ID",
			inputXML: `<w:tblGridChange w:id="1"></w:tblGridChange>`,
			expected: GridChange{ID: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := xml.NewDecoder(strings.NewReader(tt.inputXML))
			var result GridChange

			err := decoder.Decode(&result)

			if tt.expectFail {
				if err == nil {
					t.Error("Expected unmarshaling to fail but it did not")
				}
				return
			}

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unmarshaled GridChange struct does not match expected:\nExpected: %+v\nActual:   %+v", tt.expected, result)
			}
		})
	}
}
