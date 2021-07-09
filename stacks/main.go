package main

import (
	"fmt"

	"github.com/jorgeAM/structures/stacks/stack"
)

func main() {
	n1 := &stack.Node{Value: 1}
	n2 := &stack.Node{Value: 23}
	n3 := &stack.Node{Value: 43}

	l := stack.Stack{}

	l.Push(n1)
	l.Push(n2)
	l.Push(n3)

	fmt.Println(l)

	l.Pop()
	fmt.Println(l)
}
