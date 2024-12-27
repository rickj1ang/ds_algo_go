package greedy

import "testing"

func TestFracKnapsack(t *testing.T) {
	wgt := []int{10, 20, 30, 40, 50}
	val := []int{50, 120, 150, 210, 240}
	cap := 50
	res := fractionalKnapsack(wgt, val, cap)
	t.Log(res)
}
