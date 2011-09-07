package tree

//import "fmt"

type Node struct {
	value int
	left *Node
	right *Node
	parent *Node
}

func (n *Node) IsLeaf() bool {
	return (n.left == nil) && (n.right == nil)
}

func (n *Node) IsRoot() bool {
	return (n.parent == nil)
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) IsNeighbor(a *Node) bool {
	return (a.parent == n.parent)
}

func (n *Node) Insert(value *Node) {
	if value.value <= n.value {
		if n.left != nil {
			n.left.Insert(value)
		} else {
			n.left = value
			value.parent = n.left
		}
	} else {
		if n.right != nil {
			n.right.Insert(value)
		} else {
			n.right = value
			value.parent = n.right
		}
	}
}

func (n *Node) Size() int {
	var s int = 1
	if n.left != nil {
		s += n.left.Size()
	}
	if n.right != nil {
		s += n.right.Size()
	}
	return s
}

func (n *Node) Search(v int) *Node {
	if n == nil || n.value == v {
		return n
	}

	if v < n.value && n.left != nil {
		return n.left.Search(v)
	} else if n.right != nil {
		return n.right.Search(v)
	}
	return nil
}

func (n *Node) Minimum() *Node {
	var p *Node = n
	for p.left != nil {
		p = p.left
	}
	return p
}

func (n *Node) Maximum() *Node {
	var p *Node = n
	for p.right != nil {
		p = p.right
	}
	return p
}

// is the node with the smallest value greater than n.value
func (n *Node) Successor() *Node {
	if n.right != nil {
		return n.right.Minimum()
	}
	//go up tree from n until we find a node that is the left child of its parent

	var y *Node = n.parent
	var x *Node = n.right
	
	for y != nil && x == y.right {
		x = y
		y = y.parent	
	}
	return y
}
