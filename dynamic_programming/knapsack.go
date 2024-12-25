package dynamicprogramming

func knapsackDFS(wgt, val []int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}

	if wgt[i-1] > c {
		return knapsackDFS(wgt, val, i-1, c)
	}

	no := knapsackDFS(wgt, val, i-1, c)
	yes := knapsackDFS(wgt, val, i-1, c-wgt[i-1]) + val[i-1]

	return max(no, yes)
}

func knapsackMem(wgt, val []int, mem [][]int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}

	if mem[i][c] != -1 {
		return mem[i][c]
	}

	if wgt[i-1] > c {
		return knapsackMem(wgt, val, mem, i-1, c)
	}

	no := knapsackMem(wgt, val, mem, i-1, c)
	yes := knapsackMem(wgt, val, mem, i-1, c-wgt[i-1]) + val[i-1]

	mem[i][c] = max(no, yes)
	return mem[i][c]
}

func knapsackDP(wgt, val []int, cap int) int {
	n := len(wgt)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}

	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if wgt[i-1] > cap {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = max(dp[i-1][c], dp[i-1][c-wgt[i-1]]+val[i-1])
			}
		}
	}
	return dp[n][cap]
}

func knapsackDPComp(wgt, val []int, cap int) int {
	n := len(wgt)
	dp := make([]int, cap+1)

	for i := 1; i < n; i++ {
		for c := cap; c > 0; c-- {
			if wgt[i-1] < c {
				dp[c] = max(dp[c], dp[c-wgt[i-1]]+val[i-1])
			}
		}
	}
	return dp[cap]
}
