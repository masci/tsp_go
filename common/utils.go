package common

import (
	"errors"
	"fmt"
	"github.com/masci/go-gnuplot"
	"math/rand"
)

// shortest_tour finds the shortest tour in the given array of tours
func ShortestTour(tours []Tour) Tour {
	min := Tour{}
	for _, t := range tours {
		if min.Length() == 0 || t.Length() < min.Length() {
			min = t
		}
	}
	return min
}

// cities returns a set of n cities, each with random coordinates
// within a (width x height) rectangle
func Cities(n int, width int, height int, seed int64) []City {
	r := rand.New(rand.NewSource(seed))
	out := make([]City, n)
	i := 0
	for i < n {
		out[i] = City{float64(r.Intn(width)), float64(r.Intn(height))}
		i++
	}

	return out
}

//
func PlotTour(tour Tour, fname string) error {
	if tour.Length() == 0 {
		return errors.New("Cannot plot a Tour with 0 length")
	}

	persist := false
	debug := false

	p, err := gnuplot.NewPlotter("", persist, debug)
	if err != nil {
		err_str := fmt.Sprintf("Error initializing gnuplot: %v", err)
		return errors.New(err_str)
	}
	defer p.Close()

	var xs, ys []float64
	for _, c := range tour.Cities() {
		xs = append(xs, c.x)
		ys = append(ys, c.y)
	}
	xs = append(xs, tour.Cities()[0].x)
	ys = append(ys, tour.Cities()[0].y)

	p.CheckedCmd("set for [i=1:5] linetype i dt i")
	p.SetStyle("lines")
	p.PlotXY(xs, ys, fname)
	// highlight starting City
	p.CheckedCmd(fmt.Sprintf("set object circle at first %f,%f radius char 0.5"+
		" fillstyle empty border lc rgb 'red' lw 2",
		tour.Cities()[0].x,
		tour.Cities()[0].y))
	p.CheckedCmd("set terminal png")
	p.CheckedCmd(fmt.Sprintf("set output '%s.png'", fname))
	p.CheckedCmd("replot")

	p.CheckedCmd("q")

	return nil
}

// Reverse a slice of City
func Reverse(cities []City) []City {
	for i := len(cities)/2 - 1; i >= 0; i-- {
		opp := len(cities) - 1 - i
		cities[i], cities[opp] = cities[opp], cities[i]
	}

	return cities
}

// Return ordered couples of indexes that form segments of tour of length N.
func AllSegments(n int) []int {
	ret := []int{}

	for l := n; l > 1; l-- {
		for start := 0; start < n-l+1; start++ {
			ret = append(ret, start)
			ret = append(ret, start+l)
		}
	}

	return ret
}

// Try to alter tour for the better by reversing segments.
func AlterTour(tour Tour) Tour {
	length := tour.Length()
	ret := NewTour(tour.Cities())

	segments := AllSegments(len(tour.Cities()))
	// we need at least 2 segments: (start, end, start, end)
	if len(segments) < 4 {
		return *ret
	}

	for index := 0; index < len(segments)-1; index += 2 {
		i := segments[index]
		j := segments[index+1]
		ret.ReverseSegmentIfBetter(i, j)
	}

	if ret.Length() < length {
		return AlterTour(*ret)
	}
	return *ret
}
