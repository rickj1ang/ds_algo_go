package backtrack

type Records struct {
	cols, diags1, diags2 []bool
}

type Situation struct {
	col, diag1, diag2 int
}

func backtrackNQ(row, n int, state [][]string, res [][][]string, r Records) [][][]string {
	if isSolutionQ(row, n) {
		res = recordSolutionQ(state, res)
		return res
	}
	for col := 0; col < n; col++ {
		//compute diag1, diag2
		diag1, diag2 := computeDiags(row, col, n)

		situation := Situation{
			col:   col,
			diag1: diag1,
			diag2: diag2,
		}

		if r.isValidQ(situation) {
			//try
			state = r.makeChoiceQ(row, state, situation)
			//recursion
			res = backtrackNQ(row+1, n, state, res, r)
			//backtrack
			state = r.undoChoiceQ(row, state, situation)
		}
	}
	return res
}

func isSolutionQ(row, n int) bool {
	return row == n
}

func recordSolutionQ(state [][]string, res [][][]string) [][][]string {
	// This is deepCopy, U can not use append only
	// I used it before, it will the res will be modify
	// Thoughts: silce is a pointer to array, maybe 2d slice
	// is a pointer to 1d slice, 3d slice is pointers to 2d slices
	// So just append a 2dslice to 3dslice will append the pointer
	// But when you want to append a int to 1dSlice this situation will
	// not happend like the code is sub_sum.go
	copiedState := make([][]string, len(state))
	for i := range state {
		copiedState[i] = make([]string, len(state[i]))
		copy(copiedState[i], state[i])
	}
	res = append(res, copiedState)
	return res
}

func computeDiags(row, col, n int) (diag1, diag2 int) {
	diag1 = row - col + n - 1
	diag2 = row + col
	return diag1, diag2
}

func (r Records) isValidQ(s Situation) bool {
	return !r.cols[s.col] && !r.diags1[s.diag1] && !r.diags2[s.diag2]
}

func (r *Records) makeChoiceQ(row int, state [][]string, s Situation) [][]string {
	state[row][s.col] = "Q"
	r.cols[s.col], r.diags1[s.diag1], r.diags2[s.diag2] = true, true, true
	return state
}

func (r *Records) undoChoiceQ(row int, state [][]string, s Situation) [][]string {
	state[row][s.col] = "#"
	r.cols[s.col] = false
	r.diags2[s.diag2] = false
	r.diags1[s.diag1] = false
	return state
}

func nQueens(n int) [][][]string {
	state := make([][]string, n)
	for i := range state {
		row := make([]string, n)
		for i := range row {
			row[i] = "#"
		}
		state[i] = row
	}
	cols := make([]bool, n)
	diags1 := make([]bool, 2*n-1)
	diags2 := make([]bool, 2*n-1)
	res := make([][][]string, 0)
	records := Records{
		cols:   cols,
		diags1: diags1,
		diags2: diags2,
	}
	res = backtrackNQ(0, n, state, res, records)
	return res
}
