[![Build Status](https://travis-ci.org/masci/tsp_go.svg)](https://travis-ci.org/masci/tsp_go)

# tsp_go

This is a Go implementantion of [this Python notebook](http://nbviewer.ipython.org/url/norvig.com/ipython/TSPv3.ipynb) 
authored by Peter Norvig about *The Traveling Salesperson Problem*.

The problem is solved using different approaches:

 * **All Tours Algorithm**, ``alltours_tsp``: this algorithm is guaranteed to solve the problem but it is highly inefficient for large input data.
 * **Improved All Tours**, ``improved_alltours_tsp``: slightly improved version of the above.

To run the tests simply go with:

    go test -v ./...

To run benchmarks without running tests, issue the command: 

    go test -run=XXX -bench=.

These are the result on an early 2011 MacBook Pro:

    BenchmarkAlltoursTsp5	   10000	    197454 ns/op
    BenchmarkAlltoursTsp8	      20	  69997179 ns/op
    BenchmarkAlltoursTsp10	       1	9183772440 ns/op
    BenchmarkImprovedAlltoursTsp5	    2000	    653940 ns/op
    BenchmarkImprovedAlltoursTsp8	     100	  15741905 ns/op
    BenchmarkImprovedAlltoursTsp10	       2	 880900858 ns/op

