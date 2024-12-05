package graphTraverse

import adjListGraph "github.com/rickj1ang/ds_algo_go/adj_list_graph"

// for the clear of test I seperate the defination of graph and it's BFS algo,
// code may be a little verbose but can be accept
func graphBFS(g *adjListGraph.AdjListGraph, stratVet adjListGraph.Vertex) []adjListGraph.Vertex {
	// the result we going to return
	res := make([]adjListGraph.Vertex, 0)

	// record the Vertex we visited and struct{} mean nothing, you can use it in channel
	// as signal, it even light weight than bool
	visited := make(map[adjListGraph.Vertex]struct{})
	visited[stratVet] = struct{}{}

	queue := make([]adjListGraph.Vertex, 0)
	queue = append(queue, stratVet)

	for len(queue) > 0 {
		//Pop from queue
		vet := queue[0]
		queue = queue[1:]
		res = append(res, vet)

		for _, adjVet := range g.AdjList[vet] {
			_, isExist := visited[adjVet]
			if !isExist {
				queue = append(queue, adjVet)
				visited[adjVet] = struct{}{}
			}
		}
	}

	return res
}
