package main

import (
	"testing"
)

func TestDistance(t *testing.T) {
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
