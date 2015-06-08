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

func benchGreedyTsp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		GreedyTsp(common.Cities(i, 900, 600, 42))
	}
}

func BenchmarkGreedyTsp5(b *testing.B) {
	benchGreedyTsp(5, b)
}

func BenchmarkGreedyTsp8(b *testing.B) {
	benchGreedyTsp(8, b)
}

func BenchmarkGreedyTsp10(b *testing.B) {
	benchGreedyTsp(10, b)
}

func BenchmarkGreedyTsp100(b *testing.B) {
	benchGreedyTsp(100, b)
}

func BenchmarkGreedyTsp1000(b *testing.B) {
	benchGreedyTsp(1000, b)
}
