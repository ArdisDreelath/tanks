package view

import (
	"image"

	"github.com/ArdisDreelath/tanks/pkg/domain/terrain"
)

type ImageFromTerrain interface {
	CreateImage(t *terrain.Terrain) image.Image
}
