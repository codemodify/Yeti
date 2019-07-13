package objectmodelstyling

import "strings"

type ElementText struct {
	Align             string `json:"align,omitempty"`
	AlignLast         string `json:"alignLast,omitempty"`
	CombineUpright    string `json:"combineUpright,omitempty"`
	Decoration        string `json:"decoration,omitempty"`
	DecorationColor   string `json:"decorationColor,omitempty"`
	DecorationLine    string `json:"decorationLine,omitempty"`
	DecorationStyle   string `json:"decorationStyle,omitempty"`
	Emphasis          string `json:"emphasis,omitempty"`
	EmphasisColor     string `json:"emphasisColor,omitempty"`
	EmphasisPosition  string `json:"emphasisPosition,omitempty"`
	EmphasisStyle     string `json:"emphasisStyle,omitempty"`
	Indent            string `json:"indent,omitempty"`
	Justify           string `json:"justify,omitempty"`
	Orientation       string `json:"orientation,omitempty"`
	Overflow          string `json:"overflow,omitempty"`
	Rendering         string `json:"rendering,omitempty"`
	Shadow            string `json:"shadow,omitempty"`
	Transform         string `json:"transform,omitempty"`
	UnderlinePosition string `json:"underlinePosition,omitempty"`
}

const (
	StylingK_ElementText_Align             = "text-align"
	StylingK_ElementText_AlignLast         = "text-align-last"
	StylingK_ElementText_CombineUpright    = "text-combine-upright"
	StylingK_ElementText_Decoration        = "text-decoration"
	StylingK_ElementText_DecorationColor   = "text-decoration-color"
	StylingK_ElementText_DecorationLine    = "text-decoration-line"
	StylingK_ElementText_DecorationStyle   = "text-decoration-style"
	StylingK_ElementText_Emphasis          = "text-emphasis"
	StylingK_ElementText_EmphasisColor     = "text-emphasis-color"
	StylingK_ElementText_EmphasisPosition  = "text-emphasis-position"
	StylingK_ElementText_EmphasisStyle     = "text-emphasis-style"
	StylingK_ElementText_Indent            = "text-indent"
	StylingK_ElementText_Justify           = "text-justify"
	StylingK_ElementText_Orientation       = "text-orientation"
	StylingK_ElementText_Overflow          = "text-overflow"
	StylingK_ElementText_Rendering         = "text-rendering"
	StylingK_ElementText_Shadow            = "text-shadow"
	StylingK_ElementText_Transform         = "text-transform"
	StylingK_ElementText_UnderlinePosition = "text-underline-position"
)

var StylingK_ElementText = map[string]string{
	StylingK_ElementText_Align:             "",
	StylingK_ElementText_AlignLast:         "",
	StylingK_ElementText_CombineUpright:    "",
	StylingK_ElementText_Decoration:        "",
	StylingK_ElementText_DecorationColor:   "",
	StylingK_ElementText_DecorationLine:    "",
	StylingK_ElementText_DecorationStyle:   "",
	StylingK_ElementText_Emphasis:          "",
	StylingK_ElementText_EmphasisColor:     "",
	StylingK_ElementText_EmphasisPosition:  "",
	StylingK_ElementText_EmphasisStyle:     "",
	StylingK_ElementText_Indent:            "",
	StylingK_ElementText_Justify:           "",
	StylingK_ElementText_Orientation:       "",
	StylingK_ElementText_Overflow:          "",
	StylingK_ElementText_Rendering:         "",
	StylingK_ElementText_Shadow:            "",
	StylingK_ElementText_Transform:         "",
	StylingK_ElementText_UnderlinePosition: "",
}

// Load -
func (thisRef *ElementText) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementText_Align) == 0 {
		thisRef.Align = value
	} else if strings.Compare(key, StylingK_ElementText_AlignLast) == 0 {
		thisRef.AlignLast = value
	} else if strings.Compare(key, StylingK_ElementText_CombineUpright) == 0 {
		thisRef.CombineUpright = value
	} else if strings.Compare(key, StylingK_ElementText_Decoration) == 0 {
		thisRef.Decoration = value
	} else if strings.Compare(key, StylingK_ElementText_DecorationColor) == 0 {
		thisRef.DecorationColor = value
	} else if strings.Compare(key, StylingK_ElementText_DecorationLine) == 0 {
		thisRef.DecorationLine = value
	} else if strings.Compare(key, StylingK_ElementText_DecorationStyle) == 0 {
		thisRef.DecorationStyle = value
	} else if strings.Compare(key, StylingK_ElementText_Emphasis) == 0 {
		thisRef.Emphasis = value
	} else if strings.Compare(key, StylingK_ElementText_EmphasisColor) == 0 {
		thisRef.EmphasisColor = value
	} else if strings.Compare(key, StylingK_ElementText_EmphasisPosition) == 0 {
		thisRef.EmphasisPosition = value
	} else if strings.Compare(key, StylingK_ElementText_EmphasisStyle) == 0 {
		thisRef.EmphasisStyle = value
	} else if strings.Compare(key, StylingK_ElementText_Indent) == 0 {
		thisRef.Indent = value
	} else if strings.Compare(key, StylingK_ElementText_Justify) == 0 {
		thisRef.Justify = value
	} else if strings.Compare(key, StylingK_ElementText_Orientation) == 0 {
		thisRef.Orientation = value
	} else if strings.Compare(key, StylingK_ElementText_Overflow) == 0 {
		thisRef.Overflow = value
	} else if strings.Compare(key, StylingK_ElementText_Rendering) == 0 {
		thisRef.Rendering = value
	} else if strings.Compare(key, StylingK_ElementText_Shadow) == 0 {
		thisRef.Shadow = value
	} else if strings.Compare(key, StylingK_ElementText_Transform) == 0 {
		thisRef.Transform = value
	} else if strings.Compare(key, StylingK_ElementText_UnderlinePosition) == 0 {
		thisRef.UnderlinePosition = value
	}
}
