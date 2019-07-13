package styling

import (
	"strings"

	numbersHelpers "github.com/brightappsllc/gohelpers/numbers"
)

type Element struct {
	Height        float64 `json:"width,omitempty"`
	Width         float64 `json:"width,omitempty"`
	VerticalAlign string  `json:"verticalAlign,omitempty"`
	Visibility    string  `json:"visibility,omitempty"`
}

const (
	StylingK_Element_Height        = "height"
	StylingK_Element_Width         = "width"
	StylingK_Element_VerticalAlign = "vertical-align"
	StylingK_Element_Visibility    = "visibility"
)

var StylingK_Element = map[string]string{
	StylingK_Element_Height:        "",
	StylingK_Element_Width:         "",
	StylingK_Element_VerticalAlign: "",
	StylingK_Element_Visibility:    "",
}

// Load -
func (thisRef *Element) Load(key, value string) {
	if strings.Compare(key, StylingK_Element_Height) == 0 {
		thisRef.Height = numbersHelpers.ParseFloat(value, 0)
	} else if strings.Compare(key, StylingK_Element_Width) == 0 {
		thisRef.Width = numbersHelpers.ParseFloat(value, 0)
	} else if strings.Compare(key, StylingK_Element_VerticalAlign) == 0 {
		thisRef.VerticalAlign = value
	} else if strings.Compare(key, StylingK_Element_Visibility) == 0 {
		thisRef.Visibility = value
	}
}
