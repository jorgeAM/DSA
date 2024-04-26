package main

import (
	"fmt"

	"github.com/jorgeAM/DSA/ds/tree"
)

func main() {
	tree := tree.BST{}
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(-3)
	tree.Insert(10)
	tree.Insert(120)

	fmt.Println(tree)
	tree.Reverse()

	fmt.Println(tree)
}
