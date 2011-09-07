package tree
import (
	"testing"
	"fmt"
)

func TestNode(t *testing.T){
	var n = Node{value:1}
	n.Insert(&Node{value:0})
	n.Insert(&Node{value:2})
	fmt.Println(n,n.IsLeaf(),n.Size())
}

func BenchmarkNode(b *testing.B){
}

