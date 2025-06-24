package stypes

import (
	"errors"

	"github.com/samuel-jimenez/xml"
)

// w_ST_PageBorderZOrder = string "front" | string "back"
type PageBorderZOrder string

const (
	PageBorderZOrderFront PageBorderZOrder = "front"
	PageBorderZOrderBack  PageBorderZOrder = "back"
)

func PageBorderZOrderFromStr(value string) (PageBorderZOrder, error) {
	switch value {
	case "front":
		return PageBorderZOrderFront, nil
	case "back":
		return PageBorderZOrderBack, nil
	default:
		return "", errors.New("Invalid Border ZOrder Input")
	}
}

func (d *PageBorderZOrder) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := PageBorderZOrderFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}

// w_ST_PageBorderDisplay =
// string "allPages" | string "firstPage" | string "notFirstPage"
type PageBorderDisplay string

const (
	PageBorderDisplayAllPages     PageBorderDisplay = "allPages"
	PageBorderDisplayFirstPage    PageBorderDisplay = "firstPage"
	PageBorderDisplayNotFirstPage PageBorderDisplay = "notFirstPage"
)

func PageBorderDisplayFromStr(value string) (PageBorderDisplay, error) {
	switch value {
	case "allPages":
		return PageBorderDisplayAllPages, nil
	case "firstPage":
		return PageBorderDisplayFirstPage, nil
	case "notFirstPage":
		return PageBorderDisplayNotFirstPage, nil
	default:
		return "", errors.New("Invalid BorderDisplay Input")
	}
}

func (d *PageBorderDisplay) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := PageBorderDisplayFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}

// w_ST_PageBorderOffset = string "page" | string "text"
type PageBorderOffset string

const (
	PageBorderOffsetPage PageBorderOffset = "page"
	PageBorderOffsetText PageBorderOffset = "text"
)

func PageBorderOffsetFromStr(value string) (PageBorderOffset, error) {
	switch value {
	case "page":
		return PageBorderOffsetPage, nil
	case "text":
		return PageBorderOffsetText, nil
	default:
		return "", errors.New("Invalid BorderOffset Input")
	}
}

func (d *PageBorderOffset) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := PageBorderOffsetFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
