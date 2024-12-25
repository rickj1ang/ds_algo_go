package dynamicprogramming

import (
	"testing"
	"time"
)

var grid = [][]int{
	{1, 3, 1, 5},
	{2, 2, 4, 2},
	{5, 3, 2, 1},
	{4, 3, 5, 2},
}

func TestDFS(t *testing.T) {
	tI := time.Now()
	res := minPathSumDFS(grid, 3, 3)
	t.Log(res)
	t.Log(time.Now().Sub(tI))
}

func TestMem(t *testing.T) {
	tII := time.Now()
	mem := make([][]int, len(grid))
	for i := range mem {
		mem[i] = make([]int, len(grid[i]))
	}
	for i := 0; i < len(mem); i++ {
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = -1
		}
	}

	res := minPathSumMem(grid, mem, 3, 3)
	t.Log(res)
	t.Log(time.Now().Sub(tII))

}

func TestDP(t *testing.T) {
	tIII := time.Now()
	res := minPathSumDP(grid)
	t.Log(res)
	t.Log(time.Now().Sub(tIII))
}
func TestDPComp(t *testing.T) {
	tIV := time.Now()
	res := minPathSumDPComp(grid)
	t.Log(res)
	t.Log(time.Now().Sub(tIV))
}
