package ctypes

import (
	"errors"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type LsdExceptionXML struct {
	Attr    xml.Attr     `xml:",any,attr,omitempty"`
	Element LsdException `xml:"LsdException"`
}

func wrapLsdExceptionXML(el LsdException) *LsdExceptionXML {
	return &LsdExceptionXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapLsdExceptionOutput(output string) string {
	return `<LsdExceptionXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</LsdExceptionXML>`
}

func TestLsdException_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    LsdException
		expected string
	}{
		{
			name: "All attributes set",
			input: LsdException{
				Name:           "Heading1",
				Locked:         internal.ToPtr(stypes.OnOffOn),
				UIPriority:     internal.ToPtr(99),
				SemiHidden:     internal.ToPtr(stypes.OnOffOn),
				UnhideWhenUsed: internal.ToPtr(stypes.OnOffOn),
				QFormat:        internal.ToPtr(stypes.OnOffOn),
			},
			expected: `<w:lsdException w:name="Heading1" w:locked="on" w:uiPriority="99" w:semiHidden="on" w:unhideWhenUsed="on" w:qFormat="on"></w:lsdException>`,
		},
		{
			name: "Only name set",
			input: LsdException{
				Name: "Heading2",
			},
			expected: `<w:lsdException w:name="Heading2"></w:lsdException>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapLsdExceptionXML(tt.input))
			expected := wrapLsdExceptionOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestLsdException_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected LsdException
	}{
		{
			name:     "All attributes set",
			inputXML: `<w:lsdException w:name="Heading1" w:locked="on" w:uiPriority="99" w:semiHidden="on" w:unhideWhenUsed="on" w:qFormat="on"></w:lsdException>`,
			expected: LsdException{
				Name:           "Heading1",
				Locked:         internal.ToPtr(stypes.OnOffOn),
				UIPriority:     internal.ToPtr(99),
				SemiHidden:     internal.ToPtr(stypes.OnOffOn),
				UnhideWhenUsed: internal.ToPtr(stypes.OnOffOn),
				QFormat:        internal.ToPtr(stypes.OnOffOn),
			},
		},
		{
			name:     "Only name set",
			inputXML: `<w:lsdException w:name="Heading2"></w:lsdException>`,
			expected: LsdException{
				Name: "Heading2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lsd LsdException
			err := xml.Unmarshal([]byte(tt.inputXML), &lsd)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if err := compareLsdExFields(lsd, tt.expected); err != nil {
				t.Error(err)
			}
		})
	}
}

func compareLsdExFields(a, b LsdException) error {
	if a.Name != b.Name {
		return errors.New("Name mismatch")
	}
	if err := internal.ComparePtr("Locked", a.Locked, b.Locked); err != nil {
		return err
	}
	if err := internal.ComparePtr("UIPriority", a.UIPriority, b.UIPriority); err != nil {
		return err
	}
	if err := internal.ComparePtr("SemiHidden", a.SemiHidden, b.SemiHidden); err != nil {
		return err
	}
	if err := internal.ComparePtr("UnhideWhenUsed", a.UnhideWhenUsed, b.UnhideWhenUsed); err != nil {
		return err
	}
	if err := internal.ComparePtr("QFormat", a.QFormat, b.QFormat); err != nil {
		return err
	}
	return nil
}
