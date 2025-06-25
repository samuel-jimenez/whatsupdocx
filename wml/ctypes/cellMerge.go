package ctypes

import (
	"errors"

	"github.com/samuel-jimenez/xml"
)

type CellMerge struct {
	ID         int               `xml:"w:id,attr"`
	Author     string            `xml:"w:author,attr"`
	Date       *string           `xml:"w:date,attr,omitempty"`
	VMerge     *AnnotationVMerge `xml:"w:vMerge,attr,omitempty"`     //Revised Vertical Merge Setting
	VMergeOrig *AnnotationVMerge `xml:"w:vMergeOrig,attr,omitempty"` //Vertical Merge Setting Removed by Revision

}

type AnnotationVMerge string

const (
	// AnnotationVMergeCont represents a vertically merged cell.
	AnnotationVMergeCont AnnotationVMerge = "cont"
	// AnnotationVMergeRest represents a vertically split cell.
	AnnotationVMergeRest AnnotationVMerge = "rest"
)

// AnnotationVMergeFromStr converts a string to AnnotationVMerge type.
func AnnotationVMergeFromStr(value string) (AnnotationVMerge, error) {
	switch value {
	case "cont":
		return AnnotationVMergeCont, nil
	case "rest":
		return AnnotationVMergeRest, nil
	default:
		return "", errors.New("invalid AnnotationVMerge value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into AnnotationVMerge.
func (a *AnnotationVMerge) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := AnnotationVMergeFromStr(attr.Value)
	if err != nil {
		return err
	}
	*a = val
	return nil
}
