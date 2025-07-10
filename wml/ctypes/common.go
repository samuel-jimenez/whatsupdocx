package ctypes

// Generic Element with Single Val attribute
type GenSingleStrVal[T ~string] struct {
	Val T `xml:"w:val,attr"`
}

func NewGenSingleStrVal[T ~string](val T) *GenSingleStrVal[T] {
	return &GenSingleStrVal[T]{
		Val: val,
	}
}

// Generic Element with Optional Single Val attribute
type GenOptStrVal[T ~string] struct {
	Val *T `xml:"w:val,attr,omitempty"`
}

func NewGenOptStrVal[T ~string](val T) *GenOptStrVal[T] {
	return &GenOptStrVal[T]{
		Val: &val,
	}
}

// CTString - Generic Element that has only one string-type attribute
// And the String type does not have validation
// dont use this if the element requires validation
type CTString struct {
	Val string `xml:"w:val,attr"`
}

func NewCTString(value string) *CTString {
	return &CTString{
		Val: value,
	}
}

type DecimalNum struct {
	Val int `xml:"w:val,attr"`
}

func NewDecimalNum(value int) *DecimalNum {
	return &DecimalNum{
		Val: value,
	}
}

// !--- DecimalNum ends here---!

// !--- Uint64Elem starts---!

// Uint64Elem - Gomplex type that contains single val attribute which is type of uint64
// can be used where w:ST_UnsignedDecimalNumber is applicable
// example: ST_HpsMeasure
type Uint64Elem struct {
	Val uint64 `xml:"w:val,attr"`
}

func NewUint64Elem(value uint64) *Uint64Elem {
	return &Uint64Elem{
		Val: value,
	}
}

// !--- Uint64Elem ends here---!

type Markup struct {
	ID int `xml:"w:id,attr"`
}
