package ctypes

import (
	"strings"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common"
	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
)

type WrapperXML struct {
	XMLName struct{}   `xml:"testwrapper"`
	Attr    []xml.Attr `xml:",any,attr,omitempty"`
	Element any
}

func wrapXML(el any) *WrapperXML {
	return &WrapperXML{
		Attr: []xml.Attr{
			constants.NameSpaceWordprocessingML,
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

// !--- Tests for CTString start here ---!

func wrapCTStringXML(el CTString) *WrapperXML {
	return wrapXML(struct {
		CTString
		XMLName struct{} `xml:"w:rStyle"`
	}{CTString: el})
}

func TestCTString_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    CTString
		expected string
	}{
		{
			name:     "With value",
			input:    CTString{Val: "example"},
			expected: `<w:rStyle w:val="example"></w:rStyle>`,
		},
		{
			name:     "Empty value",
			input:    CTString{Val: ""},
			expected: `<w:rStyle w:val=""></w:rStyle>`,
		},
	}

	for _, tt := range tests {
		object := wrapCTStringXML(tt.input)
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
}

func TestCTString_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected CTString
	}{
		{
			name:     "With value",
			inputXML: `<w:rStyle w:val="example"></w:rStyle>`,
			expected: CTString{Val: "example"},
		},
		{
			name:     "Empty value",
			inputXML: `<w:rStyle w:val=""></w:rStyle>`,
			expected: CTString{Val: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result CTString

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, result.Val)
			}
		})
	}
}

// !--- Tests for CTString end here ---!

// !--- Tests for DecimalNum start here ---!

func wrapDecimalNumXML(el DecimalNum) *WrapperXML {
	return wrapXML(struct {
		DecimalNum
		XMLName struct{} `xml:"w:outlineLvl"`
	}{DecimalNum: el})
}

func TestDecimalNum_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    DecimalNum
		expected string
	}{
		{
			name:     "With value",
			input:    DecimalNum{Val: 10},
			expected: `<w:outlineLvl w:val="10"></w:outlineLvl>`,
		},
		{
			name:     "Empty value",
			input:    DecimalNum{Val: -1},
			expected: `<w:outlineLvl w:val="-1"></w:outlineLvl>`,
		},
	}

	for _, tt := range tests {
		object := wrapDecimalNumXML(tt.input)
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
}

func TestDecimalNum_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected DecimalNum
	}{
		{
			name:     "With value",
			inputXML: `<w:outlineLvl w:val="00122"></w:outlineLvl>`,
			expected: DecimalNum{Val: 122},
		},
		{
			name:     "Empty value",
			inputXML: `<w:outlineLvl w:val="+3"></w:outlineLvl>`,
			expected: DecimalNum{Val: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result DecimalNum

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %d but got %d", tt.expected.Val, result.Val)
			}
		})
	}
}

// !--- Tests for DecimalNum end here ---!

// !--- Tests for Uint64 start here ---!

func wrapUint64ElemXML(el Uint64Elem) *WrapperXML {
	return wrapXML(struct {
		Uint64Elem
		XMLName struct{} `xml:"w:kern"`
	}{Uint64Elem: el})
}

func TestUint64Elem_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Uint64Elem
		expected string
	}{
		{
			name:     "With value",
			input:    Uint64Elem{Val: 10},
			expected: `<w:kern w:val="10"></w:kern>`,
		},
		{
			name:     "Empty value",
			input:    Uint64Elem{Val: 18446744073709551615},
			expected: `<w:kern w:val="18446744073709551615"></w:kern>`,
		},
	}

	for _, tt := range tests {
		object := wrapUint64ElemXML(tt.input)
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
}

func TestUint64Elem_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Uint64Elem
	}{
		{
			name:     "With value",
			inputXML: `<w:kern w:val="00122"></w:kern>`,
			expected: Uint64Elem{Val: 122},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Uint64Elem

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %d but got %d", tt.expected.Val, result.Val)
			}
		})
	}
}

// !--- Tests for Uint64 end here ---!

// !--- Tests for GenSingleStrVal start here ---!

func wrapGenSingleStrValXML(el string) *WrapperXML {
	return wrapXML(struct {
		*GenSingleStrVal[string]
		XMLName struct{} `xml:"GenSingleStrVal"`
	}{GenSingleStrVal: NewGenSingleStrVal(el)})
}

func TestGenSingleStrVal_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test1", "Hello", `<GenSingleStrVal w:val="Hello"></GenSingleStrVal>`},
		{"Test2", "World", `<GenSingleStrVal w:val="World"></GenSingleStrVal>`},
	}
	for _, tt := range tests {
		object := wrapGenSingleStrValXML(tt.input)
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
			}) //TODO UnmarshalXML_1TOO
		})
	}
}

