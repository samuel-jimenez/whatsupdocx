package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type ShadingXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element Shading  `xml:"w:shd"`
}

func wrapShadingXML(el Shading) *ShadingXML {
	return &ShadingXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapShadingOutput(output string) string {
	return `<ShadingXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</ShadingXML>`
}

func TestShd_MarshalXML(t *testing.T) {
	themeColor1 := stypes.ThemeColorAccent2
	themeFill1 := stypes.ThemeColorAccent1
	tests := []struct {
		name     string
		input    Shading
		expected string
	}{
		{
			name: "Basic Shd with Val",
			input: Shading{
				Val: stypes.ShdClear,
			},
			expected: `<w:shd w:val="clear"></w:shd>`,
		},
		{
			name: "Shd with all attributes",
			input: Shading{
				Val:            stypes.ShdSolid,
				Color:          stringPtr("FFFFFF"),
				ThemeColor:     &themeColor1,
				ThemeFill:      &themeFill1,
				ThemeTint:      stringPtr("500"),
				ThemeShade:     stringPtr("200"),
				Fill:           stringPtr("000000"),
				ThemeFillTint:  stringPtr("600"),
				ThemeFillShade: stringPtr("300"),
			},
			expected: `<w:shd w:val="solid" w:color="FFFFFF" w:themeColor="accent2" w:themeFill="accent1" w:themeTint="500" w:themeShade="200" w:fill="000000" w:themeFillTint="600" w:themeFillShade="300"></w:shd>`,
		},
		{
			name: "Shd with some attributes nil",
			input: Shading{
				Val:            stypes.ShdDiagStripe,
				ThemeColor:     &themeColor1,
				ThemeFill:      &themeFill1,
				ThemeTint:      stringPtr("500"),
				ThemeFillTint:  stringPtr("600"),
				ThemeFillShade: stringPtr("300"),
			},
			expected: `<w:shd w:val="diagStripe" w:themeColor="accent2" w:themeFill="accent1" w:themeTint="500" w:themeFillTint="600" w:themeFillShade="300"></w:shd>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapShadingXML(tt.input))
			expected := wrapShadingOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestShd_UnmarshalXML(t *testing.T) {
	themeColor1 := stypes.ThemeColorAccent2
	themeFill1 := stypes.ThemeColorAccent1
	tests := []struct {
		name        string
		inputXML    string
		expectedShd Shading
	}{
		{
			name:     "Basic Shd with Val",
			inputXML: `<w:shd w:val="clear"></w:shd>`,
			expectedShd: Shading{
				Val: stypes.ShdClear,
			},
		},
		{
			name:     "Shd with all attributes",
			inputXML: `<w:shd w:val="solid" w:color="FFFFFF" w:themeColor="accent2" w:themeFill="accent1" w:themeTint="500" w:themeShade="200" w:fill="000000" w:themeFillTint="600" w:themeFillShade="300"></w:shd>`,
			expectedShd: Shading{
				Val:            stypes.ShdSolid,
				Color:          stringPtr("FFFFFF"),
				ThemeColor:     &themeColor1,
				ThemeFill:      &themeFill1,
				ThemeTint:      stringPtr("500"),
				ThemeShade:     stringPtr("200"),
				Fill:           stringPtr("000000"),
				ThemeFillTint:  stringPtr("600"),
				ThemeFillShade: stringPtr("300"),
			},
		},
		{
			name:     "Shd with some attributes missing",
			inputXML: `<w:shd w:val="diagStripe" w:themeColor="accent2" w:themeFill="accent1" w:themeTint="500" w:themeFillTint="600" w:themeFillShade="300"></w:shd>`,
			expectedShd: Shading{
				Val:            stypes.ShdDiagStripe,
				ThemeColor:     &themeColor1,
				ThemeFill:      &themeFill1,
				ThemeTint:      stringPtr("500"),
				ThemeFillTint:  stringPtr("600"),
				ThemeFillShade: stringPtr("300"),
			},
		},
		{
			name:        "Invalid XML",
			inputXML:    `<w:shd invalidattr="invalid"></w:shd>`,
			expectedShd: Shading{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Shading
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			// Check Shd struct fields
			if result.Val != tt.expectedShd.Val {
				t.Errorf("Expected Val %s but got %s", tt.expectedShd.Val, result.Val)
			}
			if result.Color == nil && tt.expectedShd.Color != nil || (result.Color != nil && *result.Color != *tt.expectedShd.Color) {
				t.Errorf("Expected Color %v but got %v", tt.expectedShd.Color, result.Color)
			}
			if result.ThemeColor == nil && tt.expectedShd.ThemeColor != nil || (result.ThemeColor != nil && *result.ThemeColor != *tt.expectedShd.ThemeColor) {
				t.Errorf("Expected ThemeColor %v but got %v", tt.expectedShd.ThemeColor, result.ThemeColor)
			}
			if result.ThemeFill == nil && tt.expectedShd.ThemeFill != nil || (result.ThemeFill != nil && *result.ThemeFill != *tt.expectedShd.ThemeFill) {
				t.Errorf("Expected ThemeFill %v but got %v", tt.expectedShd.ThemeFill, result.ThemeFill)
			}
			if result.ThemeTint == nil && tt.expectedShd.ThemeTint != nil || (result.ThemeTint != nil && *result.ThemeTint != *tt.expectedShd.ThemeTint) {
				t.Errorf("Expected ThemeTint %v but got %v", tt.expectedShd.ThemeTint, result.ThemeTint)
			}
			if result.ThemeShade == nil && tt.expectedShd.ThemeShade != nil || (result.ThemeShade != nil && *result.ThemeShade != *tt.expectedShd.ThemeShade) {
				t.Errorf("Expected ThemeShade %v but got %v", tt.expectedShd.ThemeShade, result.ThemeShade)
			}
			if result.Fill == nil && tt.expectedShd.Fill != nil || (result.Fill != nil && *result.Fill != *tt.expectedShd.Fill) {
				t.Errorf("Expected Fill %v but got %v", tt.expectedShd.Fill, result.Fill)
			}
			if result.ThemeFillTint == nil && tt.expectedShd.ThemeFillTint != nil || (result.ThemeFillTint != nil && *result.ThemeFillTint != *tt.expectedShd.ThemeFillTint) {
				t.Errorf("Expected ThemeFillTint %v but got %v", tt.expectedShd.ThemeFillTint, result.ThemeFillTint)
			}
			if result.ThemeFillShade == nil && tt.expectedShd.ThemeFillShade != nil || (result.ThemeFillShade != nil && *result.ThemeFillShade != *tt.expectedShd.ThemeFillShade) {
				t.Errorf("Expected ThemeFillShade %v but got %v", tt.expectedShd.ThemeFillShade, result.ThemeFillShade)
			}
		})
	}
}

// Helper function to convert string literals to pointers
func stringPtr(s string) *string {
	return &s
}
