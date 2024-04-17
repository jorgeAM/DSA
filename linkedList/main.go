package main

import (
	"github.com/jorgeAM/DSA/linkedList/list"
)

func main() {
	l := list.LinkedList{}

	n1 := &list.Node{Value: 43}
	n2 := &list.Node{Value: 30}
	n3 := &list.Node{Value: 90}
	n4 := &list.Node{Value: 10}
	n5 := &list.Node{Value: 78}

	l.Push(n1)
	l.Push(n2)
	l.Push(n3)
	l.Push(n4)
	l.UnShift(n5)

	l.Print()

	l.Set(100, 3)
	l.Print()
}
