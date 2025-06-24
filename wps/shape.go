package wps

import (
	"fmt"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlpic"
)

// This element specifies a WordprocessingShape. (CT_WordprocessingShape)
// www.iso.org/standard/71691.html
// wp_CT_WordprocessingShape =
type Shape struct {

	// ## default value: false
	// attribute normalEastAsianFlow { xsd:boolean }?,
	//TODO

	// element cNvPr { a_CT_NonVisualDrawingProps }?,
	NonVisualDrawingProps *dmlct.CNvPr `xml:"cNvPr,omitempty"` //element ([ISO/IEC29500-1:2016] section A.4.1)  that specifies non-visual properties. This element MUST NOT be present when the CT_WordprocessingShape is contained directly by a graphicData (Graphic Object Data) element as specified in [ISO/IEC29500-1:2016] section 20.1.2.2.17. This element MUST be present when the CT_WordprocessingShape is not contained directly by a graphicData element as specified in [ISO/IEC29500-1:2016] section 20.1.2.2.17.<120>

	// (element cNvSpPr { a_CT_NonVisualDrawingShapeProps }
	// | element cNvCnPr { a_CT_NonVisualConnectorProperties }),
	NonVisualProps NonVisualProps

	// 3.Shape Properties
	// element spPr { a_CT_ShapeProperties },
	PicShapeProp dmlpic.PicShapeProp `xml:"spPr"`

	// element style { a_CT_ShapeStyle }?,
	Style *ShapeStyle `xml:"style,omitempty"` //element ([ISO/IEC29500-1:2016] section A.4.1) that specifies the style information for a shape.

	// element extLst { a_CT_OfficeArtExtensionList }?,
	ExtLst *OfficeArtExtensionList `xml:"extLst,omitempty"` //element ([ISO/IEC29500-1:2016] section A.4.1) to hold future extensions to the parent element of this extLst element.

	// (element txbx { wp_CT_TextboxInfo }
	// | element linkedTxbx { wp_CT_LinkedTextboxInformation })?,
	ShapeTextboxInfo *ShapeTextboxInfo

	// element bodyPr { a_CT_TextBodyProperties}
	BodyPr TextBodyProperties `xml:"bodyPr"` //element ([ISO/IEC29500-1:2016] section A.4.1) that specifies the body properties for the text body in a shape.

}

// NonVisualProps represents shape properties
// (element cNvSpPr { a_CT_NonVisualDrawingShapeProps }
// | element cNvCnPr { a_CT_NonVisualConnectorProperties }),
type NonVisualProps struct { /*
		NonVisualDrawingShapeProps   *dmlct.CNvSpPr `xml:"cNvSpPr,omitempty"` //element ([ISO/IEC29500-1:2016] section A.4.1) that specifies non-visual shape properties.
		NonVisualConnectorProperties *dmlct.CNvCnPr `xml:"cNvCnPr,omitempty"` //element ([ISO/IEC29500-1:2016] section A.4.1) that specifies non-visual connector properties.*/
	NonVisualDrawingShapeProps   *dmlct.CNvSpPr `xml:"cNvSpPr"` //element ([ISO/IEC29500-1:2016] section A.4.1) that specifies non-visual shape properties.
	NonVisualConnectorProperties *dmlct.CNvCnPr `xml:"cNvCnPr"` //element ([ISO/IEC29500-1:2016] section A.4.1) that specifies non-visual connector properties.
}

func (props NonVisualProps) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if props.NonVisualDrawingShapeProps != nil {
		return e.EncodeElement(props.NonVisualDrawingShapeProps, xml.StartElement{Name: xml.Name{Local: "wps:cNvSpPr"}})
	}
	return e.EncodeElement(props.NonVisualConnectorProperties, xml.StartElement{Name: xml.Name{Local: "wps:cNvCnPr"}})

}

// ShapeStyle specifies the style information for a shape.
// CT_ShapeStyle
// a_CT_ShapeStyle =
type ShapeStyle struct {
	// element lnRef { a_CT_StyleMatrixReference },
	LnRef dmlct.StyleMatrixReference `xml:"a:lnRef"`
	// element fillRef { a_CT_StyleMatrixReference },
	FillRef dmlct.StyleMatrixReference `xml:"a:fillRef"`
	// element effectRef { a_CT_StyleMatrixReference },
	EffectRef dmlct.StyleMatrixReference `xml:"a:effectRef"`
	// element fontRef { a_CT_FontReference }
	FontRef dmlct.FontReference `xml:"a:fontRef"`
}

