# WhatsupDOCX

[![Go CI](https://github.com/samuel-jimenez/whatsupdocx/actions/workflows/go.yml/badge.svg)](https://github.com/samuel-jimenez/whatsupdocx/actions/workflows/go.yml) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/samuel-jimenez/whatsupdocx) [![Go Reference](https://pkg.go.dev/badge/github.com/samuel-jimenez/whatsupdocx.svg)](https://pkg.go.dev/github.com/samuel-jimenez/whatsupdocx)
 [![Go Report Card](https://goreportcard.com/badge/github.com/samuel-jimenez/whatsupdocx)](https://goreportcard.com/report/github.com/samuel-jimenez/whatsupdocx) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)


<p align="center"><img width="650" src="./whatsupdocx.png" alt="WhatsupDOCX logo"></p>


WhatsupDOCX is a library written in pure Go providing a set of functions that allow you to write to and read from Docx file. 
It should be correct, supporting reading, writing, and editing. Roundtrip reading, then writing should not break.
It is intended to be streamlined, offloading most functionality to the xml parser.

This library needs Go version 1.18 or later. The usage documentation for WhatsupDOCX can be accessed via Go's built-in documentation tool, or online at [go.dev](https://pkg.go.dev/github.com/samuel-jimenez/whatsupdocx). 
Please refer the [subpackage docx](https://pkg.go.dev/github.com/samuel-jimenez/whatsupdocx/docx) for the list of functions that can be used.


## Usage
Here's a simple example of how you can use WhatsupDOCX to create and modify DOCX documents:

## Installation
Use whatsupdocx in your project
```bash
go get github.com/samuel-jimenez/whatsupdocx
```


### Examples
Explore additional examples and use cases over at GitHub repository dedicated to showcasing the capabilities of Golang Docx:
https://github.com/samuel-jimenez/whatsupdocx-examples


```go
// More examples in separate repository
// https://github.com/samuel-jimenez/whatsupdocx-examples

package main

import (
	"log"

	"github.com/samuel-jimenez/whatsupdocx"
)

func main() {
		// Open an existing DOCX document
	// document, err := whatsupdocx.OpenDocument("./testdata/test.docx")

	// Create New Document
	document, err := whatsupdocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	document.AddHeading("Document Title", 0)

	// Add a new paragraph to the document
	p := document.AddParagraph("A plain paragraph having some ")
	p.AddText("bold").Bold(true)
	p.AddText(" and some ")
	p.AddText("italic.").Italic(true)

	document.AddHeading("Heading, level 1", 1)
	document.AddParagraph("Intense quote").Style("Intense Quote")
	document.AddParagraph("first item in unordered list").Style("List Bullet")
	document.AddParagraph("first item in ordered list").Style("List Number")

	records := []struct{ Qty, ID, Desc string }{{"5", "A001", "Laptop"}, {"10", "B202", "Smartphone"}, {"2", "E505", "Smartwatch"}}

	table := document.AddTable()
	table.Style("LightList-Accent4")
	hdrRow := table.AddRow()
	hdrRow.AddCell().AddParagraph("Qty")
	hdrRow.AddCell().AddParagraph("ID")
	hdrRow.AddCell().AddParagraph("Description")

	for _, record := range records {
		row := table.AddRow()
		row.AddCell().AddParagraph(record.Qty)
		row.AddCell().AddParagraph(record.ID)
		row.AddCell().AddParagraph(record.Desc)
	}

	// Save the modified document to a new file
	err = document.SaveTo("demo.docx")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Demo Output

This is screenshot of demo document generated from the whatsupdocx library. 

![Screenshot of the demo output](https://github.com/samuel-jimenez/whatsupdocx-examples/raw/main/demo.png)


## Feature addition request

If you need a feature that's missing in WhatsupDOCX, feel free to raise an issue describing what you want to achieve, along with a sample DOCX. While I can't promise immediate implementation, I'll review your request and work on it if it's valid.



## Inspiration

The WhatsupDOCX library is forked from [Godocx](https://github.com/gomutex/godocx).

## Licenses

The WhatsupDOCX library is licensed under the [MIT License](https://opensource.org/licenses/MIT).
