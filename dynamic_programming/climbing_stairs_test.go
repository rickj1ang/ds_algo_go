package dynamicprogramming

import "testing"

func TestClimbingStair(t *testing.T) {
	if climbingStair(3)+climbingStair(4) != climbingStair(5) {
		t.Log("error")
	}
	t.Log("successful climbed the stairs")

}
