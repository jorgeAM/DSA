package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func contains(vertices []*Vertex, key int) bool {
	for _, v := range vertices {
		if v.key == key {
			return true
		}
	}

	return false
}

func (g *Graph) AddVertex(key int) {
	if contains(g.vertices, key) {
		return
	}

	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) GetVertex(key int) *Vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}

	return nil
}

func (g *Graph) AddEdge(from, to int) error {
	fromVertex, toVertex := g.GetVertex(from), g.GetVertex(to)

	var err error

	if fromVertex == nil || toVertex == nil {
		err = fmt.Errorf("Invalid edge (%v -> %v)", from, to)

		return err
	}

	if contains(fromVertex.adjacent, to) {
		err = fmt.Errorf("Existing edge (%v -> %v)", from, to)

		return err
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

	return nil
}

func main() {
	graph := new(Graph)

	vertices := []int{1, 2, 3, 4, 5}

	for _, v := range vertices {
		graph.AddVertex(v)
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 5)
	graph.AddEdge(2, 5)
	graph.AddEdge(5, 1)
	graph.AddEdge(5, 2)
	graph.AddEdge(4, 1)

	fmt.Println(graph)

}
