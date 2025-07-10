package docx

import (
	"log"
	"testing"

	"github.com/samuel-jimenez/whatsupdocx/wml/ctypes"
)

func setupRootDoc(_ *testing.T) *RootDoc {
	log.Println("setting root doc")

	return &RootDoc{
		Path:        "/tmp/test",
		RootRels:    Relationships{},
		ContentType: ContentTypes{},
		Document: &Document{
			Body: &ctypes.Body{},
		},
		DocStyles:  &ctypes.Styles{},
		rID:        1,
		ImageCount: 1,
	}
}
