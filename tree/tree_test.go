package tree
import (
	"testing"
	"fmt"
)

func TestNode(t *testing.T){
	var n = Node{1,nil,nil}
	fmt.Println(n)
}

func BenchmarkNode(b *testing.B){
}

