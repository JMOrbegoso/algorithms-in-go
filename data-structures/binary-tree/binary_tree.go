package binary_tree

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type queueNode struct {
	value *Node
	next  *queueNode
}

type queue struct {
	front  *queueNode
	rear   *queueNode
	length uint
}

func (q *queue) enqueue(node *queueNode) {
	node.next = nil

	if q.rear == nil {
		q.front = node
	} else {
		q.rear.next = node
	}

	q.rear = node
	q.length++
}

func (q *queue) dequeue() *queueNode {
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

func (n *Node) printByLevel() {
	if n == nil {
		return
	}

	auxQueue := queue{}
	newQueueNode := queueNode{value: n}
	auxQueue.enqueue(&newQueueNode)
	n.print(auxQueue)
}

func (n *Node) print(queue queue) {
	if queue.length == 0 {
		return
	}

	node := queue.dequeue().value

	fmt.Println(node)

	if node.left != nil {
		leftQueueNode := queueNode{value: node.left}
		queue.enqueue(&leftQueueNode)
	}

	if node.right != nil {
		rightQueueNode := queueNode{value: node.right}
		queue.enqueue(&rightQueueNode)
	}

	n.print(queue)
}

func (n *Node) include(value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	return n.left.include(value) || n.right.include(value)
}

func (n *Node) totalSum() int {
	if n == nil {
		return 0
	}

	return n.value + n.left.totalSum() + n.right.totalSum()
}

func (n *Node) minValue() int {
	if n == nil {
		return math.MaxInt
	}

	aux := n.value
	auxLeft := n.left.minValue()
	auxRight := n.right.minValue()

	if aux > auxLeft {
		aux = auxLeft
	}
	if aux > auxRight {
		aux = auxRight
	}

	return aux
}

func (n *Node) maxValue() int {
	if n == nil {
		return math.MinInt
	}

	aux := n.value
	auxLeft := n.left.maxValue()
	auxRight := n.right.maxValue()

	if aux < auxLeft {
		aux = auxLeft
	}
	if aux < auxRight {
		aux = auxRight
	}

	return aux
}

func (n *Node) maxSumPath(accumulated int) int {
	if n == nil {
		return accumulated
	}

	accumulated += n.value

	auxLeft := n.left.maxSumPath(accumulated)
	auxRight := n.right.maxSumPath(accumulated)

	if auxRight > auxLeft {
		return auxRight
	}
	return auxLeft
}

func (n *Node) invertTree() {
	if n == nil {
		return
	}

	auxQueue := queue{}

	auxQueueNode := queueNode{value: n}
	auxQueue.enqueue(&auxQueueNode)

	invert(auxQueue)
}

func invert(queue queue) {
	if queue.length == 0 {
		return
	}

	node := queue.dequeue().value

	temp := node.right
	node.right = node.left
	node.left = temp

	if node.left != nil {
		leftQueueNode := queueNode{value: node.left}
		queue.enqueue(&leftQueueNode)
	}

	if node.right != nil {
		rightQueueNode := queueNode{value: node.right}
		queue.enqueue(&rightQueueNode)
	}

	invert(queue)
}

// Also called Height, is the number of nodes along the longest path from the root node down to the farthest leaf node.
func (n *Node) getMaxDepth() int {
	if n == nil {
		return -1
	}

	return n.getHeight(-1)
}

func (n *Node) getHeight(accumulated int) int {
	if n == nil {
		return accumulated
	}

	accumulated++

	auxLeft := n.left.getHeight(accumulated)
	auxRight := n.right.getHeight(accumulated)

	if auxRight > auxLeft {
		return auxRight
	}
	return auxLeft
}

func (n *Node) getMaxWidth() int {
	height := n.getMaxDepth()
	maxLevelWidth := 0

	for i := 1; i <= height; i++ {
		levelWidth := n.getLevelWidth(i)

		if maxLevelWidth < levelWidth {
			maxLevelWidth = levelWidth
		}
	}

	return maxLevelWidth
}

func (n *Node) getLevelWidth(level int) int {
	if n == nil {
		return 0
	}

	if level == 1 {
		return 1
	}

	return n.left.getLevelWidth(level-1) + n.right.getLevelWidth(level-1)
}

func (n *Node) insert(value int) {
	newNode := Node{value: value}
	auxQueue := queue{}

	rootQueueNode := queueNode{value: n}
	auxQueue.enqueue(&rootQueueNode)

	insertNewNode(auxQueue, &newNode)
}

func insertNewNode(queue queue, newNode *Node) {
	if queue.length == 0 {
		return
	}

	node := queue.dequeue().value

	if node.left == nil {
		node.left = newNode
		return
	}
	if node.right == nil {
		node.right = newNode
		return
	}

	leftQueueNode := queueNode{value: node.left}
	queue.enqueue(&leftQueueNode)
	rightQueueNode := queueNode{value: node.right}
	queue.enqueue(&rightQueueNode)

	insertNewNode(queue, newNode)
}
