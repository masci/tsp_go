package common

import (
	"reflect"
	"testing"
)

func TestCityDistance(t *testing.T) {
	a := City{3, 0}
	b := City{0, 4}
	d, err := a.Distance(&b)
	if err != nil && d != 5.0 {
		t.Error("Expected 5.0, found", d)
	}

	c := City{-1, 0}
	d, err = a.Distance(&c)
	if err == nil {
		t.Error("Expected an error")
	}
	if d != 0.0 {
		t.Error("Expected result 0.0, found", d)
	}
}

func TestTourSwap(t *testing.T) {
	tour := NewTour(Cities(3, 100, 100, 42))
	c1 := tour.cities[1]
	c2 := tour.cities[2]
	tour.Swap(1, 2)
	if tour.cities[1] != c2 {
		t.Error("Expected", c2, "found", tour.cities[1])
	}
	if tour.cities[2] != c1 {
		t.Error("Expected", c1, "found", tour.cities[2])
	}
}

func TestTourAppend(t *testing.T) {
	tour := NewTour(Cities(2, 100, 100, 42))
	c := City{0, 0}
	tour.Append(c)
	if tour.cities[2] != c {
		t.Error("Expected", c, "found", tour.cities[2])
	}
}

func TestTourLength(t *testing.T) {
	c1 := City{2, 4}
	c2 := City{2, 8}
	c3 := City{7, 8}
	c4 := City{7, 4}

	tour := NewTour([]City{c1, c2, c3, c4})
	if tour.Length() != 18 {
		t.Error("Expected", 18, "found", tour.Length())
	}

	// an empty Tour should have length 0
	var empty Tour
	if empty.Length() != 0 {
		t.Error("Expected", 0, "found", empty.Length())
	}
}

func TestTourContains(t *testing.T) {
	c := Cities(2, 100, 100, 42)
	tour := NewTour(c)

	if !tour.Contains(c[1]) {
		t.Error("Expected", true, "found", false)
	}

	if tour.Contains(City{101, 101}) {
		t.Error("Expected", false, "found", true)
	}
}

func TestTourReverseSegmentIfBetter(t *testing.T) {
	tour := NewTour([]City{*NewCity(9, 3), *NewCity(3, 10), *NewCity(2, 16), *NewCity(3, 21), *NewCity(9, 28),
		*NewCity(26, 3), *NewCity(32, 10), *NewCity(33, 16), *NewCity(32, 21), *NewCity(26, 28)})

	expected := []City{*NewCity(9, 3), *NewCity(3, 10), *NewCity(2, 16), *NewCity(3, 21), *NewCity(9, 28),
		*NewCity(26, 28), *NewCity(32, 21), *NewCity(33, 16), *NewCity(32, 10), *NewCity(26, 3)}

	tour.ReverseSegmentIfBetter(5, 9)

	if !reflect.DeepEqual(tour.Cities(), expected) {
		t.Error("Expected", expected, "found", tour.Cities())
	}
}
