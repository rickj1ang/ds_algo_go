package mysort

func countingSort(nums []int) {
	// find the maxValue of nums
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// count the num
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num] += 1
	}
	// change counter[num] to sum of prefix
	for i := 0; i < m; i++ {
		counter[i+1] += counter[i]
	}
	// iterate the nums from tail & declare the res
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		res[counter[num]-1] = num
		counter[num]--
	}
	copy(nums, res)
}
