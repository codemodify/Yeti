package renderers

import "github.com/codemodify/Yeti/dom"

// Renderer -
type Renderer interface {
	Render(node *dom.NodeModel)
}
