package main

import ( "fmt" )

//Closure for calculating sum & average over []int
func stats() (func(a []int) int, func() float64){
	sum := 0
	entries := 0
	log := func(){
		fmt.Printf("sum: %d, entries: %d\n",sum,entries)
	}
	return func(l []int) int {
		defer log()
		for i:=0;i<len(l);i++ {
			sum += l[i]
			entries += 1
		}
		return sum
	},
	func() float64 {
		return float64(sum)/float64(entries)
	}
}

fun main() {
	sum, avg = stats([]int{1,2,1,2})
	sum([]int{1,1,1})
	fmt.Println(avg())
}
