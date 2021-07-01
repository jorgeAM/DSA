package main

import "fmt"

type MaxHeap struct {
	slice []int
}

func (m *MaxHeap) Insert(key int) {
	m.slice = append(m.slice, key)
	m.maxHeapifyUp(len(m.slice) - 1)
}

func (m *MaxHeap) Extract() int {
	extract := m.slice[0]
	l := len(m.slice) - 1

	if len(m.slice) == 0 {
		return -1
	}

	m.slice[0] = m.slice[l]
	m.slice = m.slice[:l]

	m.maxHeapifyDown(0)

	return extract
}

func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.slice[parent(index)] < m.slice[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(m.slice) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l < lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.slice[l] > m.slice[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.slice[index] < m.slice[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (index + 1) * 2
}

func right(index int) int {
	return (index + 1) * 2
}

func (m *MaxHeap) swap(a, b int) {
	m.slice[a], m.slice[b] = m.slice[b], m.slice[a]
}

func main() {
	heap := new(MaxHeap)
	fmt.Println(heap)

	buildHeap := []int{10, 20, 30, 5, 7, 11, 1, 50, 2}

	for _, v := range buildHeap {
		heap.Insert(v)
		fmt.Println(heap)
	}

	for i := 0; i < 5; i++ {
		heap.Extract()
		fmt.Println(heap)
	}
}
