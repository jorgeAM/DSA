package main

import (
	"fmt"

	"github.com/jorgeAM/DSA/binarySearchTree/tree"
)

func main() {
	t := new(tree.Tree)

	n1 := &tree.Node{Value: 3}
	n2 := &tree.Node{Value: 31}
	n3 := &tree.Node{Value: 45}
	n4 := &tree.Node{Value: 1}
	n5 := &tree.Node{Value: 33}
	n6 := &tree.Node{Value: 0}

	t.Insert(n1)
	t.Insert(n2)
	t.Insert(n3)
	t.Insert(n4)
	t.Insert(n5)
	t.Insert(n6)

	fmt.Println(t.Search(4))
}
