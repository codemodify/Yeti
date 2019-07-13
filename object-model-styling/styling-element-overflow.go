package objectmodelstyling

import "strings"

type ElementOverflow struct {
	Wrap string `json:"wrap,omitempty"`
	X    string `json:"x,omitempty"`
	Y    string `json:"y,omitempty"`
}

const (
	StylingK_ElementOverflow_Wrap = "overflow-wrap"
	StylingK_ElementOverflow_X    = "overflow-x"
	StylingK_ElementOverflow_Y    = "overflow-y"
)

var StylingK_ElementOverflow = map[string]string{
	StylingK_ElementOverflow_Wrap: "",
	StylingK_ElementOverflow_X:    "",
	StylingK_ElementOverflow_Y:    "",
}

// Load -
func (thisRef *ElementOverflow) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementOverflow_Wrap) == 0 {
		thisRef.Wrap = value
	} else if strings.Compare(key, StylingK_ElementOverflow_X) == 0 {
		thisRef.X = value
	} else if strings.Compare(key, StylingK_ElementOverflow_Y) == 0 {
		thisRef.Y = value
	}
}