//TODO use https://github.com/samuel-jimenez/xml
// func (props ShapeStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
//
// 	err = e.EncodeToken(start)
// 	if err != nil {
// 		return err
// 	}
// 		propsElement := xml.StartElement{
// 			Name: xml.Name{Local: "a:lnRef"},
// 		}
// 		if err = e.EncodeElement(props.LnRef, propsElement); err != nil {
// 			return err
// 		}
//
// 		propsElement := xml.StartElement{
// 			Name: xml.Name{Local: "a:lnRef"},
// 		}
// 		if err = e.EncodeElement(props.LnRef, propsElement); err != nil {
// 			return err
// 		}
// 	return e.EncodeToken(xml.EndElement{Name: start.Name})
// }

// OfficeArtExtensionList holds future extensions
// CT_OfficeArtExtensionList
// a_EG_OfficeArtExtensionList = element ext { a_CT_OfficeArtExtension }*
// a_CT_OfficeArtExtensionList = a_EG_OfficeArtExtensionList
type OfficeArtExtensionList struct {
	// TODO
}

// ShapeTextboxInfo represents textual contents of the shape
// (element txbx { wp_CT_TextboxInfo }
// | element linkedTxbx { wp_CT_LinkedTextboxInformation })?,
type ShapeTextboxInfo struct {
	Txbx *TextboxInfo `xml:"txbx,omitempty"` //element that specifies the textual contents of the shape if the shape is the first in the series of shapes for the same text box story.

	LinkedTxbx *LinkedTextboxInformation `xml:"linkedTxbx,omitempty"` //element that specifies the textual contents of the shape if the shape is not the first in the series of shapes for the indicated text box story.
}

// TextboxInfo specifies the textual contents of the shape.
// wp_CT_TextboxInfo =
type TextboxInfo struct {
	// TODO
	// ## default value:
	// attribute id { xsd:unsignedShort }?,
	// element txbxContent { wp_CT_TxbxContent },
	// element extLst { a_CT_OfficeArtExtensionList }?
}

// LinkedTextboxInformation specifies the textual contents of the shape.
// wp_CT_LinkedTextboxInformation =
type LinkedTextboxInformation struct {
	// TODO
	// attribute id { xsd:unsignedShort },
	// attribute seq { xsd:unsignedShort },
	// element extLst { a_CT_OfficeArtExtensionList }?
}

// NonVisualProps represents the body properties for the text body in a shape.
// a:bodyPr
// CT_TextBodyProperties
// a_CT_TextBodyProperties =
type TextBodyProperties struct {
	// TODO

	// attribute rot { a_ST_Angle }?,
	// attribute spcFirstLastPara { xsd:boolean }?,
	// attribute vertOverflow { a_ST_TextVertOverflowType }?,
	// attribute horzOverflow { a_ST_TextHorzOverflowType }?,
	// attribute vert { a_ST_TextVerticalType }?,
	// attribute wrap { a_ST_TextWrappingType }?,
	// attribute lIns { a_ST_Coordinate32 }?,
	// attribute tIns { a_ST_Coordinate32 }?,
	// attribute rIns { a_ST_Coordinate32 }?,
	// attribute bIns { a_ST_Coordinate32 }?,
	// attribute numCol { a_ST_TextColumnCount }?,
	// attribute spcCol { a_ST_PositiveCoordinate32 }?,
	// attribute rtlCol { xsd:boolean }?,
	// attribute fromWordArt { xsd:boolean }?,
	// attribute anchor { a_ST_TextAnchoringType }?,
	// attribute anchorCtr { xsd:boolean }?,
	// attribute forceAA { xsd:boolean }?,
	//
	// ## default value: false
	// attribute upright { xsd:boolean }?,
	// attribute compatLnSpc { xsd:boolean }?,
	// element prstTxWarp { a_CT_PresetTextShape }?,
	// a_EG_TextAutofit?,
	// element scene3d { a_CT_Scene3D }?,
	// a_EG_Text3D?,
	// element extLst { a_CT_OfficeArtExtensionList }?
}

