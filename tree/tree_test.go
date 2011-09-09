package tree
import (
	"testing"
	"fmt"
)

var root = BinaryNode{value:10}

func init(){
	BuildTree()
}

func BuildTree(){
	for i:=0;i < 20; i++ {
		root.Insert(&BinaryNode{value:i})
	}
}

func TestMin(t *testing.T){
	fmt.Println(root.Minimum())
}

func TestMax(t *testing.T){
	fmt.Println(root.Maximum())
}

func TestSize(t *testing.T){
	fmt.Println(root.Size())
}

func printNode(n *BinaryNode){
	fmt.Print(n.value," ")
}

func TestInorderWalk(t *testing.T){
	root.InorderWalk(printNode)
}

func TestPostorderWalk(t *testing.T){
	root.PostorderWalk(printNode)
}

func TestPreorderWalk(t *testing.T){
	root.PreorderWalk(printNode)
}

func BenchmarkNode(b *testing.B){
}

