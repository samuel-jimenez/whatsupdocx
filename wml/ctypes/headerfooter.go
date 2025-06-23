package ctypes

import (
	"github.com/samuel-jimenez/whatsupdocx/wml/stypes"
)

// Header Reference
// w_CT_HdrFtrRef =
type HeaderFooterReference struct {
	// w_CT_Rel,
	// w_CT_Rel = r_id
	// r_id = attribute r:id { r_ST_RelationshipId }
	ID string `xml:"id,attr,omitempty"` //Relationship to Part
	// attribute w:type { w_ST_HdrFtr }
	Type stypes.HdrFtrType `xml:"type,attr,omitempty"` //Header or Footer Type

}


