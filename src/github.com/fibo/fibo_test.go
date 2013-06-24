package fibo
import (
	"testing"
	"fmt"
)
// 0 1 2 3 4 5 6 7  8  9
// 0 1 1 2 3 5 8 13 21 34

func TestFibo(t *testing.T){
	expected := [10]int64{0,1,1,2,3,5,8,13,21,34}
	actual := [10]int64{0,0,0,0,0,0,0,0,0,0}
	oper := func(a int64, b int64) int64{
		return a + b
	}
	f := Fibo(oper)

	for i:=0;i<10;i++ {
		actual[i] = f()
		fmt.Print(actual[i])
	}

	for j:=0;j<10;j++ {
		if actual[j] == expected[j]{
			fmt.Printf("Index: %d, expected: %d, actual: %d\n",j,actual[j], expected[j])
		}
	}
}
