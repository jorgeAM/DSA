package graph

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

	if fromVertex == nil || toVertex == nil {
		return fmt.Errorf("invalid edge (%v -> %v)", from, to)
	}

	if contains(fromVertex.adjacent, to) {
		return fmt.Errorf("existing edge (%v -> %v)", from, to)
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

	return nil
}

func (g *Graph) BFS(startKey int) []int {
	visited, result, queue := make(map[int]struct{}), make([]int, 0), make([]*Vertex, 0)

	vertex := g.GetVertex(startKey)
	if vertex == nil {
		return result
	}

	queue = append(queue, vertex)
	visited[vertex.key] = struct{}{}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		result = append(result, first.key)

		for _, neighbor := range first.adjacent {
			_, ok := visited[neighbor.key]
			if !ok {
				visited[neighbor.key] = struct{}{}
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func (g *Graph) DFSIterative(starKey int) []int {
	visited, result := make(map[int]struct{}), make([]int, 0)

	vertex := g.GetVertex(starKey)
	if vertex == nil {
		return result
	}

	stack := make([]*Vertex, 0)
	stack = append(stack, vertex)

	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		last := stack[lastIdx]
		stack = stack[:lastIdx]

		_, ok := visited[last.key]
		if !ok {
			visited[last.key] = struct{}{}
			result = append(result, last.key)

			for _, neighbor := range last.adjacent {
				_, ok := visited[neighbor.key]
				if !ok {
					stack = append(stack, neighbor)
				}
			}
		}
	}

	return result
}

func (g *Graph) DFSRecursive(starKey int) []int {
	visited, result := make(map[int]struct{}), make([]int, 0)

	vertex := g.GetVertex(starKey)
	if vertex == nil {
		return result
	}

	DFSRecursiveHelper(vertex, &visited, &result)

	return result
}

func DFSRecursiveHelper(vertex *Vertex, visited *map[int]struct{}, result *[]int) {
	v := *visited
	v[vertex.key] = struct{}{}
	*result = append(*result, vertex.key)

	for _, neighbor := range vertex.adjacent {
		_, ok := v[neighbor.key]
		if !ok {
			DFSRecursiveHelper(neighbor, visited, result)
		}
	}
}
