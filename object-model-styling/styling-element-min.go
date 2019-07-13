package objectmodelstyling

import "strings"

type ElementMin struct {
	BlockSize  string `json:"blockSize,omitempty"`
	Height     string `json:"height,omitempty"`
	InlineSize string `json:"inlineSize,omitempty"`
	Width      string `json:"width,omitempty"`
	Zoom       string `json:"zoom,omitempty"`
}

const (
	StylingK_ElementMin_BlockSize  = "min-block-size"
	StylingK_ElementMin_Height     = "min-height"
	StylingK_ElementMin_InlineSize = "min-inline-size"
	StylingK_ElementMin_Width      = "max-width"
	StylingK_ElementMin_Zoom       = "max-zoom"
)

var StylingK_ElementMin = map[string]string{
	StylingK_ElementMin_BlockSize:  "",
	StylingK_ElementMin_Height:     "",
	StylingK_ElementMin_InlineSize: "",
	StylingK_ElementMin_Width:      "",
	StylingK_ElementMin_Zoom:       "",
}

// Load -
func (thisRef *ElementMin) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementMin_BlockSize) == 0 {
		thisRef.BlockSize = value
	} else if strings.Compare(key, StylingK_ElementMin_Height) == 0 {
		thisRef.Height = value
	} else if strings.Compare(key, StylingK_ElementMin_InlineSize) == 0 {
		thisRef.InlineSize = value
	} else if strings.Compare(key, StylingK_ElementMin_Width) == 0 {
		thisRef.Width = value
	} else if strings.Compare(key, StylingK_ElementMin_Zoom) == 0 {
		thisRef.Zoom = value
	}
}
