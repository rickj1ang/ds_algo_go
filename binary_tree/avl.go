package binarytree

import "container/list"

// in the same package there is a ordinary TreeNode, be care
type nodeAVL struct {
	Val    int
	Height int
	Left   *nodeAVL
	Right  *nodeAVL
}

// get Height of node, if node == nil return -1
func (node *nodeAVL) height() int {
	if node != nil {
		return node.Height
	}
	return -1
}

func (node *nodeAVL) updateHeight() {
	lh := node.Left.height()
	rh := node.Right.height()
	// height of a node = max(left subtree, right subtree) + 1
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

// get BF of the node BF = height(left subNode) - height(right subNode)
// BF of any AVL tree should >= -1 && <= 1
func (node *nodeAVL) balanceFactor() int {
	if node == nil {
		return 0
	}
	return node.Left.height() - node.Right.height()
}

// right rotate is a kind of abstraction
// It just a way to balance when the unbalanced side is Left subTree
func (node *nodeAVL) rightRotate() *nodeAVL {
	child := node.Left
	grandChild := child.Right

	child.Right = node
	node.Left = grandChild

	// update height of nodes we have changed
	node.updateHeight()
	child.updateHeight()

	return child
}

// same as right rotate but the different unnbalanced subTree
func (node *nodeAVL) leftRotate() *nodeAVL {
	child := node.Right
	grandChild := child.Left

	child.Left = node
	node.Right = grandChild

	// update height of nodes we have changed
	node.updateHeight()
	child.updateHeight()

	return child
}

func (node *nodeAVL) rotate() *nodeAVL {
	bf := node.balanceFactor()

	if bf > 1 {
		leftNodeBF := node.Left.balanceFactor()
		if leftNodeBF >= 0 {
			node = node.rightRotate()
			return node
		} else {
			node.Left = node.Left.leftRotate()
			node = node.rightRotate()
			return node
		}
	}
	if bf < -1 {
		rightNodeBF := node.Right.balanceFactor()
		if rightNodeBF <= 0 {
			node = node.leftRotate()
			return node
		} else {
			node.Right = node.Right.rightRotate()
			node = node.leftRotate()
			return node
		}
	}
	return node
}

type avlTree struct {
	root *nodeAVL
}

func (t *avlTree) insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

func NewAvlNode(val int) *nodeAVL {
	return &nodeAVL{
		Val:    val,
		Height: 0,
		Left:   nil,
		Right:  nil,
	}
}

// insert a node recursive, we will update the height and rotate every node
// the path we go to the right postion of new node
func (t *avlTree) insertHelper(node *nodeAVL, val int) *nodeAVL {
	// break condition, node == nil means find the postion of new node
	if node == nil {
		return NewAvlNode(val)
	}

	// like search process, find the position of new node
	if val < node.Val {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		return node
	}

	node.updateHeight()
	node = node.rotate()
	return node
}

func (t *avlTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

func (t *avlTree) removeHelper(node *nodeAVL, val int) *nodeAVL {
	if node == nil {
		return nil
	}

	if val < node.Val {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				// The node we want to delete has no child, so just delete it
				return nil
			} else {
				// has 1 child, replace the node with child
				node = child
			}
		} else {
			// has 2 child repalce The node with next node in InOrder
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}

			// delete the temp node
			node.Right = t.removeHelper(node.Right, temp.Val)
			// replace node with temp
			node.Val = temp.Val
		}

	}

	node.updateHeight()
	node.rotate()
	return node
}

// It is not a good practice, this just as same as BFS with normal tree
// I just change the object of this method, actual we need to abstract an
// interface to have The method BFS() []any, TBD
func (aT *avlTree) avlBFS() []any {
	// list.New() return a boubly link list,
	// a little waste of storage
	queue := list.New()
	queue.PushBack(aT.root)

	res := make([]any, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*nodeAVL)

		res = append(res, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return res
}
