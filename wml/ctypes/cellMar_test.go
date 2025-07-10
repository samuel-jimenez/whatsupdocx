package ctypes

import (
	"reflect"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

func wrapCellMarginsXML(el CellMargins) *WrapperXML {
	return wrapXML(struct {
		CellMargins
		XMLName struct{} `xml:"w:tblCellMar"`
	}{CellMargins: el})
}

func TestCellMarginsMarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    CellMargins
		expected string
	}{
		{
			input: CellMargins{
				Top:    NewTableWidth(0, stypes.TableWidthDxa),
				Left:   NewTableWidth(55, stypes.TableWidthDxa),
				Bottom: NewTableWidth(0, stypes.TableWidthDxa),
				Right:  NewTableWidth(55, stypes.TableWidthDxa),
			},
			expected: `<w:tblCellMar><w:top w:w="0" w:type="dxa"></w:top><w:left w:w="55" w:type="dxa"></w:left><w:bottom w:w="0" w:type="dxa"></w:bottom><w:right w:w="55" w:type="dxa"></w:right></w:tblCellMar>`,
		},
		{
			input:    DefaultCellMargins(),
			expected: `<w:tblCellMar></w:tblCellMar>`,
		},
	}

	for _, tt := range tests {
		object := wrapCellMarginsXML(tt.input)
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

func TestCellMarginsUnmarshalXML(t *testing.T) {
	testCases := []struct {
		input    string
		expected CellMargins
	}{
		{
			input: `<w:tblCellMar xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"><w:top w:w="0" w:type="dxa"></w:top><w:left w:w="55" w:type="dxa"></w:left><w:bottom w:w="0" w:type="dxa"></w:bottom><w:right w:w="55" w:type="dxa"></w:right></w:tblCellMar>`,
			expected: CellMargins{
				Top:    NewTableWidth(0, stypes.TableWidthDxa),
				Left:   NewTableWidth(55, stypes.TableWidthDxa),
				Bottom: NewTableWidth(0, stypes.TableWidthDxa),
				Right:  NewTableWidth(55, stypes.TableWidthDxa),
			},
		},
	}

	for _, tc := range testCases {
		var result CellMargins

		err := xml.Unmarshal([]byte(tc.input), &result)
		if err != nil {
			t.Fatalf("Error unmarshaling XML: %v", err)
		}

		if result.Top == nil {
			t.Errorf("Got nil value for Top")
		}

		if result.Bottom == nil {
			t.Errorf("Got nil value for Bottom")
		}

		if result.Left == nil {
			t.Errorf("Got nil value for Left")
		}

		if result.Right == nil {
			t.Errorf("Got nil value for Right")
		}

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %+v, but got %+v", tc.expected, result)
		}
	}
}
