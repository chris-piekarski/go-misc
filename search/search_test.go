package search

import "testing"

func TestBinarySearchInts(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 11}

	cases := []struct {
		target    int
		wantIndex int
		wantFound bool
	}{
		{5, 2, true},   // middle
		{1, 0, true},   // first
		{11, 5, true},  // last
		{0, 0, false},  // below all -> insert at front
		{12, 6, false}, // above all -> insert at end
		{6, 3, false},  // between -> insertion point
	}

	for _, c := range cases {
		gotIndex, gotFound := BinarySearch(a, c.target)
		if gotIndex != c.wantIndex || gotFound != c.wantFound {
			t.Errorf("BinarySearch(%v, %d) = (%d, %t), want (%d, %t)",
				a, c.target, gotIndex, gotFound, c.wantIndex, c.wantFound)
		}
	}
}

func TestBinarySearchEmpty(t *testing.T) {
	if i, found := BinarySearch([]int{}, 42); i != 0 || found {
		t.Errorf("BinarySearch(empty, 42) = (%d, %t), want (0, false)", i, found)
	}
}

func TestBinarySearchStrings(t *testing.T) {
	a := []string{"apple", "banana", "cherry"}
	if i, found := BinarySearch(a, "banana"); i != 1 || !found {
		t.Errorf(`BinarySearch(%v, "banana") = (%d, %t), want (1, true)`, a, i, found)
	}
	if i, found := BinarySearch(a, "date"); i != 3 || found {
		t.Errorf(`BinarySearch(%v, "date") = (%d, %t), want (3, false)`, a, i, found)
	}
}

func TestBinarySearchMatchesLinearScan(t *testing.T) {
	a := make([]int, 0, 256)
	for i := 0; i < 256; i++ {
		a = append(a, i*2) // even numbers 0..510
	}
	for target := -1; target <= 512; target++ {
		gotIndex, gotFound := BinarySearch(a, target)
		wantIndex, wantFound := linearSearch(a, target)
		if gotIndex != wantIndex || gotFound != wantFound {
			t.Fatalf("BinarySearch(_, %d) = (%d, %t), want (%d, %t)",
				target, gotIndex, gotFound, wantIndex, wantFound)
		}
	}
}

// linearSearch is an independent oracle: first index >= target, and whether it equals.
func linearSearch(a []int, target int) (int, bool) {
	for i, v := range a {
		if v == target {
			return i, true
		}
		if v > target {
			return i, false
		}
	}
	return len(a), false
}

func BenchmarkBinarySearch(b *testing.B) {
	a := make([]int, 1<<16)
	for i := range a {
		a[i] = i
	}
	for i := 0; i < b.N; i++ {
		BinarySearch(a, i&0xFFFF)
	}
}
