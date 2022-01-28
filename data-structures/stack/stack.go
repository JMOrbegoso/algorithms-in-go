package stack

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top    *Node
	length uint
}

func (s *Stack) push(node *Node) {
	node.next = s.top
	s.top = node
	s.length++
}

func (s *Stack) pop() *Node {
	if s.top == nil {
		return nil
	}

	poppedNode := s.top
	s.top = s.top.next
	s.length--

	return poppedNode
}

func (s *Stack) print() {
	print(s.top)
}

func print(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.value)
	print(node.next)
}
