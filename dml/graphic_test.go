package dml

import (
	"testing"

	"github.com/samuel-jimenez/xml"

	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlpic"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlprops"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlst"
	"github.com/samuel-jimenez/whatsupdocx/dml/shapes"
)

type WrapperXML struct {
	XMLName struct{}   `xml:"testwrapper"`
	Attr    []xml.Attr `xml:",any,attr,omitempty"`
	Element any
}

func wrapXML(el any) *WrapperXML {
	return &WrapperXML{
		Attr: []xml.Attr{
			constants.NameSpaceDrawingMLPic,
			constants.NameSpaceR,
		},
		Element: el,
	}
}

func wrapXMLOutput(output string) string {
	return `<testwrapper` +
		` xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture"` +
		` xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"` +
		`>` + output + `</testwrapper>`
}

func wrapGraphicXML(el *Graphic) *WrapperXML {
	return wrapXML(struct {
		*Graphic
		XMLName struct{} `xml:"a:graphic"`
	}{Graphic: el})
}

func TestMarshalGraphic(t *testing.T) {
	tests := []struct {
		name        string
		input       *Graphic
		expectedXML string
		xmlName     string
	}{
		{
			name: "All Attributes",
			input: NewPicGraphic(&dmlpic.Pic{
				Attr: []xml.Attr{constants.NameSpaceDrawingMLPic},
				NonVisualPicProp: dmlpic.NonVisualPicProp{
					CNvPr: dmlct.CNvPr{
						ID:          1,
						Name:        "Pic 1",
						Description: "Description",
					},
					CNvPicPr: dmlpic.CNvPicPr{
						PicLocks: &dmlprops.PicLocks{
							NoChangeAspect:     dmlst.NewOptBool(true),
							NoChangeArrowheads: dmlst.NewOptBool(true),
						},
					},
				},
				BlipFill: dmlpic.BlipFill{
					Blip: &dmlpic.Blip{
						EmbedID: "rId1",
					},
					FillModeProps: &dmlpic.FillModeProps{
						Stretch: &shapes.Stretch{
							FillRect: &dmlct.RelativeRect{},
						},
					},
				},
				PicShapeProp: dmlpic.PicShapeProp{
					TransformGroup: &dmlpic.TransformGroup{
						Offset: &dmlpic.Offset{
							X: 0,
							Y: 0,
						},
						Extent: &dmlct.PSize2D{
							Width:  100000,
							Height: 100000,
						},
					},
					PresetGeometry: &dmlpic.PresetGeometry{
						Preset: "rect",
					},
				},
			}),
			expectedXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"><a:graphicData uri="` + constants.DrawingMLPicNS + `"><pic:pic xmlns:pic="` + constants.DrawingMLPicNS + `"><pic:nvPicPr><pic:cNvPr id="1" name="Pic 1" descr="Description"></pic:cNvPr><pic:cNvPicPr><a:picLocks noChangeAspect="1" noChangeArrowheads="1"></a:picLocks></pic:cNvPicPr></pic:nvPicPr><pic:blipFill><a:blip r:embed="rId1"></a:blip><a:stretch><a:fillRect></a:fillRect></a:stretch></pic:blipFill><pic:spPr><a:xfrm><a:off x="0" y="0"></a:off><a:ext cx="100000" cy="100000"></a:ext></a:xfrm><a:prstGeom prst="rect"></a:prstGeom></pic:spPr></pic:pic></a:graphicData></a:graphic>`,
			xmlName:     "a:graphic",
		},
		{
			name:        "DefaultGraphic",
			input:       DefaultGraphic(),
			expectedXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"></a:graphic>`,
			xmlName:     "a:graphic",
		},
	}

	for _, tt := range tests {
		object := wrapGraphicXML(tt.input)
		expected := wrapXMLOutput(tt.expectedXML)
		t.Run(tt.name, func(t *testing.T) {
			t.Run("MarshalXML", func(t *testing.T) {
				output, err := xml.Marshal(object)
				if err != nil {
					t.Fatalf("Error marshaling to XML: %v", err)
				}
				if got := string(output); got != expected {
					t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, got)
				}
			})
			//TODO UnmarshalXML_TOO_3
			//                 // Expected:
			//              // {"DrawingMLMainNS":"http://schemas.openxmlformats.org/drawingml/2006/main","Data":{"URI":"http://schemas.openxmlformats.org/drawingml/2006/picture","Pic":{"Attr":[{"Name":{"Space":"","Local":"xmlns:pic"},"Value":"http://schemas.openxmlformats.org/drawingml/2006/picture"}],"NonVisualPicProp":{"CNvPr":{"ID":1,"Name":"Pic 1","Description":"Description","Hidden":null},"CNvPicPr":{"PreferRelativeResize":null,"PicLocks":{"DisallowShadowGrouping":{"Bool":false,"Valid":false},"NoSelect":{"Bool":false,"Valid":false},"NoRot":{"Bool":false,"Valid":false},"NoChangeAspect":{"Bool":true,"Valid":true},"NoMove":{"Bool":false,"Valid":false},"NoResize":{"Bool":false,"Valid":false},"NoEditPoints":{"Bool":false,"Valid":false},"NoAdjustHandles":{"Bool":false,"Valid":false},"NoChangeArrowheads":{"Bool":true,"Valid":true},"NoChangeShapeType":{"Bool":false,"Valid":false},"NoCrop":{"Bool":false,"Valid":false}}}},"BlipFill":{"Blip":{"EmbedID":"rId1"},"SrcRect":null,"FillModeProps":{"Stretch":{"FillRect":{"Top":null,"Left":null,"Bottom":null,"Right":null}},"Tile":null},"DPI":null,"RotWithShape":null},"PicShapeProp":{"BwMode":null,"TransformGroup":{"Offset":{"X":0,"Y":0},"Extent":{"Width":100000,"Height":100000}},"PresetGeometry":{"Preset":"rect","AdjustValues":null},"LineProperties":null}},"Shape":null}}
			//              // Actual:
			//              // {"DrawingMLMainNS":"http://schemas.openxmlformats.org/drawingml/2006/main","Data":{"URI":"http://schemas.openxmlformats.org/drawingml/2006/picture","Pic":{"Attr":[{"Name":{"Space":"http://www.w3.org/2000/xmlns/","Local":"pic"},"Value":"http://schemas.openxmlformats.org/drawingml/2006/picture"}],"NonVisualPicProp":{"CNvPr":{"ID":1,"Name":"Pic 1","Description":"Description","Hidden":null},"CNvPicPr":{"PreferRelativeResize":null,"PicLocks":{"DisallowShadowGrouping":{"Bool":false,"Valid":false},"NoSelect":{"Bool":false,"Valid":false},"NoRot":{"Bool":false,"Valid":false},"NoChangeAspect":{"Bool":true,"Valid":true},"NoMove":{"Bool":false,"Valid":false},"NoResize":{"Bool":false,"Valid":false},"NoEditPoints":{"Bool":false,"Valid":false},"NoAdjustHandles":{"Bool":false,"Valid":false},"NoChangeArrowheads":{"Bool":true,"Valid":true},"NoChangeShapeType":{"Bool":false,"Valid":false},"NoCrop":{"Bool":false,"Valid":false}}}},"BlipFill":{"Blip":{"EmbedID":"rId1"},"SrcRect":null,"FillModeProps":{"Stretch":null,"Tile":null},"DPI":null,"RotWithShape":null},"PicShapeProp":{"BwMode":null,"TransformGroup":{"Offset":{"X":0,"Y":0},"Extent":{"Width":100000,"Height":100000}},"PresetGeometry":{"Preset":"rect","AdjustValues":null},"LineProperties":null}},"Shape":null}}
			// t.Run("UnMarshalXML", func(t *testing.T) {
			// 	object := tt.input
			// 	expected = tt.expectedXML
			// 	vt := reflect.TypeOf(object)
			// 	dest := reflect.New(vt.Elem()).Interface()
			// 	err := xml.Unmarshal([]byte(expected), dest)
			// 	if err != nil {
			// 		t.Fatalf("Error unmarshaling from XML: %v", err)
			// 	}
			// 	//TODO UnmarshalXML_TOO
			//
			// 	if got, want := dest, object; !reflect.DeepEqual(got, want) {
			// 		//TODO
			// 		// t.Errorf("XML mismatch unmarshal(%s):\nExpected:\n%#v\nActual:\n%#+v", tt.expectedXML, want, got)
			//
			// 		// TODO spew.Dump(raspberry)
			// 		// "github.com/davecgh/go-spew/spew"
			// 		// gosamples.dev/print-struct-variables/
			// 		want_j, _ := json.Marshal(want)
			// 		got_j, _ := json.Marshal(got)
			// 		t.Errorf("XML mismatch unmarshal(%s):\nExpected:\n%s\nActual:\n%s", tt.expectedXML, want_j, got_j)
			// 	}
			//
			// })
		})
	}
}

