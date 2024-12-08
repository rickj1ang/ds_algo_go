package binarytree

import (
	"testing"
)

// need a new root node
var bst *binarySearchTree = &binarySearchTree{
	root: nil,
}

func TestInsert(t *testing.T) {
	// build a BST
	bst.insert(10)
	bst.insert(19)
	bst.insert(29)
	bst.insert(23)
	bst.insert(89)
	bst.insert(100)
	bst.insert(53)
	bst.insert(1)
	bst.insert(40)

	t.Log(bst.root.Left.Val, bst.root.Val, bst.root.Right.Val)

}

func TestSearch(t *testing.T) {
	t.Log(bst.root.Right.Right.Right.Val)
	res1 := bst.search(29)
	t.Log(res1.Val)
	res2 := bst.search(1)
	t.Log(res2.Val)
	res3 := bst.search(89)
	t.Log(res3.Val)
}

func TestRemove(t *testing.T) {
	t.Log("start remove\n")
	bst.remove(89)
	t.Log("remove success\n")
	res4 := bst.search(89)
	t.Log("get the res success")
	if res4 != nil {
		t.Error("delete fail")
	}
}
