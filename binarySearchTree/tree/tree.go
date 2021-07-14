package tree

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(node *Node) {
	if t.Root == nil {
		t.Root = node
		return
	}

	t.Root.Insert(node)
}
