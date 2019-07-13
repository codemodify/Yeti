package styling

import "strings"

type ElementInset struct {
	Block       string `json:"block,omitempty"`
	BlockEnd    string `json:"blockEnd,omitempty"`
	BlockStart  string `json:"blockStart,omitempty"`
	Inline      string `json:"inline,omitempty"`
	InlineEnd   string `json:"inlineEnd,omitempty"`
	InlineStart string `json:"inlineStart,omitempty"`
}

const (
	StylingK_ElementInset_Block       = "inset-block"
	StylingK_ElementInset_BlockEnd    = "inset-block-end"
	StylingK_ElementInset_BlockStart  = "inset-block-start"
	StylingK_ElementInset_Inline      = "inset-inline"
	StylingK_ElementInset_InlineEnd   = "inset-inline-end"
	StylingK_ElementInset_InlineStart = "inset-inline-start"
)

var StylingK_ElementInset = map[string]string{
	StylingK_ElementInset_Block:       "",
	StylingK_ElementInset_BlockEnd:    "",
	StylingK_ElementInset_BlockStart:  "",
	StylingK_ElementInset_Inline:      "",
	StylingK_ElementInset_InlineEnd:   "",
	StylingK_ElementInset_InlineStart: "",
}

// Load -
func (thisRef *ElementInset) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementInset_Block) == 0 {
		thisRef.Block = value
	} else if strings.Compare(key, StylingK_ElementInset_BlockEnd) == 0 {
		thisRef.BlockEnd = value
	} else if strings.Compare(key, StylingK_ElementInset_BlockStart) == 0 {
		thisRef.BlockStart = value
	} else if strings.Compare(key, StylingK_ElementInset_Inline) == 0 {
		thisRef.Inline = value
	} else if strings.Compare(key, StylingK_ElementInset_InlineEnd) == 0 {
		thisRef.InlineEnd = value
	} else if strings.Compare(key, StylingK_ElementInset_InlineStart) == 0 {
		thisRef.InlineStart = value
	}
}
