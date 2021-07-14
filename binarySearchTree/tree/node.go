package tree

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

func (n *Node) Insert(node *Node) {
	if n.Value < node.Value {
		if n.Right == nil {
			n.Right = node
		} else {
			n.Right.Insert(node)
		}
	} else {
		if n.Left == nil {
			n.Left = node
		} else {
			n.Left.Insert(node)
		}
	}
}
