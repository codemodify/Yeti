package objectmodelstyling

import "strings"

type Element struct {
	Width         string `json:"width,omitempty"`
	VerticalAlign string `json:"verticalAlign,omitempty"`
	Visibility    string `json:"visibility,omitempty"`
}

const (
	StylingK_Element_Width         = "width"
	StylingK_Element_VerticalAlign = "vertical-align"
	StylingK_Element_Visibility    = "visibility"
)

var StylingK_Element = map[string]string{
	StylingK_Element_Width:         "",
	StylingK_Element_VerticalAlign: "",
	StylingK_Element_Visibility:    "",
}

// Load -
func (thisRef *Element) Load(key, value string) {
	if strings.Compare(key, StylingK_Element_Width) == 0 {
		thisRef.Width = value
	} else if strings.Compare(key, StylingK_Element_VerticalAlign) == 0 {
		thisRef.VerticalAlign = value
	} else if strings.Compare(key, StylingK_Element_Visibility) == 0 {
		thisRef.Visibility = value
	}
}
