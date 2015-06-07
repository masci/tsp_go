package greedy

import (
	"github.com/masci/tsp_go/common"
	"sort"
)

// Go through edges, shortest first. Use edge to join segments if possible.
func GreedyTsp(cities []common.City) common.Tour {
	return *(common.NewTour(cities))
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
