package dmlpic

// a_CT_Blip =
// a_AG_Blob,
//
// ## default value: none
// attribute cstate { a_ST_BlipCompression }?,
// (element alphaBiLevel { a_CT_AlphaBiLevelEffect }
// | element alphaCeiling { a_CT_AlphaCeilingEffect }
// | element alphaFloor { a_CT_AlphaFloorEffect }
// | element alphaInv { a_CT_AlphaInverseEffect }
// | element alphaMod { a_CT_AlphaModulateEffect }
// | element alphaModFix { a_CT_AlphaModulateFixedEffect }
// | element alphaRepl { a_CT_AlphaReplaceEffect }
// | element biLevel { a_CT_BiLevelEffect }
// | element blur { a_CT_BlurEffect }
// | element clrChange { a_CT_ColorChangeEffect }
// | element clrRepl { a_CT_ColorReplaceEffect }
// | element duotone { a_CT_DuotoneEffect }
// | element fillOverlay { a_CT_FillOverlayEffect }
// | element grayscl { a_CT_GrayscaleEffect }
// | element hsl { a_CT_HSLEffect }
// | element lum { a_CT_LuminanceEffect }
// | element tint { a_CT_TintEffect })*,
// element extLst { a_CT_OfficeArtExtensionList }?
// a_AG_Blob = r_embed?, r_link?

// Binary large image or picture
// a_CT_Blip =
type Blip struct {

	// a_AG_Blob,
	// a_AG_Blob = r_embed?, r_link?
	EmbedID string `xml:"r:embed,attr,omitempty"`

	// ## default value: none
	// attribute cstate { a_ST_BlipCompression }?,
	// CState string `xml:"cstate,attr,omitempty"`

	// (element alphaBiLevel { a_CT_AlphaBiLevelEffect }
	// | element alphaCeiling { a_CT_AlphaCeilingEffect }
	// | element alphaFloor { a_CT_AlphaFloorEffect }
	// | element alphaInv { a_CT_AlphaInverseEffect }
	// | element alphaMod { a_CT_AlphaModulateEffect }
	// | element alphaModFix { a_CT_AlphaModulateFixedEffect }
	// | element alphaRepl { a_CT_AlphaReplaceEffect }
	// | element biLevel { a_CT_BiLevelEffect }
	// | element blur { a_CT_BlurEffect }
	// | element clrChange { a_CT_ColorChangeEffect }
	// | element clrRepl { a_CT_ColorReplaceEffect }
	// | element duotone { a_CT_DuotoneEffect }
	// | element fillOverlay { a_CT_FillOverlayEffect }
	// | element grayscl { a_CT_GrayscaleEffect }
	// | element hsl { a_CT_HSLEffect }
	// | element lum { a_CT_LuminanceEffect }
	// | element tint { a_CT_TintEffect })*,

	// element extLst { a_CT_OfficeArtExtensionList }?
}
