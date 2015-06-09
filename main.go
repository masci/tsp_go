package main

import (
	"fmt"
	"github.com/masci/tsp_go/alltours"
	"github.com/masci/tsp_go/common"
	"github.com/masci/tsp_go/greedy"
	"github.com/masci/tsp_go/improved_alltours"
	"github.com/masci/tsp_go/nearest_neighbor"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"sync"
	"time"
)

var (
	app     = kingpin.New("tsp_go", "The TSP problem solver.")
	verbose = app.Flag("verbose", "Verbose mode").Short('v').Bool()

	list = app.Command("list", "List available alogrithms")

	run       = app.Command("run", "Run solution algorithms")
	name      = run.Arg("name", "Algorithm name").Required().String()
	numcities = run.Arg("cities", "Number of cities").Default("8").Int()
	plot      = run.Flag("plot", "Produce plots if gunplot available").Short('p').Bool()
	cpus      = run.Arg("cpus", "Number of cpus to use").Default("1").Int()

	algos = map[string]func([]common.City) common.Tour{
		"AlltoursTsp":          alltours.AlltoursTsp,
		"AlltoursTspQp":        alltours.AlltoursTspQp,
		"ImprovedAlltoursTsp":  improved_alltours.AlltoursTsp,
		"NnTsp":                nearest_neighbor.NnTsp,
		"RepeatedNnTsp":        nearest_neighbor.RepeatedNnTsp,
		"AlteredNnTsp":         nearest_neighbor.AlteredNnTsp,
		"GreedyTsp":            greedy.GreedyTsp,
		"SampledRepeatedNnTsp": nearest_neighbor.AutoSampledRepeatedNnTsp,
	}
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case list.FullCommand():
		println("Available algorithms:\n")
		for k, _ := range algos {
			println("\t", k)
		}
		println("\t", "SampledRepeatedNnTsp\n")
		os.Exit(0)

	// Post message
	case run.FullCommand():
		var wg sync.WaitGroup

		var scheduled []string
		if *name == "all" {
			for k := range algos {
				scheduled = append(scheduled, k)
			}
		} else {
			scheduled = append(scheduled, *name)
		}

		for _, s := range scheduled {

			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				out := ""

				var res common.Tour
				cs := common.Cities(*numcities, 200, 100, 42)
				if *verbose {
					out += fmt.Sprintln("Cities:", cs)
				}

				fun, ok := algos[s]
				if !ok {
					println(s, "algorithm not found, exiting...")
					os.Exit(1)
				}

				out += fmt.Sprintf("Best tour for %d cities with %s algorithm\n", len(cs), s)
				t0 := time.Now()
				res = fun(cs)
				t1 := time.Now()
				out += fmt.Sprintf("Took %v to run.\n", t1.Sub(t0))

				if *verbose {
					out += fmt.Sprintln(res)
				} else {
					out += fmt.Sprintln("Lenght:", res.Length())
				}

				if *plot {
					common.PlotTour(res, s)
					out += fmt.Sprintf("Tour plotted to %s.png\n", s)
				}

				fmt.Println(out)
			}(s)

		}

		wg.Wait()

	}
}
