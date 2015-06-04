package main

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

// solve the tsp problem for the given array of City
func alltours_tsp(cities []City) Tour {
	tours := alltours(cities)
	return shortest_tour(tours)
}
