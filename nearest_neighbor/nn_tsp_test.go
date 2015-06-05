package nearest_neighbor

import (
	"github.com/masci/tsp_go/common"
	"reflect"
	"testing"
)

func TestDeleteFrom(t *testing.T) {
	cities := common.Cities(3, 200, 100, 42)
	c := cities[1]
	deleteFrom(c, &cities)
	tour := common.NewTour(cities)
	if tour.Contains(c) {
		t.Error("City", c, "should have been removed")
	}
}

func TestNearestNeighbor(t *testing.T) {
	c1 := common.NewCity(2, 4)
	c2 := common.NewCity(70, 80)
	c3 := common.NewCity(3, 6)

	res := nearestNeighbor(*c1, []common.City{*c2, *c3})
	if !reflect.DeepEqual(res, *c3) {
		t.Error("Expected", *c3, "found", res)
	}
}

func benchNnTsp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		NnTsp(common.Cities(i, 900, 600, 42))
	}
}

func BenchmarkNnTsp5(b *testing.B) {
	benchNnTsp(5, b)
}

func BenchmarkNnTsp8(b *testing.B) {
	benchNnTsp(8, b)
}

func BenchmarkNnTsp10(b *testing.B) {
	benchNnTsp(10, b)
}

func BenchmarkNnTsp1000(b *testing.B) {
	benchNnTsp(1000, b)
}
