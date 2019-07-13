package objectmodelstyling

import "strings"

type ElementSpahe struct {
	ImageThreshold string `json:"imageThreshold,omitempty"`
	Margin         string `json:"margin,omitempty"`
	Outside        string `json:"outside,omitempty"`
}

const (
	StylingK_ElementSpahe_ImageThreshold = "shape-image-threshold"
	StylingK_ElementSpahe_Margin         = "shape-margin"
	StylingK_ElementSpahe_Outside        = "shape-outside"
)

var StylingK_ElementSpahe = map[string]string{
	StylingK_ElementSpahe_ImageThreshold: "",
	StylingK_ElementSpahe_Margin:         "",
	StylingK_ElementSpahe_Outside:        "",
}

// Load -
func (thisRef *ElementSpahe) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementSpahe_ImageThreshold) == 0 {
		thisRef.ImageThreshold = value
	} else if strings.Compare(key, StylingK_ElementSpahe_Margin) == 0 {
		thisRef.Margin = value
	} else if strings.Compare(key, StylingK_ElementSpahe_Outside) == 0 {
		thisRef.Outside = value
	}
}
