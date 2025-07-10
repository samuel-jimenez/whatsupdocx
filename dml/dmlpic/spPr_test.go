package dmlpic

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"
)

func wrapPicShapePropXML(el *PicShapeProp) *WrapperXML {
	return wrapXML(struct {
		*PicShapeProp
		XMLName struct{} `xml:"pic:spPr"`
	}{PicShapeProp: el})
}

func TestPicShapeProp_MarshalXML(t *testing.T) {
	bwMode := BlackWhiteModeGray

	tests := []struct {
		name     string
		input    *PicShapeProp
		expected string
	}{{
		name: "With attributes",
		input: &PicShapeProp{
			BwMode:         &bwMode,
			TransformGroup: &TransformGroup{},
			PresetGeometry: &PresetGeometry{
				Preset: "rect",
			},
		},
		expected: `<pic:spPr bwMode="gray"><a:xfrm></a:xfrm><a:prstGeom prst="rect"></a:prstGeom></pic:spPr>`,
	},
		{
			name:     "Empty",
			input:    &PicShapeProp{},
			expected: `<pic:spPr></pic:spPr>`,
		},
	}

	for _, tt := range tests {
		object := wrapPicShapePropXML(tt.input)
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

func TestUnmarshalPicShapeProp(t *testing.T) {
	bwMode := BlackWhiteModeGray

	tests := []struct {
		inputXML       string
		expectedResult PicShapeProp
	}{
		{
			inputXML: `<pic:spPr bwMode="gray"><a:xfrm></a:xfrm><a:prstGeom></a:prstGeom></pic:spPr>`,
			expectedResult: PicShapeProp{
				BwMode:         &bwMode,
				TransformGroup: &TransformGroup{},
				PresetGeometry: &PresetGeometry{},
			},
		},
		{
			inputXML:       `<pic:spPr></pic:spPr>`,
			expectedResult: PicShapeProp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var result PicShapeProp

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.BwMode == nil && tt.expectedResult.BwMode != nil {
				t.Errorf("Expected bwMode to be %v, but got nil", *tt.expectedResult.BwMode)
			} else if result.BwMode != nil && tt.expectedResult.BwMode == nil {
				t.Errorf("Expected bwMode to be nil, but got %v", *result.BwMode)
			} else if result.BwMode != nil && tt.expectedResult.BwMode != nil && *result.BwMode != *tt.expectedResult.BwMode {
				t.Errorf("Expected bwMode %v, but got %v", *tt.expectedResult.BwMode, *result.BwMode)
			}

			// Check TransformGroup and PresetGeometry
			if (result.TransformGroup == nil) != (tt.expectedResult.TransformGroup == nil) {
				t.Errorf("Expected TransformGroup to be %v, but got %v", tt.expectedResult.TransformGroup, result.TransformGroup)
			}
			if (result.PresetGeometry == nil) != (tt.expectedResult.PresetGeometry == nil) {
				t.Errorf("Expected PresetGeometry to be %v, but got %v", tt.expectedResult.PresetGeometry, result.PresetGeometry)
			}
		})
	}
}