func TestUnmarshalGraphic(t *testing.T) {
	tests := []struct {
		inputXML        string
		expectedGraphic Graphic
	}{
		{
			inputXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"><a:graphicData uri="` + constants.DrawingMLPicNS + `"><pic:pic></pic:pic></a:graphicData></a:graphic>`,
			expectedGraphic: Graphic{
				Data: &GraphicData{
					URI: constants.DrawingMLPicNS,
					Pic: &dmlpic.Pic{},
				},
			},
		},
		{
			inputXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"></a:graphic>`,
			expectedGraphic: Graphic{
				Data: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var graphic Graphic

			err := xml.Unmarshal([]byte(tt.inputXML), &graphic)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if graphic.Data == nil && tt.expectedGraphic.Data != nil {
				t.Errorf("Expected Data to be %v, but got nil", tt.expectedGraphic.Data)
			} else if graphic.Data != nil && tt.expectedGraphic.Data == nil {
				t.Errorf("Expected Data to be nil, but got %v", graphic.Data)
			} else if graphic.Data != nil && tt.expectedGraphic.Data != nil {
				if graphic.Data.URI != tt.expectedGraphic.Data.URI {
					t.Errorf("Expected URI %s, but got %s", tt.expectedGraphic.Data.URI, graphic.Data.URI)
				}
				if graphic.Data.Pic == nil && tt.expectedGraphic.Data.Pic != nil {
					t.Errorf("Expected Pic to be %v, but got nil", tt.expectedGraphic.Data.Pic)
				} else if graphic.Data.Pic != nil && tt.expectedGraphic.Data.Pic == nil {
					t.Errorf("Expected Pic to be nil, but got %v", graphic.Data.Pic)
				}
			}
		})
	}
}
