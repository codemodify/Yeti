package objectmodelstyling

import "strings"

type ElementScroll struct {
	Behavior           string `json:"behavior,omitempty"`
	MarginBlock        string `json:"marginBlock,omitempty"`
	MarginBlockEnd     string `json:"marginBlockEnd,omitempty"`
	MarginBlockStart   string `json:"marginBlockStart,omitempty"`
	MarginBottom       string `json:"marginBottom,omitempty"`
	MarginInline       string `json:"marginInline,omitempty"`
	MarginInlineEnd    string `json:"marginInlineEnd,omitempty"`
	MarginInlineStart  string `json:"marginInlineStart,omitempty"`
	MarginLeft         string `json:"marginLeft,omitempty"`
	MarginRight        string `json:"marginRight,omitempty"`
	MarginTop          string `json:"marginTop,omitempty"`
	PaddingBlock       string `json:"paddingBlock,omitempty"`
	PaddingBlockEnd    string `json:"paddingBlockEnd,omitempty"`
	PaddingBlockStart  string `json:"paddingBlockStart,omitempty"`
	PaddingBottom      string `json:"paddingBottom,omitempty"`
	PaddingInline      string `json:"paddingInline,omitempty"`
	PaddingInlineEnd   string `json:"paddingInlineEnd,omitempty"`
	PaddingInlineStart string `json:"paddingInlineStart,omitempty"`
	PaddingLeft        string `json:"paddingLeft,omitempty"`
	PaddingRight       string `json:"paddingRight,omitempty"`
	PaddingTop         string `json:"paddingTop,omitempty"`
	SnapAlign          string `json:"snapAlign,omitempty"`
	SnapStop           string `json:"snapStop,omitempty"`
	SnapType           string `json:"snapType,omitempty"`
}

const (
	StylingK_ElementScroll_Behavior           = "scroll-behavior"
	StylingK_ElementScroll_MarginBlock        = "scroll-margin-block"
	StylingK_ElementScroll_MarginBlockEnd     = "scroll-margin-block-end"
	StylingK_ElementScroll_MarginBlockStart   = "scroll-margin-block-start"
	StylingK_ElementScroll_MarginBottom       = "scroll-margin-bottom"
	StylingK_ElementScroll_MarginInline       = "scroll-margin-inline"
	StylingK_ElementScroll_MarginInlineEnd    = "scroll-margin-inline-end"
	StylingK_ElementScroll_MarginInlineStart  = "scroll-margin-inline-start"
	StylingK_ElementScroll_MarginLeft         = "scroll-margin-left"
	StylingK_ElementScroll_MarginRight        = "scroll-margin-right"
	StylingK_ElementScroll_MarginTop          = "scroll-margin-top"
	StylingK_ElementScroll_PaddingBlock       = "scroll-padding-block"
	StylingK_ElementScroll_PaddingBlockEnd    = "scroll-padding-block-end"
	StylingK_ElementScroll_PaddingBlockStart  = "scroll-padding-block-start"
	StylingK_ElementScroll_PaddingBottom      = "scroll-padding-bottom"
	StylingK_ElementScroll_PaddingInline      = "scroll-padding-inline"
	StylingK_ElementScroll_PaddingInlineEnd   = "scroll-padding-inline-end"
	StylingK_ElementScroll_PaddingInlineStart = "scroll-padding-inline-start"
	StylingK_ElementScroll_PaddingLeft        = "scroll-padding-left"
	StylingK_ElementScroll_PaddingRight       = "scroll-padding-right"
	StylingK_ElementScroll_PaddingTop         = "scroll-padding-top"
	StylingK_ElementScroll_SnapAlign          = "scroll-snap-align"
	StylingK_ElementScroll_SnapStop           = "scroll-snap-stop"
	StylingK_ElementScroll_SnapType           = "scroll-snap-type"
)

var StylingK_ElementScroll = map[string]string{
	StylingK_ElementScroll_Behavior:           "",
	StylingK_ElementScroll_MarginBlock:        "",
	StylingK_ElementScroll_MarginBlockEnd:     "",
	StylingK_ElementScroll_MarginBlockStart:   "",
	StylingK_ElementScroll_MarginBottom:       "",
	StylingK_ElementScroll_MarginInline:       "",
	StylingK_ElementScroll_MarginInlineEnd:    "",
	StylingK_ElementScroll_MarginInlineStart:  "",
	StylingK_ElementScroll_MarginLeft:         "",
	StylingK_ElementScroll_MarginRight:        "",
	StylingK_ElementScroll_MarginTop:          "",
	StylingK_ElementScroll_PaddingBlock:       "",
	StylingK_ElementScroll_PaddingBlockEnd:    "",
	StylingK_ElementScroll_PaddingBlockStart:  "",
	StylingK_ElementScroll_PaddingBottom:      "",
	StylingK_ElementScroll_PaddingInline:      "",
	StylingK_ElementScroll_PaddingInlineEnd:   "",
	StylingK_ElementScroll_PaddingInlineStart: "",
	StylingK_ElementScroll_PaddingLeft:        "",
	StylingK_ElementScroll_PaddingRight:       "",
	StylingK_ElementScroll_PaddingTop:         "",
	StylingK_ElementScroll_SnapAlign:          "",
	StylingK_ElementScroll_SnapStop:           "",
	StylingK_ElementScroll_SnapType:           "",
}

// Load -
func (thisRef *ElementScroll) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementScroll_Behavior) == 0 {
		thisRef.Behavior = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginBlock) == 0 {
		thisRef.MarginBlock = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginBlockEnd) == 0 {
		thisRef.MarginBlockEnd = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginBlockStart) == 0 {
		thisRef.MarginBlockStart = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginBottom) == 0 {
		thisRef.MarginBottom = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginInline) == 0 {
		thisRef.MarginInline = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginInlineEnd) == 0 {
		thisRef.MarginInlineEnd = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginInlineStart) == 0 {
		thisRef.MarginInlineStart = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginLeft) == 0 {
		thisRef.MarginLeft = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginRight) == 0 {
		thisRef.MarginRight = value
	} else if strings.Compare(key, StylingK_ElementScroll_MarginTop) == 0 {
		thisRef.MarginTop = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingBlock) == 0 {
		thisRef.PaddingBlock = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingBlockEnd) == 0 {
		thisRef.PaddingBlockEnd = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingBlockStart) == 0 {
		thisRef.PaddingBlockStart = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingBottom) == 0 {
		thisRef.PaddingBottom = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingInline) == 0 {
		thisRef.PaddingInline = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingInlineEnd) == 0 {
		thisRef.PaddingInlineEnd = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingInlineStart) == 0 {
		thisRef.PaddingInlineStart = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingLeft) == 0 {
		thisRef.PaddingLeft = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingRight) == 0 {
		thisRef.PaddingRight = value
	} else if strings.Compare(key, StylingK_ElementScroll_PaddingTop) == 0 {
		thisRef.PaddingTop = value
	} else if strings.Compare(key, StylingK_ElementScroll_SnapAlign) == 0 {
		thisRef.SnapAlign = value
	} else if strings.Compare(key, StylingK_ElementScroll_SnapStop) == 0 {
		thisRef.SnapStop = value
	} else if strings.Compare(key, StylingK_ElementScroll_SnapType) == 0 {
		thisRef.SnapType = value
	}
}
