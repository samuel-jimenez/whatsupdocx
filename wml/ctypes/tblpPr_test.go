package ctypes

import (
	"fmt"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type FloatPosXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element FloatPos `xml:"w:tblpPr"`
}

func wrapFloatPosXML(el FloatPos) *FloatPosXML {
	return &FloatPosXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapFloatPosOutput(output string) string {
	return `<FloatPosXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</FloatPosXML>`
}

func TestFloatPos_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FloatPos
		expected string
	}{
		{
			name: "With all attributes",
			input: FloatPos{
				LeftFromText:   internal.ToPtr(uint64(10)),
				RightFromText:  internal.ToPtr(uint64(20)),
				TopFromText:    internal.ToPtr(uint64(30)),
				BottomFromText: internal.ToPtr(uint64(40)),
				HAnchor:        internal.ToPtr(stypes.AnchorPage),
				VAnchor:        internal.ToPtr(stypes.AnchorText),
				XAlign:         internal.ToPtr(stypes.XAlignCenter),
				YAlign:         internal.ToPtr(stypes.YAlignInside),
				AbsXDist:       internal.ToPtr(int(100)),
				AbsYDist:       internal.ToPtr(int(200)),
			},
			expected: `<w:tblpPr w:leftFromText="10" w:rightFromText="20" w:topFromText="30" w:bottomFromText="40" w:vertAnchor="text" w:horzAnchor="page" w:tblpXSpec="center" w:tblpX="100" w:tblpYSpec="inside" w:tblpY="200"></w:tblpPr>`,
		},
		{
			name: "Without optional attributes",
			input: FloatPos{
				LeftFromText: internal.ToPtr(uint64(5)),
				HAnchor:      internal.ToPtr(stypes.AnchorMargin),
				XAlign:       internal.ToPtr(stypes.XAlignLeft),
				AbsXDist:     internal.ToPtr(int(50)),
			},
			expected: `<w:tblpPr w:leftFromText="5" w:horzAnchor="margin" w:tblpXSpec="left" w:tblpX="50"></w:tblpPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapFloatPosXML(tt.input))
			expected := wrapFloatPosOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestFloatPos_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FloatPos
	}{
		{
			name: "With all attributes",
			inputXML: `<w:tblpPr w:leftFromText="10" w:rightFromText="20" w:topFromText="30" w:bottomFromText="40" ` +
				`w:horzAnchor="margin" w:vertAnchor="page" w:tblpXSpec="center" w:tblpYSpec="inside" w:tblpX="100" w:tblpY="200"></w:tblpPr>`,
			expected: FloatPos{
				LeftFromText:   internal.ToPtr(uint64(10)),
				RightFromText:  internal.ToPtr(uint64(20)),
				TopFromText:    internal.ToPtr(uint64(30)),
				BottomFromText: internal.ToPtr(uint64(40)),
				HAnchor:        internal.ToPtr(stypes.AnchorMargin),
				VAnchor:        internal.ToPtr(stypes.AnchorPage),
				XAlign:         internal.ToPtr(stypes.XAlignCenter),
				YAlign:         internal.ToPtr(stypes.YAlignInside),
				AbsXDist:       internal.ToPtr(int(100)),
				AbsYDist:       internal.ToPtr(int(200)),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<w:tblpPr w:leftFromText="5" w:horzAnchor="margin" w:tblpXSpec="left" w:tblpX="50"></w:tblpPr>`,
			expected: FloatPos{
				LeftFromText: internal.ToPtr(uint64(5)),
				HAnchor:      internal.ToPtr(stypes.AnchorMargin),
				XAlign:       internal.ToPtr(stypes.XAlignLeft),
				AbsXDist:     internal.ToPtr(int(50)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FloatPos

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare each field individually due to pointer comparisons
			err = compareFloatPos(tt.expected, result)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

// Helper function to compare FloatPos structs
func compareFloatPos(expected, result FloatPos) error {
	compareUint64Ptr := func(fieldName string, a, b *uint64) error {
		if a == nil || b == nil {
			if a != b {
				return fmt.Errorf("%s: expected %v but got %v", fieldName, internal.FormatPtr(a), internal.FormatPtr(b))
			}
		} else if *a != *b {
			return fmt.Errorf("%s: expected %v but got %v", fieldName, *a, *b)
		}
		return nil
	}

	compareIntPtr := func(fieldName string, a, b *int) error {
		if a == nil || b == nil {
			if a != b {
				return fmt.Errorf("%s: expected %v but got %v", fieldName, internal.FormatPtr(a), internal.FormatPtr(b))
			}
		} else if *a != *b {
			return fmt.Errorf("%s: expected %v but got %v", fieldName, *a, *b)
		}
		return nil
	}

	// Compare each field
	if err := compareUint64Ptr("LeftFromText", expected.LeftFromText, result.LeftFromText); err != nil {
		return err
	}
	if err := compareUint64Ptr("RightFromText", expected.RightFromText, result.RightFromText); err != nil {
		return err
	}
	if err := compareUint64Ptr("TopFromText", expected.TopFromText, result.TopFromText); err != nil {
		return err
	}
	if err := compareUint64Ptr("BottomFromText", expected.BottomFromText, result.BottomFromText); err != nil {
		return err
	}
	if err := internal.ComparePtr("HAnchor", expected.HAnchor, result.HAnchor); err != nil {
		return err
	}
	if err := internal.ComparePtr("VAnchor", expected.VAnchor, result.VAnchor); err != nil {
		return err
	}
	if err := internal.ComparePtr("XAlign", expected.XAlign, result.XAlign); err != nil {
		return err
	}
	if err := internal.ComparePtr("YAlign", expected.YAlign, result.YAlign); err != nil {
		return err
	}
	if err := compareIntPtr("AbsXDist", expected.AbsXDist, result.AbsXDist); err != nil {
		return err
	}
	if err := compareIntPtr("AbsYDist", expected.AbsYDist, result.AbsYDist); err != nil {
		return err
	}

	return nil
}
