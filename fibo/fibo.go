// Copyright 2011. No rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fibo

// 0 1 2 3 4 5 6 7  8  9
// 0 1 1 2 3 5 8 13 21 34

// Fn = Fn-1 + Fn-2
// Seed values Fo=0 and F1=1
func Fibo(operator (func(a int64, b int64) int64)) (func() int64){
	var fn_1 int64 = 1
	var fn_2 int64 = -1
	operation := operator
	nextValue := func() int64{
		// yield 0 -- python generators would work good here
		// yield 1
		next := operation(fn_1,fn_2)
		fn_2 = fn_1
		fn_1 = next
		return next
	}
	return nextValue
}
