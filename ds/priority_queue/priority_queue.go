package priority_queue

type Node struct {
	Value    string
	Priority int
}

type PriorityQueue struct {
	Nodes []Node
}

func (q *PriorityQueue) Enqueue(value string, priority int) {
	node := Node{Value: value, Priority: priority}

	q.Nodes = append(q.Nodes, node)
	q.bubbleUp(len(q.Nodes) - 1)
}

func (q *PriorityQueue) Dequeue() Node {
	if len(q.Nodes) == 0 {
		return Node{}
	}

	lastIdx := len(q.Nodes) - 1
	q.swap(0, lastIdx)
	last := q.Nodes[lastIdx]
	q.Nodes = q.Nodes[:lastIdx]
	q.bubbleDown(0)

	return last
}

func (q *PriorityQueue) bubbleDown(index int) {
	l, r := left(index), right(index)
	var childToCompare int

	for l < len(q.Nodes) || r < len(q.Nodes) {
		if l == len(q.Nodes)-1 {
			childToCompare = l
		} else if q.Nodes[l].Priority > q.Nodes[r].Priority {
			childToCompare = r
		} else {
			childToCompare = l
		}

		if q.Nodes[index].Priority > q.Nodes[childToCompare].Priority {
			q.swap(index, childToCompare)

			index = childToCompare
		} else {
			return
		}
	}
}

func (q *PriorityQueue) bubbleUp(index int) {
	for q.Nodes[parent(index)].Priority > q.Nodes[index].Priority {
		q.swap(index, parent(index))

		index = parent(index)
	}
}

func (q *PriorityQueue) swap(a, b int) {
	q.Nodes[a], q.Nodes[b] = q.Nodes[b], q.Nodes[a]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (2 * index) + 1
}

func right(index int) int {
	return (2 * index) + 2
}
