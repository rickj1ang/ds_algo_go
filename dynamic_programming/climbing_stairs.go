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
