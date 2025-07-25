// Package docx provides a comprehensive set of functions and structures
// for manipulating DOCX documents. It allows for the creation, modification,
// and retrieval of document elements such as paragraphs, styles, and images.
// The package is designed to be accessed through the RootDoc element or
// instances of inner elements, providing a flexible and intuitive API for
// working with Office Open XML (OOXML) documents.
//
// // The RootDoc structure is initialized from the main whatsupdocx package,
// which provides methods for creating a new document from a default template or
// opening an existing document.
package docx

import "github.com/samuel-jimenez/whatsupdocx/wml/ctypes"

// w_CT_P
type Paragraph = ctypes.Paragraph

// w_CT_R
type Run = ctypes.Run

// w_CT_Tbl
type Table = ctypes.Table

// w_CT_Row
type Row = ctypes.Row

// w_CT_Tc
type Cell = ctypes.Cell
