package adjListGraph

import (
	"fmt"
	"testing"
)

var edges = [][]Vertex{
	{1, 2},
	{2, 3},
	{1, 4},
	{3, 4},
}

// This function can test New, AddEdge, AddVertex and print
func TestNew(t *testing.T) {
	g := NewAdjListGraph(edges)
	if g.AdjList[1][0] != 2 {
		t.Errorf("Expect: 2; Actual: %d", g.AdjList[1][0])
	}
	if g.size() != 4 {
		t.Errorf("Expect: len = 4; Actual: %d", len(g.AdjList))
	}
	g.print()
}

func TestRemoveVertex(t *testing.T) {
	g := NewAdjListGraph(edges)
	fmt.Println("The Graph:")
	g.print()
	g.removeVertex(1)
	fmt.Println("After Remove Vertex:")
	g.print()
}

func TestDel(t *testing.T) {
	v := []Vertex{1, 2, 3, 4}
	fmt.Print(v)
	v = delVertex(v, 2)
	fmt.Print(v)
}

func TestRemoveEdge(t *testing.T) {
	g := NewAdjListGraph(edges)
	fmt.Println("The Graph:")
	g.print()
	g.removeEdge(1, 2)
	fmt.Println("After Remove:")
	g.print()
}
