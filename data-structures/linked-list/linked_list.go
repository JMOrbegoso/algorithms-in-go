package linked_list

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	length uint
}

func (l *LinkedList) preppend(node *Node) {
	node.next = l.head
	l.head = node
	l.length++
}

func (l *LinkedList) append(node *Node) {
	node.next = nil

	if l.head == nil {
		l.head = node
	} else {
		lastNode := getLastNode(l.head)
		lastNode.next = node
	}

	l.length++
}

func getLastNode(node *Node) *Node {
	if node.next == nil {
		return node
	}

	return getLastNode(node.next)
}

func (l *LinkedList) print() {
	print(l.head)
}

func print(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.value)
	print(node.next)
}

func (l *LinkedList) getIndexOf(value int) int {
	return findIndexOf(l.head, 0, value)
}

func findIndexOf(node *Node, index int, value int) int {
	if node == nil {
		return -1
	}

	if node.value == value {
		return index
	}

	return findIndexOf(node.next, index+1, value)
}

func (l *LinkedList) deleteByValue(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
	} else {
		previousNode := getPreviousNodeByValue(l.head, value)
		if previousNode == nil {
			return
		}
		previousNode.next = previousNode.next.next
	}

	l.length--
}

func getPreviousNodeByValue(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if node.next == nil {
		return nil
	}

	if node.next.value == value {
		return node
	}

	return getPreviousNodeByValue(node.next, value)
}

func (l *LinkedList) reverse() {
	var previous *Node = nil
	var current *Node = l.head
	var following *Node = l.head

	for current != nil {
		following = following.next
		current.next = previous
		previous = current
		current = following
	}

	l.head = previous
}
