package pell

import "testing"

// index:    0 1 2 3 4  5  6  7   8   9
// Pell num: 0 1 2 5 12 29 70 169 408 985

func TestPell(t *testing.T) {
	expected := []uint64{0, 1, 2, 5, 12, 29, 70, 169, 408, 985}

	c := make(chan uint64)
	go Pell(c)

	for i, want := range expected {
		if got := <-c; got != want {
			t.Errorf("Pell[%d] = %d, want %d", i, got, want)
		}
	}
}
