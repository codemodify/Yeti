package objectmodelstyling

import "strings"

type ElementTransition struct {
	Delay          string `json:"delay,omitempty"`
	Duration       string `json:"duration,omitempty"`
	Property       string `json:"property,omitempty"`
	TimingFunction string `json:"timingFunction,omitempty"`
}

const (
	StylingK_ElementTransition_Delay          = "transition-delay"
	StylingK_ElementTransition_Duration       = "transition-duration"
	StylingK_ElementTransition_Property       = "transition-property"
	StylingK_ElementTransition_TimingFunction = "transition-timing-function"
)

var StylingK_ElementTransition = map[string]string{
	StylingK_ElementTransition_Delay:          "",
	StylingK_ElementTransition_Duration:       "",
	StylingK_ElementTransition_Property:       "",
	StylingK_ElementTransition_TimingFunction: "",
}

// Load -
func (thisRef *ElementTransition) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementTransition_Delay) == 0 {
		thisRef.Delay = value
	} else if strings.Compare(key, StylingK_ElementTransition_Duration) == 0 {
		thisRef.Duration = value
	} else if strings.Compare(key, StylingK_ElementTransition_Property) == 0 {
		thisRef.Property = value
	} else if strings.Compare(key, StylingK_ElementTransition_TimingFunction) == 0 {
		thisRef.TimingFunction = value
	}
}
