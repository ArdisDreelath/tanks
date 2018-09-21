package view

import (
	"image"
	"image/color"

	"github.com/ArdisDreelath/tanks/pkg/domain/terrain"
)

type TerrainToImage struct {
	skyColor     color.Color
	terrainColor color.Color
}

func NewTerrainToImage(skyColor, terrainColor color.Color) *TerrainToImage {
	return &TerrainToImage{skyColor, terrainColor}
}

func (tti *TerrainToImage) CreateImage(t *terrain.Terrain) image.Image {
	img := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 2000, Y: 1200}})
	for y := 0; y < 1200; y++ {
		for x := 0; x < 2000; x++ {
			if t.Heights[x] < 1200-y {
				img.Set(x, y, tti.skyColor)
			} else {
				img.Set(x, y, tti.terrainColor)
			}
		}
	}

	return img
}
