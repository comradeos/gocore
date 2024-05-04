package chapters

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path *Path) Distance() float64 {
	var sum float64
	for i := range *path {
		if i > 0 {
			sum += Distance((*path)[i-1], (*path)[i])
		}
	}
	return sum
}

func Methods() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}
