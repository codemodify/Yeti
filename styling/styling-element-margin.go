package styling

import "strings"

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

const (
	StylingK_ElementMargin_Block       = "margin-block"
	StylingK_ElementMargin_BlockEnd    = "margin-block-end"
	StylingK_ElementMargin_BlockStart  = "margin-block-start"
	StylingK_ElementMargin_Bottom      = "margin-bottom"
	StylingK_ElementMargin_Inline      = "margin-inline"
	StylingK_ElementMargin_InlineEnd   = "margin-inline-end"
	StylingK_ElementMargin_InlineStart = "margin-inline-start"
	StylingK_ElementMargin_Left        = "margin-left"
	StylingK_ElementMargin_Right       = "margin-right"
	StylingK_ElementMargin_Top         = "margin-top"
)

var StylingK_ElementMargin = map[string]string{
	StylingK_ElementMargin_Block:       "",
	StylingK_ElementMargin_BlockEnd:    "",
	StylingK_ElementMargin_BlockStart:  "",
	StylingK_ElementMargin_Bottom:      "",
	StylingK_ElementMargin_Inline:      "",
	StylingK_ElementMargin_InlineEnd:   "",
	StylingK_ElementMargin_InlineStart: "",
	StylingK_ElementMargin_Left:        "",
	StylingK_ElementMargin_Right:       "",
	StylingK_ElementMargin_Top:         "",
}

// Load -
func (thisRef *ElementMargin) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementMargin_Block) == 0 {
		thisRef.Block = value
	} else if strings.Compare(key, StylingK_ElementMargin_BlockEnd) == 0 {
		thisRef.BlockEnd = value
	} else if strings.Compare(key, StylingK_ElementMargin_BlockStart) == 0 {
		thisRef.BlockStart = value
	} else if strings.Compare(key, StylingK_ElementMargin_Bottom) == 0 {
		thisRef.Bottom = value
	} else if strings.Compare(key, StylingK_ElementMargin_Inline) == 0 {
		thisRef.Inline = value
	} else if strings.Compare(key, StylingK_ElementMargin_InlineEnd) == 0 {
		thisRef.InlineEnd = value
	} else if strings.Compare(key, StylingK_ElementMargin_InlineStart) == 0 {
		thisRef.InlineStart = value
	} else if strings.Compare(key, StylingK_ElementMargin_Left) == 0 {
		thisRef.Left = value
	} else if strings.Compare(key, StylingK_ElementMargin_Right) == 0 {
		thisRef.Right = value
	} else if strings.Compare(key, StylingK_ElementMargin_Top) == 0 {
		thisRef.Top = value
	}
}
