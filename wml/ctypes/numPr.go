package ctypes

// Numbering Definition Instance Reference
type NumProp struct {
	//Numbering Level Reference
	ILvl *DecimalNum `xml:"w:ilvl,omitempty"`

	//Numbering Definition Instance Reference
	NumID *DecimalNum `xml:"w:numId,omitempty"`

	//Previous Paragraph Numbering Properties
	NumChange *TrackChangeNum `xml:"w:numberingChange,omitempty"`

	//Inserted Numbering Properties
	Ins *TrackChange `xml:"w:ins,omitempty"`
}

// NewNumberingProperty creates a new NumberingProperty instance.
func NewNumberingProperty() *NumProp {
	return &NumProp{}
}
