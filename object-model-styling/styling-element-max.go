package objectmodelstyling

type ElementMax struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Zoom   string `json:"zoom,omitempty"`
}

// const (
// 	StylingK_ElementMax_Height = "max-height"
// 	StylingK_ElementMax_Width  = "max-width"
// 	StylingK_ElementMax_Zoom   = "max-zoom"
// )

var StylingK_ElementMax = map[string]string{
	"max-height": "",
	"max-width":  "",
	"max-zoom":   "",
}
