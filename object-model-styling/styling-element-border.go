package objectmodelstyling

import "strings"

type ElementBorder struct {
	BlockColor        string `json:"blockColor,omitempty"`
	BlockEnd          string `json:"blockEnd,omitempty"`
	BlockEndColor     string `json:"blockEndColor,omitempty"`
	BlockEndStyle     string `json:"blockEndStyle,omitempty"`
	BlockEndWidth     string `json:"blockEndWidth,omitempty"`
	BlockStart        string `json:"blockStart,omitempty"`
	BlockStartColor   string `json:"blockStartColor,omitempty"`
	BlockStartStyle   string `json:"blockStartStyle,omitempty"`
	BlockStartWidth   string `json:"blockStartWidth,omitempty"`
	BlockStyle        string `json:"blockStyle,omitempty"`
	BlockWidth        string `json:"blockWidth,omitempty"`
	BottomColor       string `json:"bottomColor,omitempty"`
	BottomLeftRadius  string `json:"bottomLeftRadius,omitempty"`
	BottomRightRadius string `json:"bottomRightRadius,omitempty"`
	BottomStyle       string `json:"bottomStyle,omitempty"`
	BottomWidth       string `json:"bottomWidth,omitempty"`
	Collapse          string `json:"collapse,omitempty"`
	Color             string `json:"color,omitempty"`
	EndEndRadius      string `json:"endEndRadius,omitempty"`
	EndStartRadius    string `json:"endStartRadius,omitempty"`
	Image             string `json:"image,omitempty"`
	ImageOutset       string `json:"imageOutset,omitempty"`
	ImageRepeat       string `json:"imageRepeat,omitempty"`
	ImageSlice        string `json:"imageSlice,omitempty"`
	ImageSource       string `json:"imageSource,omitempty"`
	ImageWidth        string `json:"imageWidth,omitempty"`
	Inline            string `json:"inline,omitempty"`
	InlineColor       string `json:"inlineColor,omitempty"`
	InlineEnd         string `json:"inlineEnd,omitempty"`
	InlineEndColor    string `json:"inlineEndColor,omitempty"`
	InlineEndStyle    string `json:"inlineEndStyle,omitempty"`
	InlineEndWidth    string `json:"inlineEndWidth,omitempty"`
	InlineStart       string `json:"inlineStart,omitempty"`
	InlineStartColor  string `json:"inlineStartColor,omitempty"`
	InlineStartStyle  string `json:"inlineStartStyle,omitempty"`
	InlineStartWidth  string `json:"inlineStartWidth,omitempty"`
	InlineStyle       string `json:"inlineStyle,omitempty"`
	InlineWidth       string `json:"inlineWidth,omitempty"`
	LeftColor         string `json:"leftColor,omitempty"`
	LeftStyle         string `json:"leftStyle,omitempty"`
	LeftWidth         string `json:"leftWidth,omitempty"`
	RightColor        string `json:"rightColor,omitempty"`
	RightStyle        string `json:"rightStyle,omitempty"`
	RightWidth        string `json:"rightWidth,omitempty"`
	Radius            string `json:"radius,omitempty"`
	Spacing           string `json:"spacing,omitempty"`
	StartEndRadius    string `json:"startEndRadius,omitempty"`
	StartStartRadius  string `json:"startStartRadius,omitempty"`
	Style             string `json:"style,omitempty"`
	TopColor          string `json:"topColor,omitempty"`
	TopLeftRadius     string `json:"topLeftRadius,omitempty"`
	TopRightRadius    string `json:"topRightRadius,omitempty"`
	TopStyle          string `json:"topStyle,omitempty"`
	TopWidth          string `json:"topWidth,omitempty"`
	Width             string `json:"width,omitempty"`
}

