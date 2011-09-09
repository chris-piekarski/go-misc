package tree

// BinaryNode structure and methods support a generic binary tree with values of 'ints'
type BinaryNode struct {
	value int
	left *BinaryNode
	right *BinaryNode
	parent *BinaryNode
}

func (n *BinaryNode) IsLeaf() bool {
	return (n.left == nil) && (n.right == nil)
}

func (n *BinaryNode) IsRoot() bool {
	return (n.parent == nil)
}

func (n *BinaryNode) Parent() *BinaryNode {
	return n.parent
}

func (n *BinaryNode) IsNeighbor(a *BinaryNode) bool {
	return (a.parent == n.parent)
}

func (n *BinaryNode) Insert(value *BinaryNode) {
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

func (n *BinaryNode) Delete(value *BinaryNode) bool {
	return true
}

func (n *BinaryNode) Size() int {
	var s int = 1
	if n.left != nil {
		s += n.left.Size()
	}
	if n.right != nil {
		s += n.right.Size()
	}
	return s
}

func (n *BinaryNode) Search(v int) *BinaryNode {
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

func (n *BinaryNode) Minimum() *BinaryNode {
	var p *BinaryNode = n
	for p.left != nil {
		p = p.left
	}
	return p
}

func (n *BinaryNode) Maximum() *BinaryNode {
	var p *BinaryNode = n
	for p.right != nil {
		p = p.right
	}
	return p
}

// is the node with the smallest value greater than n.value
func (n *BinaryNode) Successor() *BinaryNode {
	if n.right != nil {
		return n.right.Minimum()
	}
	//go up tree from n until we find a node that is the left child of its parent

	var y *BinaryNode = n.parent
	var x *BinaryNode = n.right
	
	for y != nil && x == y.right {
		x = y
		y = y.parent	
	}
	return y
}

func (n *BinaryNode) InorderWalk(process (func (n *BinaryNode))) {
	if n.left != nil {
		n.left.InorderWalk(process)
	}

	process(n)

	if n.right != nil {
		n.right.InorderWalk(process)
	}
}

func (n *BinaryNode) PreorderWalk(process (func (n *BinaryNode))) {
	process(n)
	
	if n.left != nil {
		n.left.PreorderWalk(process)
	}

	if n.right != nil {
		n.right.PreorderWalk(process)
	}
}

func (n *BinaryNode) PostorderWalk(process (func (n *BinaryNode))) {
	if n.left != nil {
		n.left.PreorderWalk(process)
	}

	if n.right != nil {
		n.right.PreorderWalk(process)
	}
	
	process(n)
}
