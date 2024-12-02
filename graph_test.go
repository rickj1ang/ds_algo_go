package main

import "testing"

// U better test a single function one time, but because all the function we impl
// are perfect work together, so maybe we will test more then 1 function once.

var v = []int{5, 2, 5, 7, 7, 4}

// the relationship like a linklist so make no sense to use graph
// but, for keep simple just do it
// 5--2--5--7--7--4
var edges = [][]int{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{4, 5},
}

func TestNew(t *testing.T) {
	g := newAdjMatGraph(v, edges)
	for i := range v {
		if v[i] != g.vertices[i] {
			t.Errorf("Nodes' Val:: expect: %d; actual: %d", v[i], g.vertices[i])
		}
	}
	// actually we should loop the g.adjMat to make sure the val is 1 or 0 in the correct location
	// but we can use g.print() to visualize the matrix check it with human eyes, it is more
	// clear for a leaning purpose
	g.print()
}

func TestSize(t *testing.T) {
	g := newAdjMatGraph(v, edges)
	if g.size() != len(v) {
		t.Errorf("Size:: expect: %d; actual: %d", len(v), g.size())
	}
}

func TestAddV(t *testing.T) {
	g := newAdjMatGraph(v, edges)
	g.addVertex(12)
	if g.vertices[g.size()-1] != 12 {
		t.Error("Add vertex failed")
	}
	g.print()
}

func TestRemoveV(t *testing.T) {
	g := newAdjMatGraph(v, edges)
	g.removeVertex(0)
	if g.vertices[0] != 2 {
		t.Errorf("val: Remove vertex failed expect: %d; actual %d", 2, g.vertices[0])
	}
	if g.size() != len(v)-1 {
		t.Error("size: Remove vertex failed")
	}
	g.print()
}

func TestAddE(t *testing.T) {
	// v changed because the test befor, but has no influence on the furthur test
	t.Log(v)
	g := newAdjMatGraph(v, edges)
	g.addEdge(0, 2)
	if g.adjMat[0][2] != 1 || g.adjMat[0][2] != g.adjMat[2][0] {
		t.Error("edges error")
	}
	g.print()
}

func TestRemoveE(t *testing.T) {
	g := newAdjMatGraph(v, edges)
	g.print()
	g.removeEdge(0, 1)
	if g.adjMat[0][1] != 0 || g.adjMat[0][1] != g.adjMat[1][0] {
		t.Error("edges error")
	}
	g.print()
}
