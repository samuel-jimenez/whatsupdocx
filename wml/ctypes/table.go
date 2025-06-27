package ctypes

import (
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
	RowContents []RowContent

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
