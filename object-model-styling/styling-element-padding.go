package objectmodelstyling

type ElementPadding struct {
	Block       string `json:"block,omitempty"`
	BlockEnd    string `json:"blockEnd,omitempty"`
	BlockStart  string `json:"blockStart,omitempty"`
	Bottom      string `json:"bottom,omitempty"`
	Inline      string `json:"inline,omitempty"`
	InlineEnd   string `json:"inlineEnd,omitempty"`
	InlineStart string `json:"inlineStart,omitempty"`
	Left        string `json:"left,omitempty"`
	Right       string `json:"right,omitempty"`
	Top         string `json:"top,omitempty"`
}

// const (
// 	StylingK_ElementPadding_Block       = "padding-block"
// 	StylingK_ElementPadding_BlockEnd    = "padding-block-end"
// 	StylingK_ElementPadding_BlockStart  = "padding-block-start"
// 	StylingK_ElementPadding_Bottom      = "padding-bottom"
// 	StylingK_ElementPadding_Inline      = "padding-inline"
// 	StylingK_ElementPadding_InlineEnd   = "padding-inline-end"
// 	StylingK_ElementPadding_InlineStart = "padding-inline-start"
// 	StylingK_ElementPadding_Left        = "padding-left"
// 	StylingK_ElementPadding_Right       = "padding-right"
// 	StylingK_ElementPadding_Top         = "padding-top"
// )

var StylingK_ElementPadding = map[string]string{
	"padding-block":        "",
	"padding-block-end":    "",
	"padding-block-start":  "",
	"padding-bottom":       "",
	"padding-inline":       "",
	"padding-inline-end":   "",
	"padding-inline-start": "",
	"padding-left":         "",
	"padding-right":        "",
	"padding-top":          "",
}
