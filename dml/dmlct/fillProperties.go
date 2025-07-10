package dmlct

import (
	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common"
)

// FillProperties specifies the properties of a line fill.
// a_EG_FillProperties =
type FillProperties struct {
	// element noFill { a_CT_NoFillProperties }
	// | element solidFill { a_CT_SolidColorFillProperties }
	// | element gradFill { a_CT_GradientFillProperties }
	// | element pattFill { a_CT_PatternFillProperties }
	LineFillProperties `xml:",group,any,omitempty"`

	//TODO
	// | element blipFill { a_CT_BlipFillProperties }
	// | element grpFill { a_CT_GroupFillProperties }
}

// LineFillProperties specifies the properties of a line fill.
// a_EG_LineFillProperties =
type LineFillProperties struct {

	// element noFill { a_CT_NoFillProperties }
	// a_CT_NoFillProperties = empty
	NoFillProperties *common.Empty `xml:"a:noFill,omitempty"`
	// | element solidFill { a_CT_SolidColorFillProperties }
	SolidColorFillProperties *SolidColorFillProperties `xml:"a:solidFill,omitempty"`

	//TODO
	// | element gradFill { a_CT_GradientFillProperties }
	// | element pattFill { a_CT_PatternFillProperties }
}

func (group *LineFillProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	switch start.Name.Local {
	case "noFill":
		group.NoFillProperties = &common.Empty{}
		if err = d.DecodeElement(group.NoFillProperties, &start); err != nil {
			return err
		}
	case "solidFill":
		group.SolidColorFillProperties = &SolidColorFillProperties{}
		if err = d.DecodeElement(group.SolidColorFillProperties, &start); err != nil {
			return err
		}
	default:
		err = d.Skip()
		return err
	}
	return nil
}

// StyleMatrixReference
// a_CT_StyleMatrixReference =
type StyleMatrixReference struct {
	// a_ST_StyleMatrixColumnIndex = xsd:unsignedInt
	// attribute idx { a_ST_StyleMatrixColumnIndex },
	Id string `xml:"idx,attr,omitempty"`
	// a_EG_ColorChoice?
	ColorChoice *ColorChoice `xml:",group,any,omitempty"`
}

const (
	FontCollectionIndexMajor = "major"
	FontCollectionIndexMinor = "minor"
	FontCollectionIndexNone  = "none"
)

// FontReference
// a_CT_FontReference =
type FontReference struct {
	// a_ST_FontCollectionIndex = "major" | "minor" | "none"
	// attribute idx { a_ST_FontCollectionIndex },
	Id string `xml:"idx,attr,omitempty"`
	// a_EG_ColorChoice?
	ColorChoice *ColorChoice `xml:",group,any,omitempty"`
}

// SolidColorFillProperties specifies the properties of solid fill.
// a_CT_SolidColorFillProperties = a_EG_ColorChoice?
// a_EG_ColorChoice =
type SolidColorFillProperties struct {
	ColorChoice *ColorChoice `xml:",group,any,omitempty"`
}

//TODO
// GradientFillProperties
// a_CT_GradientFillProperties =
// attribute flip { a_ST_TileFlipMode }?,
// attribute rotWithShape { xsd:boolean }?,
// element gsLst { a_CT_GradientStopList }?,
// a_EG_ShadeProperties?,
// element tileRect { a_CT_RelativeRect }?

//TODO
// PatternFillProperties
// a_CT_PatternFillProperties =
// attribute prst { a_ST_PresetPatternVal }?,
// element fgClr { a_CT_Color }?,
// element bgClr { a_CT_Color }?
