package dynamicprogramming

import "testing"

var coins = []int{
	1, 2, 5,
}

func TestCoinChangeDP(t *testing.T) {
	res := coinChangeDP(coins, 4)
	t.Log(res)

}

func TestCoinChangeComp(t *testing.T) {
	res := coinChangeDPComp(coins, 11)
	t.Log(res)
}

func TestCoinChangeII(t *testing.T) {
	res := coinChangeDPII(coins, 4)
	t.Log(res)
}
