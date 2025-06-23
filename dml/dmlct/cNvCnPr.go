package dmlct

// Non-Visual Connector Properties
// CT_NonVisualConnectorProperties
// a_CT_NonVisualConnectorProperties =

type CNvCnPr struct {
	//TODO
	// element cxnSpLocks { a_CT_ConnectorLocking }?,
	// element stCxn { a_CT_Connection }?,
	// element endCxn { a_CT_Connection }?,
	// element extLst { a_CT_OfficeArtExtensionList }?
}

func NewNonVisualConnectorProperties() *CNvCnPr {
	return &CNvCnPr{}
}

//TODO
/*
func (c CNvCnPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
