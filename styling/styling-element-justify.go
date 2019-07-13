package styling

import "strings"

type ElementJustify struct {
	Content string `json:"content,omitempty"`
	Items   string `json:"items,omitempty"`
	Self    string `json:"self,omitempty"`
}

const (
	StylingK_ElementJustify_Content = "justify-content"
	StylingK_ElementJustify_Items   = "justify-items"
	StylingK_ElementJustify_Self    = "justify-self"
)

var StylingK_ElementJustify = map[string]string{
	StylingK_ElementJustify_Content: "",
	StylingK_ElementJustify_Items:   "",
	StylingK_ElementJustify_Self:    "",
}

// Load -
func (thisRef *ElementJustify) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementJustify_Content) == 0 {
		thisRef.Content = value
	} else if strings.Compare(key, StylingK_ElementJustify_Items) == 0 {
		thisRef.Items = value
	} else if strings.Compare(key, StylingK_ElementJustify_Self) == 0 {
		thisRef.Self = value
	}
}
