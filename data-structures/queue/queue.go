package queue

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front  *Node
	rear   *Node
	length uint
}

func (q *Queue) enqueue(node *Node) {
	node.next = nil

	if q.rear == nil {
		q.front = node
	} else {
		q.rear.next = node
	}

	q.rear = node
	q.length++
}

func (q *Queue) dequeue() *Node {
	if q.front == nil {
		return nil
	}

	dequeuedNode := q.front

	if q.front == q.rear {
		q.front = nil
		q.rear = nil
	} else {
		q.front = q.front.next
	}

	q.length--

	return dequeuedNode
}

func (q *Queue) print() {
	print(q.front)
}

func print(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.value)
	print(node.next)
}
