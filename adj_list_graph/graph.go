package adjListGraph

import (
	"fmt"
	"strconv"
)

// adjacency list graph impl
type Vertex int

// looks like this :
//
//	Vertex-Vertex-Vertex
//	Vertex-Vertex-Vertex
//	Vertex-Vertex-Vertex
type adjListGraph struct {
	adjList map[Vertex][]Vertex
}

func newAdjListGraph(edges [][]Vertex) *adjListGraph {
	g := &adjListGraph{
		adjList: make(map[Vertex][]Vertex),
	}

	for _, edge := range edges {
		g.addVertex(edge[0])
		g.addVertex(edge[1])
		g.addEdge(edge[0], edge[1])
	}
	return g
}

func (g *adjListGraph) size() int {
	return len(g.adjList)
}

func (g *adjListGraph) addVertex(vet Vertex) {
	// if there exist a same vet, just do nothing and return
	_, ok := g.adjList[vet]
	if ok {
		return
	}
	g.adjList[vet] = make([]Vertex, 0)
}

func (g *adjListGraph) removeVertex(vet Vertex) {
	// delete the line
	delete(g.adjList, vet)
	// delete the relationships
	for k := range g.adjList {
		g.adjList[k] = delVertex(g.adjList[k], vet)
	}
}

func (g *adjListGraph) addEdge(vet1, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("Edge want to add is illegal")
	}
	g.adjList[vet1] = append(g.adjList[vet1], vet2)
	g.adjList[vet2] = append(g.adjList[vet2], vet1)
}

func (g *adjListGraph) removeEdge(vet1, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("Edge want to remove is illegal")
	}
	g.adjList[vet1] = delVertex(g.adjList[vet1], vet2)
	g.adjList[vet2] = delVertex(g.adjList[vet2], vet1)
}

// delete a element from slice
func delVertex(v []Vertex, val Vertex) []Vertex {
	index := -1
	for i, v := range v {
		if v == val {
			index = i
			break
		}
	}
	if index == -1 {
		return v
	}
	return append(v[:index], v[index+1:]...)
}

func (g *adjListGraph) print() {
	for k, v := range g.adjList {
		fmt.Printf(strconv.Itoa(int(k)) + ": ")
		for _, vet := range v {
			fmt.Printf("--" + strconv.Itoa(int(vet)))
		}
		fmt.Println()
	}
}
