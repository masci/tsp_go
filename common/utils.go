package common

import "math/rand"

// shortest_tour finds the shortest tour in the given array of tours
func ShortestTour(tours []Tour) Tour {
	min := Tour{}
	for _, t := range tours {
		if min.Length() == 0 || t.Length() < min.Length() {
			min = t
		}
	}
	return min
}

// cities returns a set of n cities, each with random coordinates
// within a (width x height) rectangle
func Cities(n int, width int, height int, seed int64) []City {
	r := rand.New(rand.NewSource(seed))
	out := make([]City, n)
	i := 0
	for i < n {
		out[i] = City{float64(r.Intn(width)), float64(r.Intn(height))}
		i++
	}

	return out
}
