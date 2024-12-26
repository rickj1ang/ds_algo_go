package greedy

func coinChangeGreedy(coins []int, amt int) int {
	// coins is ordered
	i := len(coins) - 1
	count := 0
	for amt > 0 {
		for i > 0 && coins[i] > amt {
			i--
		}
		amt -= coins[i]
		count++
	}
	if amt != 0 {
		return -1
	}
	return count
}
