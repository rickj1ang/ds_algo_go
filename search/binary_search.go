package search

// U can do so many things about binary search especially,
// when the container allow duplicated values binary search
// can search the boundary of the same values and so on
// there so many details so I just can give U a little
// examples

// search target in a slice which has no duplicate element
// but U can determine there is or not a such element
func binarySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m - 1
		} else {
			return m
		}
	}
	return -1
}

func binarySearchInsertion(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else {
			// if = or > the j go left from right
			// with i together go to the first index of targets(if s)
			j = m - 1
		}
	}
	return i
}
