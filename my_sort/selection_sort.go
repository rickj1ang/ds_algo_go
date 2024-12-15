package mysort

func selectionSort(nums []int) {
	// len of the slice
	n := len(nums) - 1
	for i := 0; i < n-1; i++ {
		// k record the min value
		k := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		// swap
		nums[i], nums[k] = nums[k], nums[i]
	}
}
