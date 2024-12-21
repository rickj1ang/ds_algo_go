package backtrack

import (
	"testing"
)

func TestSubSum(t *testing.T) {
	nums := []int{3, 4, 5}
	target := 9
	res := subsetSum(nums, target)
	for _, subRes := range res {
		t.Log(subRes)
	}
}