const (
	StylingK_ElementBorder_BlockColor        = "border-block-color"
	StylingK_ElementBorder_BlockEnd          = "border-block-end"
	StylingK_ElementBorder_BlockEndColor     = "border-block-end-color"
	StylingK_ElementBorder_BlockEndStyle     = "border-block-end-style"
	StylingK_ElementBorder_BlockEndWidth     = "border-block-end-width"
	StylingK_ElementBorder_BlockStart        = "border-block-start"
	StylingK_ElementBorder_BlockStartColor   = "border-block-start-color"
	StylingK_ElementBorder_BlockStartStyle   = "border-block-start-style"
	StylingK_ElementBorder_BlockStartWidth   = "border-block-start-width"
	StylingK_ElementBorder_BlockStyle        = "border-block-style"
	StylingK_ElementBorder_BlockWidth        = "border-block-width"
	StylingK_ElementBorder_BottomColor       = "border-bottom-color"
	StylingK_ElementBorder_BottomLeftRadius  = "border-bottom-left-radius"
	StylingK_ElementBorder_BottomRightRadius = "border-bottom-right-radius"
	StylingK_ElementBorder_BottomStyle       = "border-bottom-style"
	StylingK_ElementBorder_BottomWidth       = "border-bottom-width"
	StylingK_ElementBorder_Collapse          = "border-collapse"
	StylingK_ElementBorder_Color             = "border-color"
	StylingK_ElementBorder_EndEndRadius      = "border-end-end-radius"
	StylingK_ElementBorder_EndStartRadius    = "border-end-start-radius"
	StylingK_ElementBorder_Image             = "border-image"
	StylingK_ElementBorder_ImageOutset       = "border-image-outset"
	StylingK_ElementBorder_ImageRepeat       = "border-image-repeat"
	StylingK_ElementBorder_ImageSlice        = "border-image-slice"
	StylingK_ElementBorder_ImageSource       = "border-image-source"
	StylingK_ElementBorder_ImageWidth        = "border-image-width"
	StylingK_ElementBorder_Inline            = "border-inline"
	StylingK_ElementBorder_InlineColor       = "border-inline-color"
	StylingK_ElementBorder_InlineEnd         = "border-inline-end"
	StylingK_ElementBorder_InlineEndColor    = "border-inline-end-color"
	StylingK_ElementBorder_InlineEndStyle    = "border-inline-end-style"
	StylingK_ElementBorder_InlineEndWidth    = "border-inline-end-width"
	StylingK_ElementBorder_InlineStart       = "border-inline-start"
	StylingK_ElementBorder_InlineStartColor  = "border-inline-start-color"
	StylingK_ElementBorder_InlineStartStyle  = "border-inline-start-style"
	StylingK_ElementBorder_InlineStartWidth  = "border-inline-start-width"
	StylingK_ElementBorder_InlineStyle       = "border-inline-style"
	StylingK_ElementBorder_InlineWidth       = "border-inline-width"
	StylingK_ElementBorder_LeftColor         = "border-left-color"
	StylingK_ElementBorder_LeftStyle         = "border-left-style"
	StylingK_ElementBorder_LeftWidth         = "border-left-width"
	StylingK_ElementBorder_RightColor        = "border-right-color"
	StylingK_ElementBorder_RightStyle        = "border-right-style"
	StylingK_ElementBorder_RightWidth        = "border-right-width"
	StylingK_ElementBorder_Radius            = "border-radius"
	StylingK_ElementBorder_Spacing           = "border-spacing"
	StylingK_ElementBorder_StartEndRadius    = "border-start-end-radius"
	StylingK_ElementBorder_StartStartRadius  = "border-start-start-radius"
	StylingK_ElementBorder_Style             = "border-style"
	StylingK_ElementBorder_TopColor          = "border-top-color"
	StylingK_ElementBorder_TopLeftRadius     = "border-top-left-radius"
	StylingK_ElementBorder_TopRightRadius    = "border-top-right-radius"
	StylingK_ElementBorder_TopStyle          = "border-top-style"
	StylingK_ElementBorder_TopWidth          = "border-top-width"
	StylingK_ElementBorder_Width             = "border-width"
)

