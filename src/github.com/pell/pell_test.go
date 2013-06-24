package pell
import (
	"testing"
	"fmt"
)
// 0 1 2 3 4  5  6  7   8   9, 10
// 0,1,2,5,12,29,70,169,408,985,2378 

func TestPell(t *testing.T){
	expected := [10]uint64{0,1,2,5,12,29,70,169,408,985}
	actual := [10]uint64{0,0,0,0,0,0,0,0,0,0}
	
	c := make(chan uint64)
	go Pell(c)

	for i:=0;i<10;i++ {
		actual[i] = <-c
		fmt.Print(actual[i])
	}

	for j:=0;j<10;j++ {
		if actual[j] == expected[j]{
			fmt.Printf("Index: %d, expected: %d, actual: %d\n",j,actual[j], expected[j])
		}
	}
}
