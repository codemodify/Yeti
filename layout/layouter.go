package layout

import (
	objectmodel "github.com/codemodify/Yeti/object-model"
	"github.com/codemodify/Yeti/styling"
)

// Layouter - FIXME
type Layouter struct {
	ParentWidth  float64
	ParentHeight float64
}

// NewLayouter -
func NewLayouter(width float64, height float64) *Layouter {
	return &Layouter{
		ParentWidth:  width,
		ParentHeight: height,
	}
}

// Layout -
func (thisRef Layouter) Layout(node *objectmodel.ObjectModel) {
	thisRef.layoutHelper(node, thisRef.ParentWidth, thisRef.ParentHeight)
}

func (thisRef Layouter) layoutHelper(node *objectmodel.ObjectModel, width float64, height float64) {
	if node.Style == nil {
		node.Style = &styling.ElementStyle{
			Element: &styling.Element{},
		}
	}

	node.Style.Element.Width = width
	node.Style.Element.Height = height

	for _, child := range node.Children {
		thisRef.layoutHelper(child, node.Style.Element.Width, node.Style.Element.Height)
	}
}
