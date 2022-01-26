package linked_list

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPreppendInEmptyLinkedList(t *testing.T) {
	// Arrange
	node1 := Node{value: 10}
	linkedList := LinkedList{}

	// Act
	linkedList.preppend(&node1)

	// Assert
	if linkedList.length != 1 {
		t.Errorf("expected length: %v, got %v", 1, linkedList.length)
	}
	if linkedList.head.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, linkedList.head.value)
	}
	if linkedList.head.next != nil {
		t.Errorf("linkedList.head.next should be nil")
	}
}

func TestPreppend(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20}
	node1 := Node{value: 10}
	linkedList := LinkedList{}

	// Act
	linkedList.preppend(&node3)
	linkedList.preppend(&node2)
	linkedList.preppend(&node1)

	// Assert
	if linkedList.length != 3 {
		t.Errorf("expected length: %v, got %v", 3, linkedList.length)
	}
	if linkedList.head.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, linkedList.head.value)
	}
	if linkedList.head.next.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, linkedList.head.next.value)
	}
	if linkedList.head.next.next.value != node3.value {
		t.Errorf("expected %v, got %v", node2.value, linkedList.head.next.next.value)
	}
	if linkedList.head.next.next.next != nil {
		t.Errorf("linkedList.head.next.next.next should be nil")
	}
}

func TestAppendInEmptyLinkedList(t *testing.T) {
	// Arrange
	node1 := Node{value: 10}
	linkedList := LinkedList{}

	// Act
	linkedList.append(&node1)

	// Assert
	if linkedList.length != 1 {
		t.Errorf("expected length: %v, got %v", 1, linkedList.length)
	}
	if linkedList.head.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, linkedList.head.value)
	}
	if linkedList.head.next != nil {
		t.Errorf("linkedList.head.next should be nil")
	}
}

func TestAppend(t *testing.T) {
	// Arrange
	node1 := Node{value: 10}
	node2 := Node{value: 20}
	node3 := Node{value: 30}
	linkedList := LinkedList{}

	// Act
	linkedList.append(&node1)
	linkedList.append(&node2)
	linkedList.append(&node3)

	// Assert
	if linkedList.length != 3 {
		t.Errorf("expected length: %v, got %v", 3, linkedList.length)
	}
	if linkedList.head.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, linkedList.head.value)
	}
	if linkedList.head.next.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, linkedList.head.next.value)
	}
	if linkedList.head.next.next.value != node3.value {
		t.Errorf("expected %v, got %v", node3.value, linkedList.head.next.next.value)
	}
	if linkedList.head.next.next.next != nil {
		t.Errorf("linkedList.head.next.next.next should be nil")
	}
}

func TestPrint(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	linkedList := LinkedList{length: 3, head: &node1}

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	linkedList.print()

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

func TestGetIndexOf(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	linkedList := LinkedList{length: 3, head: &node1}

	testsValues := []struct {
		input int
		want  int
	}{
		{50, -1},
		{10, 0},
		{20, 1},
		{30, 2},
	}

	// Act
	for _, testValue := range testsValues {
		output := linkedList.getIndexOf(testValue.input)

		// Assert
		if output != testValue.want {
			t.Errorf("expected %v, got %v", testValue.want, output)
		}
	}
}

func TestDeleteByValueNotFound(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	linkedList := LinkedList{length: 3, head: &node1}

	// Act
	linkedList.deleteByValue(50)

	// Assert
	if linkedList.length != 3 {
		t.Errorf("expected length: %v, got %v", 3, linkedList.length)
	}
	if linkedList.head.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, linkedList.head.value)
	}
	if linkedList.head.next.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, linkedList.head.next.value)
	}
	if linkedList.head.next.next.value != node3.value {
		t.Errorf("expected %v, got %v", node3.value, linkedList.head.next.next.value)
	}
	if linkedList.head.next.next.next != nil {
		t.Errorf("linkedList.head.next.next.next should be nil")
	}
}

func TestDeleteHeadByValue(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	linkedList := LinkedList{length: 3, head: &node1}

	// Act
	linkedList.deleteByValue(10)

	// Assert
	if linkedList.length != 2 {
		t.Errorf("expected length: %v, got %v", 2, linkedList.length)
	}
	if linkedList.head.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, linkedList.head.value)
	}
	if linkedList.head.next.value != node3.value {
		t.Errorf("expected %v, got %v", node3.value, linkedList.head.next.value)
	}
	if linkedList.head.next.next != nil {
		t.Errorf("linkedList.head.next.next should be nil")
	}
}

func TestDeleteTailByValue(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	linkedList := LinkedList{length: 3, head: &node1}

	// Act
	linkedList.deleteByValue(30)

	// Assert
	if linkedList.length != 2 {
		t.Errorf("expected length: %v, got %v", 2, linkedList.length)
	}
	if linkedList.head.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, linkedList.head.value)
	}
	if linkedList.head.next.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, linkedList.head.next.value)
	}
	if linkedList.head.next.next != nil {
		t.Errorf("linkedList.head.next.next should be nil")
	}
}

func TestReverseEmptyLinkedList(t *testing.T) {
	// Arrange

	linkedList := LinkedList{}

	// Act
	linkedList.reverse()

	// Assert
	if linkedList.length != 0 {
		t.Errorf("expected length: %v, got %v", 0, linkedList.length)
	}
	if linkedList.head != nil {
		t.Errorf("linkedList.head should be nil")
	}
}

func TestReverse(t *testing.T) {
	// Arrange
	node3 := Node{value: 30}
	node2 := Node{value: 20, next: &node3}
	node1 := Node{value: 10, next: &node2}
	linkedList := LinkedList{length: 3, head: &node1}

	// Act
	linkedList.reverse()

	// Assert
	if linkedList.length != 3 {
		t.Errorf("expected length: %v, got %v", 3, linkedList.length)
	}
	if linkedList.head.value != node3.value {
		t.Errorf("expected %v, got %v", node3.value, linkedList.head.value)
	}
	if linkedList.head.next.value != node2.value {
		t.Errorf("expected %v, got %v", node2.value, linkedList.head.next.value)
	}
	if linkedList.head.next.next.value != node1.value {
		t.Errorf("expected %v, got %v", node1.value, linkedList.head.next.next.value)
	}
	if linkedList.head.next.next.next != nil {
		t.Errorf("linkedList.head.next.next.next should be nil")
	}
}
