package alltours

import (
	"github.com/masci/tsp_go/common"
)

// Return all possible tours for the City instances contained in the array
// http://en.wikipedia.org/wiki/Heap%27s_algorithm
func alltours(cities []common.City) []common.Tour {
	var helper func(cities []common.City, index int)
	res := []common.Tour{}

	helper = func(cities []common.City, n int) {
		tmp := make([]common.City, len(cities))
		copy(tmp, cities)
		tour := common.NewTour(tmp)

		if n == 1 {
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

// solve the tsp problem for the given array of City
func AlltoursTsp(cities []common.City) common.Tour {
	tours := alltours(cities)
	return common.ShortestTour(tours)
}
