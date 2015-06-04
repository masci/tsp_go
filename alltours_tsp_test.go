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

func benchAllTsp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		alltours_tsp(cities(i, 900, 600, 42))
	}
}

func BenchmarkAllTsp5(b *testing.B) {
	benchAllTsp(5, b)
}

func BenchmarkAllTsp8(b *testing.B) {
	benchAllTsp(8, b)
}

func BenchmarkAllTsp10(b *testing.B) {
	benchAllTsp(10, b)
}
