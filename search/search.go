// Package search provides generic search algorithms over sorted slices.
package search

import "cmp"

// BinarySearch looks for target in the sorted slice a. It returns the index of
// target and true if present. If target is absent it returns the index where
// target would be inserted to keep a sorted, and false.
//
// a must be sorted in ascending order; the behavior is undefined otherwise.
// It runs in O(log n) time.
func BinarySearch[T cmp.Ordered](a []T, target T) (int, bool) {
	lo, hi := 0, len(a)
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch {
		case a[mid] < target:
			lo = mid + 1
		case a[mid] > target:
			hi = mid
		default:
			return mid, true
		}
	}
	return lo, false
}
