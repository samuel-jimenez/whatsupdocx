package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// TableLayout represents the layout of a table in a document.
type TableLayout struct {
	LayoutType *stypes.TableLayout `xml:"w:type,attr,omitempty"`
}

func DefaultTableLayout() *TableLayout {
	return &TableLayout{}
}

// NewTableLayout creates a new TableLayout instance.
func NewTableLayout(t stypes.TableLayout) *TableLayout {
	return &TableLayout{LayoutType: &t}
}
