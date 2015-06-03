package main

import (
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
}

func TestAlltours(t *testing.T) {
	var tours []Tour
	alltours(&tours, cities(3, 100, 100, 42), 0)
	if len(tours) != 6 {
		t.Error("Expected", 6, "found", len(tours))
	}
}
