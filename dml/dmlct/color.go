package dmlct

import "encoding/xml"

// ColorChoice selects a color
// a_EG_ColorChoice =
type ColorChoice struct {
	//TODO
	// element scrgbClr { a_CT_ScRgbColor }
	// | element srgbClr { a_CT_SRgbColor }
	// | element hslClr { a_CT_HslColor }
	// | element sysClr { a_CT_SystemColor }
	// | element schemeClr { a_CT_SchemeColor }
	SchemeColor *SchemeColor
	// | element prstClr { a_CT_PresetColor }
}

// group marshal
func (group ColorChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if group.SchemeColor != nil {
		return e.EncodeElement(group.SchemeColor, xml.StartElement{Name: xml.Name{Local: "a:schemeClr"}})
	}
	return nil
}

func (group *ColorChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	switch start.Name.Local {
	case "schemeClr":
		group.SchemeColor = &SchemeColor{}
		if err = d.DecodeElement(group.SchemeColor, &start); err != nil {
			return err
		}
	default:
		err = d.Skip()
		return err
	}
	return nil
}

// SchemeColorVal

// a_ST_SchemeColorVal =
// "bg1"
// | "tx1"
// | "bg2"
// | "tx2"
// | "accent1"
// | "accent2"
// | "accent3"
// | "accent4"
// | "accent5"
// | "accent6"
// | "hlink"
// | "folHlink"
// | "phClr"
// | "dk1"
// | "lt1"
// | "dk2"
// | "lt2"

// SchemeColor
// a_CT_SchemeColor =
type SchemeColor struct {
	// attribute val { a_ST_SchemeColorVal },
	Val string `xml:"val,attr"`
	// a_EG_ColorTransform*
	ColorTransform []ColorTransform
}

func (props *SchemeColor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "val":
			props.Val = attr.Value
		}
	}

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "lumOff", "lumMod":
				colorTransform := ColorTransform{}
				colorTransform.UnmarshalXML(d, elem)
				props.ColorTransform = append(props.ColorTransform, colorTransform)

			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// a_ST_Percentage = s_ST_Percentage
// a_CT_Percentage = attribute val { a_ST_Percentage }
type Percentage struct {
	Val string `xml:"val,attr"`
}

// ColorTransform
// a_ST_PositivePercentage = s_ST_PositivePercentage
// a_CT_PositivePercentage = attribute val { a_ST_PositivePercentage }
// a_ST_FixedPercentage = s_ST_FixedPercentage
// a_CT_FixedPercentage = attribute val { a_ST_FixedPercentage }
// a_ST_PositiveFixedPercentage = s_ST_PositiveFixedPercentage
// a_CT_PositiveFixedPercentage =
// attribute val { a_ST_PositiveFixedPercentage }
// a_CT_ComplementTransform = empty
// a_CT_InverseTransform = empty
// a_CT_GrayscaleTransform = empty
// a_CT_GammaTransform = empty
// a_CT_InverseGammaTransform = empty
// a_EG_ColorTransform =
type ColorTransform struct {
	//TODO
	// element tint { a_CT_PositiveFixedPercentage }
	// | element shade { a_CT_PositiveFixedPercentage }
	// | element comp { a_CT_ComplementTransform }
	// | element inv { a_CT_InverseTransform }
	// | element gray { a_CT_GrayscaleTransform }
	// | element alpha { a_CT_PositiveFixedPercentage }
	// | element alphaOff { a_CT_FixedPercentage }
	// | element alphaMod { a_CT_PositivePercentage }
	// | element hue { a_CT_PositiveFixedAngle }
	// | element hueOff { a_CT_Angle }
	// | element hueMod { a_CT_PositivePercentage }
	// | element sat { a_CT_Percentage }
	// | element satOff { a_CT_Percentage }
	// | element satMod { a_CT_Percentage }
	// | element lum { a_CT_Percentage }
	// | element lumOff { a_CT_Percentage }
	// | element lumMod { a_CT_Percentage }
	LumMod *Percentage `xml:"lumMod,omitempty"`

	// | element red { a_CT_Percentage }
	// | element redOff { a_CT_Percentage }
	// | element redMod { a_CT_Percentage }
	// | element green { a_CT_Percentage }
	// | element greenOff { a_CT_Percentage }
	// | element greenMod { a_CT_Percentage }
	// | element blue { a_CT_Percentage }
	// | element blueOff { a_CT_Percentage }
	// | element blueMod { a_CT_Percentage }
	// | element gamma { a_CT_GammaTransform }
	// | element invGamma { a_CT_InverseGammaTransform }

}

// // group marshal
func (group ColorTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {

	if group.LumMod != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "a:lumMod"}}
		if err = e.EncodeElement(group.LumMod, propsElement); err != nil {
			return err
		}
	}

	return nil
}

func (group *ColorTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	switch start.Name.Local {
	case "lumMod":
		group.LumMod = &Percentage{}
		if err = d.DecodeElement(group.LumMod, &start); err != nil {
			return err
		}
	default:
		err = d.Skip()
		return err
	}
	return nil
}

//TODO
// COLOR
//

//
//
// a_CT_ScRgbColor =
// attribute r { a_ST_Percentage },
// attribute g { a_ST_Percentage },
// attribute b { a_ST_Percentage },
// a_EG_ColorTransform*
// a_CT_SRgbColor =
// attribute val { s_ST_HexColorRGB },
// a_EG_ColorTransform*
// a_CT_HslColor =
// attribute hue { a_ST_PositiveFixedAngle },
// attribute sat { a_ST_Percentage },
// attribute lum { a_ST_Percentage },
// a_EG_ColorTransform*

