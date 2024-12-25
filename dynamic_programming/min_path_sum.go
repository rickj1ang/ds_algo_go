package dynamicprogramming

import "math"

func minPathSumDFS(grid [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}

	if i < 0 || j < 0 {
		return math.MaxInt
	}

	up := minPathSumDFS(grid, i-1, j)
	left := minPathSumDFS(grid, i, j-1)

	return min(up, left) + grid[i][j]
}

func minPathSumMem(grid, mem [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}

	if i < 0 || j < 0 {
		return math.MaxInt
	}

	if mem[i][j] != -1 {
		return mem[i][j]
	}

	up := minPathSumMem(grid, mem, i-1, j)
	left := minPathSumMem(grid, mem, i, j-1)

	mem[i][j] = min(up, left) + grid[i][j]

	return mem[i][j]
}

func minPathSumDP(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]

	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}

func minPathSumDPComp(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	dp := make([]int, m)
	dp[0] = grid[0][0]
	for j := 1; j < m; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < m; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[m-1]
}
