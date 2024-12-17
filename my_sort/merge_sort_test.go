package mysort

import "testing"

// U can see This test is not same as other, because some bugs I need to
// handle with when write merge sort
func TestMergeSort(t *testing.T) {
	// declare a slice
	nums := []int{12, 42, 5, 23, 50, 22, 8, 19}
	//print the slice
	t.Log(nums)
	// call function
	mergeSort(nums, 0, 7)

	//print again
	t.Log(nums)
	// check result
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			t.Error("sort wrong")
		}
	}
}

func TestMerge(t *testing.T) {
	num := []int{1, 8, 4, 7}
	t.Log(num)
	merge(num, 0, 1, 3)
	t.Log(num)
}
