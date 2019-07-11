package objectmodelstyling

type ElementLine struct {
	Break  string `json:"break,omitempty"`
	Height string `json:"height,omitempty"`
}

// const (
// 	StylingK_ElementLine_Break  = "line-break"
// 	StylingK_ElementLine_Height = "line-height"
// )

var StylingK_ElementLine = map[string]string{
	"line-break":  "",
	"line-height": "",
}
