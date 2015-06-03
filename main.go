package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type City struct {
	x, y float64
}

func (c *City) distance(other *City) (float64, error) {
	if other.x < 0 || other.y < 0 {
		return 0.0, errors.New("Distance only allows positive coordinates")
	}
	sx := math.Pow(other.x-c.x, 2)
	sy := math.Pow(other.y-c.y, 2)
	return math.Sqrt(sx + sy), nil
}

type Tour struct {
	cities []City
}

// swap two cities from a tour
func (t *Tour) swap(i int, j int) {
	tmp := t.cities[i]
	t.cities[i] = t.cities[j]
	t.cities[j] = tmp
}

// append a new City to this Tour
func (t *Tour) append(c City) {
	t.cities = append(t.cities, c)
}

// compute Tour length
func (t *Tour) length() float64 {
	l := 0.0
	for i, c := range t.cities {
		if i > 0 {
			d, _ := c.distance(&(t.cities[i-1]))
			l += d
		} else {
			d, _ := c.distance(&(t.cities[len(t.cities)-1]))
			l += d
		}
	}
	return l
}

// solve the tsp problem for the given array of City
func alltours_tsp(cities []City) Tour {
	var tours []Tour
	alltours(&tours, cities, 0)
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
func alltours(tours *[]Tour, cities []City, index int) {
	t := Tour{cities}
	if index == len(cities)-1 {
		*tours = append(*tours, t)
	} else {
		for j := index; j < len(t.cities); j++ {
			t.swap(index, j)
			alltours(tours, t.cities, index+1)
			t.swap(index, j)
		}
	}
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
	fmt.Printf("%+v\n", alltours_tsp(cities(3, 200, 100, 42)))
}
