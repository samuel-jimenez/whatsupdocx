package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// TableWidth represents the width of a table in a document.
type TableWidth struct {
	Width     *int               `xml:"w:w,attr,omitempty"`
	WidthType *stypes.TableWidth `xml:"w:type,attr,omitempty"`
}

func NewTableWidth(width int, widthType stypes.TableWidth) *TableWidth {
	return &TableWidth{
		Width:     &width,
		WidthType: &widthType,
	}
}
