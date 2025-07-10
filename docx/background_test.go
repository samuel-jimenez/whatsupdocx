package docx

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// TODO move
type WrapperXML struct {
	XMLName struct{}   `xml:"testwrapper"`
	Attr    []xml.Attr `xml:",any,attr,omitempty"`
	Element any
}

func wrapXML(el any) *WrapperXML {
	return &WrapperXML{
		Attr: []xml.Attr{constants.NameSpaceWordprocessingML,
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

func wrapBackgroundXML(el *Background) *WrapperXML {
	return wrapXML(struct {
		*Background
		XMLName struct{} `xml:"w:background"`
	}{Background: el})
}

func TestBackground_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    *Background
		expected string
	}{
		{
			name: "With all attributes",
			input: &Background{
				Color:      StringPtr("FFFFFF"),
				ThemeColor: ThemeColorPtr(stypes.ThemeColorAccent1),
				ThemeTint:  StringPtr("500"),
				ThemeShade: StringPtr("200"),
			},
			expected: `<w:background w:color="FFFFFF" w:themeColor="accent1" w:themeTint="500" w:themeShade="200"></w:background>`,
		},
		{
			name: "Without optional attributes",
			input: &Background{
				Color: StringPtr("000000"),
			},
			expected: `<w:background w:color="000000"></w:background>`,
		},
	}

	for _, tt := range tests {
		object := wrapBackgroundXML(tt.input)
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
			// //TODO
			// 	//TestDocGrid_UnmarshalXML_NoAttributes     Expected:
			//              // &{{} [{{ xmlns:w} http://schemas.openxmlformats.org/wordprocessingml/2006/main} {{ xmlns:r} http://schemas.openxmlformats.org/officeDocument/2006/relationships}] {{ %!s(*int=<nil>) %!s(*int=<nil>)} {}}}
			//              // Actual:
			//              // &{{} [{{http://www.w3.org/2000/xmlns/ w} http://schemas.openxmlformats.org/wordprocessingml/2006/main} {{http://www.w3.org/2000/xmlns/ r} http://schemas.openxmlformats.org/officeDocument/2006/relationships}] <nil>}
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

func TestBackground_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Background
	}{
		{
			name:     "With all attributes",
			inputXML: `<w:background w:color="FFFFFF" w:themeColor="accent1" w:themeTint="500" w:themeShade="200"></w:background>`,
			expected: Background{
				Color:      StringPtr("FFFFFF"),
				ThemeColor: ThemeColorPtr(stypes.ThemeColorAccent1),
				ThemeTint:  StringPtr("500"),
				ThemeShade: StringPtr("200"),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<w:background w:color="000000"></w:background>`,
			expected: Background{
				Color: StringPtr("000000"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Background

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if tt.expected.Color != nil {
				if result.Color == nil {
					t.Errorf("Expected Color %s but got nil", *tt.expected.Color)
				} else if *result.Color != *tt.expected.Color {
					t.Errorf("Expected Color %s but got %s", *tt.expected.Color, *result.Color)
				}
			} else if result.Color != nil {
				t.Errorf("Expected nil but got %s", *result.Color)
			}

			if tt.expected.ThemeColor != nil {
				if result.ThemeColor == nil {
					t.Errorf("Expected ThemeColor %s but got nil", *tt.expected.ThemeColor)
				} else if *result.ThemeColor != *tt.expected.ThemeColor {
					t.Errorf("Expected ThemeColor %s but got %s", *tt.expected.ThemeColor, *result.ThemeColor)
				}
			} else if result.ThemeColor != nil {
				t.Errorf("Expected nil but got %s", *result.ThemeColor)
			}

			if tt.expected.ThemeTint != nil {
				if result.ThemeTint == nil {
					t.Errorf("Expected ThemeTint %s but got nil", *tt.expected.ThemeTint)
				} else if *result.ThemeTint != *tt.expected.ThemeTint {
					t.Errorf("Expected ThemeTint %s but got %s", *tt.expected.ThemeTint, *result.ThemeTint)
				}
			} else if result.ThemeTint != nil {
				t.Errorf("Expected nil but got %s", *result.ThemeTint)
			}

			if tt.expected.ThemeShade != nil {
				if result.ThemeShade == nil {
					t.Errorf("Expected ThemeShade %s but got nil", *tt.expected.ThemeShade)
				} else if *result.ThemeShade != *tt.expected.ThemeShade {
					t.Errorf("Expected ThemeShade %s but got %s", *tt.expected.ThemeShade, *result.ThemeShade)
				}
			} else if result.ThemeShade != nil {
				t.Errorf("Expected nil but got %s", *result.ThemeShade)
			}
		})
	}
}

func StringPtr(s string) *string {
	return &s
}

func ThemeColorPtr(t stypes.ThemeColor) *stypes.ThemeColor {
	return &t
}
