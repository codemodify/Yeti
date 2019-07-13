package objectmodelstyling

import "strings"

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

const (
	StylingK_ElementFont_Family            = "font-family"
	StylingK_ElementFont_FeatureSettings   = "font-feature-settings"
	StylingK_ElementFont_Kerning           = "font-kerning"
	StylingK_ElementFont_LanguageOverride  = "font-language-override"
	StylingK_ElementFont_OpticalSizing     = "font-optical-sizing"
	StylingK_ElementFont_Size              = "font-size"
	StylingK_ElementFont_SizeAdjust        = "font-size-adjust"
	StylingK_ElementFont_Stretch           = "font-stretch"
	StylingK_ElementFont_Style             = "font-style"
	StylingK_ElementFont_Synthesis         = "font-synthesis"
	StylingK_ElementFont_Variant           = "font-variant"
	StylingK_ElementFont_VariantAlternates = "font-variant-alternates"
	StylingK_ElementFont_VariantCaps       = "font-variant-caps"
	StylingK_ElementFont_VariantEastAsian  = "font-variant-east-asian"
	StylingK_ElementFont_VariantLigatures  = "font-variant-ligatures"
	StylingK_ElementFont_VariantNumeric    = "font-variant-numeric"
	StylingK_ElementFont_VariantPosition   = "font-variant-position"
	StylingK_ElementFont_VariationSettings = "font-variation-settings"
	StylingK_ElementFont_Weight            = "font-weight"
)

var StylingK_ElementFont = map[string]string{
	StylingK_ElementFont_Family:            "",
	StylingK_ElementFont_FeatureSettings:   "",
	StylingK_ElementFont_Kerning:           "",
	StylingK_ElementFont_LanguageOverride:  "",
	StylingK_ElementFont_OpticalSizing:     "",
	StylingK_ElementFont_Size:              "",
	StylingK_ElementFont_SizeAdjust:        "",
	StylingK_ElementFont_Stretch:           "",
	StylingK_ElementFont_Style:             "",
	StylingK_ElementFont_Synthesis:         "",
	StylingK_ElementFont_Variant:           "",
	StylingK_ElementFont_VariantAlternates: "",
	StylingK_ElementFont_VariantCaps:       "",
	StylingK_ElementFont_VariantEastAsian:  "",
	StylingK_ElementFont_VariantLigatures:  "",
	StylingK_ElementFont_VariantNumeric:    "",
	StylingK_ElementFont_VariantPosition:   "",
	StylingK_ElementFont_VariationSettings: "",
	StylingK_ElementFont_Weight:            "",
}

// Load -
func (thisRef *ElementFont) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementFont_Family) == 0 {
		thisRef.Family = value
	} else if strings.Compare(key, StylingK_ElementFont_FeatureSettings) == 0 {
		thisRef.FeatureSettings = value
	} else if strings.Compare(key, StylingK_ElementFont_Kerning) == 0 {
		thisRef.Kerning = value
	} else if strings.Compare(key, StylingK_ElementFont_LanguageOverride) == 0 {
		thisRef.LanguageOverride = value
	} else if strings.Compare(key, StylingK_ElementFont_OpticalSizing) == 0 {
		thisRef.OpticalSizing = value
	} else if strings.Compare(key, StylingK_ElementFont_Size) == 0 {
		thisRef.Size = value
	} else if strings.Compare(key, StylingK_ElementFont_SizeAdjust) == 0 {
		thisRef.SizeAdjust = value
	} else if strings.Compare(key, StylingK_ElementFont_Stretch) == 0 {
		thisRef.Stretch = value
	} else if strings.Compare(key, StylingK_ElementFont_Style) == 0 {
		thisRef.Style = value
	} else if strings.Compare(key, StylingK_ElementFont_Synthesis) == 0 {
		thisRef.Synthesis = value
	} else if strings.Compare(key, StylingK_ElementFont_Variant) == 0 {
		thisRef.Variant = value
	} else if strings.Compare(key, StylingK_ElementFont_VariantAlternates) == 0 {
		thisRef.VariantAlternates = value
	} else if strings.Compare(key, StylingK_ElementFont_VariantCaps) == 0 {
		thisRef.VariantCaps = value
	} else if strings.Compare(key, StylingK_ElementFont_VariantEastAsian) == 0 {
		thisRef.VariantEastAsian = value
	} else if strings.Compare(key, StylingK_ElementFont_VariantLigatures) == 0 {
		thisRef.VariantLigatures = value
	} else if strings.Compare(key, StylingK_ElementFont_VariantNumeric) == 0 {
		thisRef.VariantNumeric = value
	} else if strings.Compare(key, StylingK_ElementFont_VariantPosition) == 0 {
		thisRef.VariantPosition = value
	} else if strings.Compare(key, StylingK_ElementFont_VariationSettings) == 0 {
		thisRef.VariationSettings = value
	} else if strings.Compare(key, StylingK_ElementFont_Weight) == 0 {
		thisRef.Weight = value
	}
}
