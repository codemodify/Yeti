package objectmodelstyling

type ElementInset struct {
	Block       string `json:"block,omitempty"`
	BlockEnd    string `json:"blockEnd,omitempty"`
	BlockStart  string `json:"blockStart,omitempty"`
	Inline      string `json:"inline,omitempty"`
	InlineEnd   string `json:"inlineEnd,omitempty"`
	InlineStart string `json:"inlineStart,omitempty"`
}

// const (
// 	StylingK_ElementInset_Block       = "inset-block"
// 	StylingK_ElementInset_BlockEnd    = "inset-block-end"
// 	StylingK_ElementInset_BlockStart  = "inset-block-start"
// 	StylingK_ElementInset_Inline      = "inset-inline"
// 	StylingK_ElementInset_InlineEnd   = "inset-inline-end"
// 	StylingK_ElementInset_InlineStart = "inset-inline-start"
// )

var StylingK_ElementInset = map[string]string{
	"inset-block":        "",
	"inset-block-end":    "",
	"inset-block-start":  "",
	"inset-inline":       "",
	"inset-inline-end":   "",
	"inset-inline-start": "",
}
