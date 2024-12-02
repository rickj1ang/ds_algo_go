package main

import "testing"

func TestAdd(t *testing.T) {
	x, y := 1, 2
	expect := 3
	actual := add(x, y)
	if actual != expect {
		t.Errorf("Expect: %d, Get: %d\n", expect, actual)
	} else {
		t.Log("Add success!\n")
	}
}
