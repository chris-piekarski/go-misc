package sort

import (
	"reflect"
	"testing"
)

func getSortedArray() []int {
	var a = make([]int, 1000)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	return a
}

func getUnsortedArray() []int {
	var a = make([]int, 1000)
	for i, j := 0, len(a)-1; i < len(a); i, j = i+1, j-1 {
		a[i] = j
	}
	return a
}

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}

// sorts is the table of in-place sort functions under test.
var sorts = map[string]func([]int){
	"SelectionSort": SelectionSort,
	"InsertionSort": InsertionSort,
	"QuickSort":     QuickSort,
	"HeapSort":      HeapSort,
	"MergeSort":     MergeSort,
}

func TestSorts(t *testing.T) {
	cases := [][]int{
		{},
		{1},
		{2, 1},
		{10, 40, 20, 30, 0, 60, 70, 50, 90, 80, 5, 15, 25},
		{3, 3, 1, 2, 2, 1},
		getUnsortedArray(),
	}

	for name, sortFn := range sorts {
		for _, in := range cases {
			input := append([]int(nil), in...)
			want := append([]int(nil), in...)
			builtinSort(want)

			sortFn(input)
			if !isSorted(input) {
				t.Errorf("%s did not sort %v -> %v", name, in, input)
			}
			if !reflect.DeepEqual(input, want) {
				t.Errorf("%s(%v) = %v, want %v", name, in, input, want)
			}
		}
	}
}

// builtinSort is a tiny insertion sort used only to compute the expected
// result independently of the implementations under test.
func builtinSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
}

func BenchmarkSelectionSort_Unsorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(getUnsortedArray())
	}
}

func BenchmarkInsertionSort_Unsorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(getUnsortedArray())
	}
}

func BenchmarkQuickSort_Unsorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(getUnsortedArray())
	}
}

func BenchmarkHeapSort_Unsorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(getUnsortedArray())
	}
}

func BenchmarkMergeSort_Unsorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(getUnsortedArray())
	}
}
