[![Build Status](https://travis-ci.org/masci/tsp_go.svg)](https://travis-ci.org/masci/tsp_go)
[![GoDoc](https://godoc.org/github.com/masci/tsp_go?status.svg)](https://godoc.org/github.com/masci/tsp_go)

# tsp_go

This is a Go implementantion of part of [this Python notebook](http://nbviewer.ipython.org/url/norvig.com/ipython/TSPv3.ipynb) 
authored by Peter Norvig about *The Traveling Salesperson Problem*.

The problem is solved using different approaches:

 * **All Tours Algorithm**, ``alltours.AlltoursTsp``: this algorithm is guaranteed to solve the problem but it is highly inefficient for large input data.
 * **Improved All Tours**, ``improved_alltours_tsp.AlltoursTsp``: slightly improved version of the above.
 * **Nearest Neighbor Algorithm**, ``nearest_neighbor.NnTsp``: basic optimization algorithm.
 * **Repeated Nearest Neighbor Algorithm**, ``nearest_neighbor.RepeatedNnTsp``: repeat *nn* starting each time from a different city and take the shortest tour.
 * **Sampled Repeated Nearest Neighbor Algorithm**, ``nearest_neighbor.SampledRepeatedNnTsp``: repeat *nn* starting each time from a subset of different cities and take the shortest tour.
 * **Altered Nearest Neighbor Algorithm**, ``nearest_neighbor.AlteredNnTsp``: try to revert segments to eliminate crossing and find better solutions
 * **Greedy Algorithm**, ``greedy.GreedyTsp``: greedily join segments until you form a complete tour

## Requirements

 * go 1.4
 * gnuplot (only if you want to plot solutions)

The ``tsp_go`` executable will exercise all the algorithm with a small set of cities and plot the computed tours if *gnuplot* is installed.

To run the tests simply go with:

    go test -v ./...

To run benchmarks without running tests, issue the command: 

    go test -run=XXX -bench=. ./...

**Warning** benchmarks are run on just one set of cities (a *Map* using Norvig's terminology) so the results are far less than accurate. These are the numbers on an early 2011 MacBook Pro:

    BenchmarkAlltoursTsp5           10000   197454 ns/op
    BenchmarkAlltoursTsp8           20      69997179 ns/op
    BenchmarkAlltoursTsp10          1       9183772440 ns/op

    BenchmarkImprovedAlltoursTsp5   2000    653940 ns/op
    BenchmarkImprovedAlltoursTsp8   100     15741905 ns/op
    BenchmarkImprovedAlltoursTsp10  2       880900858 ns/op

    BenchmarkNnTsp5                 50000   38159 ns/op
    BenchmarkNnTsp8                 30000   45066 ns/op
    BenchmarkNnTsp10                30000   54629 ns/op
    BenchmarkNnTsp100               500     2475007 ns/op
    BenchmarkNnTsp1000              5       230256556 ns/op

    BenchmarkRepeatedNnTsp5         20000   85580 ns/op
    BenchmarkRepeatedNnTsp8         10000   200088 ns/op
    BenchmarkRepeatedNnTsp10        5000    358738 ns/op
    BenchmarkRepeatedNnTsp100       5       230028522 ns/op

    BenchmarkRepeatedNnTsp10        5000    360743 ns/op
    BenchmarkRepeatedNnTsp100       5       229716757 ns/op
    BenchmarkRepeatedNnTsp5_3       20000   88453 ns/op
    BenchmarkRepeatedNnTsp8_5       10000   167354 ns/op
    BenchmarkRepeatedNnTsp10_5      10000   209789 ns/op
    BenchmarkRepeatedNnTsp100_10    50      24378681 ns/op
    BenchmarkRepeatedNnTsp300_100   1       1882211378 ns/op

    BenchmarkAlteredNnTsp5          50000   39136 ns/op
    BenchmarkAlteredNnTsp8          30000   44988 ns/op
    BenchmarkAlteredNnTsp10         30000   54651 ns/op
    BenchmarkAlteredNnTsp100        500     2549538 ns/op
    BenchmarkAlteredNnTsp1000       5       228328117 ns/op

    BenchmarkGreedyTsp5             30000   53022 ns/op
    BenchmarkGreedyTsp8             10000   105980 ns/op
    BenchmarkGreedyTsp10            10000   180379 ns/op
    BenchmarkGreedyTsp100           100     17568525 ns/op
    BenchmarkGreedyTsp1000          1       4146182727 ns/op
