package main

import (
	"fmt"
	"math"
	"math/rand"
)

type City struct {
	x, y float64
}

func (c *City) distance(other *City) float64 {
	sx := math.Pow(other.x-c.x, 2)
	sy := math.Pow(other.y-c.y, 2)
	return math.Sqrt(sx + sy)
}

func alltours_tsp( /*cities*/ ) {

}

func shortest_tour( /*tours*/ ) {

}

// Make a set of n cities, each with random coordinates within a (width x height) rectangle
func cities(n int, width int, height int, seed int64) []City {
	r := rand.New(rand.NewSource(seed))
	out := make([]City, n)
	i := 0
	for i < n {
		out[i] = City{float64(r.Intn(width)), float64(r.Intn(height))}
		i += 1
	}

	return out
}

func main() {

	fmt.Println("foo")

}
