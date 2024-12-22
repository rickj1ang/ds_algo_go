package backtrack

import "testing"

func TestQueens(t *testing.T) {
	n := 4
	res := nQueens(n)
	t.Log(res)
}

func TestRecordSolutionQ(t *testing.T) {
	resI := make([][][]string, 0)
	state := make([][]string, 0)
	state = append(state, []string{"A", "B"})
	resI = recordSolutionQ(state, resI)
	t.Log(resI)
}
