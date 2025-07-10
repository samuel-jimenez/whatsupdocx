package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type TabXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element Tab      `xml:"w:tab"`
}

func wrapTabXML(el Tab) *TabXML {
	return &TabXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapTabOutput(output string) string {
	return `<TabXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</TabXML>`
}

type TabsXML struct {
	Attr    xml.Attr `xml:",any,attr,omitempty"`
	Element Tabs     `xml:"w:tabs"`
}

func wrapTabsXML(el Tabs) *TabsXML {
	return &TabsXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapTabsOutput(output string) string {
	return `<TabsXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</TabsXML>`
}

func TestTab_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Tab
		expected string
	}{
		{
			name: "With all attributes",
			input: Tab{
				Val:        stypes.CustTabStopCenter,
				Position:   720,
				LeaderChar: internal.ToPtr(stypes.CustLeadCharDot),
			},
			expected: `<w:tab w:val="center" w:pos="720" w:leader="dot"></w:tab>`,
		},
		{
			name: "Without optional attributes",
			input: Tab{
				Val:      stypes.CustTabStopRight,
				Position: 1440,
			},
			expected: `<w:tab w:val="right" w:pos="1440"></w:tab>`,
		},
		{
			name:     "Empty struct",
			input:    Tab{},
			expected: `<w:tab w:val="" w:pos="0"></w:tab>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTabXML(tt.input))
			expected := wrapTabOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestTab_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Tab
	}{
		{
			name:     "With all attributes",
			inputXML: `<w:tab w:val="center" w:pos="720" w:leader="dot"></w:tab>`,
			expected: Tab{
				Val:        stypes.CustTabStopCenter,
				Position:   720,
				LeaderChar: internal.ToPtr(stypes.CustLeadCharDot),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<w:tab w:val="right" w:pos="1440"></w:tab>`,
			expected: Tab{
				Val:      stypes.CustTabStopRight,
				Position: 1440,
			},
		},
		{
			name:     "Empty struct",
			inputXML: `<w:tab></w:tab>`,
			expected: Tab{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Tab

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare Val
			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %v", tt.expected.Val, result.Val)
			}

			// Compare Position
			if result.Position != tt.expected.Position {
				t.Errorf("Expected Position %d but got %v", tt.expected.Position, result.Position)
			}

			// Compare LeaderChar
			if tt.expected.LeaderChar != nil {
				if result.LeaderChar == nil || *result.LeaderChar != *tt.expected.LeaderChar {
					t.Errorf("Expected LeaderChar %s but got %v", *tt.expected.LeaderChar, *result.LeaderChar)
				}
			} else if result.LeaderChar != nil {
				t.Errorf("Expected LeaderChar nil but got %v", *result.LeaderChar)
			}
		})
	}
}
func TestTabs_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Tabs
		expected string
	}{
		{
			name:     "Empty Tabs",
			input:    Tabs{},
			expected: ``,
		},
		{
			name: "Tabs with Multiple Tab elements",
			input: Tabs{
				Tab: []Tab{
					{Val: stypes.CustTabStopCenter, Position: 100, LeaderChar: internal.ToPtr(stypes.CustLeadCharDot)},
					{Val: stypes.CustTabStopLeft, Position: 200, LeaderChar: internal.ToPtr(stypes.CustLeadCharHyphen)},
				},
			},
			expected: `<w:tabs><w:tab w:val="center" w:pos="100" w:leader="dot"></w:tab><w:tab w:val="left" w:pos="200" w:leader="hyphen"></w:tab></w:tabs>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTabsXML(tt.input))
			expected := wrapTabsOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}
func TestTabs_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected Tabs
	}{
		{
			name:     "Empty Tabs",
			xmlInput: `<w:tabs></w:tabs>`,
			expected: Tabs{},
		},
		{
			name: "Tabs with Multiple Tab elements",
			xmlInput: `<tabs>
                <w:tab val="center" pos="100" leader="dot"/>
                <w:tab val="left" pos="200" leader="hyphen"/>
            </tabs>`,
			expected: Tabs{
				Tab: []Tab{
					{Val: stypes.CustTabStopCenter, Position: 100, LeaderChar: internal.ToPtr(stypes.CustLeadCharDot)},
					{Val: stypes.CustTabStopLeft, Position: 200, LeaderChar: internal.ToPtr(stypes.CustLeadCharHyphen)},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tabs Tabs
			err := xml.Unmarshal([]byte(tt.xmlInput), &tabs)
			if err != nil {
				t.Fatalf("Unexpected error during Unmarshal: %v", err)
			}

			// Compare individual fields of Tabs struct
			if !reflect.DeepEqual(tabs, tt.expected) {
				t.Errorf("Unmarshaled Tabs struct does not match expected:\nExpected: %+v\nActual:   %+v", tt.expected, tabs)
			}
		})
	}
}
