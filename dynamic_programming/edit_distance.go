package dynamicprogramming

func editDistance(s, t string) int {
	n := len(s)
	m := len(t)
	// init dp
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	//state transfer: first col & row
	for i := 1; i < n; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// fill the dp table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}

func editDistanceComp(s, t string) int {
	n := len(s)
	m := len(t)
	// init dp
	dp := make([]int, m+1)
	// state transfer: first row
	for i := 1; i <= m; i++ {
		dp[i] = i
	}

	for i := 1; i <= n; i++ {
		// first col
		leftUp := dp[0]
		dp[0] = i
		// other col
		for j := 1; j <= m; j++ {
			temp := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = leftUp
			} else {
				dp[j] = min(dp[j-1], dp[j], leftUp) + 1
			}
			leftUp = temp
		}
	}
	return dp[m]
}
