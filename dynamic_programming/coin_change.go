package dynamicprogramming

func coinChangeDP(coins []int, amt int) int {
	n := len(coins)
	maxVal := amt + 1
	//init dp slice
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amt+1)
	}
	//state transfer: first row
	for a := 1; a <= amt; a++ {
		dp[0][a] = maxVal
	}
	// first col, U do not need to do this, default int num is 0
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	//state transfer: other cols & rows
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				dp[i][a] = dp[i-1][a]
			} else {
				dp[i][a] = min(dp[i-1][a], dp[i][a-coins[i-1]]+1)
			}
		}
	}
	if dp[n][amt] != maxVal {
		return dp[n][amt]
	}
	return -1
}

func coinChangeDPComp(coins []int, amt int) int {
	n := len(coins)
	maxVal := amt + 1
	// init dp slice
	dp := make([]int, amt+1)
	//state transfer: first row
	for i := 1; i <= amt; i++ {
		dp[i] = maxVal
	}
	// state transfer: other rows
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				dp[a] = dp[a]
			} else {
				dp[a] = min(dp[a], dp[a-coins[i-1]]+1)
			}
		}
	}
	if dp[amt] != maxVal {
		return dp[amt]
	}

	return -1
}

func coinChangeDPII(coins []int, amt int) int {
	n := len(coins)
	// init dp slice
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amt+1)
	}
	//init first col
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	//init first row
	//again U do not need to do this, it just for clear
	for i := 1; i <= amt; i++ {
		dp[0][i] = 0
	}

	// fill the dp table
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				dp[i][a] = dp[i-1][a]
			} else {
				dp[i][a] = dp[i-1][a] + dp[i][a-coins[i-1]]
			}

		}
	}
	return dp[n][amt]
}
