package dmlct

// LineProperties specifies the properties of a line.
// a_CT_LineProperties =
type LineProperties struct {
	// attribute w { a_ST_LineWidth }?,
	LineWidth *string `xml:"w,attr,omitempty"`

	//TODO
	// attribute cap { a_ST_LineCap }?,
	// attribute cmpd { a_ST_CompoundLine }?,
	// attribute algn { a_ST_PenAlignment }?,

	// a_EG_LineFillProperties?,
	LineFillProperties *LineFillProperties `xml:",group,any,omitempty"`

	//TODO
	// a_EG_LineDashProperties?,
	// a_EG_LineJoinProperties?,
	// element headEnd { a_CT_LineEndProperties }?,
	// element tailEnd { a_CT_LineEndProperties }?,
	// element extLst { a_CT_OfficeArtExtensionList }?
}
