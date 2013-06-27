// Implements a Go concurrent Pell number generator

package main

import ( "fmt"
	 "flag"
	 "github.com/pell"
)

// Pell Numbers
// 0, 1, 2, 5, 12, 29, 70, 169, 408, 985, 2378
// P0 = 0, P1 = 1, PN = 2PN-1 + PN-2

func main() {
	c := make(chan uint64)
	var iterations *int = flag.Int("i", 10, "number of iterations to run")
	flag.Parse()

	go pell.Pell(c)
	for i:=0;i < *iterations; i++ {
		var num uint64 = <-c
		fmt.Println(num)
	}

}
