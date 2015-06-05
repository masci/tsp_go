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

func TestPopulation(t *testing.T) {
	cities := common.Cities(50, 200, 100, 42)
	out := sample(cities, 50, 42)

	if !reflect.DeepEqual(out, cities) {
		t.Error("Expected", cities, "found", out)
	}

	out = sample(cities, 3, 42)
	expected := []common.City{
		*common.NewCity(31, 7),
		*common.NewCity(129, 67),
		*common.NewCity(107, 79)}
	if !reflect.DeepEqual(out, expected) {
		t.Error("Expected", expected, "found", out)
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

func BenchmarkNnTsp100(b *testing.B) {
	benchNnTsp(100, b)
}

func BenchmarkNnTsp1000(b *testing.B) {
	benchNnTsp(1000, b)
}

func benchRepeatedNnTsp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		RepeatedNnTsp(common.Cities(i, 900, 600, 42))
	}
}

func BenchmarkRepeatedNnTsp5(b *testing.B) {
	benchRepeatedNnTsp(5, b)
}

func BenchmarkRepeatedNnTsp8(b *testing.B) {
	benchRepeatedNnTsp(8, b)
}

func BenchmarkRepeatedNnTsp10(b *testing.B) {
	benchRepeatedNnTsp(10, b)
}

func BenchmarkRepeatedNnTsp100(b *testing.B) {
	benchRepeatedNnTsp(100, b)
}

func benchSampledRepeatedNnTsp(i int, k int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SampledRepeatedNnTsp(common.Cities(i, 900, 600, 42), k)
	}
}

func BenchmarkRepeatedNnTsp5_3(b *testing.B) {
	benchSampledRepeatedNnTsp(5, 3, b)
}

func BenchmarkRepeatedNnTsp8_5(b *testing.B) {
	benchSampledRepeatedNnTsp(8, 5, b)
}

func BenchmarkRepeatedNnTsp10_5(b *testing.B) {
	benchSampledRepeatedNnTsp(10, 5, b)
}

func BenchmarkRepeatedNnTsp100_10(b *testing.B) {
	benchSampledRepeatedNnTsp(100, 10, b)
}

func BenchmarkRepeatedNnTsp300_100(b *testing.B) {
	benchSampledRepeatedNnTsp(300, 100, b)
}

func benchAlteredNnTsp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		NnTsp(common.Cities(i, 900, 600, 42))
	}
}

func BenchmarkAlteredNnTsp5(b *testing.B) {
	benchAlteredNnTsp(5, b)
}

func BenchmarkAlteredNnTsp8(b *testing.B) {
	benchAlteredNnTsp(8, b)
}

func BenchmarkAlteredNnTsp10(b *testing.B) {
	benchAlteredNnTsp(10, b)
}

func BenchmarkAlteredNnTsp100(b *testing.B) {
	benchAlteredNnTsp(100, b)
}

func BenchmarkAlteredNnTsp1000(b *testing.B) {
	benchAlteredNnTsp(1000, b)
}
