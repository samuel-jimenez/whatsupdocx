package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapDocGridXML(el DocGrid) *WrapperXML {
	return wrapXML(struct {
		DocGrid
		XMLName struct{} `xml:"w:docGrid"`
	}{DocGrid: el})
}

func TestDocGrid_MarshalXML(t *testing.T) {
	linePitch := 240
	charSpace := 120
	tests := []struct {
		name     string
		input    DocGrid
		expected string
	}{{
		name: "AllAttributes",
		input: DocGrid{
			Type:      stypes.DocGridLinesAndChars,
			LinePitch: &linePitch,
			CharSpace: &charSpace,
		},
		expected: `<w:docGrid w:type="linesAndChars" w:linePitch="240" w:charSpace="120"></w:docGrid>`,
	}, {
		name: "OmitEmptyAttributes",
		input: DocGrid{
			Type: stypes.DocGridLines,
		},
		expected: `<w:docGrid w:type="lines"></w:docGrid>`,
	}, {
		name:     "NoAttributes",
		input:    DocGrid{},
		expected: `<w:docGrid></w:docGrid>`,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			object := wrapDocGridXML(tt.input)
			expected := wrapXMLOutput(tt.expected)

			t.Run("MarshalXML", func(t *testing.T) {
				output, err := xml.Marshal(object)
				if err != nil {
					t.Fatalf("Error marshaling to XML: %v", err)
				}
				if got := string(output); got != expected {
					t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
				}
			}) //TODO UnmarshalXML_TOO
			// //TODO
			// 	//TestDocGrid_UnmarshalXML_NoAttributes     Expected:
			//              // &{{} [{{ xmlns:w} http://schemas.openxmlformats.org/wordprocessingml/2006/main} {{ xmlns:r} http://schemas.openxmlformats.org/officeDocument/2006/relationships}] {{ %!s(*int=<nil>) %!s(*int=<nil>)} {}}}
			//              // Actual:
			//              // &{{} [{{http://www.w3.org/2000/xmlns/ w} http://schemas.openxmlformats.org/wordprocessingml/2006/main} {{http://www.w3.org/2000/xmlns/ r} http://schemas.openxmlformats.org/officeDocument/2006/relationships}] <nil>}
			// t.Run("UnMarshalXML", func(t *testing.T) {
			// 	vt := reflect.TypeOf(object)
			// 	dest := reflect.New(vt.Elem()).Interface()
			// 	err := xml.Unmarshal([]byte(expected), dest)
			// 	if err != nil {
			// 		t.Fatalf("Error unmarshaling from XML: %v", err)
			// 	}
			// 	if got, want := dest, object; !reflect.DeepEqual(got, want) {
			// 		t.Errorf("XML mismatch unmarshal(%s):\nExpected:\n%s\nActual:\n%s", tt.expected, want, got)
			// 	}
			//
			// })
		})
	}

}

func TestDocGrid_UnmarshalXML_AllAttributes(t *testing.T) {
	xmlInput := `<w:docGrid xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" w:type="linesAndChars" w:linePitch="240" w:charSpace="120"></w:docGrid>`

	linePitch := 240
	charSpace := 120
	expectedDocGrid := DocGrid{
		Type:      stypes.DocGridLinesAndChars,
		LinePitch: &linePitch,
		CharSpace: &charSpace,
	}

	checkUnmarshalXML(t, xmlInput, expectedDocGrid)
}

func TestDocGrid_UnmarshalXML_MinimalAttributes(t *testing.T) {
	xmlInput := `<w:docGrid xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" w:type="lines"></w:docGrid>`

	expectedDocGrid := DocGrid{
		Type: stypes.DocGridLines,
	}

	checkUnmarshalXML(t, xmlInput, expectedDocGrid)
}

func TestDocGrid_UnmarshalXML_NoAttributes(t *testing.T) {
	xmlInput := `<w:docGrid xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"></w:docGrid>`

	expectedDocGrid := DocGrid{}

	checkUnmarshalXML(t, xmlInput, expectedDocGrid)
}

func checkUnmarshalXML(t *testing.T, xmlInput string, expectedDocGrid DocGrid) {
	t.Helper()

	var docGrid DocGrid

	err := xml.Unmarshal([]byte(xmlInput), &docGrid)
	if err != nil {
		t.Fatalf("Error unmarshaling XML: %v", err)
	}

	if docGrid.Type != expectedDocGrid.Type {
		t.Errorf("Expected Type %s but got %s", expectedDocGrid.Type, docGrid.Type)
	}

	if expectedDocGrid.LinePitch == nil {
		if docGrid.LinePitch != nil {
			t.Errorf("Expected LinePitch to be nil but got %d", *docGrid.LinePitch)
		}
	} else {
		if docGrid.LinePitch == nil {
			t.Errorf("Expected LinePitch %d but got nil", *expectedDocGrid.LinePitch)
		} else if *docGrid.LinePitch != *expectedDocGrid.LinePitch {
			t.Errorf("Expected LinePitch %d but got %d", *expectedDocGrid.LinePitch, *docGrid.LinePitch)
		}
	}

	if expectedDocGrid.CharSpace == nil {
		if docGrid.CharSpace != nil {
			t.Errorf("Expected CharSpace to be nil but got %d", *docGrid.CharSpace)
		}
	} else {
		if docGrid.CharSpace == nil {
			t.Errorf("Expected CharSpace %d but got nil", *expectedDocGrid.CharSpace)
		} else if *docGrid.CharSpace != *expectedDocGrid.CharSpace {
			t.Errorf("Expected CharSpace %d but got %d", *expectedDocGrid.CharSpace, *docGrid.CharSpace)
		}
	}
}
