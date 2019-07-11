package objectmodelstyling

type Element struct {
	Width         string `json:"width,omitempty"`
	VerticalAlign string `json:"verticalAlign,omitempty"`
	Visibility    string `json:"visibility,omitempty"`
}

// const (
// 	StylingK_Element_Width         = "width"
// 	StylingK_Element_VerticalAlign = "vertical-align"
// 	StylingK_Element_Visibility    = "visibility"
// )

var StylingK_Element = map[string]string{
	"width":          "",
	"vertical-align": "",
	"visibility":     "",
}
