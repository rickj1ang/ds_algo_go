package dynamicprogramming

import "testing"

var s = "bag"
var b = "pack"

func TestEditDistance(t *testing.T) {
	res := editDistance(s, b)
	t.Log(res)
}

func TestEditDistanceComp(t *testing.T) {
	res := editDistanceComp(s, b)
	t.Log(res)
}
