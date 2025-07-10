package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// .Expanded/Compressed Text
// This element specifies the amount by which each character shall be expanded or when the character is rendered in the document
//
// This property has an of stretching or compressing each character in the run, as opposed to the spacing element (§2.3.2.33) which expands/compresses the text by adding additional character pitch but not changing the width of the actual characters displayed on the line.
// w_CT_TextScale = attribute w:val { w_ST_TextScale }?

type ExpaComp struct {
	Val *stypes.TextScale `xml:"w:val,attr,omitempty"`
}
