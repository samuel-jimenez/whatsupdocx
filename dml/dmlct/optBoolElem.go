package dmlct

import (
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
)

// Optional Bool Element: Helper element that only has one attribute which is optional
type OptBoolElem struct {
	Val dmlst.OptBool `xml:"w:val,attr,omitempty"`
}

func NewOptBoolElem(value bool) *OptBoolElem {
	return &OptBoolElem{
		Val: dmlst.NewOptBool(value),
	}
}

// Disable sets the value to false and valexists true
func (n *OptBoolElem) Disable() {
	n.Val = dmlst.NewOptBool(false)
}
