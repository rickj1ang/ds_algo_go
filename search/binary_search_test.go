package search

import "testing"

func TestBinarySearch(t *testing.T) {
	// binary search must use on a sorted slice
	nums := []int{5, 19, 23, 53, 78, 123, 245, 432, 633}
	nilTarget := 60
	goodTarget := 53

	nilRes := binarySearch(nums, nilTarget)
	goodRes := binarySearch(nums, goodTarget)

	if nilRes != -1 {
		t.Errorf("Expect: %d; ACtual: %d\n", -1, nilRes)
	}
	if nums[goodRes] != goodTarget {
		t.Errorf("Expect: %d; ACtual: %d\n", goodTarget, nums[goodRes])
	}
}
