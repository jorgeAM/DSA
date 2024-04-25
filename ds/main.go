package main

import (
	"fmt"

	"github.com/jorgeAM/DSA/ds/list"
)

func main() {
	list := list.SinglyLinkedList{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)

	list.Reverse()
	fmt.Println(list)
}
