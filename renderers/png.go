package renderers

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	objectmodel "github.com/codemodify/Yeti/object-model"
)

// PngRenderer - FIXME
type PngRenderer struct {
	fileName string
}

// NewPngRenderer -
func NewPngRenderer(fileName string) Renderer {
	return &PngRenderer{
		fileName: fileName,
	}
}

// Render - `Renderer` interface
func (thisRef PngRenderer) Render(node *objectmodel.ObjectModel) {
	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, int(node.Style.Element.Width), int(node.Style.Element.Height)))

	for y := 0; y < int(node.Style.Element.Height); y++ {
		for x := 0; x < int(node.Style.Element.Width); x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create(thisRef.fileName)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
