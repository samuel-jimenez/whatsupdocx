package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type RunFontsXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element RunFonts `xml:"w:val"`
}

func wrapRunFontsXML(el RunFonts) *RunFontsXML {
	return &RunFontsXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapRunFontsOutput(output string) string {
	return `<RunFontsXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</RunFontsXML>`
}

func TestRunFonts_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    RunFonts
		expected string
	}{
		{
			name: "Base",
			input: RunFonts{
				Hint:          stypes.FontTypeHintDefault,
				Ascii:         "Arial",
				HAnsi:         "Calibri",
				EastAsia:      "SimSun",
				CS:            "Arial",
				AsciiTheme:    stypes.ThemeFontMajorAscii,
				HAnsiTheme:    stypes.ThemeFontMajorHAnsi,
				EastAsiaTheme: stypes.ThemeFontMajorEastAsia,
				CSTheme:       stypes.ThemeFontMajorBidi,
			},
			expected: `<w:rFonts w:eastAsia="SimSun" w:hint="default" w:ascii="Arial" w:hAnsi="Calibri" w:cs="Arial" w:asciiTheme="majorAscii" w:hAnsiTheme="majorHAnsi" w:eastAsiaTheme="majorEastAsia" w:cstheme="majorBidi"></w:rFonts>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapRunFontsXML(tt.input))
			expected := wrapRunFontsOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestRunFontsUnmarshalXML(t *testing.T) {
	input := `<w:rFonts w:eastAsia="SimSun" w:hint="default" w:ascii="Arial" w:hAnsi="Calibri" w:cs="Arial" w:asciiTheme="majorAscii" w:hAnsiTheme="majorHAnsi" w:eastAsiaTheme="majorEastAsia" w:cstheme="majorBidi"></w:rFonts>`

	var rf RunFonts
	err := xml.Unmarshal([]byte(input), &rf)
	if err != nil {
		t.Fatalf("Error unmarshaling XML: %v", err)
	}

	expected := RunFonts{
		Hint:          stypes.FontTypeHintDefault,
		Ascii:         "Arial",
		HAnsi:         "Calibri",
		EastAsia:      "SimSun",
		CS:            "Arial",
		AsciiTheme:    stypes.ThemeFontMajorAscii,
		HAnsiTheme:    stypes.ThemeFontMajorHAnsi,
		EastAsiaTheme: stypes.ThemeFontMajorEastAsia,
		CSTheme:       stypes.ThemeFontMajorBidi,
	}

	if rf != expected {
		t.Errorf("Expected %+v but got %+v", expected, rf)
	}
}
