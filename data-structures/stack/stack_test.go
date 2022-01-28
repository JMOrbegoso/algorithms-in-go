package stack

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPushInEmptyStack(t *testing.T) {
	// Arrange
	node1 := Node{value: 10}
	stack := Stack{}

	// Act
	stack.push(&node1)

	// Assert
	if stack.length != 1 {
		t.Errorf("expected %v, got %v", 1, stack.length)
	}
	if stack.top.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, stack.top.value)
	}
	if stack.top.next != nil {
		t.Errorf("expected %v, got %v", nil, stack.top.next)
	}
}

func TestPush(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20}
	node1 := Node{value: 10}
	stack := Stack{}

	// Act
	stack.push(&node3)
	stack.push(&node2)
	stack.push(&node1)

	// Assert
	if stack.length != 3 {
		t.Errorf("expected %v, got %v", 3, stack.length)
	}
	if stack.top.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, stack.top.value)
	}
	if stack.top.next.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, stack.top.next.value)
	}
	if stack.top.next.next.value != node3.value {
		t.Errorf("expected %v, got %v", node2.value, stack.top.next.next.value)
	}
	if stack.top.next.next.next != nil {
		t.Errorf("expected %v, got %v", nil, stack.top.next.next.next)
	}
}

func TestPopEmptyStack(t *testing.T) {
	// Arrange
	stack := Stack{}

	// Act
	output := stack.pop()

	// Assert
	if stack.length != 0 {
		t.Errorf("expected %v, got %v", 0, stack.length)
	}
	if output != nil {
		t.Errorf("expected %v, got %v", nil, output)
	}
}

func TestPop(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	stack := Stack{length: 3, top: &node1}

	// Act
	output := stack.pop()

	// Assert
	if stack.length != 2 {
		t.Errorf("expected %v, got %v", 2, stack.length)
	}
	if stack.top.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, stack.top.value)
	}
	if stack.top.next.value != node3.value {
		t.Errorf("expected %v, got %v", node3.value, stack.top.next.value)
	}
	if stack.top.next.next != nil {
		t.Errorf("expected %v, got %v", nil, stack.top.next.next)
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
	stack := Stack{length: 3, top: &node1}

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	stack.print()

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
