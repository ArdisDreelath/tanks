package main

import (
	"image/color"
	"image/png"
	"os"

	"github.com/ArdisDreelath/tanks/pkg/adapters/view"
	"github.com/ArdisDreelath/tanks/pkg/domain/terrain"
)

func main() {
	t := terrain.Generate()
	gen := view.NewTerrainToImage(color.RGBA{0, 0, 255, 255}, color.RGBA{0, 255, 0, 255})
	img := gen.CreateImage(t)
	f, err := os.Create("terrain.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
