package main

import (
	"testing"
)

func TestAlltours(t *testing.T) {
	c := cities(4, 100, 100, 42)
	tours := alltours(c)
	if len(tours) != 24 {
		t.Error("Expected", 24, "found", len(tours))
	}

	for _, tour := range tours {
		for _, c := range c {
			if !tour.contains(c) {
				t.Error("Expected", c, "in tour", tour, "but was not found")
			}
		}
	}
}

func benchAlltoursTsp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		alltours_tsp(cities(i, 900, 600, 42))
	}
}

func BenchmarkAlltoursTsp5(b *testing.B) {
	benchAlltoursTsp(5, b)
}

func BenchmarkAlltoursTsp8(b *testing.B) {
	benchAlltoursTsp(8, b)
}

func BenchmarkAlltoursTsp10(b *testing.B) {
	benchAlltoursTsp(10, b)
}
