package sort

//import "fmt"

func swap(a []int, x int, y int){
	var temp = a[x]
	a[x] = a[y]
	a[y] = temp
}

func SelectionSort(a []int) {
	for i, _ := range a {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				swap(a, i, j)
			}
		}
	}
}

func InsertionSort(a []int) {
	for j := 1; j < len(a); j++ {
		var key = a[j]
		var i = j-1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i -1
		}
		a[i+1] = key
	}
}

func partition(a []int, p int, r int) int {
	var x = a[r]
	var i = p - 1
	for j := p; j <= r - 1; j++ {
		if a[j] <= x {
			i = i + 1
			swap(a,i,j)
		}
	}
	swap(a,i+1,r)
	return i + 1
}

func quickSort(a []int, p int, r int) {
	if p < r {
		var q = partition(a, p, r)
		quickSort(a,p,q-1)
		quickSort(a,q+1,r)
	}
}

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

// in Python these would be a lambda and C++ would be an inline
func parent(i int) int {
	return ((i-1)/2)
}

func left(i int) int {
	return ((2*i)+1)
}

func right(i int) int {
	return ((2*i)+2)
}

func MaxHeapify(a []int, i int, size int) {
	var left = left(i)
	var right = right(i)
	var largest = i
	
	if left < size && a[left] > a[i] {
		largest = left
	}
	
	if right < size && a[right] > a[largest] {
		largest = right
	}

	if largest != i {
		swap(a,i,largest)
		MaxHeapify(a, largest, size)
	}
}

func BuildMaxHeap(a []int) {
	for i:=(len(a)/2); i >= 0; i-- {
		MaxHeapify(a,i,len(a))
	}
}

func HeapSort(a []int) {
	BuildMaxHeap(a)
	var size = len(a)
	for i:=len(a)-1; i >=1; i-- {
		swap(a,0,i)
		size = size - 1
		MaxHeapify(a,0, size)
	}
}
