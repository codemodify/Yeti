package objectmodelstyling

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

// const (
// 	StylingK_ElementMask_Clip      = "mask-clip"
// 	StylingK_ElementMask_Composite = "mask-composite"
// 	StylingK_ElementMask_Image     = "mask-image"
// 	StylingK_ElementMask_Mode      = "mask-mode"
// 	StylingK_ElementMask_Origin    = "mask-origin"
// 	StylingK_ElementMask_Position  = "mask-position"
// 	StylingK_ElementMask_Repeat    = "mask-repeat"
// 	StylingK_ElementMask_Size      = "mask-size"
// 	StylingK_ElementMask_Type      = "mask-type"
// )

var StylingK_ElementMask = map[string]string{
	"mask-clip":      "",
	"mask-composite": "",
	"mask-image":     "",
	"mask-mode":      "",
	"mask-origin":    "",
	"mask-position":  "",
	"mask-repeat":    "",
	"mask-size":      "",
	"mask-type":      "",
}
