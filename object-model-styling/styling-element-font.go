package objectmodelstyling

type ElementFont struct {
	Family            string `json:"family,omitempty"`
	FeatureSettings   string `json:"featureSettings,omitempty"`
	Kerning           string `json:"kerning,omitempty"`
	LanguageOverride  string `json:"languageOverride,omitempty"`
	OpticalSizing     string `json:"opticalSizing,omitempty"`
	Size              string `json:"size,omitempty"`
	SizeAdjust        string `json:"sizeAdjust,omitempty"`
	Stretch           string `json:"stretch,omitempty"`
	Style             string `json:"style,omitempty"`
	Synthesis         string `json:"synthesis,omitempty"`
	Variant           string `json:"variant,omitempty"`
	VariantAlternates string `json:"variantAlternates,omitempty"`
	VariantCaps       string `json:"variantCaps,omitempty"`
	VariantEastAsian  string `json:"variantEastAsian,omitempty"`
	VariantLigatures  string `json:"variantLigatures,omitempty"`
	VariantNumeric    string `json:"variantNumeric,omitempty"`
	VariantPosition   string `json:"variantPosition,omitempty"`
	VariationSettings string `json:"variationSettings,omitempty"`
	Weight            string `json:"weight,omitempty"`
}

// const (
// 	StylingK_ElementFont_Family            = "font-family"
// 	StylingK_ElementFont_FeatureSettings   = "font-feature-settings"
// 	StylingK_ElementFont_Kerning           = "font-kerning"
// 	StylingK_ElementFont_LanguageOverride  = "font-language-override"
// 	StylingK_ElementFont_OpticalSizing     = "font-optical-sizing"
// 	StylingK_ElementFont_Size              = "font-size"
// 	StylingK_ElementFont_SizeAdjust        = "font-size-adjust"
// 	StylingK_ElementFont_Stretch           = "font-stretch"
// 	StylingK_ElementFont_Style             = "font-style"
// 	StylingK_ElementFont_Synthesis         = "font-synthesis"
// 	StylingK_ElementFont_Variant           = "font-variant"
// 	StylingK_ElementFont_VariantAlternates = "font-variant-alternates"
// 	StylingK_ElementFont_VariantCaps       = "font-variant-caps"
// 	StylingK_ElementFont_VariantEastAsian  = "font-variant-east-asian"
// 	StylingK_ElementFont_VariantLigatures  = "font-variant-ligatures"
// 	StylingK_ElementFont_VariantNumeric    = "font-variant-numeric"
// 	StylingK_ElementFont_VariantPosition   = "font-variant-position"
// 	StylingK_ElementFont_VariationSettings = "font-variation-settings"
// 	StylingK_ElementFont_Weight            = "font-weight"
// )

var StylingK_ElementFont = map[string]string{
	"font-family":             "",
	"font-feature-settings":   "",
	"font-kerning":            "",
	"font-language-override":  "",
	"font-optical-sizing":     "",
	"font-size":               "",
	"font-size-adjust":        "",
	"font-stretch":            "",
	"font-style":              "",
	"font-synthesis":          "",
	"font-variant":            "",
	"font-variant-alternates": "",
	"font-variant-caps":       "",
	"font-variant-east-asian": "",
	"font-variant-ligatures":  "",
	"font-variant-numeric":    "",
	"font-variant-position":   "",
	"font-variation-settings": "",
	"font-weight":             "",
}