func TestGenSingleStrVal_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name string
		xml  string
		want string
	}{
		{"Test1", `<GenSingleStrVal w:val="Hello"></GenSingleStrVal>`, "Hello"},
		{"Test2", `<GenSingleStrVal w:val="World"></GenSingleStrVal>`, "World"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gen GenSingleStrVal[string]

			err := xml.Unmarshal([]byte(tt.xml), &gen)
			if err != nil {
				t.Errorf("UnmarshalXML() error = %v", err)
				return
			}

			if gen.Val != tt.want {
				t.Errorf("UnmarshalXML() = %v, want %v", gen.Val, tt.want)
			}
		})
	}
}

// !--- Tests for GenSingleStrVal end here ---!

// !--- Tests for Empty start here ---!

func TestEmpty_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    common.Empty
		expected string
	}{
		{
			name:     "Empty element",
			input:    common.Empty{},
			expected: `<w:tab></w:tab>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:tab"}}

			err := encoder.EncodeElement(tt.input, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestEmpty_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
	}{
		{
			name:     "Empty element",
			inputXML: `<w:tab></w:tab>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result common.Empty

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}
		})
	}
}

// !--- Tests for Empty end here ---!

// !--- Tests for Markup start here ---!

func wrapMarkupXML(el Markup) *WrapperXML {
	return wrapXML(struct {
		Markup
		XMLName struct{} `xml:"Markup"`
	}{Markup: el})
}

func TestMarkup_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Markup
		expected string
	}{
		{
			name:     "With ID",
			input:    Markup{ID: 42},
			expected: `<Markup w:id="42"></Markup>`,
		},
		{
			name:     "Zero ID",
			input:    Markup{ID: 0},
			expected: `<Markup w:id="0"></Markup>`,
		},
	}

	for _, tt := range tests {
		object := wrapMarkupXML(tt.input)
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
}

func TestMarkup_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Markup
	}{
		{
			name:     "With ID",
			inputXML: `<Markup w:id="42"></Markup>`,
			expected: Markup{ID: 42},
		},
		{
			name:     "Zero ID",
			inputXML: `<Markup w:id="0"></Markup>`,
			expected: Markup{ID: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Markup

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.ID != tt.expected.ID {
				t.Errorf("Expected ID %d but got %d", tt.expected.ID, result.ID)
			}
		})
	}
}

// !--- Tests for Markup end here ---!

// !--- Tests for GenOptStrVal start here ---!

func wrapGenOptStrValstringXML(el GenOptStrVal[string]) *WrapperXML {
	return wrapXML(struct {
		GenOptStrVal[string]
		XMLName struct{} `xml:"element"`
	}{GenOptStrVal: el})
}

// Test function for MarshalXML method
func TestMarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    GenOptStrVal[string]
		expected string
	}{
		{
			name:     "WithValue",
			input:    GenOptStrVal[string]{Val: internal.ToPtr("test")},
			expected: `<element w:val="test"></element>`,
		},
		{
			name:     "WithNilValue",
			input:    GenOptStrVal[string]{Val: nil},
			expected: `<element></element>`,
		},
		{
			name:     "EmptyValue",
			input:    GenOptStrVal[string]{Val: internal.ToPtr("")},
			expected: `<element w:val=""></element>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapGenOptStrValstringXML(tt.input))
			expected := wrapXMLOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestUnmarshalXML(t *testing.T) {
	tests := []struct {
		name           string
		xmlString      string
		expectedStruct GenOptStrVal[string]
	}{
		{
			name:      "WithValAttribute",
			xmlString: `<element w:val="test"></element>`,
			expectedStruct: GenOptStrVal[string]{
				Val: internal.ToPtr("test"),
			},
		},
		{
			name:      "WithoutValAttribute",
			xmlString: `<element></element>`,
			expectedStruct: GenOptStrVal[string]{
				Val: nil,
			},
		},
		{
			name:      "EmptyValAttribute",
			xmlString: `<element w:val=""></element>`,
			expectedStruct: GenOptStrVal[string]{
				Val: internal.ToPtr(""),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new XML decoder for the test XML string
			decoder := xml.NewDecoder(strings.NewReader(tt.xmlString))

			// Create an empty instance of GenOptStrVal to store the decoded result
			var result GenOptStrVal[string]

			// Decode XML into result
			err := decoder.Decode(&result)
			if err != nil {
				t.Errorf("Decode error: %v", err)
				return
			}

			// Compare the decoded result with the expectedStruct using ComparePtr
			err = internal.ComparePtr("Val", tt.expectedStruct.Val, result.Val)
			if err != nil {
				t.Errorf("Comparison error: %v", err)
			}
		})
	}
}

// !--- Tests for GenOptStrVal[string] end here ---!
