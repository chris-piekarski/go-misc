// Package mathx provides small number-theory helpers.
package mathx

// GCD returns the greatest common divisor of a and b using Euclid's algorithm.
// It operates on absolute values, so signs are ignored. GCD(0, 0) is 0.
func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a and b. It is 0 if either is 0.
// The result is non-negative.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	l := a / GCD(a, b) * b
	if l < 0 {
		l = -l
	}
	return l
}

// ModExp computes (base^exp) mod m using fast (binary) exponentiation.
// m must be greater than 0.
func ModExp(base, exp, m uint64) uint64 {
	if m == 1 {
		return 0
	}
	result := uint64(1)
	base %= m
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % m
		}
		exp >>= 1
		base = base * base % m
	}
	return result
}

// IsPrime reports whether n is a prime number using 6k±1 trial division.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	if n%3 == 0 {
		return n == 3
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Sieve returns all primes <= n using the Sieve of Eratosthenes.
// It returns nil for n < 2.
func Sieve(n int) []int {
	if n < 2 {
		return nil
	}
	composite := make([]bool, n+1)
	var primes []int
	for i := 2; i <= n; i++ {
		if composite[i] {
			continue
		}
		primes = append(primes, i)
		for j := i * i; j <= n; j += i {
			composite[j] = true
		}
	}
	return primes
}
