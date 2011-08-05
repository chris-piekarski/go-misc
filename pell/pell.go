package pell

// Pell Numbers
// 0, 1, 2, 5, 12, 29, 70, 169, 408, 985, 2378
// P0 = 0, P1 = 1, PN = 2PN-1 + PN-2

// Day 3: Do not communicate by sharing memory.
// Instead, share memory by communicating.

func Pell(ch chan<- uint64) {
	var fn_1, fn_2 uint64 = 1, 0
	ch <- 0
	ch <- 1
	for {
		fn_2, fn_1 = fn_1, ((2*fn_1) + fn_2)
		ch <- fn_1
	}
}
