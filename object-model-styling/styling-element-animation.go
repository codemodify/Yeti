package objectmodelstyling

import "strings"

type ElementAnimation struct {
	Delay          string `json:"delay,omitempty"`
	Direction      string `json:"direction,omitempty"`
	Duration       string `json:"duration,omitempty"`
	FillMode       string `json:"fillMode,omitempty"`
	IterationCount string `json:"iterationCount,omitempty"`
	Name           string `json:"name,omitempty"`
	PlayState      string `json:"playState,omitempty"`
	TimingFunction string `json:"timingFunction,omitempty"`
}

const (
	StylingK_ElementAnimation_Delay          = "animation-delay"
	StylingK_ElementAnimation_Direction      = "animation-direction"
	StylingK_ElementAnimation_Duration       = "animation-duration"
	StylingK_ElementAnimation_FillMode       = "animation-fill-mode"
	StylingK_ElementAnimation_IterationCount = "animation-iteration-count"
	StylingK_ElementAnimation_Name           = "animation-name"
	StylingK_ElementAnimation_PlayState      = "animation-play-state"
	StylingK_ElementAnimation_TimingFunction = "animation-timing-function"
)

var StylingK_ElementAnimation = map[string]string{
	StylingK_ElementAnimation_Delay:          "",
	StylingK_ElementAnimation_Direction:      "",
	StylingK_ElementAnimation_Duration:       "",
	StylingK_ElementAnimation_FillMode:       "",
	StylingK_ElementAnimation_IterationCount: "",
	StylingK_ElementAnimation_Name:           "",
	StylingK_ElementAnimation_PlayState:      "",
	StylingK_ElementAnimation_TimingFunction: "",
}

// Load -
func (thisRef *ElementAnimation) Load(key, value string) {
	if strings.Compare(key, StylingK_ElementAnimation_Delay) == 0 {
		thisRef.Delay = value
	} else if strings.Compare(key, StylingK_ElementAnimation_Direction) == 0 {
		thisRef.Direction = value
	} else if strings.Compare(key, StylingK_ElementAnimation_Duration) == 0 {
		thisRef.Duration = value
	} else if strings.Compare(key, StylingK_ElementAnimation_FillMode) == 0 {
		thisRef.FillMode = value
	} else if strings.Compare(key, StylingK_ElementAnimation_IterationCount) == 0 {
		thisRef.IterationCount = value
	} else if strings.Compare(key, StylingK_ElementAnimation_Name) == 0 {
		thisRef.Name = value
	} else if strings.Compare(key, StylingK_ElementAnimation_PlayState) == 0 {
		thisRef.PlayState = value
	} else if strings.Compare(key, StylingK_ElementAnimation_TimingFunction) == 0 {
		thisRef.TimingFunction = value
	}
}
