package terrain

type Generator interface {
	Generate() *Terrain
}

type Terrain struct {
	Heights []int
}
