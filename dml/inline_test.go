package dml

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
)

func wrapInlineXML(el *Inline) *WrapperXML {
	return wrapXML(struct {
		*Inline
		XMLName struct{} `xml:"wp:inline"`
	}{Inline: el})
}

//TODO generic
// type inlineXML = struct {
// 	*Inline
// 	XMLName xml.Name
// }
// generatedXML, err := xml.Marshal(inlineXML{Inline: tt.input, XMLName: xml.Name{Local: tt.xmlName}})
// xmlName:     "wp:inline",

func TestInline_MarshalXML(t *testing.T) {
	tests := []struct {
		name        string
		input       *Inline
		expectedXML string
	}{{
		name: "All Attributes",
		input: &Inline{
			Attr:  constants.DefaultNamespacesInline,
			DistT: 2,
			DistB: 3,
			DistL: 4,
			DistR: 5,
			Extent: dmlct.PSize2D{
				Width:  100,
				Height: 200,
			},
			DocProp: DocProp{
				ID:          1,
				Name:        "Document Property",
				Description: "This is a document property",
			},
			CNvGraphicFramePr: &NonVisualGraphicFrameProp{
				GraphicFrameLocks: &GraphicFrameLocks{
					NoChangeAspect: dmlst.NewOptBool(true),
				},
			},
			Graphic: *DefaultGraphic(),
		},
		expectedXML: `<wp:inline xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture" distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"></a:graphic></wp:inline>`,
	}}

	for _, tt := range tests {
		object := wrapInlineXML(tt.input)
		expectedXML := wrapXMLOutput(tt.expectedXML)
		t.Run(tt.name, func(t *testing.T) {
			t.Run("MarshalXML", func(t *testing.T) {
				output, err := xml.Marshal(object)
				if err != nil {
					t.Fatalf("Error marshaling to XML: %v", err)
				}
				if got := string(output); got != expectedXML {
					t.Errorf("XML mismatch\nexpectedXML:\n%s\nActual:\n%s", expectedXML, got)
				}
			})
			//TODO UnmarshalXML_TOO_3
			//              // expectedXML:
			//              // &dml.Inline{Attr:[]xml.Attr{xml.Attr{Name:xml.Name{Space:"", Local:"xmlns:a"}, Value:"http://schemas.openxmlformats.org/drawingml/2006/main"}, xml.Attr{Name:xml.Name{Space:"", Local:"xmlns:pic"}, Value:"http://schemas.openxmlformats.org/drawingml/2006/picture"}}, DistT:0x2, DistB:0x3, DistL:0x4, DistR:0x5, Extent:dmlct.PSize2D{Width:0x64, Height:0xc8}, EffectExtent:(*dml.EffectExtent)(nil), DocProp:dml.DocProp{ID:0x1, Name:"Document Property", Description:"This is a document property"}, CNvGraphicFramePr:(*dml.NonVisualGraphicFrameProp)(0xc000048448), Graphic:dml.Graphic{DrawingMLMainNS:"http://schemas.openxmlformats.org/drawingml/2006/main", Data:(*dml.GraphicData)(nil)}}
			//              // Actual:
			//              // &dml.Inline{Attr:[]xml.Attr{xml.Attr{Name:xml.Name{Space:"http://www.w3.org/2000/xmlns/", Local:"a"}, Value:"http://schemas.openxmlformats.org/drawingml/2006/main"}, xml.Attr{Name:xml.Name{Space:"http://www.w3.org/2000/xmlns/", Local:"pic"}, Value:"http://schemas.openxmlformats.org/drawingml/2006/picture"}}, DistT:0x2, DistB:0x3, DistL:0x4, DistR:0x5, Extent:dmlct.PSize2D{Width:0x64, Height:0xc8}, EffectExtent:(*dml.EffectExtent)(nil), DocProp:dml.DocProp{ID:0x1, Name:"Document Property", Description:"This is a document property"}, CNvGraphicFramePr:(*dml.NonVisualGraphicFrameProp)(0xc000048450), Graphic:dml.Graphic{DrawingMLMainNS:"http://schemas.openxmlformats.org/drawingml/2006/main", Data:(*dml.GraphicData)(nil)}}
			// t.Run("UnMarshalXML", func(t *testing.T) {
			// 	object := tt.input
			// 	expectedXML = tt.expectedXML
			// 	vt := reflect.TypeOf(object)
			// 	dest := reflect.New(vt.Elem()).Interface()
			// 	err := xml.Unmarshal([]byte(expectedXML), dest)
			// 	if err != nil {
			// 		t.Fatalf("Error unmarshaling from XML: %v", err)
			// 	}
			// 	if got, want := dest, object; !reflect.DeepEqual(got, want) {
			// 		t.Errorf("XML mismatch unmarshal(%s):\nexpectedXML:\n%#v\nActual:\n%#v", tt.expectedXML, want, got)
			// 	}
			//
			// })
		})
	}
}

func TestUnmarshalInline(t *testing.T) {

	tests := []struct {
		inputXML string
		expected Inline
	}{
		{
			inputXML: `<wp:inline xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture" distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"></a:graphic></wp:inline>`,
			expected: Inline{
				DistT: 2,
				DistB: 3,
				DistL: 4,
				DistR: 5,
				Extent: dmlct.PSize2D{
					Width:  100,
					Height: 200,
				},
				DocProp: DocProp{
					ID:          1,
					Name:        "Document Property",
					Description: "This is a document property",
				},
				CNvGraphicFramePr: &NonVisualGraphicFrameProp{
					GraphicFrameLocks: &GraphicFrameLocks{
						NoChangeAspect: dmlst.NewOptBool(true),
					},
				},
				Graphic: *DefaultGraphic(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var inline Inline

			err := xml.Unmarshal([]byte(tt.inputXML), &inline)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if inline.DistT != tt.expected.DistT {
				t.Errorf("Expected DistT %d, but got %v", tt.expected.DistT, inline.DistT)
			}
			if inline.DistB != tt.expected.DistB {
				t.Errorf("Expected DistB %d, but got %v", tt.expected.DistB, inline.DistB)
			}
			if inline.DistL != tt.expected.DistL {
				t.Errorf("Expected DistL %d, but got %v", tt.expected.DistL, inline.DistL)
			}
			if inline.DistR != tt.expected.DistR {
				t.Errorf("Expected DistR %d, but got %v", tt.expected.DistR, inline.DistR)
			}

			if inline.Extent != tt.expected.Extent {
				t.Errorf("Expected Extent %+v, but got %+v", tt.expected.Extent, inline.Extent)
			}

			if inline.DocProp != tt.expected.DocProp {
				t.Errorf("Expected DocProp %+v, but got %+v", tt.expected.DocProp, inline.DocProp)
			}

			if inline.CNvGraphicFramePr == nil || inline.CNvGraphicFramePr.GraphicFrameLocks == nil ||
				inline.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect != tt.expected.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect {
				t.Errorf("Expected CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect %v, but got %v",
					tt.expected.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect,
					inline.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect)
			}

			if inline.Graphic != tt.expected.Graphic {
				t.Errorf("Expected Graphic %+v, but got %+v", tt.expected.Graphic, inline.Graphic)
			}
		})
	}
}
