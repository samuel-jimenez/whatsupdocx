package dml

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"
)

func TestNewEffectExtent(t *testing.T) {
	var left int64 = 10
	var top int64 = -2
	var right int64 = -1
	var bottom int64 = 10
	effExt := NewEffectExtent(left, top, right, bottom)

	if effExt.LeftEdge != left {
		t.Errorf("Expected left edge %d, but got %d", left, effExt.LeftEdge)
	}

	if effExt.TopEdge != top {
		t.Errorf("Expected top edge %d, but got %d", top, effExt.TopEdge)
	}

	if effExt.RightEdge != right {
		t.Errorf("Expected right edge %d, but got %d", right, effExt.RightEdge)
	}

	if effExt.BottomEdge != bottom {
		t.Errorf("Expected bottom edge %d, but got %d", bottom, effExt.BottomEdge)
	}
}

func wrapEffectExtentXML(el *EffectExtent) *WrapperXML {
	return wrapXML(struct {
		*EffectExtent
		XMLName struct{} `xml:"wp:effectExtent"`
	}{EffectExtent: el})
}

func TestEffectExtent_MarshalXML(t *testing.T) {
	tests := []struct {
		name        string
		input       *EffectExtent
		expectedXML string
	}{
		{
			name:        "NewEffectExtent",
			input:       NewEffectExtent(-27273042329600, -1, -2, 27273042316900),
			expectedXML: `<wp:effectExtent l="-27273042329600" t="-1" r="-2" b="27273042316900"></wp:effectExtent>`,
		},
		{
			name:        "Default",
			input:       &EffectExtent{},
			expectedXML: `<wp:effectExtent l="0" t="0" r="0" b="0"></wp:effectExtent>`,
		},
	}

	for _, tt := range tests {
		object := wrapEffectExtentXML(tt.input)
		expectedXML := wrapXMLOutput(tt.expectedXML)
		t.Run(tt.name, func(t *testing.T) {
			t.Run("MarshalXML", func(t *testing.T) {
				output, err := xml.Marshal(object)
				if err != nil {
					t.Fatalf("Error marshaling to XML: %v", err)
				}
				if got := string(output); got != expectedXML {
					t.Errorf("XML mismatch\nexpectedXML:\n%s\nActual:\n%s", expectedXML, got)
				}
			})
			t.Run("UnMarshalXML", func(t *testing.T) {
				object := tt.input
				expectedXML = tt.expectedXML
				vt := reflect.TypeOf(object)
				dest := reflect.New(vt.Elem()).Interface()
				err := xml.Unmarshal([]byte(expectedXML), dest)
				if err != nil {
					t.Fatalf("Error unmarshaling from XML: %v", err)
				}
				if got, want := dest, object; !reflect.DeepEqual(got, want) {
					t.Errorf("XML mismatch unmarshal(%s):\nexpectedXML:\n%#v\nActual:\n%#v", tt.expectedXML, want, got)
				}

			})
		})
	}
}

func TestEffectExtentXMLUnmarshal(t *testing.T) {
	tests := []struct {
		name   string
		xmlStr string
		effect EffectExtent
	}{
		{
			name:   "Empty XML",
			xmlStr: `<wp:extent></wp:extent>`,
			effect: EffectExtent{},
		},
		{
			name:   "XML with values",
			xmlStr: `<wp:extent l="27273042316900" t="20" r="-1" b="-27273042329600"></wp:extent>`,
			effect: EffectExtent{
				LeftEdge:   27273042316900,
				TopEdge:    20,
				RightEdge:  -1,
				BottomEdge: -27273042329600,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e EffectExtent
			err := xml.Unmarshal([]byte(tt.xmlStr), &e)
			if err != nil {
				t.Errorf("error unmarshalling XML: %v", err)
				return
			}
			if !reflect.DeepEqual(e, tt.effect) {
				t.Errorf("expected EffectExtent:\n%+v\ngot EffectExtent:\n%+v", tt.effect, e)
			}
		})
	}
}
