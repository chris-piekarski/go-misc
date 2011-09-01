package tree

//import "fmt"

type Node struct {
	value int
	left *Node
	right *Node	
}

func (n *Node) isLeaf() bool {
	return false
}
