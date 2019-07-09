package parsers2dom

import (
	"bytes"
	"errors"
	"strings"

	"github.com/codemodify/Yeti/dom"
	"golang.org/x/net/html"
)

// HTMLParser - FIXME
type HTMLParser struct{}

// NewHTMLParser - FIXME
func NewHTMLParser() dom.Parser {
	return HTMLParser{}
}

// Parse - `Parser` interface
func (thisRef HTMLParser) Parse(data []byte) (*dom.NodeModel, error) {

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

	// &dom.NodeModel{
	// 	Tag: "html",
	// 	Children: []*dom.NodeModel{
	// 		&dom.NodeModel{
	// 			Tag: "body",
	// 			Children: []*dom.NodeModel{
	// 				&dom.NodeModel{
	// 					Tag: "h1",
	// 					Children: []*dom.NodeModel{
	// 						&dom.NodeModel{
	// 							Tag:  dom.FiniteTextTag,
	// 							Text: "Title",
	// 						},
	// 					},
	// 				},
	// 				&dom.NodeModel{
	// 					Tag: "div",
	// 					Attributes: map[string]string{
	// 						"id":    "main",
	// 						"class": "test",
	// 					},
	// 					Children: []*dom.NodeModel{
	// 						&dom.NodeModel{
	// 							Tag: "p",
	// 							Children: []*dom.NodeModel{
	// 								&dom.NodeModel{
	// 									Tag:  dom.FiniteTextTag,
	// 									Text: "Hello",
	// 								},
	// 								&dom.NodeModel{
	// 									Tag: "em",
	// 									Children: []*dom.NodeModel{
	// 										&dom.NodeModel{
	// 											Tag:  dom.FiniteTextTag,
	// 											Text: "world",
	// 										},
	// 									},
	// 								},
	// 								&dom.NodeModel{
	// 									Tag:  dom.FiniteTextTag,
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

func (thisRef HTMLParser) parseHelper(n *html.Node) (*dom.NodeModel, error) {

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

	var node *dom.NodeModel

	if n.Type == html.TextNode {
		trimmedData := strings.TrimSpace(n.Data)
		if len(trimmedData) > 0 {
			node = &dom.NodeModel{
				Tag:  dom.FiniteTextTag,
				Text: strings.TrimSpace(n.Data),
			}
		}

		return node, nil
	}

	node = &dom.NodeModel{}

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
		node.Children = []*dom.NodeModel{}

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
