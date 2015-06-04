package main

import (
	"fmt"
	"github.com/masci/tsp_go/alltours"
	"github.com/masci/tsp_go/common"
	"github.com/masci/tsp_go/improved_alltours"
)

func main() {
	cs := common.Cities(8, 200, 100, 42)
	fmt.Println("Cities:", cs)
	fmt.Println()

	fmt.Println("Best tour with AlltoursTsp:")
	t := alltours.AlltoursTsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
	fmt.Println()
	common.PlotTour(t, "alltours")

	fmt.Println("Best tour with improved AlltoursTsp:")
	t = improved_alltours.AlltoursTsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
	fmt.Println()
	common.PlotTour(t, "improved-alltours")
}
