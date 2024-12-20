package mysort

import "testing"

func TestBucketSort(t *testing.T) {
	// declare a slice
	nums := []int{12, 42, 5, 23, 50, 32, 78, 69, 89, 93}
	// call function
	bucketSort(nums)
	// check result
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			t.Error("sort wrong")
		}
	}
}