// MarshalXML implements the xml.Marshaler interface for the Shape type.
// It encodes the Shape to its corresponding XML representation.
func (shape Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "wps:wsp"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	// ## default value: false
	// attribute normalEastAsianFlow { xsd:boolean }?,
	//TODO

	// element cNvPr { a_CT_NonVisualDrawingProps }?,
	if shape.NonVisualDrawingProps != nil {
		if err = shape.NonVisualDrawingProps.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	// (element cNvSpPr { a_CT_NonVisualDrawingShapeProps }
	// | element cNvCnPr { a_CT_NonVisualConnectorProperties })
	if err = shape.NonVisualProps.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	// element spPr { a_CT_ShapeProperties },
	if err = shape.PicShapeProp.MarshalXML(e, xml.StartElement{
		Name: xml.Name{Local: "wps:spPr"},
	}); err != nil {
		return fmt.Errorf("marshalling PicShapeProp: %w", err)
	}

	// element style { a_CT_ShapeStyle }?,
	if shape.Style != nil {
		if err = e.EncodeElement(shape.Style, xml.StartElement{Name: xml.Name{Local: "wps:style"}}); err != nil {
			return fmt.Errorf("marshalling Style: %w", err)
		}
	}

	// TODO

	// element extLst { a_CT_OfficeArtExtensionList }?,
	// ExtLst *OfficeArtExtensionList `xml:"extLst,omitempty"` //element ([ISO/IEC29500-1:2016] section A.4.1) to hold future extensions to the parent element of this extLst element.

	// (element txbx { wp_CT_TextboxInfo }
	// | element linkedTxbx { wp_CT_LinkedTextboxInformation })?,
	// ShapeTextboxInfo *ShapeTextboxInfo

	// element bodyPr { a_CT_TextBodyProperties}
	if err = e.EncodeElement(shape.BodyPr, xml.StartElement{Name: xml.Name{Local: "wps:bodyPr"}}); err != nil {
		return fmt.Errorf("marshalling Style: %w", err)
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Shape type.
// It decodes the XML representation of the Shape.
func (shape *Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		// ## default value: false
		// attribute normalEastAsianFlow { xsd:boolean }?,
		//TODO

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "cNvPr": // element cNvPr { a_CT_NonVisualDrawingProps }?,
				props := &dmlct.CNvPr{}
				if err = d.DecodeElement(props, &elem); err != nil {
					return err
				}

			case "cNvSpPr": // (element cNvSpPr { a_CT_NonVisualDrawingShapeProps }
				shape.NonVisualProps.NonVisualDrawingShapeProps = dmlct.NewNonVisualDrawingShapeProps()
				if err = d.DecodeElement(shape.NonVisualProps.NonVisualDrawingShapeProps, &elem); err != nil {
					return err
				}

			case "cNvCnPr": // | element cNvCnPr { a_CT_NonVisualConnectorProperties }),
				shape.NonVisualProps.NonVisualConnectorProperties = dmlct.NewNonVisualConnectorProperties()
				if err = d.DecodeElement(shape.NonVisualProps.NonVisualConnectorProperties, &elem); err != nil {
					return err
				}

			case "spPr": // element spPr { a_CT_ShapeProperties },
				shape.PicShapeProp = *dmlpic.NewPicShapeProp()
				// PicShapeProp
				if err = d.DecodeElement(&shape.PicShapeProp, &elem); err != nil {
					return err
				}

			case "style": // element style { a_CT_ShapeStyle }?,
				shape.Style = &ShapeStyle{}
				if err = d.DecodeElement(&shape.Style, &elem); err != nil {
					return err
				}

			// element extLst { a_CT_OfficeArtExtensionList }?,
			// ExtLst *OfficeArtExtensionList `xml:"extLst,omitempty"` //element ([ISO/IEC29500-1:2016] section A.4.1) to hold future extensions to the parent element of this extLst element.

			// (element txbx { wp_CT_TextboxInfo }
			// | element linkedTxbx { wp_CT_LinkedTextboxInformation })?,
			// ShapeTextboxInfo *ShapeTextboxInfo

			case "bodyPr": // element bodyPr { a_CT_TextBodyProperties}
				props := TextBodyProperties{}
				if err = d.DecodeElement(&props, &elem); err != nil {
					// if err :=shape.PicShapeProp.UnmarshalXML(d, elem); err != nil {
					return err
				}

			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}
