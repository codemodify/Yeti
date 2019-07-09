package dom

import (
	"encoding/json"
	"log"
	"strings"
)

const (
	FiniteTextTag string = "__FiniteTextTag__"
)

// NodeModel - FIXME
type NodeModel struct {
	Tag        string            `json:"tag,omitempty"`
	Text       string            `json:"text,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Children   []*NodeModel      `json:"children,omitempty"`
}

// ID - `Stringer` interface
func (thisRef NodeModel) String() string {
	data, err := json.Marshal(thisRef)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

// ID - FIXME
func (thisRef NodeModel) ID() string {
	if value, keyExists := thisRef.Attributes["id"]; keyExists {
		return value
	}

	return ""
}

// Classes - FIXME
func (thisRef NodeModel) Classes() []string {
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
