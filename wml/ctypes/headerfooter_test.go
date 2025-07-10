package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapHeaderReferenceXML(el HeaderFooterReference) *WrapperXML {
	return wrapXML(struct {
		HeaderFooterReference
		XMLName struct{} `xml:"w:headerReference"`
	}{HeaderFooterReference: el})
}

func wrapFooterReferenceXML(el HeaderFooterReference) *WrapperXML {
	return wrapXML(struct {
		HeaderFooterReference
		XMLName struct{} `xml:"w:footerReference"`
	}{HeaderFooterReference: el})
}

func TestHeaderReference_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    HeaderFooterReference
		expected string
	}{
		{
			name: "Marshal with ID and Type",
			input: HeaderFooterReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
			expected: `<w:headerReference r:id="rId1" w:type="first"></w:headerReference>`,
		},
		{
			name: "Marshal with Type only",
			input: HeaderFooterReference{
				Type: stypes.HdrFtrEven,
			},
			expected: `<w:headerReference w:type="even"></w:headerReference>`,
		},
		{
			name: "Marshal with ID only",
			input: HeaderFooterReference{
				ID: "rId2",
			},
			expected: `<w:headerReference r:id="rId2"></w:headerReference>`,
		},
		{
			name:     "Marshal with neither ID nor Type",
			input:    HeaderFooterReference{},
			expected: `<w:headerReference></w:headerReference>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapHeaderReferenceXML(tt.input))
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

func TestHeaderReference_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected HeaderFooterReference
	}{
		{
			name:     "Unmarshal with ID and Type",
			inputXML: `<w:headerReference w:type="first" r:id="rId1"></w:headerReference>`,
			expected: HeaderFooterReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
		},
		{
			name:     "Unmarshal with Type only",
			inputXML: `<w:headerReference w:type="even"></w:headerReference>`,
			expected: HeaderFooterReference{
				Type: stypes.HdrFtrEven,
			},
		},
		{
			name:     "Unmarshal with ID only",
			inputXML: `<w:headerReference r:id="rId2"></w:headerReference>`,
			expected: HeaderFooterReference{
				ID: "rId2",
			},
		},
		{
			name:     "Unmarshal with neither ID nor Type",
			inputXML: `<w:headerReference></w:headerReference>`,
			expected: HeaderFooterReference{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result HeaderFooterReference

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			if result.ID != tt.expected.ID || result.Type != tt.expected.Type {
				t.Errorf("Expected %+v but got %+v", tt.expected, result)
			}
		})
	}
}

func TestFooterReference_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    HeaderFooterReference
		expected string
	}{
		{
			name: "Marshal with ID and Type",
			input: HeaderFooterReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
			expected: `<w:footerReference r:id="rId1" w:type="first"></w:footerReference>`,
		},
		{
			name: "Marshal with Type only",
			input: HeaderFooterReference{
				Type: stypes.HdrFtrEven,
			},
			expected: `<w:footerReference w:type="even"></w:footerReference>`,
		},
		{
			name: "Marshal with ID only",
			input: HeaderFooterReference{
				ID: "rId2",
			},
			expected: `<w:footerReference r:id="rId2"></w:footerReference>`,
		},
		{
			name:     "Marshal with neither ID nor Type",
			input:    HeaderFooterReference{},
			expected: `<w:footerReference></w:footerReference>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapFooterReferenceXML(tt.input))
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

func TestFooterReference_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected HeaderFooterReference
	}{
		{
			name:     "Unmarshal with ID and Type",
			inputXML: `<w:footerReference w:type="first" r:id="rId1"></w:footerReference>`,
			expected: HeaderFooterReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
		},
		{
			name:     "Unmarshal with Type only",
			inputXML: `<w:footerReference w:type="even"></w:footerReference>`,
			expected: HeaderFooterReference{
				Type: stypes.HdrFtrEven,
			},
		},
		{
			name:     "Unmarshal with ID only",
			inputXML: `<w:footerReference r:id="rId2"></w:footerReference>`,
			expected: HeaderFooterReference{
				ID: "rId2",
			},
		},
		{
			name:     "Unmarshal with neither ID nor Type",
			inputXML: `<w:footerReference></w:footerReference>`,
			expected: HeaderFooterReference{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result HeaderFooterReference

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			if result.ID != tt.expected.ID || result.Type != tt.expected.Type {
				t.Errorf("Expected %+v but got %+v", tt.expected, result)
			}
		})
	}
}
