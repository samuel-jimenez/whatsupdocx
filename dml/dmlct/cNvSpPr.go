package dmlct

// Non-Visual Shape Properties
// CT_NonVisualDrawingShapeProps
// a_CT_NonVisualDrawingShapeProps =
type CNvSpPr struct {
	//TODO
	// ## default value: false
	// attribute txBox { xsd:boolean }?,
	// element spLocks { a_CT_ShapeLocking }?,
	// element extLst { a_CT_OfficeArtExtensionList }?
}

func NewNonVisualDrawingShapeProps() *CNvSpPr {
	return &CNvSpPr{}
}

//TODO
/*
func (c CNvSpPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// ! NOTE: Disabling the empty name check for the Picture
	//  since popular docx tools allow them
	// if c.Name == "" {
	// 	return fmt.Errorf("invalid Name for Non-Visual Drawing Properties when marshaling")
	// }

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: strconv.FormatUint(uint64(c.ID), 10)},
		{Name: xml.Name{Local: "name"}, Value: c.Name},
	}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "descr"}, Value: c.Description})

	if c.Hidden != nil {
		if *c.Hidden {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"}, Value: "true"})
		} else {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"}, Value: "false"})
		}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}*/
