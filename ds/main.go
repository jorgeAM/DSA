package main

import (
	"fmt"

	"github.com/jorgeAM/DSA/ds/tries"
)

func main() {
	tries := tries.Trie{Root: &tries.Node{}}

	tries.Insert("baby")
	tries.Insert("babysitter")

	fmt.Println(tries.Search("babe"))
	fmt.Println(tries.Search("babysitter"))
}
