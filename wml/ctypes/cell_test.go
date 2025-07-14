package ctypes

import (
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/internal"
	"github.com/samuel-jimenez/whatsupdocx/internal/testsuite"
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/stretchr/testify/suite"
)

func wrapCellXML(el any) *testsuite.WrapperXML {
	return wrapXML(struct {
		*Cell
		XMLName struct{} `xml:"w:tc"`
	}{Cell: el.(*Cell)})
}

func TestCell(t *testing.T) {
	xmlTester := new(testsuite.XMLTester)
	xmlTester.WrapXMLInput = wrapCellXML
	xmlTester.WrapXMLOutput = wrapXMLOutput

	xmlTester.Tests = []testsuite.XMLTestData{
		{
			Name:        "Empty Cell",
			Input:       &Cell{},
			ExpectedXML: `<w:tc></w:tc>`,
		},
		{
			Name: "Cell with Property and Paragraph Content",
			Input: &Cell{
				Property: &CellProperty{
					NoWrap: &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				},
				Contents: []BlockLevel{
					{
						Paragraph: AddParagraph("Test paragraph content"),
					},
				},
			},
			ExpectedXML: `<w:tc><w:tcPr><w:noWrap w:val="true"></w:noWrap></w:tcPr>` +
				`<w:p><w:r><w:t>Test paragraph content</w:t></w:r></w:p></w:tc>`,
		},
		{
			Name: "Cell with Table Content",
			Input: &Cell{
				Contents: []BlockLevel{
					{
						Table: &Table{},
					},
				},
			},
			ExpectedXML: `<w:tc><w:tbl><w:tblPr></w:tblPr><w:tblGrid></w:tblGrid></w:tbl></w:tc>`,
		},
		{
			Name: "Cell with Property and Paragraph Content 2",
			Input: &Cell{
				Property: &CellProperty{
					CnfStyle: &CTString{"001000000000"},
					Width:    NewTableWidth(5000, "dxa"), //" stypes.TableWidth)
				},
				Contents: []BlockLevel{
					{
						Paragraph: AddParagraph("Hello, World!"),
					},
					{
						Table: &Table{
							RowContents: []RowContent{{
								Row: &Row{Contents: []TRCellContent{{
									Cell: &Cell{Contents: []BlockLevel{{
										Paragraph: AddParagraph("Nested Table Cell")}}}}}}}}},
					},
				},
			},
			ExpectedXML: `<w:tc><w:tcPr><w:cnfStyle w:val="001000000000"></w:cnfStyle><w:tcW w:w="5000" w:type="dxa"></w:tcW></w:tcPr>` +
				`<w:p><w:r><w:t>Hello, World!</w:t></w:r></w:p>` +
				`<w:tbl><w:tblPr></w:tblPr><w:tblGrid></w:tblGrid><w:tr><w:tc><w:p><w:r><w:t>Nested Table Cell</w:t></w:r></w:p></w:tc></w:tr></w:tbl></w:tc>`,
		},
	}
	suite.Run(t, xmlTester)
	if !xmlTester.Stats.Passed() {
		xmlTester.FailNow("XML Failure")
	}
}
