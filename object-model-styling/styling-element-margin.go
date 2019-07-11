package objectmodelstyling

type ElementMargin struct {
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
// 	StylingK_ElementMargin_Block       = "margin-block"
// 	StylingK_ElementMargin_BlockEnd    = "margin-block-end"
// 	StylingK_ElementMargin_BlockStart  = "margin-block-start"
// 	StylingK_ElementMargin_Bottom      = "margin-bottom"
// 	StylingK_ElementMargin_Inline      = "margin-inline"
// 	StylingK_ElementMargin_InlineEnd   = "margin-inline-end"
// 	StylingK_ElementMargin_InlineStart = "margin-inline-start"
// 	StylingK_ElementMargin_Left        = "margin-left"
// 	StylingK_ElementMargin_Right       = "margin-right"
// 	StylingK_ElementMargin_Top         = "margin-top"
// )

var StylingK_ElementMargin = map[string]string{
	"margin-block":        "",
	"margin-block-end":    "",
	"margin-block-start":  "",
	"margin-bottom":       "",
	"margin-inline":       "",
	"margin-inline-end":   "",
	"margin-inline-start": "",
	"margin-left":         "",
	"margin-right":        "",
	"margin-top":          "",
}
