package binarytree

import "testing"

func TestBFS(t *testing.T) {
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
	res := root.BFS()
	t.Log(res)
}
