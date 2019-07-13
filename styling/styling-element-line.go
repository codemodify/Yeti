package styling

import "strings"

type ElementLine struct {
	Break  string `json:"break,omitempty"`
	Height string `json:"height,omitempty"`
}

const (
	StylingK_ElementLine_Break  = "line-break"
	StylingK_ElementLine_Height = "line-height"
)

var StylingK_ElementLine = map[string]string{
	StylingK_ElementLine_Break:  "",
	StylingK_ElementLine_Height: "",
}

// Load -
func (thisRef *ElementLine) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementLine_Break) == 0 {
		thisRef.Break = value
	} else if strings.Compare(key, StylingK_ElementLine_Height) == 0 {
		thisRef.Height = value
	}
}
