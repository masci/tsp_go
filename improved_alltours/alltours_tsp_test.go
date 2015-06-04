package improved_alltours

import (
	"testing"
)

func TestImprovedAlltours(t *testing.T) {
	c := cities(4, 100, 100, 42)
	tours := improved_alltours(c)

	if len(tours) != 6 {
		t.Error("Expected", 6, "found", len(tours))
	}

	for _, tour := range tours {
		for _, c := range c {
			if !tour.contains(c) {
				t.Error("Expected", c, "in tour", tour, "but was not found")
			}
		}
	}
}

func benchImprovedAlltoursTsp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		improved_alltours_tsp(cities(i, 900, 600, 42))
	}
}

func BenchmarkImprovedAlltoursTsp5(b *testing.B) {
	benchImprovedAlltoursTsp(5, b)
}

func BenchmarkImprovedAlltoursTsp8(b *testing.B) {
	benchImprovedAlltoursTsp(8, b)
}

func BenchmarkImprovedAlltoursTsp10(b *testing.B) {
	benchImprovedAlltoursTsp(10, b)
}
