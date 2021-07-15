package trie

import (
	"strings"
)

type Trie struct {
	Root *Node
}

func InitializeTrie() *Trie {
	return &Trie{Root: &Node{}}
}

func (t *Trie) Insert(word string) {
	currentNode := t.Root

	for _, w := range strings.ToLower(word) {
		charIndex := w - 'a'

		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.Root

	for _, w := range strings.ToLower(word) {
		charIndex := w - 'a'

		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	if !currentNode.isEnd {
		return false
	}

	return true
}
