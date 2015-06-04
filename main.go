package main

import (
	"fmt"
	"math/rand"
)

// solve the tsp problem for the given array of City
func alltours_tsp(cities []City) Tour {
	tours := alltours(cities)
	return shortest_tour(tours)
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

// Return all possible tours for the City instances contained in the array
// http://en.wikipedia.org/wiki/Heap%27s_algorithm
func alltours(cities []City) []Tour {
	var helper func(cities []City, index int)
	res := []Tour{}

	helper = func(cities []City, n int) {
		tmp := make([]City, len(cities))
		copy(tmp, cities)
		tour := NewTour(tmp)

		if n == 1 {
			res = append(res, *tour)
		} else {
			for i := 0; i < n; i++ {
				helper(tour.cities, n-1)
				if n%2 == 1 {
					tour.swap(i, n-1)
				} else {
					tour.swap(0, n-1)
				}
			}
		}
	}

	helper(cities, len(cities))
	return res
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
	cs := cities(9, 200, 100, 42)
	fmt.Println("Cities:", cs)
	fmt.Println("Best tour:")
	t := alltours_tsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.length())
}
