package trie

const ALPHABET_SIZE = 26

type Node struct {
	children [ALPHABET_SIZE]*Node
	isEnd    bool
}
