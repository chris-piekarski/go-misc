package tree

import (
	"reflect"
	"testing"
)

// buildTree returns a small BST:
//
//	      10
//	     /  \
//	    5    15
//	   / \   / \
//	  3   7 12  20
//	 /
//	1
func buildTree() *BinaryNode {
	root := &BinaryNode{value: 10}
	for _, v := range []int{5, 15, 3, 7, 12, 20, 1} {
		root.Insert(&BinaryNode{value: v})
	}
	return root
}

func collect(root *BinaryNode, walk func(*BinaryNode, func(*BinaryNode))) []int {
	var out []int
	walk(root, func(n *BinaryNode) { out = append(out, n.value) })
	return out
}

func TestInsertAndSearch(t *testing.T) {
	root := buildTree()

	for _, v := range []int{1, 3, 5, 7, 10, 12, 15, 20} {
		if got := root.Search(v); got == nil || got.value != v {
			t.Errorf("Search(%d) = %v, want a node with that value", v, got)
		}
	}
	if got := root.Search(99); got != nil {
		t.Errorf("Search(99) = %v, want nil", got)
	}
}

func TestMinMax(t *testing.T) {
	root := buildTree()
	if got := root.Minimum(); got.value != 1 {
		t.Errorf("Minimum() = %d, want 1", got.value)
	}
	if got := root.Maximum(); got.value != 20 {
		t.Errorf("Maximum() = %d, want 20", got.value)
	}
}

func TestSize(t *testing.T) {
	if got := buildTree().Size(); got != 8 {
		t.Errorf("Size() = %d, want 8", got)
	}
}

func TestIsLeafAndIsRoot(t *testing.T) {
	root := buildTree()
	if !root.IsRoot() {
		t.Error("root.IsRoot() = false, want true")
	}
	if root.IsLeaf() {
		t.Error("root.IsLeaf() = true, want false")
	}
	if leaf := root.Search(1); leaf == nil || !leaf.IsLeaf() {
		t.Errorf("node(1).IsLeaf() = false, want true")
	}
}

func TestInorderWalkIsSorted(t *testing.T) {
	root := buildTree()
	got := collect(root, (*BinaryNode).InorderWalk)
	want := []int{1, 3, 5, 7, 10, 12, 15, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InorderWalk = %v, want %v", got, want)
	}
}

func TestPreorderWalk(t *testing.T) {
	root := buildTree()
	got := collect(root, (*BinaryNode).PreorderWalk)
	want := []int{10, 5, 3, 1, 7, 15, 12, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PreorderWalk = %v, want %v", got, want)
	}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root := &BinaryNode{value: 0}
		for j := 1; j < 100; j++ {
			root.Insert(&BinaryNode{value: j})
		}
	}
}

func inorder(root *BinaryNode) []int {
	return collect(root, (*BinaryNode).InorderWalk)
}

// makeTree builds a BST from the given root value and inserted values.
func makeTree(root int, rest ...int) *BinaryNode {
	r := &BinaryNode{value: root}
	for _, v := range rest {
		r.Insert(&BinaryNode{value: v})
	}
	return r
}

func TestParent(t *testing.T) {
	root := buildTree()
	if root.Parent() != nil {
		t.Error("root.Parent() != nil")
	}
	if p := root.Search(5).Parent(); p == nil || p.value != 10 {
		t.Errorf("node(5).Parent() = %v, want node(10)", p)
	}
	if p := root.Search(1).Parent(); p == nil || p.value != 3 {
		t.Errorf("node(1).Parent() = %v, want node(3)", p)
	}
}

func TestIsNeighbor(t *testing.T) {
	root := buildTree()
	three, seven, twelve := root.Search(3), root.Search(7), root.Search(12)
	if !three.IsNeighbor(seven) {
		t.Error("3 and 7 should be neighbors (both children of 5)")
	}
	if three.IsNeighbor(twelve) {
		t.Error("3 and 12 should not be neighbors")
	}
}

func TestSuccessor(t *testing.T) {
	root := buildTree()
	cases := []struct {
		of, want int
		nilWant  bool
	}{
		{5, 7, false},   // has a right subtree
		{10, 12, false}, // has a right subtree
		{1, 3, false},   // no right child: immediate parent
		{7, 10, false},  // no right child: climbs up
		{20, 0, true},   // maximum -> nil
	}
	for _, c := range cases {
		got := root.Search(c.of).Successor()
		if c.nilWant {
			if got != nil {
				t.Errorf("Successor(%d) = %v, want nil", c.of, got)
			}
			continue
		}
		if got == nil || got.value != c.want {
			t.Errorf("Successor(%d) = %v, want %d", c.of, got, c.want)
		}
	}
}

