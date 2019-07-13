package styling

import "strings"

type ElementMax struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Zoom   string `json:"zoom,omitempty"`
}

const (
	StylingK_ElementMax_Height = "max-height"
	StylingK_ElementMax_Width  = "max-width"
	StylingK_ElementMax_Zoom   = "max-zoom"
)

var StylingK_ElementMax = map[string]string{
	StylingK_ElementMax_Height: "",
	StylingK_ElementMax_Width:  "",
	StylingK_ElementMax_Zoom:   "",
}

// Load -
func (thisRef *ElementMax) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementMax_Height) == 0 {
		thisRef.Height = value
	} else if strings.Compare(key, StylingK_ElementMax_Width) == 0 {
		thisRef.Width = value
	} else if strings.Compare(key, StylingK_ElementMax_Zoom) == 0 {
		thisRef.Zoom = value
	}
}
