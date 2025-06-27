package ctypes

// Document Default Paragraph and Run Properties
type DocDefault struct {
	//Sequence

	//1.Default Run Properties
	RunProp *RunPropDefault `xml:"w:rPrDefault,omitempty"`

	//2.Default Paragraph Properties
	ParaProp *ParaPropDefault `xml:"w:pPrDefault,omitempty"`
}

type RunPropDefault struct {
	RunProp *RunProperty `xml:"w:rPr,omitempty"`
}

type ParaPropDefault struct {
	ParaProp *ParagraphProp `xml:"w:pPr,omitempty"`
}
