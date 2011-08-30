package sort
import (
	"testing"
	"fmt"
)

func getSortedArray() []int {
	var a = make([]int, 1000)
	for i:=0; i < len(a); i++ {
		a[i] = i
	}
	return a
}

func getUnsortedArray() []int {
	var a = make([]int, 1000)
	for i,j:=0,len(a)-1; i < len(a); i,j = i+1, j-1 {
		a[i] = j
	}
	return a
}

func TestSelectionSort(t *testing.T){
	unsorted := []int{10,40,20,30,0,60,70,50,90,80}
	fmt.Println(unsorted)
	SelectionSort(unsorted)
	fmt.Println(unsorted)
}

func BenchmarkSelectionSort_Unsorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		SelectionSort(getUnsortedArray())
	}
}

func BenchmarkSelectionSort_Sorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		SelectionSort(getSortedArray())
	}
}


func TestInsertionSort(t *testing.T){
	unsorted := []int{10,40,20,30,0,60,70,50,90,80}
	fmt.Println(unsorted)
	InsertionSort(unsorted)
	fmt.Println(unsorted)
}

func BenchmarkInsertionSort_Unsorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		InsertionSort(getUnsortedArray())
	}
}

func BenchmarkInsertionSort_Sorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		InsertionSort(getSortedArray())
	}
}


func TestQuickSort(t *testing.T){
	unsorted := []int{10,40,20,30,0,60,70,50,90,80,5,15,25}
	fmt.Println(unsorted)
	QuickSort(unsorted)
	fmt.Println(unsorted)
}

func BenchmarkQuickSort_Unsorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		QuickSort(getUnsortedArray())
	}
}

func BenchmarkQuickSort_Sorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		QuickSort(getSortedArray())
	}
}

func TestHeapSort(t *testing.T){
	unsorted := []int{10,40,20,30,0,60,70,50,90,80,5,15,25}
	fmt.Println(unsorted)
	HeapSort(unsorted)
	fmt.Println(unsorted)
}

func BenchmarkHeapSort_Unsorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		HeapSort(getUnsortedArray())
	}
}

func BenchmarkHeapSort_Sorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		HeapSort(getSortedArray())
	}
}
