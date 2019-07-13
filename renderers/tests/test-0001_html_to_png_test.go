package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/codemodify/Yeti/layout"
	parsers2objectmodel "github.com/codemodify/Yeti/object-model/parsers-to-object-model"
	"github.com/codemodify/Yeti/renderers"
)

func TestHTML_to_PNG(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(fmt.Sprintf("WORKING-DIR: %s", workingDir))

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/../../object-model/tests/sample-files/file-0001.html", workingDir))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	parser := parsers2objectmodel.NewHTMLParser()
	parsedNode, err := parser.Parse(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if parsedNode == nil {
		t.Error(err)
		t.FailNow()
	}

	// layout node and children in a box
	layouter := layout.NewLayouter(640, 480)
	layouter.Layout(parsedNode)

	// render node using existing sizes
	renderer := renderers.NewPngRenderer("sample.png")
	renderer.Render(parsedNode)
}
