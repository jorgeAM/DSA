package trie

const ALPHABET_SIZE = 26

type Node struct {
	children [26]*Node
	isEnd    bool
}
