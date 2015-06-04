// Everything lives in the main package
package main

import (
	"fmt"
	"github.com/masci/tsp_go/alltours"
	"github.com/masci/tsp_go/common"
)

func main() {
	cs := common.Cities(8, 200, 100, 42)
	fmt.Println("Cities:", cs)

	fmt.Println("Best tour with alltours_tsp:")
	t := alltours.AlltoursTsp(cs)
	fmt.Printf("%+v, length:%f\n", t, t.Length())
}
