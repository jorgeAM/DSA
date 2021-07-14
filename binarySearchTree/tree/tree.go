package tree

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(node *Node) {
	if t.Root == nil {
		t.Root = node
		return
	}

	t.Root.insert(node)
}

func (t *Tree) Search(value int) bool {
	return t.Root.search(value)
}
