package binarytree

import "container/list"

func (node *TreeNode) BFS() []any {
	// list.New() return a boubly link list,
	// a little waste of storage
	queue := list.New()
	queue.PushBack(node)

	res := make([]any, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)

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