var StylingK_ElementBorder = map[string]string{
	StylingK_ElementBorder_BlockColor:        "",
	StylingK_ElementBorder_BlockEnd:          "",
	StylingK_ElementBorder_BlockEndColor:     "",
	StylingK_ElementBorder_BlockEndStyle:     "",
	StylingK_ElementBorder_BlockEndWidth:     "",
	StylingK_ElementBorder_BlockStart:        "",
	StylingK_ElementBorder_BlockStartColor:   "",
	StylingK_ElementBorder_BlockStartStyle:   "",
	StylingK_ElementBorder_BlockStartWidth:   "",
	StylingK_ElementBorder_BlockStyle:        "",
	StylingK_ElementBorder_BlockWidth:        "",
	StylingK_ElementBorder_BottomColor:       "",
	StylingK_ElementBorder_BottomLeftRadius:  "",
	StylingK_ElementBorder_BottomRightRadius: "",
	StylingK_ElementBorder_BottomStyle:       "",
	StylingK_ElementBorder_BottomWidth:       "",
	StylingK_ElementBorder_Collapse:          "",
	StylingK_ElementBorder_Color:             "",
	StylingK_ElementBorder_EndEndRadius:      "",
	StylingK_ElementBorder_EndStartRadius:    "",
	StylingK_ElementBorder_Image:             "",
	StylingK_ElementBorder_ImageOutset:       "",
	StylingK_ElementBorder_ImageRepeat:       "",
	StylingK_ElementBorder_ImageSlice:        "",
	StylingK_ElementBorder_ImageSource:       "",
	StylingK_ElementBorder_ImageWidth:        "",
	StylingK_ElementBorder_Inline:            "",
	StylingK_ElementBorder_InlineColor:       "",
	StylingK_ElementBorder_InlineEnd:         "",
	StylingK_ElementBorder_InlineEndColor:    "",
	StylingK_ElementBorder_InlineEndStyle:    "",
	StylingK_ElementBorder_InlineEndWidth:    "",
	StylingK_ElementBorder_InlineStart:       "",
	StylingK_ElementBorder_InlineStartColor:  "",
	StylingK_ElementBorder_InlineStartStyle:  "",
	StylingK_ElementBorder_InlineStartWidth:  "",
	StylingK_ElementBorder_InlineStyle:       "",
	StylingK_ElementBorder_InlineWidth:       "",
	StylingK_ElementBorder_LeftColor:         "",
	StylingK_ElementBorder_LeftStyle:         "",
	StylingK_ElementBorder_LeftWidth:         "",
	StylingK_ElementBorder_RightColor:        "",
	StylingK_ElementBorder_RightStyle:        "",
	StylingK_ElementBorder_RightWidth:        "",
	StylingK_ElementBorder_Radius:            "",
	StylingK_ElementBorder_Spacing:           "",
	StylingK_ElementBorder_StartEndRadius:    "",
	StylingK_ElementBorder_StartStartRadius:  "",
	StylingK_ElementBorder_Style:             "",
	StylingK_ElementBorder_TopColor:          "",
	StylingK_ElementBorder_TopLeftRadius:     "",
	StylingK_ElementBorder_TopRightRadius:    "",
	StylingK_ElementBorder_TopStyle:          "",
	StylingK_ElementBorder_TopWidth:          "",
	StylingK_ElementBorder_Width:             "",
}

