package objectmodelstyling

import "strings"

type ElementWord struct {
	Break   string `json:"break,omitempty"`
	Spacing string `json:"spacing,omitempty"`
	Wrap    string `json:"wrap,omitempty"`
}

const (
	StylingK_ElementWord_Break   = "word-break"
	StylingK_ElementWord_Spacing = "word-spacing"
	StylingK_ElementWord_Wrap    = "word-wrap"
)

var StylingK_ElementWord = map[string]string{
	StylingK_ElementWord_Break:   "",
	StylingK_ElementWord_Spacing: "",
	StylingK_ElementWord_Wrap:    "",
}

// Load -
func (thisRef *ElementWord) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementWord_Break) == 0 {
		thisRef.Break = value
	} else if strings.Compare(key, StylingK_ElementWord_Spacing) == 0 {
		thisRef.Spacing = value
	} else if strings.Compare(key, StylingK_ElementWord_Wrap) == 0 {
		thisRef.Wrap = value
	}
}
