package objectmodelstyling

import "strings"

type ElementMask struct {
	Clip      string `json:"clip,omitempty"`
	Composite string `json:"composite,omitempty"`
	Image     string `json:"image,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Origin    string `json:"origin,omitempty"`
	Position  string `json:"position,omitempty"`
	Repeat    string `json:"repeat,omitempty"`
	Size      string `json:"size,omitempty"`
	Type      string `json:"type,omitempty"`
}

const (
	StylingK_ElementMask_Clip      = "mask-clip"
	StylingK_ElementMask_Composite = "mask-composite"
	StylingK_ElementMask_Image     = "mask-image"
	StylingK_ElementMask_Mode      = "mask-mode"
	StylingK_ElementMask_Origin    = "mask-origin"
	StylingK_ElementMask_Position  = "mask-position"
	StylingK_ElementMask_Repeat    = "mask-repeat"
	StylingK_ElementMask_Size      = "mask-size"
	StylingK_ElementMask_Type      = "mask-type"
)

var StylingK_ElementMask = map[string]string{
	StylingK_ElementMask_Clip:      "",
	StylingK_ElementMask_Composite: "",
	StylingK_ElementMask_Image:     "",
	StylingK_ElementMask_Mode:      "",
	StylingK_ElementMask_Origin:    "",
	StylingK_ElementMask_Position:  "",
	StylingK_ElementMask_Repeat:    "",
	StylingK_ElementMask_Size:      "",
	StylingK_ElementMask_Type:      "",
}

// Load -
func (thisRef *ElementMask) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementMask_Clip) == 0 {
		thisRef.Clip = value
	} else if strings.Compare(key, StylingK_ElementMask_Composite) == 0 {
		thisRef.Composite = value
	} else if strings.Compare(key, StylingK_ElementMask_Image) == 0 {
		thisRef.Image = value
	} else if strings.Compare(key, StylingK_ElementMask_Mode) == 0 {
		thisRef.Mode = value
	} else if strings.Compare(key, StylingK_ElementMask_Origin) == 0 {
		thisRef.Origin = value
	} else if strings.Compare(key, StylingK_ElementMask_Position) == 0 {
		thisRef.Position = value
	} else if strings.Compare(key, StylingK_ElementMask_Repeat) == 0 {
		thisRef.Repeat = value
	} else if strings.Compare(key, StylingK_ElementMask_Size) == 0 {
		thisRef.Size = value
	} else if strings.Compare(key, StylingK_ElementMask_Type) == 0 {
		thisRef.Type = value
	}
}
