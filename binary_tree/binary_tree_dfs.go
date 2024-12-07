package binarytree

// a slice to store the result of DFS
type dfsRes []int

func (res *dfsRes) preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	res.preOrder(node.Left)
	res.preOrder(node.Right)
}

func (res *dfsRes) inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	res.inOrder(node.Left)
	*res = append(*res, node.Val)
	res.inOrder(node.Right)
}

func (res *dfsRes) postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	res.postOrder(node.Left)
	res.postOrder(node.Right)
	*res = append(*res, node.Val)
}
