package objectmodelstyling

type ElementMin struct {
	BlockSize  string `json:"blockSize,omitempty"`
	Height     string `json:"height,omitempty"`
	InlineSize string `json:"inlineSize,omitempty"`
	Width      string `json:"width,omitempty"`
	Zoom       string `json:"zoom,omitempty"`
}

// const (
// 	StylingK_ElementMin_BlockSize  = "min-block-size"
// 	StylingK_ElementMin_Height     = "min-height"
// 	StylingK_ElementMin_InlineSize = "min-inline-size"
// 	StylingK_ElementMin_Width      = "max-width"
// 	StylingK_ElementMin_Zoom       = "max-zoom"
// )

var StylingK_ElementMin = map[string]string{
	"min-block-size":  "",
	"min-height":      "",
	"min-inline-size": "",
	"max-width":       "",
	"max-zoom":        "",
}
