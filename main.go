package main

import (
	"fmt"
	"math/rand"
)

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

// find the shortest tour in the given array of tours
func shortest_tour(tours []Tour) Tour {
	min := Tour{}
	for _, t := range tours {
		if min.length() == 0 || t.length() < min.length() {
			min = t
		}
	}
	return min
}

func main() {
	cs := cities(8, 200, 100, 42)
	fmt.Println("Cities:", cs)

	fmt.Println("Best tour with alltours_tsp:")
	t := alltours_tsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.length())
}
