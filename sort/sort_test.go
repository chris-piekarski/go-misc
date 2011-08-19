package sort
import (
	"testing"
	"fmt"
)

func TestSelectionSort(t *testing.T){
	unsorted := []int{10,40,20,30,0,60,70,50,90,80}
	fmt.Println(unsorted)
	SelectionSort(unsorted)
	fmt.Println(unsorted)
}

func BenchmarkSelectionSort_Unsorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		SelectionSort([]int{130,20,0,120,110,10,30,90,40,80,100,70,50,60})
	}
}

func BenchmarkSelectionSort_Sorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		SelectionSort([]int{0,10,20,30,40,50,60,70,80,90,100,110,120,130})
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
		InsertionSort([]int{130,20,0,120,110,10,30,90,40,80,100,70,50,60})
	}
}

func BenchmarkInsertionSort_Sorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		InsertionSort([]int{0,10,20,30,40,50,60,70,80,90,100,110,120,130})
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
		QuickSort([]int{130,20,0,120,110,10,30,90,40,80,100,70,50,60})
	}
}

func BenchmarkQuickSort_Sorted(b *testing.B){
	for i:=0; i < b.N; i++ {
		QuickSort([]int{0,10,20,30,40,50,60,70,80,90,100,110,120,130})
	}
}

