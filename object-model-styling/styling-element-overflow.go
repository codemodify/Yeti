package objectmodelstyling

type ElementOverflow struct {
	Wrap string `json:"wrap,omitempty"`
	X    string `json:"x,omitempty"`
	Y    string `json:"y,omitempty"`
}

// const (
// 	StylingK_ElementOverflow_Wrap = "overflow-wrap"
// 	StylingK_ElementOverflow_X    = "overflow-x"
// 	StylingK_ElementOverflow_Y    = "overflow-y"
// )

var StylingK_ElementOverflow = map[string]string{
	"overflow-wrap": "",
	"overflow-x":    "",
	"overflow-y":    "",
}
