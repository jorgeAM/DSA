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

func (h *MaxBinaryHeap) bubbleDown(index int) {
	r, l := right(index), left(index)
	var childToCompareIdX int

	for r < len(h.Values) || l < len(h.Values) {
		if l == len(h.Values)-1 {
			childToCompareIdX = l
		} else if h.Values[l] > h.Values[r] {
			childToCompareIdX = l
		} else {
			childToCompareIdX = r
		}

		if h.Values[index] < h.Values[childToCompareIdX] {
			h.swap(index, childToCompareIdX)
			index = childToCompareIdX
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxBinaryHeap) ExtractMax() int {
	if len(h.Values) == 0 {
		return -1
	}

	lastIdX := len(h.Values) - 1
	h.swap(0, lastIdX)
	last := h.Values[lastIdX]

	h.Values = h.Values[:lastIdX]
	h.bubbleDown(0)

	return last
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
