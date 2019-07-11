package objectmodelstyling

type ElementWord struct {
	Break   string `json:"break,omitempty"`
	Spacing string `json:"spacing,omitempty"`
	Wrap    string `json:"wrap,omitempty"`
}

// const (
// 	StylingK_ElementWord_Break   = "word-break"
// 	StylingK_ElementWord_Spacing = "word-spacing"
// 	StylingK_ElementWord_Wrap    = "word-wrap"
// )

var StylingK_ElementWord = map[string]string{
	"word-break":   "",
	"word-spacing": "",
	"word-wrap":    "",
}
