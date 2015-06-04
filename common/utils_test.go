package common

import (
	"reflect"
	"testing"
)

func TestShortestTour(t *testing.T) {
	c1 := []City{City{0, 0}, City{1, 1}, City{2, 2}}
	c2 := []City{City{0, 0}, City{2, 2}, City{8, 2}}
	t1 := NewTour(c1)
	t2 := NewTour(c2)
	tours := []Tour{*t1, *t2}
	out := ShortestTour(tours)
	if !reflect.DeepEqual(out.Cities(), (*t1).Cities()) {
		t.Error("Expected", t1, "found", out)
	}
}

func TestCities(t *testing.T) {
	c := Cities(3, 200, 100, 42)
	expected := []City{City{105, 87}, City{68, 50}, City{23, 45}}

	for i, v := range c {
		if v != expected[i] {
			t.Error("Expected", expected[i], "found", v)
		}
	}
}
