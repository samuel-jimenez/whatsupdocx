package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/internal"
)

func wrapCellMergeXML(el *CellMerge) *WrapperXML {
	return wrapXML(struct {
		*CellMerge
		XMLName struct{} `xml:"CellMerge"`
	}{CellMerge: el})
}

func TestCellMerge_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    *CellMerge
		expected string
	}{{
		name: "Default",
		input: &CellMerge{
			ID:         1,
			Author:     "John Doe",
			Date:       nil,
			VMerge:     internal.ToPtr(AnnotationVMergeCont),
			VMergeOrig: internal.ToPtr(AnnotationVMergeRest),
		},
		expected: `<CellMerge w:id="1" w:author="John Doe" w:vMerge="cont" w:vMergeOrig="rest"></CellMerge>`,
	}}

	for _, tt := range tests {
		object := wrapCellMergeXML(tt.input)
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
}

func TestCellMerge_UnmarshalXML(t *testing.T) {
	// Define a sample XML data corresponding to CellMerge structure
	xmlData := []byte(`<CellMerge w:id="2" w:author="Jane Smith" w:date="2024-06-25" w:vMerge="rest"></CellMerge>`)

	// Define the expected CellMerge instance after unmarshaling
	expectedCellMerge := &CellMerge{
		ID:         2,
		Author:     "Jane Smith",
		Date:       xmlStrPtr("2024-06-25"), // Helper function to get pointer to string
		VMerge:     internal.ToPtr(AnnotationVMergeRest),
		VMergeOrig: nil,
	}

	// Variable to hold unmarshaled CellMerge instance
	var unmarshaledCellMerge CellMerge

	// Unmarshal the XML into the CellMerge instance
	err := xml.Unmarshal(xmlData, &unmarshaledCellMerge)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to CellMerge: %v", err)
	}

	// Compare the unmarshaled CellMerge instance with the expected instance
	if !reflect.DeepEqual(&unmarshaledCellMerge, expectedCellMerge) {
		t.Errorf("Expected %#v, got %#v", expectedCellMerge, unmarshaledCellMerge)
	}
}

// Helper function to return pointer to string
func xmlStrPtr(s string) *string {
	return &s
}
