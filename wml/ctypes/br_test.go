package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapBreakXML(el *Break) *WrapperXML {
	return wrapXML(struct {
		*Break
		XMLName struct{} `xml:"w:br"`
	}{Break: el})
}

func TestBreak_MarshalXML(t *testing.T) {
	breakTypePage := stypes.BreakTypePage
	breakTypeColumn := stypes.BreakTypeColumn
	breakClearAll := stypes.BreakClearAll

	tests := []struct {
		name     string
		input    *Break
		expected string
	}{
		{
			name: "With all attributes",
			input: &Break{
				BreakType: &breakTypeColumn,
				Clear:     &breakClearAll,
			},
			expected: `<w:br w:type="column" w:clear="all"></w:br>`,
		}, {
			name:     "BreakTypePage",
			input:    NewBreak(breakTypePage),
			expected: `<w:br w:type="page"></w:br>`,
		}}

	for _, tt := range tests {
		object := wrapBreakXML(tt.input)
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
	br := NewBreak(breakTypePage)

	xmlData, err := xml.Marshal(br)
	if err != nil {
		t.Fatalf("Error marshaling Break to XML: %v", err)
	}

	var unmarshalledBreak Break
	err = xml.Unmarshal(xmlData, &unmarshalledBreak)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Break: %v", err)
	}

	if *unmarshalledBreak.BreakType != *br.BreakType {
		t.Errorf("Expected BreakType %s, got %s", *br.BreakType, *unmarshalledBreak.BreakType)
	}
}
