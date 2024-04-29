package tree

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

func (n *Node) Insert(value int) {
	if n.Value > value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Node) Search(value int) bool {
	if n.Value == value {
		return true
	}

	if n.Value > value {
		if n.Left == nil {
			return false
		} else {
			return n.Left.Search(value)
		}
	}

	if n.Right == nil {
		return false
	}

	return n.Right.Search(value)
}

func (n *Node) Reverse() {
	if n == nil {
		return
	}

	n.Left, n.Right = n.Right, n.Left

	n.Left.Reverse()
	n.Right.Reverse()
}

type BST struct {
	Root *Node
}

func (t *BST) Insert(value int) {
	node := &Node{Value: value}

	if t.Root == nil {
		t.Root = node

		return
	}

	t.Root.Insert(value)
}

func (t *BST) Search(value int) bool {
	if t.Root == nil {
		return false
	}

	return t.Root.Search(value)
}

func (t *BST) Reverse() {
	t.Root.Reverse()
}

func (t *BST) BFS() []int {
	queue := make([]*Node, 0)
	viewed := make([]int, 0)
	var node *Node

	if t.Root == nil {
		return viewed
	}

	queue = append(queue, t.Root)

	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		viewed = append(viewed, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return viewed
}
