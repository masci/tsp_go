package nearest_neighbor

import (
	"github.com/masci/tsp_go/common"
	"math"
	"reflect"
)

// Find the nearest to city in cities
func nearestNeighbor(city common.City, neighbors []common.City) common.City {
	var nearest common.City
	min_dist := math.MaxFloat64

	for _, c := range neighbors {
		d, _ := city.Distance(&c)
		if d != 0.0 && d < min_dist {
			nearest = c
			min_dist = d
		}
	}

	return nearest
}

// remove city from cities in place
func deleteFrom(city common.City, cities *[]common.City) {
	for i, c := range *cities {
		if reflect.DeepEqual(c, city) {
			*cities = append((*cities)[:i], (*cities)[i+1:]...)
			break
		}
	}
}

// Solve the tsp problem for the given array of City.
// Start the tour at the first city; at each step extend the tour
// by moving from the previous city to its nearest neighbor
// that has not yet been visited.
func NnTsp(cities []common.City) common.Tour {
	start := cities[0]
	unvisited := cities[1:]
	t := common.NewTour([]common.City{start})
	for len(unvisited) > 0 {
		c := nearestNeighbor(t.Cities()[len(t.Cities())-1], unvisited)
		t.Append(c)
		deleteFrom(c, &unvisited)
	}

	return *t
}

func NnTspFrom(cities []common.City, start common.City) common.Tour {
	// make a copy of the original cities slice
	my_cities := make([]common.City, len(cities))
	copy(my_cities, cities)

	// create a tour to work with
	t := common.NewTour(my_cities)
	if !t.Contains(start) {
		return *t
	}

	// move start city in the first place
	for i, c := range my_cities {
		if reflect.DeepEqual(c, start) {
			t.Swap(0, i)
			break
		}
	}

	return NnTsp(t.Cities())
}

// Repeat the nn_tsp algorithm starting from each city; return the shortest tour.
func RepeatedNnTsp(cities []common.City) common.Tour {
	tours := []common.Tour{}

	for _, c := range cities {
		t := NnTspFrom(cities, c)
		tours = append(tours, t)
	}

	return common.ShortestTour(tours)
}
