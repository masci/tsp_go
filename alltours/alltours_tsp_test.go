package alltours

import (
	"github.com/masci/tsp_go/common"
	"reflect"
	"testing"
)

func TestAlltours(t *testing.T) {
	c := common.Cities(4, 100, 100, 42)
	tours := alltours(c)
	if len(tours) != 24 {
		t.Error("Expected", 24, "found", len(tours))
	}

	for _, tour := range tours {
		for _, c := range c {
			if !tour.Contains(c) {
				t.Error("Expected", c, "in tour", tour, "but was not found")
			}
		}
	}

	for i, tour := range tours {
		for j := i + 1; j < len(tours)-1; j++ {
			if reflect.DeepEqual(tour.Cities(), tours[j].Cities()) {
				t.Error("Found double tour", i, j)
			}
		}
	}
}

func TestAlltoursQp(t *testing.T) {
	c := common.Cities(4, 100, 100, 42)
	res := alltoursQp(c)
	exp := alltours(c)

	for _, tres := range res {
		found := false
		for _, texp := range exp {
			if reflect.DeepEqual(tres.Cities(), texp.Cities()) {
				found = true
				break
			}
		}
		if !found {
			t.Error("Tour", tres, "not found in expected", exp)
		}
	}
}

func benchAlltoursTsp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		AlltoursTsp(common.Cities(i, 900, 600, 42))
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

func benchAlltoursTspQp(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		AlltoursTspQp(common.Cities(i, 900, 600, 42))
	}
}

func BenchmarkAlltoursTspQp5(b *testing.B) {
	benchAlltoursTspQp(5, b)
}

func BenchmarkAlltoursTspQp8(b *testing.B) {
	benchAlltoursTspQp(8, b)
}

func BenchmarkAlltoursTspQp10(b *testing.B) {
	benchAlltoursTspQp(10, b)
}
