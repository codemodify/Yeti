package objectmodelstyling

type ElementSpahe struct {
	ImageThreshold string `json:"imageThreshold,omitempty"`
	Margin         string `json:"margin,omitempty"`
	Outside        string `json:"outside,omitempty"`
}

// const (
// 	StylingK_ElementSpahe_ImageThreshold = "shape-image-threshold"
// 	StylingK_ElementSpahe_Margin         = "shape-margin"
// 	StylingK_ElementSpahe_Outside        = "shape-outside"
// )

var StylingK_ElementSpahe = map[string]string{
	"shape-image-threshold": "",
	"shape-margin":          "",
	"shape-outside":         "",
}
