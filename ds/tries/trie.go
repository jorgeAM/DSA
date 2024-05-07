package tries

import "strings"

const ALPHABET_SIZE = 26

type Node struct {
	children [ALPHABET_SIZE]*Node
	isEnd    bool
}

type Trie struct {
	Root *Node
}

func (t *Trie) Insert(s string) {
	curr := t.Root

	for _, char := range strings.ToLower(s) {
		index := char - 'a'

		if curr.children[index] == nil {
			curr.children[index] = &Node{}
		}

		curr = curr.children[index]
	}

	curr.isEnd = true
}

func (t *Trie) Search(s string) bool {
	curr := t.Root

	for _, char := range strings.ToLower(s) {
		index := char - 'a'

		if curr.children[index] == nil {
			return false
		}

		curr = curr.children[index]
	}

	return curr.isEnd
}
