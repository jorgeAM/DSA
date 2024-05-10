package main

import (
	"fmt"

	"github.com/jorgeAM/DSA/ds/graph"
)

func main() {
	graph := new(graph.Graph)

	vertices := []int{1, 2, 3, 4, 5}

	for _, v := range vertices {
		graph.AddVertex(v)
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(2, 5)

	fmt.Println(graph.BFS(1))
	fmt.Println(graph.DFSRecursive(1))
	fmt.Println(graph.DFSIterative(1))
}
