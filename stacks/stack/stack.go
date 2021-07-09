package stack

type Stack struct {
	First  *Node
	Last   *Node
	Length int
}

func (s *Stack) Push(node *Node) int {
	if s.First == nil {
		s.First = node
	}

	if s.Last != nil {
		s.Last.Next = node
	}

	s.Last = node
	s.Length++

	return s.Length
}

func (s *Stack) Pop() *Node {
	if s.Length == 0 {
		return nil
	}

	current := s.First

	if s.Length == 1 {
		s.First = nil
		s.Last = nil
		s.Length = 0

		return current
	}

	pre := new(Node)

	for current.Next != nil {
		pre = current
		current = current.Next
	}

	pre.Next = nil
	s.Last = pre
	s.Length--

	return current
}
