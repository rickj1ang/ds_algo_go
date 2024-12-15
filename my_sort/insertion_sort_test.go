package mysort

import "testing"

func TestInsertionSort(t *testing.T) {
	// declare a slice
	nums := []int{12, 42, 5, 23, 50, 22, 8, 19}
	// call function
	insertionSort(nums)
	// check result
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			t.Error("sort wrong")
		}
	}
}
