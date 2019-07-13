package styling

import "strings"

type GlobalScrollbar struct {
	Color string
	Width string
}

const (
	StylingK_GlobalScrollbar_Color = "scrollbar-color"
	StylingK_GlobalScrollbar_Width = "scrollbar-width"
)

var StylingK_GlobalScrollbar = map[string]string{
	StylingK_GlobalScrollbar_Color: "",
	StylingK_GlobalScrollbar_Width: "",
}

// Load -
func (thisRef *GlobalScrollbar) Load(key, value string) {
	if strings.Compare(key, StylingK_GlobalScrollbar_Color) == 0 {
		thisRef.Color = value
	} else if strings.Compare(key, StylingK_GlobalScrollbar_Width) == 0 {
		thisRef.Width = value
	}
}
