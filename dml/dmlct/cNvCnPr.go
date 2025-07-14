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
