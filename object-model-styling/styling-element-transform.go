package objectmodelstyling

type ElementTransform struct {
	Box    string `json:"box,omitempty"`
	Origin string `json:"origin,omitempty"`
	Style  string `json:"style,omitempty"`
}

// const (
// 	StylingK_ElementTransform_Box    = "transform-box"
// 	StylingK_ElementTransform_Origin = "transform-origin"
// 	StylingK_ElementTransform_Style  = "transform-style"
// )

var StylingK_ElementTransform = map[string]string{
	"transform-box":    "",
	"transform-origin": "",
	"transform-style":  "",
}
