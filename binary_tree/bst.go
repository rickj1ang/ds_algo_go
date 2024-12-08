package binarytree

type binarySearchTree struct {
	root *TreeNode
}

func (bst *binarySearchTree) search(val int) *TreeNode {
	node := bst.root
	for node != nil {
		if node.Val < val {
			node = node.Right
		} else if node.Val > val {
			node = node.Left
		} else {
			break
		}
	}
	return node
}

func (bst *binarySearchTree) insert(val int) {
	cur := bst.root
	if cur == nil {
		bst.root = NewTreeNode(val)
		return
	}

	var pre *TreeNode = nil
	for cur != nil {
		if cur.Val == val {
			return
		}

		pre = cur

		if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	newNode := NewTreeNode(val)
	if pre.Val < val {
		pre.Right = newNode
	} else {
		pre.Left = newNode
	}
}

func (bst *binarySearchTree) remove(val int) {
	// get the root
	cur := bst.root
	// root should not be nil
	if cur == nil {
		return
	}

	// need to record the parent node of cur
	var pre *TreeNode = nil
	// find the Node we need to delete
	for cur != nil {
		if cur.Val == val {
			break
		}
		pre = cur
		if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	if cur == nil {
		return
	}

	// The node we plan to delete has 0 or 1 subNode
	if cur.Left == nil || cur.Right == nil {
		// find the child node of cur which side
		var child *TreeNode = nil
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}

		if child != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			// The node we need to delete is root ndoe
			bst.root = child
		}
	} else {
		// THe node we need to delete has 2 subNodes
		// replace the node we need to delete with
		// the next node in inOrder traverse
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}

		// remove the next Node first and replace
		bst.remove(tmp.Val)
		cur.Val = tmp.Val
	}
}
