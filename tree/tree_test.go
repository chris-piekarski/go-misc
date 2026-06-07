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

// NOTE: Delete, Successor, and PostorderWalk are known-buggy (see AGENTS.md)
// and are intentionally left untested here; they are slated for a separate
// bug-fix PR.
