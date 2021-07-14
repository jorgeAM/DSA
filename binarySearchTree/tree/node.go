package tree

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

func (n *Node) insert(node *Node) {
	if n.Value < node.Value {
		if n.Right == nil {
			n.Right = node
		} else {
			n.Right.insert(node)
		}
	} else {
		if n.Left == nil {
			n.Left = node
		} else {
			n.Left.insert(node)
		}
	}
}

func (n *Node) search(value int) bool {
	if n.Value < value {
		if n.Right == nil {
			return false
		}

		return n.Right.search(value)
	} else if n.Value > value {
		if n.Left == nil {
			return false
		}

		return n.Left.search(value)
	} else {
		return true
	}
}

func (n *Node) isLeaf() bool {
	if n.Right == nil && n.Left == nil {
		return true
	}

	return false
}
