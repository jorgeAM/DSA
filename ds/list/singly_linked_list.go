package list

type Node struct {
	Value int
	Next  *Node
}

type SinglyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *SinglyLinkedList) Push(value int) {
	node := &Node{Value: value}
	if l.Head == nil {
		l.Head = node
	}

	if l.Tail != nil {
		l.Tail.Next = node
	}

	l.Tail = node
	l.Length++
}

func (l *SinglyLinkedList) Pop() *Node {
	if l.Length == 0 {
		return nil
	}

	curr := l.Head

	if l.Length == 1 {
		l.Head = nil
		l.Tail = nil
		l.Length = 0

		return curr
	}

	prev := new(Node)

	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil
	l.Tail = prev
	l.Length--

	return curr
}

func (l *SinglyLinkedList) Shift() *Node {
	if l.Head == nil {
		return nil
	}

	head := l.Head
	l.Head = head.Next
	l.Length--

	if l.Length == 0 {
		l.Tail = nil
	}

	return head
}

func (l *SinglyLinkedList) Unshift(value int) {
	node := &Node{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node

		return
	}

	node.Next = l.Head
	l.Head = node
	l.Length++
}

func (l *SinglyLinkedList) Get(index int) *Node {
	if index >= l.Length || index < 0 {
		return nil
	}

	curr := l.Head

	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	return curr
}

func (l *SinglyLinkedList) Set(value, index int) bool {
	node := l.Get(index)

	if node == nil {
		return false
	}

	node.Value = value

	return true
}

func (l *SinglyLinkedList) Insert(value, index int) bool {
	if index > l.Length || index < 0 {
		return false
	}

	if index == 0 {
		l.Unshift(value)

		return true
	}

	if index == l.Length {
		l.Push(value)

		return true
	}

	newNode := &Node{Value: value}
	prev := l.Get(index - 1)
	next := prev.Next

	prev.Next = newNode
	newNode.Next = next
	l.Length++

	return true
}

func (l *SinglyLinkedList) Remove(index int) *Node {
	if index >= l.Length || index < 0 {
		return nil
	}

	if index == 0 {
		return l.Shift()
	}

	if index == l.Length-1 {
		return l.Pop()
	}

	prev := l.Get(index - 1)
	removedNode := prev.Next
	prev.Next = removedNode.Next
	removedNode.Next = nil

	l.Length--

	return removedNode
}

func (l *SinglyLinkedList) Reverse() {
	curr := l.Head
	l.Head = l.Tail
	l.Tail = curr

	var prev *Node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}
