package dmlct

// a_ST_Coordinate = a_ST_CoordinateUnqualified | s_ST_UniversalMeasure
//
// a_ST_CoordinateUnqualified =
// xsd:long {
// minInclusive = "-27273042329600"
// maxInclusive = "27273042316900"
// }
//
// s_ST_UniversalMeasure =
// xsd:string { pattern = "-?[0-9]+(\.[0-9]+)?(mm|cm|in|pt|pc|pi)" }
// http://www.datypic.com/sc/ooxml/t-a_ST_Coordinate.html

// Wrapping Polygon Point2D
// a_CT_Point2D =
// attribute x { a_ST_Coordinate },
// attribute y { a_ST_Coordinate }
type Point2D struct {
	XAxis int64 `xml:"x,attr"`
	YAxis int64 `xml:"y,attr"`
}

func NewPoint2D(x, y int64) Point2D {
	return Point2D{
		XAxis: int64(x),
		YAxis: int64(y),
	}
}
