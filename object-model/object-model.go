package objectmodel

import (
	"encoding/json"
	"log"
	"strings"
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
