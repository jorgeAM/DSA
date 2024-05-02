package binary_heap

type MaxBinaryHeap struct {
	Values []int
}

func (h *MaxBinaryHeap) swap(a, b int) {
	h.Values[a], h.Values[b] = h.Values[b], h.Values[a]
}

func (h *MaxBinaryHeap) bubbleUp(index int) {
	for h.Values[parent(index)] < h.Values[index] {
		h.swap(parent(index), index)

		index = parent(index)
	}
}

func (h *MaxBinaryHeap) Insert(value int) {
	h.Values = append(h.Values, value)

	h.bubbleUp(len(h.Values) - 1)
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
