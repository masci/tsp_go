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
