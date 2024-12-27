package greedy

import (
	"math"
	"testing"
)

func TestMPC(t *testing.T) {
	n := 15
	res := maxProductCutting(n)
	t.Log(res)
	t.Log(math.Pow(3, 5))
}
