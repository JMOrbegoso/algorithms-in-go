package queue

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestEnqueueInEmptyQueue(t *testing.T) {
	// Arrange
	node1 := Node{value: 10}
	queue := Queue{}

	// Act
	queue.enqueue(&node1)

	// Assert
	if queue.length != 1 {
		t.Errorf("expected %v, got %v", 1, queue.length)
	}
	if queue.front.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, queue.front.value)
	}
	if queue.rear.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, queue.rear.value)
	}
	if queue.front.next != nil {
		t.Errorf("expected %v, got %v", nil, queue.front.next)
	}
}

func TestEnqueue(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20}
	node1 := Node{value: 10}
	queue := Queue{}

	// Act
	queue.enqueue(&node1)
	queue.enqueue(&node2)
	queue.enqueue(&node3)

	// Assert
	if queue.length != 3 {
		t.Errorf("expected %v, got %v", 3, queue.length)
	}
	if queue.front.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, queue.front.value)
	}
	if queue.front.next.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, queue.front.next.value)
	}
	if queue.front.next.next.value != node3.value {
		t.Errorf("expected %v, got %v", node2.value, queue.front.next.next.value)
	}
	if queue.front.next.next.next != nil {
		t.Errorf("expected %v, got %v", nil, queue.front.next.next.next)
	}
	if queue.rear.value != node3.value {
		t.Errorf("expected %v, got %v", node3.value, queue.rear.value)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	// Arrange
	queue := Queue{}

	// Act
	output := queue.dequeue()

	// Assert
	if queue.length != 0 {
		t.Errorf("expected %v, got %v", 0, queue.length)
	}
	if queue.front != nil {
		t.Errorf("expected %v, got %v", nil, queue.front)
	}
	if queue.rear != nil {
		t.Errorf("expected %v, got %v", nil, queue.front)
	}
	if output != nil {
		t.Errorf("expected %v, got %v", nil, output)
	}
}

func TestDequeueUniqueElement(t *testing.T) {
	// Arrange
	node1 := Node{value: 10}
	queue := Queue{length: 1, front: &node1, rear: &node1}

	// Act
	output := queue.dequeue()

	// Assert
	if queue.length != 0 {
		t.Errorf("expected %v, got %v", 0, queue.length)
	}
	if queue.front != nil {
		t.Errorf("expected %v, got %v", nil, queue.front)
	}
	if queue.rear != nil {
		t.Errorf("expected %v, got %v", nil, queue.rear)
	}
	if output.value != 10 {
		t.Errorf("expected %v, got %v", 10, output.value)
	}
}

func TestDequeue(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	queue := Queue{length: 3, front: &node1, rear: &node3}

	// Act
	output := queue.dequeue()

	// Assert
	if queue.length != 2 {
		t.Errorf("expected %v, got %v", 2, queue.length)
	}
	if queue.front.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, queue.front.value)
	}
	if queue.front.next.value != node3.value {
		t.Errorf("expected %v, got %v", node3.value, queue.front.next.value)
	}
	if queue.front.next.next != nil {
		t.Errorf("expected %v, got %v", nil, queue.front.next.next)
	}
	if queue.rear.value != node3.value {
		t.Errorf("expected %v, got %v", node3.value, queue.rear.value)
	}
	if output.value != 10 {
		t.Errorf("expected %v, got %v", 10, output.value)
	}
}

func TestPrint(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	queue := Queue{length: 3, front: &node1, rear: &node3}

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	queue.print()

	w.Close()
	out, _ := ioutil.ReadAll(r)

	// Assert
	if !strings.Contains(string(out), strconv.Itoa(node1.value)) {
		t.Errorf("expected %v", node1.value)
	}
	if !strings.Contains(string(out), strconv.Itoa(node2.value)) {
		t.Errorf("expected %v", node2.value)
	}
	if !strings.Contains(string(out), strconv.Itoa(node3.value)) {
		t.Errorf("expected %v", node3.value)
	}
}
