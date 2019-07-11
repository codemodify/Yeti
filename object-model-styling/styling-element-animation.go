package objectmodelstyling

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

// const (
// 	StylingK_ElementAnimation_Delay          = "animation-delay"
// 	StylingK_ElementAnimation_Direction      = "animation-direction"
// 	StylingK_ElementAnimation_Duration       = "animation-duration"
// 	StylingK_ElementAnimation_FillMode       = "animation-fill-mode"
// 	StylingK_ElementAnimation_IterationCount = "animation-iteration-count"
// 	StylingK_ElementAnimation_Name           = "animation-name"
// 	StylingK_ElementAnimation_PlayState      = "animation-play-state"
// 	StylingK_ElementAnimation_TimingFunction = "animation-timing-function"
// )

var StylingK_ElementAnimation = map[string]string{
	"animation-delay":           "",
	"animation-direction":       "",
	"animation-duration":        "",
	"animation-fill-mode":       "",
	"animation-iteration-count": "",
	"animation-name":            "",
	"animation-play-state":      "",
	"animation-timing-function": "",
}