func TestPostorderWalk(t *testing.T) {
	got := collect(buildTree(), (*BinaryNode).PostorderWalk)
	want := []int{1, 3, 7, 5, 12, 20, 15, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PostorderWalk = %v, want %v", got, want)
	}
}

func TestDeleteMissing(t *testing.T) {
	root := buildTree()
	if root.Delete(99) {
		t.Error("Delete(99) = true, want false")
	}
	if got := inorder(root); len(got) != 8 {
		t.Errorf("tree changed after deleting a missing value: %v", got)
	}
}

func TestDeleteLeaves(t *testing.T) {
	root := buildTree()
	if !root.Delete(20) { // right-child leaf
		t.Error("Delete(20) = false")
	}
	if !root.Delete(1) { // left-child leaf
		t.Error("Delete(1) = false")
	}
	want := []int{3, 5, 7, 10, 12, 15}
	if got := inorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("after deleting leaves = %v, want %v", got, want)
	}
}

func TestDeleteOneChild(t *testing.T) {
	root := buildTree() // node 3 has a single left child (1)
	if !root.Delete(3) {
		t.Error("Delete(3) = false")
	}
	want := []int{1, 5, 7, 10, 12, 15, 20}
	if got := inorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("after Delete(3) = %v, want %v", got, want)
	}
	if root.Search(3) != nil {
		t.Error("Search(3) != nil after delete")
	}
	if p := root.Search(1).Parent(); p == nil || p.value != 5 {
		t.Errorf("node(1).Parent() = %v, want node(5) after relink", p)
	}
}

func TestDeleteTwoChildren(t *testing.T) {
	root := buildTree() // node 5 has children 3 and 7
	if !root.Delete(5) {
		t.Error("Delete(5) = false")
	}
	want := []int{1, 3, 7, 10, 12, 15, 20}
	if got := inorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("after Delete(5) = %v, want %v", got, want)
	}
	if root.Search(5) != nil {
		t.Error("Search(5) != nil after delete")
	}
}

func TestDeleteRootTwoChildren(t *testing.T) {
	root := buildTree()
	if !root.Delete(10) {
		t.Error("Delete(10) = false")
	}
	want := []int{1, 3, 5, 7, 12, 15, 20}
	if got := inorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("after Delete(10) = %v, want %v", got, want)
	}
	if !root.IsRoot() {
		t.Error("root pointer should remain valid (IsRoot) after deleting root value")
	}
}

func TestDeleteRootOneChild(t *testing.T) {
	// Root with only a left child, which itself has a left child.
	left := makeTree(10, 5, 3)
	if !left.Delete(10) {
		t.Error("Delete(10) = false")
	}
	if got, want := inorder(left), []int{3, 5}; !reflect.DeepEqual(got, want) {
		t.Errorf("after Delete(10) = %v, want %v", got, want)
	}
	if left.value != 5 || !left.IsRoot() {
		t.Errorf("root = %d (IsRoot=%t), want value 5 root", left.value, left.IsRoot())
	}
	if p := left.Search(3).Parent(); p == nil || p.value != 5 {
		t.Errorf("node(3).Parent() = %v, want node(5)", p)
	}

	// Root with only a right child, which itself has a right child.
	right := makeTree(10, 15, 20)
	if !right.Delete(10) {
		t.Error("Delete(10) = false")
	}
	if got, want := inorder(right), []int{15, 20}; !reflect.DeepEqual(got, want) {
		t.Errorf("after Delete(10) = %v, want %v", got, want)
	}
	if p := right.Search(20).Parent(); p == nil || p.value != 15 {
		t.Errorf("node(20).Parent() = %v, want node(15)", p)
	}
}

func TestDeleteRootOnlyNode(t *testing.T) {
	root := &BinaryNode{value: 42}
	if !root.Delete(42) {
		t.Error("Delete(42) = false")
	}
	if root.Search(42) != nil {
		t.Error("Search(42) != nil after deleting the only node")
	}
}
