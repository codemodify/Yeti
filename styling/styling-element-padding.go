package styling

import "strings"

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

const (
	StylingK_ElementPadding_Block       = "padding-block"
	StylingK_ElementPadding_BlockEnd    = "padding-block-end"
	StylingK_ElementPadding_BlockStart  = "padding-block-start"
	StylingK_ElementPadding_Bottom      = "padding-bottom"
	StylingK_ElementPadding_Inline      = "padding-inline"
	StylingK_ElementPadding_InlineEnd   = "padding-inline-end"
	StylingK_ElementPadding_InlineStart = "padding-inline-start"
	StylingK_ElementPadding_Left        = "padding-left"
	StylingK_ElementPadding_Right       = "padding-right"
	StylingK_ElementPadding_Top         = "padding-top"
)

var StylingK_ElementPadding = map[string]string{
	StylingK_ElementPadding_Block:       "",
	StylingK_ElementPadding_BlockEnd:    "",
	StylingK_ElementPadding_BlockStart:  "",
	StylingK_ElementPadding_Bottom:      "",
	StylingK_ElementPadding_Inline:      "",
	StylingK_ElementPadding_InlineEnd:   "",
	StylingK_ElementPadding_InlineStart: "",
	StylingK_ElementPadding_Left:        "",
	StylingK_ElementPadding_Right:       "",
	StylingK_ElementPadding_Top:         "",
}

// Load -
func (thisRef *ElementPadding) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementPadding_Block) == 0 {
		thisRef.Block = value
	} else if strings.Compare(key, StylingK_ElementPadding_BlockEnd) == 0 {
		thisRef.BlockEnd = value
	} else if strings.Compare(key, StylingK_ElementPadding_BlockStart) == 0 {
		thisRef.BlockStart = value
	} else if strings.Compare(key, StylingK_ElementPadding_Bottom) == 0 {
		thisRef.Bottom = value
	} else if strings.Compare(key, StylingK_ElementPadding_Inline) == 0 {
		thisRef.Inline = value
	} else if strings.Compare(key, StylingK_ElementPadding_InlineEnd) == 0 {
		thisRef.InlineEnd = value
	} else if strings.Compare(key, StylingK_ElementPadding_InlineStart) == 0 {
		thisRef.InlineStart = value
	} else if strings.Compare(key, StylingK_ElementPadding_Left) == 0 {
		thisRef.Left = value
	} else if strings.Compare(key, StylingK_ElementPadding_Right) == 0 {
		thisRef.Right = value
	} else if strings.Compare(key, StylingK_ElementPadding_Top) == 0 {
		thisRef.Top = value
	}
}
