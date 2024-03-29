package objectmodel

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/codemodify/Yeti/styling"
)

const (
	FiniteTextTag string = "__FiniteTextTag__"
)

// ObjectModel - FIXME
type ObjectModel struct {
	Tag        string            `json:"tag,omitempty"`
	Text       string            `json:"text,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Children   []*ObjectModel    `json:"children,omitempty"`

	Style *styling.ElementStyle `json:"style,omitempty"` // `json:"-"`
}

// ID - `Stringer` interface
func (thisRef ObjectModel) String() string {
	data, err := json.Marshal(thisRef)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

// ID - FIXME
func (thisRef ObjectModel) ID() string {
	if value, keyExists := thisRef.Attributes["id"]; keyExists {
		return value
	}

	return ""
}

// Classes - FIXME
func (thisRef ObjectModel) Classes() []string {
	resut := []string{}

	if value, keyExists := thisRef.Attributes["class"]; keyExists {
		classes := strings.Split(value, ".")
		for _, class := range classes {
			resut = append(resut, class)
		}
	} else {
		return nil
	}

	return resut
}

// Styles - FIXME
func (thisRef ObjectModel) Styles() map[string]string {
	resut := map[string]string{}

	if value, keyExists := thisRef.Attributes["style"]; keyExists {
		styles := strings.Split(value, ";")
		for _, style := range styles {
			if len(strings.TrimSpace(style)) > 0 {
				styleKeyValue := strings.Split(style, ":")
				resut[styleKeyValue[0]] = styleKeyValue[1]
			}
		}
	}

	return resut
}

// NodeReady - FIXME
func (thisRef *ObjectModel) NodeReady() {
	// 1. Create `ElementStyle` based on Attributes
	var styles = thisRef.Styles()
	if len(styles) > 0 {
		thisRef.Style = &styling.ElementStyle{}

		for key, value := range styles {
			if _, keyExists := styling.StylingK_ElementAlign[key]; keyExists {
				if thisRef.Style.Align == nil {
					thisRef.Style.Align = &styling.ElementAlign{}
				}
				thisRef.Style.Align.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementAnimation[key]; keyExists {
				if thisRef.Style.Animation == nil {
					thisRef.Style.Animation = &styling.ElementAnimation{}
				}
				thisRef.Style.Animation.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementBackground[key]; keyExists {
				if thisRef.Style.Background == nil {
					thisRef.Style.Background = &styling.ElementBackground{}
				}
				thisRef.Style.Background.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementBorder[key]; keyExists {
				if thisRef.Style.Border == nil {
					thisRef.Style.Border = &styling.ElementBorder{}
				}
				thisRef.Style.Border.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementFont[key]; keyExists {
				if thisRef.Style.Font == nil {
					thisRef.Style.Font = &styling.ElementFont{}
				}
				thisRef.Style.Font.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementInset[key]; keyExists {
				if thisRef.Style.Inset == nil {
					thisRef.Style.Inset = &styling.ElementInset{}
				}
				thisRef.Style.Inset.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementJustify[key]; keyExists {
				if thisRef.Style.Justify == nil {
					thisRef.Style.Justify = &styling.ElementJustify{}
				}
				thisRef.Style.Justify.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementLine[key]; keyExists {
				if thisRef.Style.Line == nil {
					thisRef.Style.Line = &styling.ElementLine{}
				}
				thisRef.Style.Line.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementMargin[key]; keyExists {
				if thisRef.Style.Margin == nil {
					thisRef.Style.Margin = &styling.ElementMargin{}
				}
				thisRef.Style.Margin.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementMask[key]; keyExists {
				if thisRef.Style.Mask == nil {
					thisRef.Style.Mask = &styling.ElementMask{}
				}
				thisRef.Style.Mask.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementMax[key]; keyExists {
				if thisRef.Style.Max == nil {
					thisRef.Style.Max = &styling.ElementMax{}
				}
				thisRef.Style.Max.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementMin[key]; keyExists {
				if thisRef.Style.Min == nil {
					thisRef.Style.Min = &styling.ElementMin{}
				}
				thisRef.Style.Min.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementOutline[key]; keyExists {
				if thisRef.Style.Outline == nil {
					thisRef.Style.Outline = &styling.ElementOutline{}
				}
				thisRef.Style.Outline.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementOverflow[key]; keyExists {
				if thisRef.Style.Overflow == nil {
					thisRef.Style.Overflow = &styling.ElementOverflow{}
				}
				thisRef.Style.Overflow.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementPadding[key]; keyExists {
				if thisRef.Style.Padding == nil {
					thisRef.Style.Padding = &styling.ElementPadding{}
				}
				thisRef.Style.Padding.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementScroll[key]; keyExists {
				if thisRef.Style.Scroll == nil {
					thisRef.Style.Scroll = &styling.ElementScroll{}
				}
				thisRef.Style.Scroll.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementSpahe[key]; keyExists {
				if thisRef.Style.Spahe == nil {
					thisRef.Style.Spahe = &styling.ElementSpahe{}
				}
				thisRef.Style.Spahe.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementText[key]; keyExists {
				if thisRef.Style.Text == nil {
					thisRef.Style.Text = &styling.ElementText{}
				}
				thisRef.Style.Text.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementTransform[key]; keyExists {
				if thisRef.Style.Transform == nil {
					thisRef.Style.Transform = &styling.ElementTransform{}
				}
				thisRef.Style.Transform.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementTransition[key]; keyExists {
				if thisRef.Style.Transition == nil {
					thisRef.Style.Transition = &styling.ElementTransition{}
				}
				thisRef.Style.Transition.Load(key, value)
			} else if _, keyExists := styling.StylingK_ElementWord[key]; keyExists {
				if thisRef.Style.Word == nil {
					thisRef.Style.Word = &styling.ElementWord{}
				}
				thisRef.Style.Word.Load(key, value)
			} else if _, keyExists := styling.StylingK_Element[key]; keyExists {
				if thisRef.Style.Element == nil {
					thisRef.Style.Element = &styling.Element{}
				}
				thisRef.Style.Element.Load(key, value)
			}
		}
	}

	// FINAL: Do the same for children
	for _, child := range thisRef.Children {
		child.NodeReady()
	}
}
