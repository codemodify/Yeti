package objectmodelstyling

type ElementTransition struct {
	Delay          string
	Duration       string
	Property       string
	TimingFunction string
}

// const (
// 	StylingK_ElementTransition_Delay          = "transition-delay"
// 	StylingK_ElementTransition_Duration       = "transition-duration"
// 	StylingK_ElementTransition_Property       = "transition-property"
// 	StylingK_ElementTransition_TimingFunction = "transition-timing-function"
// )

var StylingK_ElementTransition = map[string]string{
	"transition-delay":           "",
	"transition-duration":        "",
	"transition-property":        "",
	"transition-timing-function": "",
}