// a_ST_SystemColorVal =
// "scrollBar"
// | "background"
// | "activeCaption"
// | "inactiveCaption"
// | "menu"
// | "window"
// | "windowFrame"
// | "menuText"
// | "windowText"
// | "captionText"
// | "activeBorder"
// | "inactiveBorder"
// | "appWorkspace"
// | "highlight"
// | "highlightText"
// | "btnFace"
// | "btnShadow"
// | "grayText"
// | "btnText"
// | "inactiveCaptionText"
// | "btnHighlight"
// | "3dDkShadow"
// | "3dLight"
// | "infoText"
// | "infoBk"
// | "hotLight"
// | "gradientActiveCaption"
// | "gradientInactiveCaption"
// | "menuHighlight"
// | "menuBar"
// a_CT_SystemColor =
// attribute val { a_ST_SystemColorVal },
// attribute lastClr { s_ST_HexColorRGB }?,
// a_EG_ColorTransform*

/*


a_ST_PresetColorVal =
"aliceBlue"
| "antiqueWhite"
| "aqua"
| "aquamarine"
| "azure"
| "beige"
| "bisque"
| "black"
| "blanchedAlmond"
| "blue"
| "blueViolet"
| "brown"
| "burlyWood"
| "cadetBlue"
| "chartreuse"
| "chocolate"
| "coral"
| "cornflowerBlue"
| "cornsilk"
| "crimson"
| "cyan"
| "darkBlue"
| "darkCyan"
| "darkGoldenrod"
| "darkGray"
| "darkGrey"
| "darkGreen"
| "darkKhaki"
| "darkMagenta"
| "darkOliveGreen"
| "darkOrange"
| "darkOrchid"
| "darkRed"
| "darkSalmon"
| "darkSeaGreen"
| "darkSlateBlue"
| "darkSlateGray"
| "darkSlateGrey"
| "darkTurquoise"
| "darkViolet"
| "dkBlue"
| "dkCyan"
| "dkGoldenrod"
| "dkGray"
| "dkGrey"
| "dkGreen"
| "dkKhaki"
| "dkMagenta"
| "dkOliveGreen"
| "dkOrange"
| "dkOrchid"
| "dkRed"
| "dkSalmon"
| "dkSeaGreen"
| "dkSlateBlue"
| "dkSlateGray"
| "dkSlateGrey"
| "dkTurquoise"
| "dkViolet"
| "deepPink"
| "deepSkyBlue"
| "dimGray"
| "dimGrey"
| "dodgerBlue"
| "firebrick"
| "floralWhite"
| "forestGreen"
| "fuchsia"
| "gainsboro"
| "ghostWhite"
| "gold"
| "goldenrod"
| "gray"
| "grey"
| "green"
| "greenYellow"
| "honeydew"
| "hotPink"
| "indianRed"
| "indigo"
| "ivory"
| "khaki"
| "lavender"
| "lavenderBlush"
| "lawnGreen"
| "lemonChiffon"
| "lightBlue"
| "lightCoral"
| "lightCyan"
| "lightGoldenrodYellow"
| "lightGray"
| "lightGrey"
| "lightGreen"
| "lightPink"
| "lightSalmon"
| "lightSeaGreen"
| "lightSkyBlue"
| "lightSlateGray"
| "lightSlateGrey"
| "lightSteelBlue"
| "lightYellow"
| "ltBlue"
| "ltCoral"
| "ltCyan"
| "ltGoldenrodYellow"
| "ltGray"
| "ltGrey"
| "ltGreen"
| "ltPink"
| "ltSalmon"
| "ltSeaGreen"
| "ltSkyBlue"
| "ltSlateGray"
| "ltSlateGrey"
| "ltSteelBlue"
| "ltYellow"
| "lime"
| "limeGreen"
| "linen"
| "magenta"
| "maroon"
| "medAquamarine"
| "medBlue"
| "medOrchid"
| "medPurple"
| "medSeaGreen"
| "medSlateBlue"
| "medSpringGreen"
| "medTurquoise"
| "medVioletRed"
| "mediumAquamarine"
| "mediumBlue"
| "mediumOrchid"
| "mediumPurple"
| "mediumSeaGreen"
| "mediumSlateBlue"
| "mediumSpringGreen"
| "mediumTurquoise"
| "mediumVioletRed"
| "midnightBlue"
| "mintCream"
| "mistyRose"
| "moccasin"
| "navajoWhite"
| "navy"
| "oldLace"
| "olive"
| "oliveDrab"
| "orange"
| "orangeRed"
| "orchid"
| "paleGoldenrod"
| "paleGreen"
| "paleTurquoise"
| "paleVioletRed"
| "papayaWhip"
| "peachPuff"
| "peru"
| "pink"
| "plum"
| "powderBlue"
| "purple"
| "red"
| "rosyBrown"
| "royalBlue"
| "saddleBrown"
| "salmon"
| "sandyBrown"
| "seaGreen"
| "seaShell"
| "sienna"
| "silver"
| "skyBlue"
| "slateBlue"
| "slateGray"
| "slateGrey"
| "snow"
| "springGreen"
| "steelBlue"
| "tan"
| "teal"
| "thistle"
| "tomato"
| "turquoise"
| "violet"
| "wheat"
| "white"
| "whiteSmoke"
| "yellow"
| "yellowGreen"
a_CT_PresetColor =
attribute val { a_ST_PresetColorVal },
a_EG_ColorTransform*
//
a_EG_OfficeArtExtensionList = element ext { a_CT_OfficeArtExtension }*
a_CT_OfficeArtExtensionList = a_EG_OfficeArtExtensionList
//
a_CT_Scale2D =
element sx { a_CT_Ratio },
element sy { a_CT_Ratio }*/
