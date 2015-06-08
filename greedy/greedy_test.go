package greedy

import (
	"fmt"
	"github.com/masci/tsp_go/common"
	"testing"
)

func TestShortestEdgeFirst(t *testing.T) {
	cities := common.Cities(3, 200, 100, 42)

	res := ShortestEdgeFirst(cities)

	e1 := common.Edge{cities[1], cities[2]}
	e2 := common.Edge{cities[0], cities[1]}
	e3 := common.Edge{cities[0], cities[2]}

	if res[0] != e1 || res[1] != e2 || res[2] != e3 {
		t.Error("Wrong sorting for", res)
	}
}

func TestGreedyTsp(t *testing.T) {
	cities := common.Cities(10, 200, 100, 42)
	res := GreedyTsp(cities)
	fmt.Println(res)
}
