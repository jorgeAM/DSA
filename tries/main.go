package main

import (
	"fmt"

	"github.com/jorgeAM/structures/tries/trie"
)

func main() {
	t := trie.InitializeTrie()

	t.Insert("baby")
	t.Insert("babysitter")

	fmt.Println(t.Search("babysitter"))
	fmt.Println(t.Search("babies"))

}
