package adjListGraph

import (
	"fmt"
	"strconv"
)

// adjacency list graph impl
// this is just a mimic of Vertex in real-world U need to have a more complex struct
// or just let package and database to worry about algo:)
type Vertex int

// looks like this :
//
//	Vertex-Vertex-Vertex
//	Vertex-Vertex-Vertex
//	Vertex-Vertex-Vertex
type AdjListGraph struct {
	AdjList map[Vertex][]Vertex
}

func newAdjListGraph(edges [][]Vertex) *AdjListGraph {
	g := &AdjListGraph{
		AdjList: make(map[Vertex][]Vertex),
	}

	for _, edge := range edges {
		g.addVertex(edge[0])
		g.addVertex(edge[1])
		g.addEdge(edge[0], edge[1])
	}
	return g
}

func (g *AdjListGraph) size() int {
	return len(g.AdjList)
}

func (g *AdjListGraph) addVertex(vet Vertex) {
	// if there exist a same vet, just do nothing and return
	_, ok := g.AdjList[vet]
	if ok {
		return
	}
	g.AdjList[vet] = make([]Vertex, 0)
}

func (g *AdjListGraph) removeVertex(vet Vertex) {
	// delete the line
	delete(g.AdjList, vet)
	// delete the relationships
	for k := range g.AdjList {
		g.AdjList[k] = delVertex(g.AdjList[k], vet)
	}
}

func (g *AdjListGraph) addEdge(vet1, vet2 Vertex) {
	_, ok1 := g.AdjList[vet1]
	_, ok2 := g.AdjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("Edge want to add is illegal")
	}
	g.AdjList[vet1] = append(g.AdjList[vet1], vet2)
	g.AdjList[vet2] = append(g.AdjList[vet2], vet1)
}

func (g *AdjListGraph) removeEdge(vet1, vet2 Vertex) {
	_, ok1 := g.AdjList[vet1]
	_, ok2 := g.AdjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("Edge want to remove is illegal")
	}
	g.AdjList[vet1] = delVertex(g.AdjList[vet1], vet2)
	g.AdjList[vet2] = delVertex(g.AdjList[vet2], vet1)
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

func (g *AdjListGraph) print() {
	for k, v := range g.AdjList {
		fmt.Printf(strconv.Itoa(int(k)) + ": ")
		for _, vet := range v {
			fmt.Printf("--" + strconv.Itoa(int(vet)))
		}
		fmt.Println()
	}
}
