package objectmodelstyling

import (
	"strings"
)

type ElementAlign struct {
	Content string `json:"content,omitempty"`
	Items   string `json:"items,omitempty"`
	Self    string `json:"self,omitempty"`
}

const (
	StylingK_ElementAlign_Content = "align-content"
	StylingK_ElementAlign_Items   = "align-items"
	StylingK_ElementAlign_Self    = "align-self"
)

var StylingK_ElementAlign = map[string]string{
	StylingK_ElementAlign_Content: "",
	StylingK_ElementAlign_Items:   "",
	StylingK_ElementAlign_Self:    "",
}

// Load -
func (thisRef *ElementAlign) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementAlign_Content) == 0 {
		thisRef.Content = value
	} else if strings.Compare(key, StylingK_ElementAlign_Items) == 0 {
		thisRef.Items = value
	} else if strings.Compare(key, StylingK_ElementAlign_Self) == 0 {
		thisRef.Self = value
	}
}
