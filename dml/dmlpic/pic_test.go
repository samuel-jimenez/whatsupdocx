package dmlpic

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/samuel-jimenez/xml"

	// "github.com/davecgh/go-spew/spew"
	"github.com/samuel-jimenez/whatsupdocx/common/constants"
	"github.com/samuel-jimenez/whatsupdocx/dml/dmlct"
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

func wrapPicXML(el *Pic) *WrapperXML {
	return wrapXML(struct {
		*Pic
		XMLName struct{} `xml:"pic:pic"`
	}{Pic: el})
}

func TestPic_MarshalXML(t *testing.T) {
	tests := []struct {
		name        string
		input       *Pic
		expectedXML string
	}{{
		name: "",
		input: &Pic{
			Attr: []xml.Attr{constants.NameSpaceDrawingMLPic},
			NonVisualPicProp: NonVisualPicProp{
				CNvPr: dmlct.CNvPr{
					ID:          1,
					Name:        "Pic 1",
					Description: "Description",
				},
				CNvPicPr: CNvPicPr{
					PicLocks: &dmlprops.PicLocks{
						NoChangeAspect:     dmlst.NewOptBool(true),
						NoChangeArrowheads: dmlst.NewOptBool(true),
					},
				},
			},
			BlipFill: BlipFill{
				Blip: &Blip{
					EmbedID: "rId1",
				},
				FillModeProps: FillModeProps{
					Stretch: &shapes.Stretch{
						FillRect: &dmlct.RelativeRect{},
					},
				},
			},
			PicShapeProp: PicShapeProp{
				TransformGroup: &TransformGroup{
					Offset: &Offset{
						X: 0,
						Y: 0,
					},
					Extent: &dmlct.PSize2D{
						Width:  100000,
						Height: 100000,
					},
				},
				PresetGeometry: &PresetGeometry{
					Preset: "rect",
				},
			},
		},
		expectedXML: `<pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture"><pic:nvPicPr><pic:cNvPr id="1" name="Pic 1" descr="Description"></pic:cNvPr><pic:cNvPicPr><a:picLocks noChangeAspect="1" noChangeArrowheads="1"></a:picLocks></pic:cNvPicPr></pic:nvPicPr><pic:blipFill><a:blip r:embed="rId1"></a:blip><a:stretch><a:fillRect></a:fillRect></a:stretch></pic:blipFill><pic:spPr><a:xfrm><a:off x="0" y="0"></a:off><a:ext cx="100000" cy="100000"></a:ext></a:xfrm><a:prstGeom prst="rect"></a:prstGeom></pic:spPr></pic:pic>`,
	}}

	for _, tt := range tests {
		object := wrapPicXML(tt.input)
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
			t.Run("UnMarshalXML", func(t *testing.T) {
				object := tt.input
				expected = tt.expectedXML
				vt := reflect.TypeOf(object)
				dest := reflect.New(vt.Elem()).Interface()
				err := xml.Unmarshal([]byte(expected), dest)
				if err != nil {
					t.Fatalf("Error unmarshaling from XML: %v", err)
				}
				if got, want := dest, object; !reflect.DeepEqual(got, want) {
					// t.Errorf("XML mismatch unmarshal(%s):\nExpected:\n%v\nActual:\n%v", tt.expectedXML, want, got)
					// t.Errorf("XML mismatch unmarshal(%s):\nExpected:\n%s\nActual:\n%s", tt.expectedXML, spew.Sdump(want), spew.Sdump(got))

					want_j, _ := json.Marshal(want)
					got_j, _ := json.Marshal(got)
					t.Errorf("XML mismatch unmarshal(%s):\nExpected:\n%s\nActual:\n%s", tt.expectedXML, want_j, got_j)
				}

			})
		})
	}
}

func TestPicUnmarshalXML(t *testing.T) {
	xmlData := `<pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture">
        	<pic:nvPicPr>
        		<pic:cNvPr id="1" name="Pic 1" descr="Description"></pic:cNvPr>
        		<pic:cNvPicPr>
        			<a:picLocks noChangeAspect="1" noChangeArrowheads="1"></a:picLocks>
        		</pic:cNvPicPr>
        	</pic:nvPicPr>
        	<pic:blipFill>
        		<a:blip r:embed="rId1"></a:blip>
        		<a:stretch>
        			<a:fillRect></a:fillRect>
        		</a:stretch>
        	</pic:blipFill>
        	<pic:spPr>
        		<a:xfrm>
        			<a:off x="0" y="0"></a:off>
        			<a:ext cx="100000" cy="100000"></a:ext>
        		</a:xfrm>
        		<a:prstGeom prst="rect"></a:prstGeom>
        	</pic:spPr>
        </pic:pic>`

	var pic Pic

	err := xml.NewDecoder(strings.NewReader(xmlData)).Decode(&pic)
	if err != nil {
		t.Errorf("Error decoding XML: %v", err)
	}

	checkNotNil := func(fieldName string, fieldValue interface{}) {
		if fieldValue == nil {
			t.Errorf("Expected field '%s' to be unmarshaled, but it was nil", fieldName)
		}
	}

	checkNotNil("NonVisualPicProp", pic.NonVisualPicProp)
	checkNotNil("BlipFill", pic.BlipFill)
	checkNotNil("PicShapeProp", pic.PicShapeProp)
}
