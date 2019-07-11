package objectmodel

import (
	"encoding/json"
	"log"
	"strings"

	objectmodelstyling "github.com/codemodify/Yeti/object-model-styling"
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

	Style *objectmodelstyling.ElementStyle `json:"style,omitempty"` // `json:"-"`
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

func (thisRef *ObjectModel) NodeReady() {
	// 1. Create `ElementStyle` based on Attributes
	var styles = thisRef.Styles()
	if len(styles) > 0 {
		thisRef.Style = &objectmodelstyling.ElementStyle{}

		for key, value := range styles {
			if _, keyExists := objectmodelstyling.StylingK_ElementAlign[key]; keyExists {
				if thisRef.Style.Align == nil {
					thisRef.Style.Align = &objectmodelstyling.ElementAlign{}
				}

				thisRef.Style.Align.Load(key, value)
			}
		}
	}

	// FINAL: Do the same for children
	for _, child := range thisRef.Children {
		child.NodeReady()
	}
}
