package styling

import "strings"

type ElementBackground struct {
	Attachment string `json:"attachment,omitempty"`
	BlendMode  string `json:"blendMode,omitempty"`
	Clip       string `json:"clip,omitempty"`
	Color      string `json:"color,omitempty"`
	Image      string `json:"image,omitempty"`
	Origin     string `json:"origin,omitempty"`
	Position   string `json:"position,omitempty"`
	Repeat     string `json:"repeat,omitempty"`
	Size       string `json:"size,omitempty"`
}

const (
	StylingK_ElementBackground_Attachment = "background-attachment"
	StylingK_ElementBackground_BlendMode  = "background-blend-mode"
	StylingK_ElementBackground_Clip       = "background-clip"
	StylingK_ElementBackground_Color      = "background-color"
	StylingK_ElementBackground_Image      = "background-image"
	StylingK_ElementBackground_Origin     = "background-origin"
	StylingK_ElementBackground_Position   = "background-position"
	StylingK_ElementBackground_Repeat     = "background-repeat"
	StylingK_ElementBackground_Size       = "background-size"
)

var StylingK_ElementBackground = map[string]string{
	StylingK_ElementBackground_Attachment: "",
	StylingK_ElementBackground_BlendMode:  "",
	StylingK_ElementBackground_Clip:       "",
	StylingK_ElementBackground_Color:      "",
	StylingK_ElementBackground_Image:      "",
	StylingK_ElementBackground_Origin:     "",
	StylingK_ElementBackground_Position:   "",
	StylingK_ElementBackground_Repeat:     "",
	StylingK_ElementBackground_Size:       "",
}

// Load -
func (thisRef *ElementBackground) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementBackground_Attachment) == 0 {
		thisRef.Attachment = value
	} else if strings.Compare(key, StylingK_ElementBackground_BlendMode) == 0 {
		thisRef.BlendMode = value
	} else if strings.Compare(key, StylingK_ElementBackground_Clip) == 0 {
		thisRef.Clip = value
	} else if strings.Compare(key, StylingK_ElementBackground_Color) == 0 {
		thisRef.Color = value
	} else if strings.Compare(key, StylingK_ElementBackground_Image) == 0 {
		thisRef.Image = value
	} else if strings.Compare(key, StylingK_ElementBackground_Origin) == 0 {
		thisRef.Origin = value
	} else if strings.Compare(key, StylingK_ElementBackground_Position) == 0 {
		thisRef.Position = value
	} else if strings.Compare(key, StylingK_ElementBackground_Position) == 0 {
		thisRef.Position = value
	} else if strings.Compare(key, StylingK_ElementBackground_Repeat) == 0 {
		thisRef.Repeat = value
	} else if strings.Compare(key, StylingK_ElementBackground_Size) == 0 {
		thisRef.Size = value
	}
}
