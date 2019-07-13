package styling

import "strings"

type ElementTransform struct {
	Box    string `json:"box,omitempty"`
	Origin string `json:"origin,omitempty"`
	Style  string `json:"style,omitempty"`
}

const (
	StylingK_ElementTransform_Box    = "transform-box"
	StylingK_ElementTransform_Origin = "transform-origin"
	StylingK_ElementTransform_Style  = "transform-style"
)

var StylingK_ElementTransform = map[string]string{
	StylingK_ElementTransform_Box:    "",
	StylingK_ElementTransform_Origin: "",
	StylingK_ElementTransform_Style:  "",
}

// Load -
func (thisRef *ElementTransform) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementTransform_Box) == 0 {
		thisRef.Box = value
	} else if strings.Compare(key, StylingK_ElementTransform_Origin) == 0 {
		thisRef.Origin = value
	} else if strings.Compare(key, StylingK_ElementTransform_Style) == 0 {
		thisRef.Style = value
	}
}
