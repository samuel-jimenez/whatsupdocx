package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
	"github.com/samuel-jimenez/xml"
)

// Table
// w_CT_Tbl =
type Table struct {
	//1.Choice: RangeMarkupElements
	// w_EG_RangeMarkupElements*,
	RngMarkupElems []RngMarkupElem

	//2. Table Properties
	// element tblPr { w_CT_TblPr },
	TableProp TableProp `xml:"w:tblPr,omitempty"`

	//3. Table Grid
	// element tblGrid { w_CT_TblGrid },
	Grid Grid `xml:"w:tblGrid,omitempty"`

	//4.1 Choice:
	// w_EG_ContentRowContent*
	RowContents []RowContent `xml:",group,any,omitempty"`

	//4.2 TODO: Remaining choices
}

func DefaultTable() *Table {
	return &Table{}
}

func (t *Table) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "tblPr":
				prop := TableProp{}
				if err = d.DecodeElement(&prop, &elem); err != nil {
					return err
				}

				t.TableProp = prop
			case "tblGrid":
				grid := Grid{}
				if err = d.DecodeElement(&grid, &elem); err != nil {
					return err
				}

				t.Grid = grid
			case "tr":
				row := Row{}
				if err = d.DecodeElement(&row, &elem); err != nil {
					return err
				}

				t.RowContents = append(t.RowContents, RowContent{
					Row: &row,
				})

			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}

// w_EG_ContentRowContent
// w_EG_ContentRowContent =
type RowContent struct {
	// element tr { w_CT_Row }*
	Row *Row `xml:"w:tr,omitempty"`
	// | element customXml { w_CT_CustomXmlRow }
	// | element sdt { w_CT_SdtRow }
	// | w_EG_RunLevelElts*
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Width(v int, u stypes.TableWidth) *Table {
	w := TableWidth{
		Width:     &v,
		WidthType: &u,
	}
	t.TableProp.Width = &w
	return t
}

/*
The table grid is a definition of the set of grid columns which define all of the shared vertical edges of the table,
as well as default widths for each of these grid columns.
These grid column widths are then used to determine the size of the table based on the table layout algorithm used
*/
func (t *Table) AppendGrid(widths ...uint64) *Table {
	for _, w := range widths {
		tw := w
		col := Column{Width: &tw}
		t.Grid.Col = append(t.Grid.Col, col)
	}
	return t
}

func (t *Table) CellMargin(top *TableWidth, left *TableWidth, bottom *TableWidth, right *TableWidth) *Table {
	t.TableProp.CellMargin = &CellMargins{
		Top:    top,
		Left:   left,
		Bottom: bottom,
		Right:  right,
	}
	return t
}

func (t *Table) Layout(layout stypes.TableLayout) *Table {
	t.TableProp.Layout = &TableLayout{
		LayoutType: &layout,
	}

	return t
}

// AddRow adds a new row to the table.
//
// It creates a new row and appends it to the table's row contents. Use this method to construct the structure
// of the table by sequentially adding rows and cells.
//
// Returns:
//   - *Row: A pointer to the newly added row.

func (t *Table) AddRow() *Row {
	row := DefaultRow()

	t.RowContents = append(t.RowContents, RowContent{
		Row: row,
	})

	return row
}

// //TODO
// func (t *Table) ensureProp() {
// }

// Indent sets the indent width for the table.
//
// Parameters:
//   - indent: An integer specifying the indent width
func (t *Table) Indent(indent int) {
	t.TableProp.Indent = NewTableWidth(indent, stypes.TableWidthAuto)
}

// Style sets the style for the table.
//
// TableStyle represents the style of a table in a document.
// This is applicable when creating a new document. When using this style in a new document, you need to ensure
// that the specified style ID exists in your document's style base or is manually created through the library.
//
// Some examples of predefined style IDs in the docx template that can be used are:
//
//   - "LightShading"
//   - "LightShading-Accent1"
//   - "LightShading-Accent2"
//   - "LightShading-Accent3"
//   - "LightShading-Accent4"
//   - "LightShading-Accent5"
//   - "LightShading-Accent6"
//   - "LightList"
//   - "LightList-Accent1"..."LightList-Accent6"
//   - "LightGrid"
//   - "LightGrid-Accent1"..."LightGrid-Accent6"
//   - "MediumShading"
//   - "MediumShading-Accent1"..."MediumShading-Accent6"
//   - "MediumShading2"
//   - "MediumShading2-Accent1"..."MediumShading2-Accent6"
//   - "MediumList1"
//   - "MediumList1-Accent1"..."MediumList1-Accent6"
//   - "MediumList2"
//   - "MediumList2-Accent1"..."MediumList2-Accent6"
//   - "TableGrid"
//   - "MediumGrid1"
//   - "MediumGrid1-Accent1"..."MediumGrid1-Accent6"
//   - "MediumGrid2"
//   - "MediumGrid2-Accent1"..."MediumGrid2-Accent6"
//   - "MediumGrid3"
//   - "MediumGrid3-Accent1"..."MediumGrid3-Accent6"
//   - "DarkList"
//   - "DarkList-Accent1"..."DarkList-Accent6"
//   - "ColorfulShading"
//   - "ColorfulShading-Accent1"..."ColorfulShading-Accent6"
//   - "ColorfulList"
//   - "ColorfulList-Accent1"..."ColorfulList-Accent6"
//   - "ColorfulGrid"
//   - "ColorfulGrid-Accent1"..."ColorfulGrid-Accent6"
//
// Parameters:
//   - value: A string representing the style value. It should match a valid table style defined in the WordprocessingML specification.
func (t *Table) Style(value string) {
	t.TableProp.Style = NewCTString(value)
}
