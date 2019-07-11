package objectmodelstyling

type ElementOutline struct {
	Color  string `json:"color,omitempty"`
	Offset string `json:"offset,omitempty"`
	Style  string `json:"style,omitempty"`
	Width  string `json:"width,omitempty"`
}

// const (
// 	StylingK_ElementOutline_Color  = "outline-color"
// 	StylingK_ElementOutline_Offset = "outline-offset"
// 	StylingK_ElementOutline_Style  = "outline-style"
// 	StylingK_ElementOutline_Width  = "outline-width"
// )

var StylingK_ElementOutline = map[string]string{
	"outline-color":  "",
	"outline-offset": "",
	"outline-style":  "",
	"outline-width":  "",
}
