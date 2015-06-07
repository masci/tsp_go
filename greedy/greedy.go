package greedy

import (
	"github.com/masci/tsp_go/common"
)

// Go through edges, shortest first. Use edge to join segments if possible.
func GreedyTsp(cities []common.City) common.Tour {
	return *(common.NewTour(cities))
}

// Return all edges between distinct cities, sorted shortest first.
func ShortestEdgeFirst(cities []common.City) []common.Edge {
	return []common.Edge{}
}
