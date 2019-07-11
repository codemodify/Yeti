package objectmodelstyling

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

// const (
// 	StylingK_ElementBorder_BlockColor      = "border-block-color"
// 	StylingK_ElementBorder_BlockEnd        = "border-block-end"
// 	StylingK_ElementBorder_BlockEndColor   = "border-block-end-color"
// 	StylingK_ElementBorder_BlockEndStyle   = "border-block-end-style"
// 	StylingK_ElementBorder_BlockEndWidth   = "border-block-end-width"
// 	StylingK_ElementBorder_BlockStart      = "border-block-start"
// 	StylingK_ElementBorder_BlockStartColor = "border-block-start-color"
// 	StylingK_ElementBorder_BlockStartStyle = "border-block-start-style"
// 	StylingK_ElementBorder_BlockStartWidth = "border-block-start-width"
// 	StylingK_ElementBorder_BlockStyle      = "border-block-style"
// 	StylingK_ElementBorder_BlockWidth      = "border-block-width"

// 	StylingK_ElementBorder_BottomColor       = "border-bottom-color"
// 	StylingK_ElementBorder_BottomLeftRadius  = "border-bottom-left-radius"
// 	StylingK_ElementBorder_BottomRightRadius = "border-bottom-right-radius"
// 	StylingK_ElementBorder_BottomStyle       = "border-bottom-style"
// 	StylingK_ElementBorder_BottomWidth       = "border-bottom-width"

// 	StylingK_ElementBorder_Collapse = "border-collapse"

// 	StylingK_ElementBorder_Color = "border-color"

// 	StylingK_ElementBorder_EndEndRadius   = "border-end-end-radius"
// 	StylingK_ElementBorder_EndStartRadius = "border-end-start-radius"

// 	StylingK_ElementBorder_Image       = "border-image"
// 	StylingK_ElementBorder_ImageOutset = "border-image-outset"
// 	StylingK_ElementBorder_ImageRepeat = "border-image-repeat"
// 	StylingK_ElementBorder_ImageSlice  = "border-image-slice"
// 	StylingK_ElementBorder_ImageSource = "border-image-source"
// 	StylingK_ElementBorder_ImageWidth  = "border-image-width"

// 	StylingK_ElementBorder_Inline           = "border-inline"
// 	StylingK_ElementBorder_InlineColor      = "border-inline-color"
// 	StylingK_ElementBorder_InlineEnd        = "border-inline-end"
// 	StylingK_ElementBorder_InlineEndColor   = "border-inline-end-color"
// 	StylingK_ElementBorder_InlineEndStyle   = "border-inline-end-style"
// 	StylingK_ElementBorder_InlineEndWidth   = "border-inline-end-width"
// 	StylingK_ElementBorder_InlineStart      = "border-inline-start"
// 	StylingK_ElementBorder_InlineStartColor = "border-inline-start-color"
// 	StylingK_ElementBorder_InlineStartStyle = "border-inline-start-style"
// 	StylingK_ElementBorder_InlineStartWidth = "border-inline-start-width"
// 	StylingK_ElementBorder_InlineStyle      = "border-inline-style"
// 	StylingK_ElementBorder_InlineWidth      = "border-inline-width"

// 	StylingK_ElementBorder_LeftColor = "border-left-color"
// 	StylingK_ElementBorder_LeftStyle = "border-left-style"
// 	StylingK_ElementBorder_LeftWidth = "border-left-width"

// 	StylingK_ElementBorder_RightColor = "border-right-color"
// 	StylingK_ElementBorder_RightStyle = "border-right-style"
// 	StylingK_ElementBorder_RightWidth = "border-right-width"

// 	StylingK_ElementBorder_Radius = "border-radius"

// 	StylingK_ElementBorder_Spacing = "border-spacing"

// 	StylingK_ElementBorder_StartEndRadius   = "border-start-end-radius"
// 	StylingK_ElementBorder_StartStartRadius = "border-start-start-radius"

// 	StylingK_ElementBorder_Style = "border-style"

// 	StylingK_ElementBorder_TopColor       = "border-top-color"
// 	StylingK_ElementBorder_TopLeftRadius  = "border-top-left-radius"
// 	StylingK_ElementBorder_TopRightRadius = "border-top-right-radius"
// 	StylingK_ElementBorder_TopStyle       = "border-top-style"
// 	StylingK_ElementBorder_TopWidth       = "border-top-width"

// 	StylingK_ElementBorder_Width = "border-width"
// )

var StylingK_ElementBorder = map[string]string{
	"border-block-color":         "",
	"border-block-end":           "",
	"border-block-end-color":     "",
	"border-block-end-style":     "",
	"border-block-end-width":     "",
	"border-block-start":         "",
	"border-block-start-color":   "",
	"border-block-start-style":   "",
	"border-block-start-width":   "",
	"border-block-style":         "",
	"border-block-width":         "",
	"border-bottom-color":        "",
	"border-bottom-left-radius":  "",
	"border-bottom-right-radius": "",
	"border-bottom-style":        "",
	"border-bottom-width":        "",
	"border-collapse":            "",
	"border-color":               "",
	"border-end-end-radius":      "",
	"border-end-start-radius":    "",
	"border-image":               "",
	"border-image-outset":        "",
	"border-image-repeat":        "",
	"border-image-slice":         "",
	"border-image-source":        "",
	"border-image-width":         "",
	"border-inline":              "",
	"border-inline-color":        "",
	"border-inline-end":          "",
	"border-inline-end-color":    "",
	"border-inline-end-style":    "",
	"border-inline-end-width":    "",
	"border-inline-start":        "",
	"border-inline-start-color":  "",
	"border-inline-start-style":  "",
	"border-inline-start-width":  "",
	"border-inline-style":        "",
	"border-inline-width":        "",
	"border-left-color":          "",
	"border-left-style":          "",
	"border-left-width":          "",
	"border-right-color":         "",
	"border-right-style":         "",
	"border-right-width":         "",
	"border-radius":              "",
	"border-spacing":             "",
	"border-start-end-radius":    "",
	"border-start-start-radius":  "",
	"border-style":               "",
	"border-top-color":           "",
	"border-top-left-radius":     "",
	"border-top-right-radius":    "",
	"border-top-style":           "",
	"border-top-width":           "",
	"border-width":               "",
}
