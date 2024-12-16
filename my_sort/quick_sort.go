package mysort

// receive slice left and right boundary of the nums we want to sort
// return the index of base num AKA middle num
func partion(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] <= nums[left] {
			j--
		}
		for i < j && nums[i] >= nums[right] {
			i--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}
