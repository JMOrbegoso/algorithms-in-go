package trie

import "unicode"

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Node struct {
	children [len(ALPHABET)]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func (t *Trie) init() {
	t.root = &Node{}
}

func getRuneNumber(rune rune) int {
	for i, char := range ALPHABET {
		if char == unicode.ToUpper(rune) {
			return i
		}
	}
	return -1
}

func (t *Trie) insert(word string) {
	currentNode := t.root

	for _, rune := range word {
		charIndex := getRuneNumber(rune)

		if charIndex == -1 {
			return
		}

		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

func (t *Trie) find(word string) bool {
	currentNode := t.root

	for _, rune := range word {
		charIndex := getRuneNumber(rune)

		if charIndex == -1 {
			return false
		}

		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}
