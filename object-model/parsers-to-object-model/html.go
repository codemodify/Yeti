package parsers2objectmodel

import (
	"bytes"
	"errors"
	"strings"

	objectmodel "github.com/codemodify/Yeti/object-model"
	"golang.org/x/net/html"
)

// HTMLParser - FIXME
type HTMLParser struct{}

// NewHTMLParser - FIXME
func NewHTMLParser() objectmodel.Parser {
	return HTMLParser{}
}

// Parse - `Parser` interface
func (thisRef HTMLParser) Parse(data []byte) (*objectmodel.ObjectModel, error) {

	// <html>
	// <body>
	// 	<h1>Title</h1>
	// 	<div id="main" class="test">
	// 		<p>Hello<em>world</em>!</p>
	// 	</div>
	// </body>
	// </html>

	//	|
	//	V

	// &objectmodel.ObjectModel{
	// 	Tag: "html",
	// 	Children: []*objectmodel.ObjectModel{
	// 		&objectmodel.ObjectModel{
	// 			Tag: "body",
	// 			Children: []*objectmodel.ObjectModel{
	// 				&objectmodel.ObjectModel{
	// 					Tag: "h1",
	// 					Children: []*objectmodel.ObjectModel{
	// 						&objectmodel.ObjectModel{
	// 							Tag:  objectmodel.FiniteTextTag,
	// 							Text: "Title",
	// 						},
	// 					},
	// 				},
	// 				&objectmodel.ObjectModel{
	// 					Tag: "div",
	// 					Attributes: map[string]string{
	// 						"id":    "main",
	// 						"class": "test",
	// 					},
	// 					Children: []*objectmodel.ObjectModel{
	// 						&objectmodel.ObjectModel{
	// 							Tag: "p",
	// 							Children: []*objectmodel.ObjectModel{
	// 								&objectmodel.ObjectModel{
	// 									Tag:  objectmodel.FiniteTextTag,
	// 									Text: "Hello",
	// 								},
	// 								&objectmodel.ObjectModel{
	// 									Tag: "em",
	// 									Children: []*objectmodel.ObjectModel{
	// 										&objectmodel.ObjectModel{
	// 											Tag:  objectmodel.FiniteTextTag,
	// 											Text: "world",
	// 										},
	// 									},
	// 								},
	// 								&objectmodel.ObjectModel{
	// 									Tag:  objectmodel.FiniteTextTag,
	// 									Text: "!",
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	node, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return thisRef.parseHelper(node.FirstChild)
}

func (thisRef HTMLParser) parseHelper(n *html.Node) (*objectmodel.ObjectModel, error) {

	// FROM `golang.org/x/net/html/node.go`
	//	ErrorNode NodeType = iota
	//	TextNode
	//	DocumentNode
	//	ElementNode
	//	CommentNode
	//	DoctypeNode
	//	scopeMarkerNode

	if n.Type == html.ErrorNode {
		return nil, errors.New("PARSE_ERROR")
	}

	var node *objectmodel.ObjectModel

	if n.Type == html.TextNode {
		trimmedData := strings.TrimSpace(n.Data)
		if len(trimmedData) > 0 {
			node = &objectmodel.ObjectModel{
				Tag:  objectmodel.FiniteTextTag,
				Text: strings.TrimSpace(n.Data),
			}
		}

		return node, nil
	}

	node = &objectmodel.ObjectModel{}

	if n.Type == html.ElementNode {
		node.Tag = n.Data
	}

	// Attributes
	if len(n.Attr) > 0 {
		node.Attributes = map[string]string{}

		for _, attr := range n.Attr {
			node.Attributes[attr.Key] = attr.Val
		}
	}

	// Children
	if n.FirstChild != nil {
		node.Children = []*objectmodel.ObjectModel{}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			childNode, err := thisRef.parseHelper(c)
			if err != nil {
				return nil, errors.New("PARSE_ERROR")
			}

			if childNode != nil {
				node.Children = append(node.Children, childNode)
			}
		}
	}

	return node, nil
}
