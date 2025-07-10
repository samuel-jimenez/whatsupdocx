package ctypes

import (
	"reflect"
	"strings"
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

type TablePropXML struct {
	Attr    xml.Attr  `xml:",any,attr,omitempty"`
	Element TableProp `xml:"w:tblPr"`
}

func wrapTablePropXML(el TableProp) *TablePropXML {
	return &TablePropXML{
		Attr:    constants.NameSpaceWordprocessingML,
		Element: el,
	}
}
func wrapTablePropOutput(output string) string {
	return `<TablePropXML xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` + output + `</TablePropXML>`
}

func TestTableProp_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    TableProp
		expected string
	}{
		{
			name:     "Empty TableProp",
			input:    TableProp{},
			expected: `<w:tblPr></w:tblPr>`,
		},
		{
			name: "With Style",
			input: TableProp{
				Style: NewCTString("TestStyle"),
			},
			expected: `<w:tblPr><w:tblStyle w:val="TestStyle"></w:tblStyle></w:tblPr>`,
		},
		{
			name: "With Justification",
			input: TableProp{
				Justification: NewGenSingleStrVal(stypes.JustificationCenter),
			},
			expected: `<w:tblPr><w:jc w:val="center"></w:jc></w:tblPr>`,
		},
		{
			name: "With All Fields",
			input: TableProp{
				Style: NewCTString("TestStyle"),
				FloatPos: &FloatPos{
					LeftFromText: internal.ToPtr(uint64(10)),
				},
				Overlap: NewGenSingleStrVal(stypes.TblOverlapNever),
				BidiVisual: &OnOff{
					Val: internal.ToPtr(stypes.OnOffOne),
				},
				RowCountInRowBand: &DecimalNum{Val: 1},
				RowCountInColBand: &DecimalNum{Val: 2},
				Width:             NewTableWidth(10, stypes.TableWidthAuto),
				Justification:     NewGenSingleStrVal(stypes.JustificationCenter),
				CellSpacing:       NewTableWidth(20, stypes.TableWidthDxa),
				Indent:            NewTableWidth(30, stypes.TableWidthPct),
				Borders: &TableBorders{
					Top: &Border{Val: stypes.BorderStyleApples},
				},
				Shading:    &Shading{Val: "clear"},
				Layout:     &TableLayout{LayoutType: internal.ToPtr(stypes.TableLayoutAutoFit)},
				CellMargin: &CellMargins{Top: NewTableWidth(40, stypes.TableWidthDxa)},
				TableLook:  &CTString{Val: "001"},
			},
			expected: `<w:tblPr>` +
				`<w:tblStyle w:val="TestStyle"></w:tblStyle>` +
				`<w:tblpPr w:leftFromText="10"></w:tblpPr>` +
				`<w:tblOverlap w:val="never"></w:tblOverlap>` +
				`<w:bidiVisual w:val="1"></w:bidiVisual>` +
				`<w:tblStyleRowBandSize w:val="1"></w:tblStyleRowBandSize>` +
				`<w:tblStyleColBandSize w:val="2"></w:tblStyleColBandSize>` +
				`<w:tblW w:w="10" w:type="auto"></w:tblW>` +
				`<w:jc w:val="center"></w:jc>` +
				`<w:blCellSpacing w:w="20" w:type="dxa"></w:blCellSpacing>` +
				`<w:tblInd w:w="30" w:type="pct"></w:tblInd>` +
				`<w:tblBorders><w:top w:val="apples"></w:top></w:tblBorders>` +
				`<w:shd w:val="clear"></w:shd>` +
				`<w:tblLayout w:type="autofit"></w:tblLayout>` +
				`<w:tblCellMar><w:top w:w="40" w:type="dxa"></w:top></w:tblCellMar>` +
				`<w:tblLook w:val="001"></w:tblLook>` +
				`</w:tblPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(wrapTablePropXML(tt.input))
			expected := wrapTablePropOutput(tt.expected)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}
			if got := string(output); got != expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
			}
		})
	}
}
func TestTableProp_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   TableProp
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Empty TableProp",
			inputXML: `<w:tblPr></w:tblPr>`,
			expected: TableProp{},
		},
		{
			name: "With Style",
			inputXML: `<w:tblPr>
						<w:tblStyle w:val="TestStyle"></w:tblStyle>
						</w:tblPr>`,
			expected: TableProp{
				Style: NewCTString("TestStyle"),
			},
		},
		{
			name: "With Justification",
			inputXML: `<w:tblPr>
						<w:jc w:val="center"></w:jc>
						</w:tblPr>`,
			expected: TableProp{
				Justification: NewGenSingleStrVal(stypes.JustificationCenter),
			},
		},
		{
			name: "With All Fields",
			inputXML: `<w:tblPr>` +
				`<w:tblStyle w:val="TestStyle"></w:tblStyle>` +
				`<w:tblpPr w:leftFromText="10"></w:tblpPr>` +
				`<w:tblOverlap w:val="never"></w:tblOverlap>` +
				`<w:bidiVisual w:val="1"></w:bidiVisual>` +
				`<w:tblStyleRowBandSize w:val="1"></w:tblStyleRowBandSize>` +
				`<w:tblStyleColBandSize w:val="2"></w:tblStyleColBandSize>` +
				`<w:tblW w:w="10" w:type="auto"></w:tblW>` +
				`<w:jc w:val="center"></w:jc>` +
				`<w:blCellSpacing w:w="20" w:type="dxa"></w:blCellSpacing>` +
				`<w:tblInd w:w="30" w:type="pct"></w:tblInd>` +
				`<w:tblBorders><w:top w:val="apples"></w:top></w:tblBorders>` +
				`<w:shd w:val="clear"></w:shd>` +
				`<w:tblLayout w:type="autofit"></w:tblLayout>` +
				`<w:tblCellMar><w:top w:w="40" w:type="dxa"></w:top></w:tblCellMar>` +
				`<w:tblLook w:val="001"></w:tblLook>` +
				`</w:tblPr>`,
			expected: TableProp{
				Style: NewCTString("TestStyle"),
				FloatPos: &FloatPos{
					LeftFromText: internal.ToPtr(uint64(10)),
				},
				Overlap: NewGenSingleStrVal(stypes.TblOverlapNever),
				BidiVisual: &OnOff{
					Val: internal.ToPtr(stypes.OnOffOne),
				},
				RowCountInRowBand: &DecimalNum{Val: 1},
				RowCountInColBand: &DecimalNum{Val: 2},
				Width:             NewTableWidth(10, stypes.TableWidthAuto),
				Justification:     NewGenSingleStrVal(stypes.JustificationCenter),
				CellSpacing:       NewTableWidth(20, stypes.TableWidthDxa),
				Indent:            NewTableWidth(30, stypes.TableWidthPct),
				Borders: &TableBorders{
					Top: &Border{Val: stypes.BorderStyleApples},
				},
				Shading:    &Shading{Val: "clear"},
				Layout:     &TableLayout{LayoutType: internal.ToPtr(stypes.TableLayoutAutoFit)},
				CellMargin: &CellMargins{Top: NewTableWidth(40, stypes.TableWidthDxa)},
				TableLook:  &CTString{Val: "001"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := xml.NewDecoder(strings.NewReader(tt.inputXML))
			var result TableProp

			err := decoder.Decode(&result)

			if tt.expectFail {
				if err == nil {
					t.Error("Expected unmarshaling to fail but it did not")
				}
				return
			}

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unmarshaled TableProp struct does not match expected:\nExpected: %+v\nActual:   %+v", tt.expected, result)
			}
		})
	}
}
