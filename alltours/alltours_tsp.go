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
		if n == 1 {
			tour := common.NewTour(cities)
			res = append(res, *tour)
		} else {
			for i := 0; i < n; i++ {
				helper(cities, n-1)
				if n%2 == 0 {
					cities[i], cities[n-1] = cities[n-1], cities[i]
				} else {
					cities[0], cities[n-1] = cities[n-1], cities[0]
				}
			}
		}
	}

	helper(cities, len(cities))
	return res
}

// Quickperm implementation
func alltoursQp(cities []common.City) []common.Tour {
	res := []common.Tour{}
	tmp := make([]common.City, len(cities))
	copy(tmp, cities)
	res = append(res, *(common.NewTour(tmp)))
	N := len(cities)
	p := make([]int, N+1)

	for i := 0; i <= N; i++ {
		p[i] = i
	}

	i := 1
	for i < N {
		p[i]--
		var j int
		if i%2 != 0 {
			j = p[i]
		} else {
			j = 0
		}

		cities[j], cities[i] = cities[i], cities[j]

		tmp = make([]common.City, len(cities))
		copy(tmp, cities)
		res = append(res, *(common.NewTour(tmp)))

		i = 1
		for p[i] == 0 {
			p[i] = i
			i++
		}
	}

	return res

}

// solve the tsp problem for the given array of City
func AlltoursTsp(cities []common.City) common.Tour {
	tours := alltours(cities)
	return common.ShortestTour(tours)
}

// solve using Quickperm algorithm
func AlltoursTspQp(cities []common.City) common.Tour {
	tours := alltoursQp(cities)
	return common.ShortestTour(tours)
}