// Load -
func (thisRef *ElementBorder) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementBorder_BlockColor) == 0 {
		thisRef.BlockColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockEnd) == 0 {
		thisRef.BlockEnd = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockEndColor) == 0 {
		thisRef.BlockEndColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockEndStyle) == 0 {
		thisRef.BlockEndStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockEndWidth) == 0 {
		thisRef.BlockEndWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockStart) == 0 {
		thisRef.BlockStart = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockStartColor) == 0 {
		thisRef.BlockStartColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockStartStyle) == 0 {
		thisRef.BlockStartStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockStartWidth) == 0 {
		thisRef.BlockStartWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockStyle) == 0 {
		thisRef.BlockStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_BlockWidth) == 0 {
		thisRef.BlockWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_BottomColor) == 0 {
		thisRef.BottomColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_BottomLeftRadius) == 0 {
		thisRef.BottomLeftRadius = value
	} else if strings.Compare(key, StylingK_ElementBorder_BottomRightRadius) == 0 {
		thisRef.BottomRightRadius = value
	} else if strings.Compare(key, StylingK_ElementBorder_BottomStyle) == 0 {
		thisRef.BottomStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_BottomWidth) == 0 {
		thisRef.BottomWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_Collapse) == 0 {
		thisRef.Collapse = value
	} else if strings.Compare(key, StylingK_ElementBorder_Color) == 0 {
		thisRef.Color = value
	} else if strings.Compare(key, StylingK_ElementBorder_EndEndRadius) == 0 {
		thisRef.EndEndRadius = value
	} else if strings.Compare(key, StylingK_ElementBorder_EndStartRadius) == 0 {
		thisRef.EndStartRadius = value
	} else if strings.Compare(key, StylingK_ElementBorder_Image) == 0 {
		thisRef.Image = value
	} else if strings.Compare(key, StylingK_ElementBorder_ImageOutset) == 0 {
		thisRef.ImageOutset = value
	} else if strings.Compare(key, StylingK_ElementBorder_ImageRepeat) == 0 {
		thisRef.ImageRepeat = value
	} else if strings.Compare(key, StylingK_ElementBorder_ImageSlice) == 0 {
		thisRef.ImageSlice = value
	} else if strings.Compare(key, StylingK_ElementBorder_ImageSource) == 0 {
		thisRef.ImageSource = value
	} else if strings.Compare(key, StylingK_ElementBorder_ImageWidth) == 0 {
		thisRef.ImageWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_Inline) == 0 {
		thisRef.Inline = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineColor) == 0 {
		thisRef.InlineColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineEnd) == 0 {
		thisRef.InlineEnd = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineEndColor) == 0 {
		thisRef.InlineEndColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineEndStyle) == 0 {
		thisRef.InlineEndStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineEndWidth) == 0 {
		thisRef.InlineEndWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineStart) == 0 {
		thisRef.InlineStart = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineStartColor) == 0 {
		thisRef.InlineStartColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineStartStyle) == 0 {
		thisRef.InlineStartStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineStartWidth) == 0 {
		thisRef.InlineStartWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineStyle) == 0 {
		thisRef.InlineStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_InlineWidth) == 0 {
		thisRef.InlineWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_LeftColor) == 0 {
		thisRef.LeftColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_LeftStyle) == 0 {
		thisRef.LeftStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_LeftWidth) == 0 {
		thisRef.LeftWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_RightColor) == 0 {
		thisRef.RightColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_RightStyle) == 0 {
		thisRef.RightStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_RightWidth) == 0 {
		thisRef.RightWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_Radius) == 0 {
		thisRef.Radius = value
	} else if strings.Compare(key, StylingK_ElementBorder_Spacing) == 0 {
		thisRef.Spacing = value
	} else if strings.Compare(key, StylingK_ElementBorder_StartEndRadius) == 0 {
		thisRef.StartEndRadius = value
	} else if strings.Compare(key, StylingK_ElementBorder_StartStartRadius) == 0 {
		thisRef.StartStartRadius = value
	} else if strings.Compare(key, StylingK_ElementBorder_Style) == 0 {
		thisRef.Style = value
	} else if strings.Compare(key, StylingK_ElementBorder_TopColor) == 0 {
		thisRef.TopColor = value
	} else if strings.Compare(key, StylingK_ElementBorder_TopLeftRadius) == 0 {
		thisRef.TopLeftRadius = value
	} else if strings.Compare(key, StylingK_ElementBorder_TopRightRadius) == 0 {
		thisRef.TopRightRadius = value
	} else if strings.Compare(key, StylingK_ElementBorder_TopStyle) == 0 {
		thisRef.TopStyle = value
	} else if strings.Compare(key, StylingK_ElementBorder_TopWidth) == 0 {
		thisRef.TopWidth = value
	} else if strings.Compare(key, StylingK_ElementBorder_Width) == 0 {
		thisRef.Width = value
	}
}
