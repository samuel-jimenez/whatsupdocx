package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/xml"
)

func optBoolElemPtr(value OnOff) *OnOff {
	return &value
}

func singleUint64ValPtr(value Uint64Elem) *Uint64Elem {
	return &value
}

func singleIntValPtr(value DecimalNum) *DecimalNum {
	return &value
}

func singleStrValPtr(value CTString) *CTString {
	return &value
}

type RunPropertyXML struct {
	Attr    xml.Attr    `xml:",any,attr,omitempty"`
	Element RunProperty `xml:"w:rPr"`
}

func wrapRunPropertyXML(el RunProperty) *RunPropertyXML {
	return &RunPropertyXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapRunPropertyOutput(output string) string {
	return `<RunPropertyXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</RunPropertyXML>`
}

func TestRunProperty_MarshalXML(t *testing.T) {
	// trueOptBool := types.NewOptBool(true)
	tests := []struct {
		name     string
		input    RunProperty
		expected string
	}{
		{
			name: "All attributes set",
			input: RunProperty{
				Bold:         optBoolElemPtr(OnOff{}),
				BoldCS:       optBoolElemPtr(OnOff{}),
				Italic:       optBoolElemPtr(OnOff{}),
				ItalicCS:     optBoolElemPtr(OnOff{}),
				Strike:       optBoolElemPtr(OnOff{}),
				DoubleStrike: optBoolElemPtr(OnOff{}),
				Outline:      optBoolElemPtr(OnOff{}),
				Shadow:       optBoolElemPtr(OnOff{}),
				Caps:         optBoolElemPtr(OnOff{}),
				SmallCaps:    optBoolElemPtr(OnOff{}),
				Emboss:       optBoolElemPtr(OnOff{}),
				Imprint:      optBoolElemPtr(OnOff{}),
				NoGrammar:    optBoolElemPtr(OnOff{}),
				SnapToGrid:   optBoolElemPtr(OnOff{}),
				Vanish:       optBoolElemPtr(OnOff{}),
				WebHidden:    optBoolElemPtr(OnOff{}),
				RightToLeft:  optBoolElemPtr(OnOff{}),
				CSFormat:     optBoolElemPtr(OnOff{}),
				SpecVanish:   optBoolElemPtr(OnOff{}),
				OMath:        optBoolElemPtr(OnOff{}),
				Kern:         singleUint64ValPtr(Uint64Elem{Val: 20}),
				Spacing:      singleIntValPtr(DecimalNum{Val: 100}),
				Style:        singleStrValPtr(CTString{Val: "Heading1"}),
				Position:     singleIntValPtr(DecimalNum{Val: 10}),
			},
			expected: `<w:rPr><w:rStyle w:val="Heading1"></w:rStyle><w:b></w:b><w:bCs></w:bCs><w:i></w:i><w:iCs></w:iCs><w:caps></w:caps><w:smallCaps></w:smallCaps><w:strike></w:strike><w:dstrike></w:dstrike><w:outline></w:outline><w:shadow></w:shadow><w:emboss></w:emboss><w:imprint></w:imprint><w:noProof></w:noProof><w:snapToGrid></w:snapToGrid><w:vanish></w:vanish><w:webHidden></w:webHidden><w:spacing w:val="100"></w:spacing><w:kern w:val="20"></w:kern><w:position w:val="10"></w:position><w:rtl></w:rtl><w:cs></w:cs><w:specVanish></w:specVanish><w:oMath></w:oMath></w:rPr>`,
		},
		{
			name: "Only Bold set",
			input: RunProperty{
				Bold: optBoolElemPtr(OnOff{}),
			},
			expected: `<w:rPr><w:b></w:b></w:rPr>`,
		},
		{
			name:     "No attributes set",
			input:    RunProperty{},
			expected: `<w:rPr></w:rPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapRunPropertyXML(tt.input))
			expected := wrapRunPropertyOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestRunProperty_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name         string
		inputXML     string
		expectedProp RunProperty
	}{
		{
			name:     "All attributes set",
			inputXML: `<w:rPr><w:rStyle w:val="Heading1"></w:rStyle><w:b></w:b><w:bCs></w:bCs><w:i></w:i><w:iCs></w:iCs><w:caps></w:caps><w:smallCaps></w:smallCaps><w:strike></w:strike><w:dstrike></w:dstrike><w:outline></w:outline><w:shadow></w:shadow><w:emboss></w:emboss><w:imprint></w:imprint><w:noProof></w:noProof><w:snapToGrid></w:snapToGrid><w:vanish></w:vanish><w:webHidden></w:webHidden><w:spacing w:val="100"></w:spacing><w:kern w:val="20"></w:kern><w:position w:val="10"></w:position><w:rtl></w:rtl><w:cs></w:cs><w:specVanish></w:specVanish><w:oMath></w:oMath></w:rPr>`,
			expectedProp: RunProperty{
				Bold:         optBoolElemPtr(OnOff{}),
				BoldCS:       optBoolElemPtr(OnOff{}),
				Italic:       optBoolElemPtr(OnOff{}),
				ItalicCS:     optBoolElemPtr(OnOff{}),
				Strike:       optBoolElemPtr(OnOff{}),
				DoubleStrike: optBoolElemPtr(OnOff{}),
				Outline:      optBoolElemPtr(OnOff{}),
				Shadow:       optBoolElemPtr(OnOff{}),
				Caps:         optBoolElemPtr(OnOff{}),
				SmallCaps:    optBoolElemPtr(OnOff{}),
				Emboss:       optBoolElemPtr(OnOff{}),
				Imprint:      optBoolElemPtr(OnOff{}),
				NoGrammar:    optBoolElemPtr(OnOff{}),
				SnapToGrid:   optBoolElemPtr(OnOff{}),
				Vanish:       optBoolElemPtr(OnOff{}),
				WebHidden:    optBoolElemPtr(OnOff{}),
				RightToLeft:  optBoolElemPtr(OnOff{}),
				CSFormat:     optBoolElemPtr(OnOff{}),
				SpecVanish:   optBoolElemPtr(OnOff{}),
				OMath:        optBoolElemPtr(OnOff{}),
				Kern:         singleUint64ValPtr(Uint64Elem{Val: 20}),
				Spacing:      singleIntValPtr(DecimalNum{Val: 100}),
				Style:        singleStrValPtr(CTString{Val: "Heading1"}),
				Position:     singleIntValPtr(DecimalNum{Val: 10}),
			},
		},
		{
			name:     "Only Bold set",
			inputXML: `<w:rPr><w:b/></w:rPr>`,
			expectedProp: RunProperty{
				Bold: optBoolElemPtr(OnOff{}),
			},
		},
		{
			name:         "No attributes set",
			inputXML:     `<w:rPr></w:rPr>`,
			expectedProp: RunProperty{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var prop RunProperty
			err := xml.Unmarshal([]byte(tt.inputXML), &prop)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare expected and actual RunProperty structs
			if !reflect.DeepEqual(prop, tt.expectedProp) {
				t.Errorf("Expected:\n%+v\nBut got:\n%+v", tt.expectedProp, prop)
			}
		})
	}
}
