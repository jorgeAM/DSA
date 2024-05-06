package main

import (
	"fmt"

	"github.com/jorgeAM/DSA/ds/priority_queue"
)

func main() {
	pQueue := priority_queue.PriorityQueue{}
	pQueue.Enqueue("gripe", 4)
	pQueue.Enqueue("fractura", 3)
	pQueue.Enqueue("dolor de barriga", 4)
	pQueue.Enqueue("infarto", 1)
	fmt.Println(pQueue.Nodes)
	fmt.Println(pQueue.Dequeue())
	fmt.Println(pQueue.Nodes)
	fmt.Println(pQueue.Dequeue())
	fmt.Println(pQueue.Nodes)
	pQueue.Enqueue("sin cabeza", 0)
	fmt.Println(pQueue.Nodes)
}
