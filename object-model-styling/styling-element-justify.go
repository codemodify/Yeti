package objectmodelstyling

type ElementJustify struct {
	Content string `json:"content,omitempty"`
	Items   string `json:"items,omitempty"`
	Self    string `json:"self,omitempty"`
}

// const (
// 	StylingK_ElementJustify_Content = "justify-content"
// 	StylingK_ElementJustify_Items   = "justify-items"
// 	StylingK_ElementJustify_Self    = "justify-self"
// )

var StylingK_ElementJustify = map[string]string{
	"justify-content": "",
	"justify-items":   "",
	"justify-self":    "",
}
