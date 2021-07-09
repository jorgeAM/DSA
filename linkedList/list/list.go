package list

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *LinkedList) Print() {
	current := l.Head

	for current.Next != nil {
		fmt.Println(current.Value)
		current = current.Next
	}

	fmt.Println(current.Value)
}

func (l *LinkedList) Push(node *Node) {
	if l.Head == nil {
		l.Head = node
	}

	if l.Tail != nil {
		l.Tail.Next = node
	}

	l.Tail = node
	l.Length++
}

// Pop method remove node at the end
func (l *LinkedList) Pop() *Node {
	if l.Length == 0 {
		return nil
	}

	current := l.Head

	if l.Length == 1 {
		l.Head = nil
		l.Tail = nil
		l.Length = 0

		return current
	}

	pre := new(Node)

	for current.Next != nil {
		pre = current
		current = current.Next
	}

	pre.Next = nil
	l.Tail = pre
	l.Length--

	return current
}

// Shift method remove node at the beginning
func (l *LinkedList) Shift() *Node {
	if l.Head == nil {
		return nil
	}

	head := l.Head

	if head.Next == nil {
		l.Head = nil
		l.Tail = nil
		l.Length = 0

		return head
	}

	l.Head = head.Next
	l.Length--

	return head
}

// UnShift method add a node at the beginning
func (l *LinkedList) UnShift(node *Node) {
	head := l.Head
	l.Head = node
	node.Next = head

	if l.Length == 0 {
		l.Tail = node
	}

	l.Length++
}

func (l *LinkedList) Get(index int) (*Node, error) {
	size := l.Length

	if index < 0 || index >= size {
		return nil, errors.New("wrong index")
	}

	node := l.Head
	var counter int

	for counter != index {
		node = node.Next
		counter++
	}

	return node, nil
}

func (l *LinkedList) Set(value, index int) error {
	node, err := l.Get(index)

	if err != nil {
		return err
	}

	node.Value = value

	return nil
}
