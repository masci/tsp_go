package improved_alltours

import (
	"github.com/masci/tsp_go/common"
)

// solve the tsp problem with the improved version of alltours_tsp
func AlltoursTsp(cities []common.City) common.Tour {
	tours := improved_alltours(cities)
	return common.ShortestTour(tours)
}

// Return all possible tours for the common.City instances contained in the array
// http://en.wikipedia.org/wiki/Heap%27s_algorithm
func improved_alltours(cities []common.City) []common.Tour {
	first := cities[0]
	cities = cities[1:]
	var helper func(cities []common.City, index int)
	res := []common.Tour{}

	helper = func(cities []common.City, n int) {
		tmp := make([]common.City, len(cities))
		copy(tmp, cities)
		tour := common.NewTour(tmp)

		if n == 1 {
			tour.Append(first)
			tour.Swap(0, len(tour.Cities())-1)
			res = append(res, *tour)
		} else {
			for i := 0; i < n; i++ {
				helper(tour.Cities(), n-1)
				if n%2 == 1 {
					tour.Swap(i, n-1)
				} else {
					tour.Swap(0, n-1)
				}
			}
		}
	}

	helper(cities, len(cities))
	return res
}
