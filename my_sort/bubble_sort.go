package mysort

func bubbleSort(nums []int) {
	n := len(nums)
	// outter loop from right to left
	for i := n - 1; i > 0; i-- {
		// inner loop from left to right boundary is i
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
