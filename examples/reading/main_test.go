package main

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/samuel-jimenez/whatsupdocx"
)

func TestMain(t *testing.T) {
	document, err := whatsupdocx.OpenDocument("hello-world.docx")
	if err != nil {
		t.Fatal(err)
	}

	xmlEncoder := xml.NewEncoder(os.Stdout)
	err = xmlEncoder.Encode(document.Document.Body)
	if err != nil {
		t.Fatal(err)
	}
}
