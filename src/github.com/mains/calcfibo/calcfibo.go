// Copyright 2011. No rights reserved.

package main

import (
	"fmt"
	"flag"
	"github.com/fibo"
)

// 0 1 2 3 4 5 6 7  8  9
// 0 1 1 2 3 5 8 13 21 34

// Fn = Fn-1 + Fn-2
// Seed values F0=0 and F1=1

func main() {
	var iterations *int = flag.Int("i", 10, "number of iterations to run")
	flag.Parse()

	oper := func(a int64, b int64) int64{
		return a + b
	}
	f := fibo.Fibo(oper)
	for i:=0; i< *iterations; i++ {
		fmt.Print(f(), "\n")
	}
}


