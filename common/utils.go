package common

import (
	"bitbucket.org/binet/go-gnuplot/pkg/gnuplot"
	"fmt"
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
func PlotTour(tour Tour, fname string) {
	persist := false
	debug := true

	p, err := gnuplot.NewPlotter("", persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
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
}
