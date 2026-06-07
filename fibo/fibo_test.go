package fibo

import "testing"

// 0 1 2 3 4 5 6 7  8  9
// 0 1 1 2 3 5 8 13 21 34

func add(a, b int64) int64 { return a + b }

func TestFibo(t *testing.T) {
	expected := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	f := Fibo(add)
	for i, want := range expected {
		if got := f(); got != want {
			t.Errorf("Fibo()[%d] = %d, want %d", i, got, want)
		}
	}
}

func BenchmarkFibo(b *testing.B) {
	f := Fibo(add)
	for i := 0; i < b.N; i++ {
		f()
	}
}
