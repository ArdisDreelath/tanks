package terrain

import (
	"math/rand"
	"time"
)

func Generate() *Terrain {
	rand.Seed(time.Now().UnixNano())
	randstep := func(currentStep int) int {
		if currentStep > 4 {
			return -1
		}
		if currentStep < -4 {
			return 1
		}
		r := rand.Intn(10)
		if r == 0 {
			return -3
		} else if r == 1 {
			return -2
		} else if r == 2 || r == 3 {
			return -1
		} else if r == 6 || r == 7 {
			return 1
		} else if r == 8 {
			return 2
		} else if r == 9 {
			return 3
		}

		return 0
	}
	min := 0.0
	max := 0.0
	currentStep := 0
	lastHeight := 0.0
	heights := make([]float64, 2000)
	for i := 0; i < 2000; i++ {
		currentStep += randstep(currentStep)
		lastHeight += float64(currentStep)
		heights[i] = lastHeight
		if lastHeight > max {
			max = lastHeight
		}
		if lastHeight < min {
			min = lastHeight
		}
	}
	t := Terrain{Heights: make([]int, 2000)}
	for i := 0; i < 2000; i++ {
		t.Heights[i] = 200 + int(((heights[i]-min)/(max-min))*300)
	}

	return &t
}
