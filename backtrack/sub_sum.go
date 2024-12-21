package backtrack

import "sort"

func backtrackSubSum(start, target int, state, choices []int, res [][]int) [][]int {
	//record condition
	if isSolutionS(target) {
		res = recordSolutionS(state, res)
		return res
	}
	for i := start; i < len(choices); i++ {
		choice := choices[i]
		if notValidI(target, choice) {
			break
		}
		if notValidII(i, start, choices) {
			continue
		}
		// try
		state = makeChoiceS(state, choice)
		//recursion
		res = backtrackSubSum(i, target-choice, state, choices, res)
		//backtrack
		state = undoChoiceS(state)
	}
	return res
}

func isSolutionS(target int) bool {
	return 0 == target
}

func recordSolutionS(state []int, res [][]int) [][]int {
	res = append(res, append([]int{}, state...))
	return res
}

func notValidI(target, choice int) bool {
	return target-choice < 0
}

func notValidII(i, start int, choices []int) bool {
	return i > start && choices[i] == choices[i-1]

}

func makeChoiceS(state []int, choice int) []int {
	state = append(state, choice)
	return state
}

func undoChoiceS(state []int) []int {
	state = state[:len(state)-1]
	return state
}

func subsetSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	state := make([]int, 0)
	start := 0
	res := make([][]int, 0)
	res = backtrackSubSum(start, target, state, nums, res)
	return res
}
