package objectmodelstyling

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

// const (
// 	StylingK_ElementBackground_Attachment = "background-attachment"
// 	StylingK_ElementBackground_BlendMode  = "background-blend-mode"
// 	StylingK_ElementBackground_Clip       = "background-clip"
// 	StylingK_ElementBackground_Color      = "background-color"
// 	StylingK_ElementBackground_Image      = "background-image"
// 	StylingK_ElementBackground_Origin     = "background-origin"
// 	StylingK_ElementBackground_Position   = "background-position"
// 	StylingK_ElementBackground_Repeat     = "background-repeat"
// 	StylingK_ElementBackground_Size       = "background-size"
// )

var StylingK_ElementBackground = map[string]string{
	"background-attachment": "",
	"background-blend-mode": "",
	"background-clip":       "",
	"background-color":      "",
	"background-image":      "",
	"background-origin":     "",
	"background-position":   "",
	"background-repeat":     "",
	"background-size":       "",
}
