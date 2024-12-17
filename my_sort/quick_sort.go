package mysort

// receive slice left and right boundary of the nums we want to sort
// return the index of base num AKA middle num
func partition(nums []int, left, right int) int {
	mid := findMid(nums, left, (left+right)/2, right)
	nums[left], nums[mid] = nums[mid], nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[right] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func quikSort(nums []int, left, right int) {
	// break recursion condition, left >= right we can not part nums anymore
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	// keep part the subSlice which already part by pivot
	quikSort(nums, left, pivot-1)
	quikSort(nums, pivot+1, right)
}

func findMid(nums []int, left, mid, right int) int {
	l, m, r := nums[left], nums[mid], nums[right]
	if (l <= m && m <= r) || (r <= m && m <= l) {
		return mid
	}
	if (m <= l && l <= r) || (r <= l && l <= m) {
		return left
	}
	return right
}

// Tip: Go do not have Tail recursive optimization
