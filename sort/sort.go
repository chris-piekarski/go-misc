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
