package main

import (
	"fmt"
	"github.com/masci/tsp_go/alltours"
	"github.com/masci/tsp_go/common"
	"github.com/masci/tsp_go/improved_alltours"
	"github.com/masci/tsp_go/nearest_neighbor"
)

func main() {
	cs := common.Cities(8, 200, 100, 42)
	fmt.Println("Cities:", cs)
	fmt.Println()

	fmt.Println("Best tour with AlltoursTsp:")
	cs = common.Cities(8, 200, 100, 42)
	t := alltours.AlltoursTsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
	fmt.Println()
	common.PlotTour(t, "alltours")

	fmt.Println("Best tour with improved AlltoursTsp:")
	cs = common.Cities(8, 200, 100, 42)
	t = improved_alltours.AlltoursTsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
	fmt.Println()
	common.PlotTour(t, "improved-alltours")

	fmt.Println("Best tour with NnTsp:")
	cs = common.Cities(8, 200, 100, 42)
	t = nearest_neighbor.NnTsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
	fmt.Println()
	common.PlotTour(t, "nearest-neighbor")

	fmt.Println("Best tour with Repeated NnTsp:")
	cs = common.Cities(8, 200, 100, 42)
	t = nearest_neighbor.RepeatedNnTsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
	fmt.Println()
	common.PlotTour(t, "rep-nearest-neighbor")

	fmt.Println("Best tour with Sampled Repeated NnTsp:")
	cs = common.Cities(8, 200, 100, 42)
	t = nearest_neighbor.SampledRepeatedNnTsp(cs, 3)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
	fmt.Println()
	common.PlotTour(t, "sampled-rep-nearest-neighbor")

	fmt.Println("Best tour with altered NnTsp:")
	cs = common.Cities(8, 200, 100, 42)
	t = nearest_neighbor.AlteredNnTsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
	fmt.Println()
	common.PlotTour(t, "altered-nearest-neighbor")
}
