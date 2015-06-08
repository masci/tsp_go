package greedy

import (
	"github.com/masci/tsp_go/common"
	"reflect"
	"sort"
)

// Go through edges, shortest first. Use edge to join segments if possible.
func GreedyTsp(cities []common.City) common.Tour {
	edges := ShortestEdgeFirst(cities)
	endpoints := make(common.Endpoints)

	for _, c := range cities {
		endpoints[c] = []common.City{c}
	}

	for _, e := range edges {
		A_val, A_ok := endpoints[e.A]
		B_val, B_ok := endpoints[e.B]
		if A_ok && B_ok && !reflect.DeepEqual(A_val, B_val) {
			new_segment := JoinEndpoints(endpoints, e.A, e.B)
			if len(new_segment) == len(cities) {
				return *(common.NewTour(new_segment))
			}
		}
	}

	return common.Tour{}
}

// Return all edges between distinct cities, sorted shortest first.
func ShortestEdgeFirst(cities []common.City) []common.Edge {
	edges := common.Edges{}

	for i, cstart := range cities {
		for j, cend := range cities {
			if j <= i {
				continue
			}
			edges = append(edges, common.Edge{cstart, cend})
		}
	}

	sort.Sort(edges)

	return edges
}

func reverse(a []common.City) []common.City {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

// Join B's segment onto the end of A's and return the segment. Maintain endpoints map.
func JoinEndpoints(endpoints common.Endpoints, A common.City, B common.City) []common.City {
	A_segment := endpoints[A]
	if A_segment[len(A_segment)-1] != A {
		A_segment = reverse(A_segment)
	}

	B_segment := endpoints[B]
	if B_segment[0] != B {
		B_segment = reverse(B_segment)
	}

	A_segment = append(A_segment, B_segment...)
	delete(endpoints, A)
	delete(endpoints, B)

	endpoints[A_segment[0]] = A_segment
	endpoints[A_segment[len(A_segment)-1]] = A_segment

	return A_segment
}
