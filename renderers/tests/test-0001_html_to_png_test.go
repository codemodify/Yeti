package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/codemodify/Yeti/dom/parsers2dom"
	"github.com/codemodify/Yeti/renderers"
)

func TestHTML_to_PNG(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(fmt.Sprintf("WORKING-DIR: %s", workingDir))

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/../../dom/tests/sample-files/file-0001.html", workingDir))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	parser := parsers2dom.NewHTMLParser()
	parsedNode, err := parser.Parse(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if parsedNode == nil {
		t.Error(err)
		t.FailNow()
	}

	renderer := renderers.NewPngRenderer("sample.png")
	renderer.Render(parsedNode)
}
