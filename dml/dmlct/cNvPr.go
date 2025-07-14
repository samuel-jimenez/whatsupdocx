package dmlct

// Non-Visual Drawing Properties
// 19.3.1.12 cNvPr (Non-Visual Drawing Properties)
// CT_NonVisualDrawingProps
// a_CT_NonVisualDrawingProps =
type CNvPr struct {
	// attribute id { a_ST_DrawingElementId },
	ID uint `xml:"id,attr"`
	// attribute name { xsd:string },
	Name string `xml:"name,attr"`

	//Alternative Text for Object - Default value is "".
	// attribute descr { xsd:string }?,
	Description string `xml:"descr,attr,omitempty"`

	// Hidden - Default value is "false".
	// ## default value: false
	// attribute hidden { xsd:boolean }?,
	Hidden *bool `xml:"hidden,attr,omitempty"`

	// attribute title { xsd:string }?,
	// element hlinkClick { a_CT_Hyperlink }?,
	// element hlinkHover { a_CT_Hyperlink }?,
	// element extLst { a_CT_OfficeArtExtensionList }?
	// TODO: implement child elements
	// Sequence [1..1]
	// a:hlinkClick [0..1]    Drawing Element On Click Hyperlink
	// a:hlinkHover [0..1]    Hyperlink for Hover
	// a:extLst [0..1]    Extension List
}

func NewNonVisProp(id uint, name string) *CNvPr {
	return &CNvPr{
		ID:   id,
		Name: name,
	}
}
