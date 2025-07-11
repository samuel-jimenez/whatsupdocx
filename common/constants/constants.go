// Package constants provides constant values related to OpenXML namespaces, relationships, and other attributes.
package constants

import "github.com/samuel-jimenez/xml"

var XMLHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")

const (
	OFFICE_DOC_TYPE    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	CORE_PROP_TYPE     = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"
	EXTENDED_PROP_TYPE = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	StylesType         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
)

var (
	//TODO collapse
	DrawingMLMainNS = "http://schemas.openxmlformats.org/drawingml/2006/main"
	DrawingMLPicNS  = "http://schemas.openxmlformats.org/drawingml/2006/picture"

	//"xmlns:vt"
	NameSpaceDocumentPropertiesVariantTypes = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "vt"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"}
	//"xmlns:asvg"
	NameSpaceDrawing2016SVG = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "asvg"}, Value: "http://schemas.microsoft.com/office/drawing/2016/SVG/main"}
	//"xmlns:a"
	NameSpaceDrawingML = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "a"}, Value: DrawingMLMainNS}
	//"xmlns:pic"
	NameSpaceDrawingMLPic = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "pic"}, Value: DrawingMLPicNS}
	//"xmlns:a14"
	NameSpaceDrawingMLA14 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "a14"}, Value: "http://schemas.microsoft.com/office/drawing/2010/main"}
	//"xmlns:c"
	NameSpaceDrawingMLChart = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "c"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/chart"}
	//"xmlns:sle"
	NameSpaceDrawingMLSlicer = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "sle"}, Value: "http://schemas.microsoft.com/office/drawing/2010/slicer"}
	//"xmlns:sle15"
	NameSpaceDrawingMLSlicerX15 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "sle15"}, Value: "http://schemas.microsoft.com/office/drawing/2012/slicer"}
	//"xmlns:r"
	SourceRelationship = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"}
	//"xmlns:c14"
	SourceRelationshipChart20070802 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "c14"}, Value: "http://schemas.microsoft.com/office/drawing/2007/8/2/chart"}
	//"xmlns:c16"
	SourceRelationshipChart2014 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "c16"}, Value: "http://schemas.microsoft.com/office/drawing/2014/chart"}
	//"xmlns:c16r2"
	SourceRelationshipChart201506 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "c16r2"}, Value: "http://schemas.microsoft.com/office/drawing/2015/06/chart"}
	//"xmlns:mc"
	SourceRelationshipCompatibility = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "mc"}, Value: "http://schemas.openxmlformats.org/markup-compatibility/2006"}

	//"xmlns:w"
	NameSpaceWordprocessingML = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"}
	//"xmlns:o"
	NameSpaceO = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "o"}, Value: "urn:schemas-microsoft-com:office:office"}
	//"xmlns:r"
	NameSpaceR = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"}
	//"xmlns:v"
	NameSpaceV = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "v"}, Value: "urn:schemas-microsoft-com:vml"}
	//"xmlns:w10"
	NameSpaceW10 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "w10"}, Value: "urn:schemas-microsoft-com:office:word"}
	//"xmlns:wp"
	NameSpaceWp = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"}
	//"xmlns:wps"
	NameSpaceWps = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "wps"}, Value: "http://schemas.microsoft.com/office/word/2010/wordprocessingShape"}
	//"xmlns:wpg"
	NameSpaceWpg = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "wpg"}, Value: "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup"}
	//"xmlns:mc"
	NameSpaceMc = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "mc"}, Value: "http://schemas.openxmlformats.org/markup-compatibility/2006"}
	//"xmlns:wp14"
	NameSpaceWp14 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "wp14"}, Value: "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing"}
	//"xmlns:w14"
	NameSpaceW14 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "w14"}, Value: "http://schemas.microsoft.com/office/word/2010/wordml"}
	//"xmlns:w15"
	NameSpaceW15 = xml.Attr{Name: xml.Name{Space: "http://www.w3.org/2000/xmlns/", Local: "w15"}, Value: "http://schemas.microsoft.com/office/word/2012/wordml"}

	//"mc:Ignorable"
	MCIgnorableDoc = xml.Attr{Name: xml.Name{Space: "http://schemas.openxmlformats.org/markup-compatibility/2006", Local: "Ignorable"}, Value: "w14 wp14 w15"}
	//"mc:Ignorable"
	MCIgnorableStyle = xml.Attr{Name: xml.Name{Space: "http://schemas.openxmlformats.org/markup-compatibility/2006", Local: "Ignorable"}, Value: "w14"}

	DefaultNamespacesDoc = []xml.Attr{
		NameSpaceWordprocessingML,
		NameSpaceO,
		NameSpaceR,
		NameSpaceV,
		NameSpaceW10,
		NameSpaceWp,
		NameSpaceWps,
		NameSpaceWpg,
		NameSpaceMc,
		NameSpaceWp14,
		NameSpaceW14,
		NameSpaceW15,
		MCIgnorableDoc,
	}

	DefaultNamespacesStyle = []xml.Attr{
		NameSpaceWordprocessingML,
		NameSpaceMc,
		NameSpaceW14,
		MCIgnorableStyle,
	}

	DefaultNamespacesInline = []xml.Attr{
		NameSpaceDrawingML,
		NameSpaceDrawingMLPic,
	}
)

const (
	XMLNS          = `http://schemas.openxmlformats.org/package/2006/relationships`
	HyperLinkStyle = "Hyperlink"
)

const (
	WMLNamespace    = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"
	AltWMLNamespace = "http://purl.oclc.org/ooxml/wordprocessingml/main"

	WMLDrawingNS = "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
)

const (
	StrictNameSpaceDocumentPropertiesVariantTypes = "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes"
	StrictNameSpaceDrawingMLMain                  = "http://purl.oclc.org/ooxml/drawingml/main"
	StrictNameSpaceExtendedProperties             = "http://purl.oclc.org/ooxml/officeDocument/extendedProperties"
	StrictSourceRelationship                      = "http://purl.oclc.org/ooxml/officeDocument/relationships"
	StrictSourceRelationshipChart                 = "http://purl.oclc.org/ooxml/officeDocument/relationships/chart"
	StrictSourceRelationshipComments              = "http://purl.oclc.org/ooxml/officeDocument/relationships/comments"
	StrictSourceRelationshipExtendProperties      = "http://purl.oclc.org/ooxml/officeDocument/relationships/extendedProperties"
	StrictSourceRelationshipImage                 = "http://purl.oclc.org/ooxml/officeDocument/relationships/image"
	StrictSourceRelationshipOfficeDocument        = "http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument"
)

const (
	NameSpaceDrawingMLMain                = "http://schemas.openxmlformats.org/drawingml/2006/main"
	NameSpaceDublinCore                   = "http://purl.org/dc/elements/1.1/"
	NameSpaceDublinCoreMetadataInitiative = "http://purl.org/dc/dcmitype/"
	NameSpaceDublinCoreTerms              = "http://purl.org/dc/terms/"
	NameSpaceExtendedProperties           = "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"
	NameSpaceXML                          = "http://www.w3.org/XML/1998/namespace"
	NameSpaceXMLSchemaInstance            = "http://www.w3.org/2001/XMLSchema-instance"
)

// Relationships
const (
	SourceRelationshipChart            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
	SourceRelationshipComments         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"
	SourceRelationshipExtendProperties = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	SourceRelationshipImage            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"
	SourceRelationshipOfficeDocument   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	SourceRelationshipHyperLink        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
)

const (
	XMLNS_W = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	XMLNS_R = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
)

const MediaPath = "word/media/"

const ContentTypeFileIdx = "[Content_Types].xml"
