package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
)

type TrackChangeXML struct {
	Attr    xml.Attr    `xml:",any,attr,omitempty"`
	Element TrackChange `xml:"w:TrackChange"`
}

func wrapTrackChangeXML(el TrackChange) *TrackChangeXML {
	return &TrackChangeXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}

func wrapTrackChangeOutput(output string) string {
	return `<TrackChangeXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</TrackChangeXML>`
}

func TestTrackChange_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    TrackChange
		expected string
	}{
		{
			name: "With all attributes",
			input: TrackChange{
				ID:     123,
				Author: "John Doe",
				Date:   internal.ToPtr("2023-06-18T12:34:56Z"),
			},
			expected: `<w:TrackChange w:id="123" w:author="John Doe" w:date="2023-06-18T12:34:56Z"></w:TrackChange>`,
		},
		{
			name: "Without date attribute",
			input: TrackChange{
				ID:     124,
				Author: "Jane Doe",
			},
			expected: `<w:TrackChange w:id="124" w:author="Jane Doe"></w:TrackChange>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTrackChangeXML(tt.input))
			expected := wrapTrackChangeOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}

func TestTrackChange_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected TrackChange
	}{
		{
			name:     "With all attributes",
			inputXML: `<w:TrackChange w:id="123" w:author="John Doe" w:date="2023-06-18T12:34:56Z"></w:TrackChange>`,
			expected: TrackChange{
				ID:     123,
				Author: "John Doe",
				Date:   internal.ToPtr("2023-06-18T12:34:56Z"),
			},
		},
		{
			name:     "Without date attribute",
			inputXML: `<w:TrackChange w:id="124" w:author="Jane Doe"></w:TrackChange>`,
			expected: TrackChange{
				ID:     124,
				Author: "Jane Doe",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TrackChange

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare ID
			if result.ID != tt.expected.ID {
				t.Errorf("Expected ID %d but got %d", tt.expected.ID, result.ID)
			}

			// Compare Author
			if result.Author != tt.expected.Author {
				t.Errorf("Expected Author %s but got %s", tt.expected.Author, result.Author)
			}

			// Compare Date
			if tt.expected.Date != nil {
				if result.Date == nil || *result.Date != *tt.expected.Date {
					t.Errorf("Expected Date %s but got %v", *tt.expected.Date, result.Date)
				}
			} else if result.Date != nil {
				t.Errorf("Expected Date nil but got %v", result.Date)
			}
		})
	}
}
