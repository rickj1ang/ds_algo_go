package GBFS

import (
	"testing"

	adjListGraph "github.com/rickj1ang/ds_algo_go/adj_list_graph"
)

// make a graph use adj list

// looks like:
//
//	  1-2-3
//		| | |
//		4-5-6
//		| | |
//		7-8-9
var edges = [][]adjListGraph.Vertex{
	{1, 2},
	{2, 3},
	{1, 4},
	{4, 5},
	{5, 6},
	{2, 5},
	{3, 6},
	{4, 7},
	{7, 8},
	{5, 8},
	{8, 9},
	{6, 9},
}

func TestBFS(t *testing.T) {
	g := adjListGraph.NewAdjListGraph(edges)
	var start adjListGraph.Vertex = 1
	res := graphBFS(g, start)
	// please check the out put fit or not the BFS by your self.
	// on my devide it works well!! :)
	t.Log(res)

}
