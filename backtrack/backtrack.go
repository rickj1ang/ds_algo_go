package backtrack

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

// this is a template of backtrack quiz
func backtrack(state []*TreeNode, choices []*TreeNode, res [][]*TreeNode) {
	if isSolution(state) {
		recordSolution(state, res)
	}
	for _, choice := range choices {
		if isValid(choice) {
			makeChoice(state, choice)
			temp := make([]*TreeNode, 0)
			temp = append(temp, choice.left, choice.right)
			backtrack(state, temp, res)
			undoChoice(state)
		}
	}
}

// find the node which value is 9
func isSolution(state []*TreeNode) bool {
	return len(state) != 0 && state[len(state)-1].val == 9
}

func recordSolution(state []*TreeNode, res [][]*TreeNode) {
	res = append(res, append([]*TreeNode{}, state...))
}

func isValid(choice *TreeNode) bool {
	return choice != nil && choice.val != 6
}

func makeChoice(state []*TreeNode, choice *TreeNode) {
	state = append(state, choice)
}

func undoChoice(state []*TreeNode) {
	state = state[:len(state)-1]
}
