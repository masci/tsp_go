package main

import (
	"reflect"
	"testing"
)

func TestCityDistance(t *testing.T) {
	a := City{3, 0}
	b := City{0, 4}
	d, err := a.distance(&b)
	if err != nil && d != 5.0 {
		t.Error("Expected 5.0, found", d)
	}

	c := City{-1, 0}
	d, err = a.distance(&c)
	if err == nil {
		t.Error("Expected an error")
	}
	if d != 0.0 {
		t.Error("Expected result 0.0, found", d)
	}
}

func TestCities(t *testing.T) {
	c := cities(3, 200, 100, 42)
	expected := []City{City{105, 87}, City{68, 50}, City{23, 45}}

	for i, v := range c {
		if v != expected[i] {
			t.Error("Expected", expected[i], "found", v)
		}
	}
}

func TestTourSwap(t *testing.T) {
	tour := Tour{cities(3, 100, 100, 42)}
	c1 := tour.cities[1]
	c2 := tour.cities[2]
	tour.swap(1, 2)
	if tour.cities[1] != c2 {
		t.Error("Expected", c2, "found", tour.cities[1])
	}
	if tour.cities[2] != c1 {
		t.Error("Expected", c1, "found", tour.cities[2])
	}
}

func TestTourAppend(t *testing.T) {
	tour := Tour{cities(2, 100, 100, 42)}
	c := City{0, 0}
	tour.append(c)
	if tour.cities[2] != c {
		t.Error("Expected", c, "found", tour.cities[2])
	}
}

func TestTourLength(t *testing.T) {
	c1 := City{2, 4}
	c2 := City{2, 8}
	c3 := City{7, 8}
	c4 := City{7, 4}

	tour := Tour{[]City{c1, c2, c3, c4}}
	if tour.length() != 18 {
		t.Error("Expected", 18, "found", tour.length())
	}

	// an empty Tour should have length 0
	var empty Tour
	if empty.length() != 0 {
		t.Error("Expected", 0, "found", empty.length())
	}
}

func TestTourContains(t *testing.T) {
	c := cities(2, 100, 100, 42)
	tour := Tour{c}

	if !tour.contains(c[1]) {
		t.Error("Expected", true, "found", false)
	}

	if tour.contains(City{101, 101}) {
		t.Error("Expected", false, "found", true)
	}
}

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

func TestShortestTour(t *testing.T) {
	t1 := Tour{[]City{City{0, 0}, City{1, 1}, City{2, 2}}}
	t2 := Tour{[]City{City{0, 0}, City{2, 2}, City{8, 2}}}
	tours := []Tour{t1, t2}
	out := shortest_tour(tours)
	if !reflect.DeepEqual(out, t1) {
		t.Error("Expected", t1, "found", out)
	}
}
