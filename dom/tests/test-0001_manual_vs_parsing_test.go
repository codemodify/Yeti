package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/codemodify/Yeti/dom"
	"github.com/codemodify/Yeti/dom/parsers2dom"
)

func TestManualVsParsing(t *testing.T) {
	// workingDir, err := os.Getwd()
	// if err != nil {
	// 	t.Error(err)
	// 	t.FailNow()
	// }

	// fmt.Println(fmt.Sprintf("WORKING-DIR: %s", workingDir))

	// data, err := ioutil.ReadFile(fmt.Sprintf("%s/sample-files/file-0001.html", workingDir))
	// if err != nil {
	// 	t.Error(err)
	// 	t.FailNow()
	// }

	htmlAsString := `<html>
	<body>
		<h1>Title</h1>
		<div id="main" class="test">
			<p>Hello<em>world</em>!</p>
		</div>
	</body>
	</html>`

	//	|
	//	V

	nodeToCompareAgainst := &dom.NodeModel{
		Tag: "html",
		Children: []*dom.NodeModel{
			&dom.NodeModel{
				Tag: "head",
			},
			&dom.NodeModel{
				Tag: "body",
				Children: []*dom.NodeModel{
					&dom.NodeModel{
						Tag: "h1",
						Children: []*dom.NodeModel{
							&dom.NodeModel{
								Tag:  dom.FiniteTextTag,
								Text: "Title",
							},
						},
					},
					&dom.NodeModel{
						Tag: "div",
						Attributes: map[string]string{
							"id":    "main",
							"class": "test",
						},
						Children: []*dom.NodeModel{
							&dom.NodeModel{
								Tag: "p",
								Children: []*dom.NodeModel{
									&dom.NodeModel{
										Tag:  dom.FiniteTextTag,
										Text: "Hello",
									},
									&dom.NodeModel{
										Tag: "em",
										Children: []*dom.NodeModel{
											&dom.NodeModel{
												Tag:  dom.FiniteTextTag,
												Text: "world",
											},
										},
									},
									&dom.NodeModel{
										Tag:  dom.FiniteTextTag,
										Text: "!",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println()
	fmt.Println("TARGET")
	fmt.Println()
	fmt.Println(nodeToCompareAgainst.String())

	parser := parsers2dom.NewHTMLParser()
	parsedNode, err := parser.Parse([]byte(htmlAsString))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println()
	fmt.Println("PARSED")
	fmt.Println()
	fmt.Println(parsedNode.String())
	fmt.Println()

	if strings.Compare(nodeToCompareAgainst.String(), parsedNode.String()) != 0 {
		t.Error(err)
		t.FailNow()
	}
}
