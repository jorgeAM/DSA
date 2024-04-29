package main

import (
	"fmt"

	"github.com/jorgeAM/DSA/ds/tree"
)

func main() {
	tree := tree.BST{}
	tree.Insert(10)
	tree.Insert(6)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(20)

	fmt.Println(tree.BFS())
}
