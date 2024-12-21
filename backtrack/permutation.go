package backtrack

func permutation(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	res, _, _ = backtrackPermutation(state, nums, selected, res)
	return res

}

// rather than input slice of pointers Golang recommand the usage as follow
// but this codes need to be test I will do it later
// if U decalre this function like func(*[]int, *[][]int) it will work well
// like in C, but the slice it self is already pointer to an array
func backtrackPermutation(state []int, choices []int, selected []bool, res [][]int) ([][]int, []int, []bool) {
	if isSolutionP(state, choices) {
		res = recordSolutionP(state, res)
	}
	for idx, choice := range choices {
		if isValidP(selected, idx) {
			selected, state = makeChoiceP(idx, selected, state, choice)
			res, state, selected = backtrackPermutation(state, choices, selected, res)
			state, selected = undoChoiceP(idx, state, selected)
		}
	}
	return res, state, selected
}

// when len of state equal to len of choices find a soliution
func isSolutionP(state []int, choices []int) bool {
	return len(state) == len(choices)
}

func recordSolutionP(state []int, res [][]int) [][]int {
	res = append(res, append([]int{}, state...))
	return res
}

func isValidP(selected []bool, idx int) bool {
	return !selected[idx]
}

func makeChoiceP(idx int, selected []bool, state []int, choice int) ([]bool, []int) {
	selected[idx] = true
	state = append(state, choice)
	return selected, state
}

func undoChoiceP(i int, state []int, selected []bool) ([]int, []bool) {
	selected[i] = false
	state = state[:len(state)-1]
	return state, selected
}
