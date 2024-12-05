package graphTraverse

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

func TestDFS(t *testing.T) {
	g := adjListGraph.NewAdjListGraph(edges)
	var start adjListGraph.Vertex = 1
	res := graphDFS(g, start)
	t.Log("DFS:\n")
	t.Log(res)
}
