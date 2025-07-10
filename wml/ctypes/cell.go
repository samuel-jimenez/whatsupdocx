package ctypes

import "github.com/samuel-jimenez/whatsupdocx/wml/stypes"

// w_CT_Tc =
type Cell struct {
	// attribute w:id { s_ST_String }?,

	// 1.Table Cell Properties
	// element tcPr { w_CT_TcPr }?,
	Property *CellProperty `xml:"w:tcPr,omitempty"`

	// 2.1 Choice: ZeroOrMore
	// Any number of elements can exists within this choice group
	// w_EG_BlockLevelElts+
	Contents []BlockLevel `xml:",group,any,omitempty"`

	//TODO: Remaining choices
}

func DefaultCell() *Cell {
	return &Cell{
		Property: &CellProperty{
			Shading: DefaultShading(),
		},
	}
}

// Adds paragraph with text and returns Paragraph
func (c *Cell) AddParagraph(text string) *Paragraph {
	p := NewParagraph(paraWithText(text))
	tblContent := BlockLevel{
		Paragraph: p,
	}

	c.Contents = append(c.Contents, tblContent)

	return p
}

// Add empty paragraph without any text and returns Paragraph
func (c *Cell) AddEmptyPara() *Paragraph {
	p := NewParagraph()
	tblContent := BlockLevel{
		Paragraph: p,
	}

	c.Contents = append(c.Contents, tblContent)

	return p
}

// ColSpan sets the number of columns a cell should span across in a table.
func (c *Cell) ColSpan(cols int) *Cell {
	if c.Property != nil {
		c.Property.GridSpan = &DecimalNum{Val: cols}
	}
	return c
}

// RowSpan sets the cell to span vertically in a table, indicating it is part of a vertically merged group of cells.
func (c *Cell) RowSpan() *Cell {
	if c.Property != nil {
		vMerge := AnnotationVMergeRest
		c.Property.CellMerge = &CellMerge{
			VMerge: &vMerge,
		}
	}
	return c
}

// VerticalAlign sets the vertical alignment of a cell based on the provided string: "top", "center", "middle", or "bottom".
func (c *Cell) VerticalAlign(valign string) *Cell {
	if c.Property != nil {
		switch valign {
		case "top":
			c.Property.VAlign = NewGenSingleStrVal(stypes.VerticalJcTop)
		case "center", "middle":
			c.Property.VAlign = NewGenSingleStrVal(stypes.VerticalJcCenter)
		case "bottom":
			c.Property.VAlign = NewGenSingleStrVal(stypes.VerticalJcBottom)
		}
	}
	return c
}

func (c *Cell) BackgroundColor(color string) *Cell {
	c.Property.Shading.Fill = &color
	return c
}

func (c *Cell) Width(width int, widthType stypes.TableWidth) *Cell {
	c.Property.Width = NewTableWidth(width, widthType)
	return c
}

func (c *Cell) Borders(top *Border, left *Border, bottom *Border, right *Border,
	insideH *Border, insideV *Border, tl2br *Border, tr2bl *Border) *Cell {
	c.Property.Borders = &CellBorders{
		Top:     top,
		Left:    left,
		Bottom:  bottom,
		Right:   right,
		InsideH: insideH,
		InsideV: insideV,
		TL2BR:   tl2br,
		TR2BL:   tr2bl,
	}
	return c
}
