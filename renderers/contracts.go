package renderers

import objectmodel "github.com/codemodify/Yeti/object-model"

// Renderer -
type Renderer interface {
	Render(node *objectmodel.ObjectModel)
}
