package binarytree

import "testing"

func TestDFS(t *testing.T) {
	/*
	          7
	        /  \
	      4     10
	    /  \    / \
	   30  27  31 18
	*/
	root := NewTreeNode(7)
	root.Left = NewTreeNode(4)
	root.Right = NewTreeNode(10)
	root.Left.Left = NewTreeNode(30)
	root.Left.Right = NewTreeNode(27)
	root.Right.Left = NewTreeNode(31)
	root.Right.Right = NewTreeNode(18)

	var preRes dfsRes
	preRes.preOrder(root)

	var inRes dfsRes
	inRes.inOrder(root)

	var postRes dfsRes
	postRes.postOrder(root)

	t.Log("pre:\n")
	t.Log(preRes)
	t.Log("in:\n")
	t.Log(inRes)
	t.Log("post:\n")
	t.Log(postRes)
}
