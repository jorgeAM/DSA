package main

import (
	"fmt"

	"github.com/jorgeAM/structures/queues/queue"
)

func main() {
	q := queue.Queue{}

	n1 := &queue.Node{Value: 1}
	n2 := &queue.Node{Value: 2}
	n3 := &queue.Node{Value: 3}
	n4 := &queue.Node{Value: 4}

	q.Push(n1)
	q.Push(n2)
	q.Push(n3)
	q.Push(n4)

	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
}
