package alltours

import (
	"fmt"
	"github.com/masci/tsp_go/common"
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
}

func TestAlltoursQp(t *testing.T) {
	c := common.Cities(3, 100, 100, 42)
	fmt.Println(alltoursQp(c))
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
