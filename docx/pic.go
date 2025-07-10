package docx

import (
	"github.com/samuel-jimenez/whatsupdocx/dml"
	"github.com/samuel-jimenez/whatsupdocx/wml/ctypes"
)

type PicMeta struct {
	Paragraph *ctypes.Paragraph
	Inline    *dml.Inline
}
