package mysort

import "math"

// exp = 10^(k-1)
// k = whcih digit we want to sort
func digit(num, exp int) int {
	return (num / exp) % 10
}

// sort num in nums by Kth digit
func countingSortDigit(nums []int, exp int) {
	counter := make([]int, 10)
	n := len(nums)

	// count
	for i := 0; i < n; i++ {
		d := digit(nums[i], exp)
		counter[d]++
	}

	// prefix
	for i := 1; i < 10; i++ {
		counter[i] += counter[i-1]
	}

	// iterate from tail
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d := digit(nums[i], exp)
		j := counter[d] - 1
		res[j] = nums[i]
		counter[d]--
	}

	for i := 0; i < n; i++ {
		nums[i] = res[i]
	}
}

func radixSort(nums []int) {
	max := math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	for exp := 1; max >= exp; exp *= 10 {
		countingSortDigit(nums, exp)
	}
}
