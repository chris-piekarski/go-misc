package mathx

import (
	"reflect"
	"testing"
)

func TestGCD(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{48, 18, 6},
		{18, 48, 6},
		{-48, 18, 6}, // a < 0
		{48, -18, 6}, // b < 0
		{0, 0, 0},    // loop never runs
		{0, 5, 5},
		{5, 0, 5},
		{17, 13, 1}, // coprime
	}
	for _, c := range cases {
		if got := GCD(c.a, c.b); got != c.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestLCM(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{4, 6, 12},
		{0, 5, 0},   // a == 0
		{5, 0, 0},   // b == 0
		{-4, 6, 12}, // negative -> non-negative result
		{21, 6, 42},
	}
	for _, c := range cases {
		if got := LCM(c.a, c.b); got != c.want {
			t.Errorf("LCM(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestModExp(t *testing.T) {
	cases := []struct {
		base, exp, m, want uint64
	}{
		{2, 10, 1000, 24}, // 1024 % 1000
		{3, 3, 7, 6},      // 27 % 7
		{5, 0, 7, 1},      // exp == 0
		{2, 10, 1, 0},     // m == 1
		{7, 256, 13, 9},
	}
	for _, c := range cases {
		if got := ModExp(c.base, c.exp, c.m); got != c.want {
			t.Errorf("ModExp(%d, %d, %d) = %d, want %d", c.base, c.exp, c.m, got, c.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	primes := map[int]bool{
		-5: false, 0: false, 1: false, // n < 2
		2: true, 4: false, // even branch
		3: true, 9: false, // divisible-by-3 branch
		25: false,          // n%i == 0 in loop
		49: false,          // n%(i+2) == 0 in loop
		37: true, 97: true, // pass the loop
	}
	for n, want := range primes {
		if got := IsPrime(n); got != want {
			t.Errorf("IsPrime(%d) = %t, want %t", n, got, want)
		}
	}
}

func TestSieve(t *testing.T) {
	if got := Sieve(1); got != nil {
		t.Errorf("Sieve(1) = %v, want nil", got)
	}
	got := Sieve(30)
	want := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sieve(30) = %v, want %v", got, want)
	}
}

func TestSieveAgreesWithIsPrime(t *testing.T) {
	const n = 200
	set := make(map[int]bool)
	for _, p := range Sieve(n) {
		set[p] = true
	}
	for i := 0; i <= n; i++ {
		if set[i] != IsPrime(i) {
			t.Errorf("disagreement at %d: Sieve=%t IsPrime=%t", i, set[i], IsPrime(i))
		}
	}
}
