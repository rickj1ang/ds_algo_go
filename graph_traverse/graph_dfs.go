package graphTraverse

import (
	adjListGraph "github.com/rickj1ang/ds_algo_go/adj_list_graph"
)

func bfs(g *adjListGraph.AdjListGraph, res *[]adjListGraph.Vertex, visited map[adjListGraph.Vertex]struct{}, vet adjListGraph.Vertex) {
	*res = append(*res, vet)
	visited[vet] = struct{}{}

	for _, adjVet := range g.AdjList[vet] {
		_, isExist := visited[adjVet]
		if !isExist {
			bfs(g, res, visited, adjVet)
		}
	}
}

func graphDFS(g *adjListGraph.AdjListGraph, startVet adjListGraph.Vertex) []adjListGraph.Vertex {
	res := make([]adjListGraph.Vertex, 0)
	visited := make(map[adjListGraph.Vertex]struct{})

	bfs(g, &res, visited, startVet)

	return res
}
