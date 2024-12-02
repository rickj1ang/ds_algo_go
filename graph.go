package main

import "fmt"

// adjacency matrix graph without direction
type adjMatGraph struct {
	vertices []int   //node value in the slice
	adjMat   [][]int //the matrix
}

// input value of nodes and edges return a adjMatGraph struct
func newAdjMatGraph(vertices []int, edges [][]int) *adjMatGraph {
	// init the matrix
	nodeNum := len(vertices)
	adjMat := make([][]int, nodeNum)
	for i := range adjMat {
		adjMat[i] = make([]int, nodeNum)
	}
	// init the struct
	g := &adjMatGraph{
		vertices: vertices,
		adjMat:   adjMat,
	}

	for i := range edges {
		g.addEdge(edges[i][0], edges[i][1])
	}
	return g

}

// return the num of nodes
func (g *adjMatGraph) size() int {
	return len(g.vertices)
}

// edges [][]int store the relationship of nodes' index convert the relation into
// g.adjMat
func (g *adjMatGraph) addEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		_ = fmt.Errorf("Index out of Bounds")
	}
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

func (g *adjMatGraph) removeEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		_ = fmt.Errorf("Index out of Bounds")
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

// add node: val as input append to g.vertices and add a newRow a newCol in g.adjMat
func (g *adjMatGraph) addVertex(val int) {
	n := g.size()
	// new nod
	g.vertices = append(g.vertices, val)
	// new row
	newRow := make([]int, n)
	g.adjMat = append(g.adjMat, newRow)

	// new col
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}

func (g *adjMatGraph) removeVertex(index int) {
	if index > g.size() {
		return
	}
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)

	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

func (g *adjMatGraph) print() {
	fmt.Printf("\tNodes:\n")
	fmt.Printf("\t\t%v\n", g.vertices)
	fmt.Printf("\tadjMat:\n")
	for i := range g.adjMat {
		fmt.Printf("\t\t%v\n", g.adjMat[i])
	}
}
