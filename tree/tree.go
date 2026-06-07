package tree

// BinaryNode structure and methods support a generic binary tree with values of 'ints'
type BinaryNode struct {
	value  int
	left   *BinaryNode
	right  *BinaryNode
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
			value.parent = n
		}
	} else {
		if n.right != nil {
			n.right.Insert(value)
		} else {
			n.right = value
			value.parent = n
		}
	}
}

// Delete removes the first node holding value v from the subtree rooted at n
// and reports whether it was found.
//
// To keep the caller's root pointer valid even when the root itself is removed,
// deletion is done by value replacement: a two-child node's value is overwritten
// with its in-order successor, reducing the problem to unlinking a node with at
// most one child.
func (n *BinaryNode) Delete(v int) bool {
	target := n.Search(v)
	if target == nil {
		return false
	}

	// Reduce to deleting a node with at most one child.
	for target.left != nil && target.right != nil {
		succ := target.right.Minimum()
		target.value = succ.value
		target = succ
	}

	var child *BinaryNode
	if target.left != nil {
		child = target.left
	} else {
		child = target.right
	}
	if child != nil {
		child.parent = target.parent
	}

	switch {
	case target.parent == nil:
		// Removing the root (which now has <= 1 child): copy the child's
		// contents up so callers holding the root pointer stay valid.
		if child != nil {
			*target = *child
			if target.left != nil {
				target.left.parent = target
			}
			if target.right != nil {
				target.right.parent = target
			}
		} else {
			*target = BinaryNode{}
		}
	case target == target.parent.left:
		target.parent.left = child
	default:
		target.parent.right = child
	}
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
	var x *BinaryNode = n

	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (n *BinaryNode) InorderWalk(process func(n *BinaryNode)) {
	if n.left != nil {
		n.left.InorderWalk(process)
	}

	process(n)

	if n.right != nil {
		n.right.InorderWalk(process)
	}
}

func (n *BinaryNode) PreorderWalk(process func(n *BinaryNode)) {
	process(n)

	if n.left != nil {
		n.left.PreorderWalk(process)
	}

	if n.right != nil {
		n.right.PreorderWalk(process)
	}
}

func (n *BinaryNode) PostorderWalk(process func(n *BinaryNode)) {
	if n.left != nil {
		n.left.PostorderWalk(process)
	}

	if n.right != nil {
		n.right.PostorderWalk(process)
	}

	process(n)
}
