package main

import (
	"fmt"

	"github.com/jorgeAM/DSA/ds/binary_heap"
)

func main() {
	heap := binary_heap.MaxBinaryHeap{}
	heap.Insert(41)
	heap.Insert(39)
	heap.Insert(33)
	heap.Insert(18)
	heap.Insert(27)
	heap.Insert(12)

	fmt.Println(heap.Values)
	heap.Insert(55)
	fmt.Println(heap.Values)
}
