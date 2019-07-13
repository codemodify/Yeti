package styling

import "strings"

type ElementOutline struct {
	Color  string `json:"color,omitempty"`
	Offset string `json:"offset,omitempty"`
	Style  string `json:"style,omitempty"`
	Width  string `json:"width,omitempty"`
}

const (
	StylingK_ElementOutline_Color  = "outline-color"
	StylingK_ElementOutline_Offset = "outline-offset"
	StylingK_ElementOutline_Style  = "outline-style"
	StylingK_ElementOutline_Width  = "outline-width"
)

var StylingK_ElementOutline = map[string]string{
	StylingK_ElementOutline_Color:  "",
	StylingK_ElementOutline_Offset: "",
	StylingK_ElementOutline_Style:  "",
	StylingK_ElementOutline_Width:  "",
}

// Load -
func (thisRef *ElementOutline) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementOutline_Color) == 0 {
		thisRef.Color = value
	} else if strings.Compare(key, StylingK_ElementOutline_Offset) == 0 {
		thisRef.Offset = value
	} else if strings.Compare(key, StylingK_ElementOutline_Style) == 0 {
		thisRef.Style = value
	} else if strings.Compare(key, StylingK_ElementOutline_Width) == 0 {
		thisRef.Width = value
	}
}
