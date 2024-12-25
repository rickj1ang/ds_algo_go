package dynamicprogramming

func climbingStair(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func minCostClimbStair(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 2 {
		return cost[n]
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	a, b := cost[1], cost[2]
	for i := 3; i <= n; i++ {
		tmp := b
		b = min(a, tmp) + cost[i]
		a = tmp
	}

	return b
}
