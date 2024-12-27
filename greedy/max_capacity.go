package greedy

func maxCapacity(ht []int) int {
	i, j := 0, len(ht)-1
	res := 0

	for i < j {
		capacity := min(ht[i], ht[j]) * (j - i)
		res = max(res, capacity)
		if ht[i] < ht[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
