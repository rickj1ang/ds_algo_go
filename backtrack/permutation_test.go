package backtrack

import "testing"

func TestPermutation(t *testing.T) {
	nums := []int{1, 7, 5}
	res := permutation(nums)
	t.Log(res)
}