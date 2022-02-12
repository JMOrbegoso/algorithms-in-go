package priority_queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	// Arrange
	priorityQueue := PriorityQueue{}

	// Act
	priorityQueue.enqueue(20, 2)
	priorityQueue.enqueue(10, 1)

	// Assert
	if priorityQueue.count != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.count)
	}

	if priorityQueue.items[0].value != 20 {
		t.Errorf("expected %v, got %v", 20, priorityQueue.items[0].value)
	}
	if priorityQueue.items[0].priority != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.items[0].priority)
	}

	if priorityQueue.items[1].value != 10 {
		t.Errorf("expected %v, got %v", 10, priorityQueue.items[1].value)
	}
	if priorityQueue.items[1].priority != 1 {
		t.Errorf("expected %v, got %v", 1, priorityQueue.items[1].priority)
	}
}

func TestSwap(t *testing.T) {
	// Arrange
	priorityQueue := PriorityQueue{}
	priorityQueue.enqueue(30, 3)
	priorityQueue.enqueue(20, 2)
	priorityQueue.enqueue(10, 1)

	// Act
	priorityQueue.swap(0, 1)

	// Assert
	if priorityQueue.count != 3 {
		t.Errorf("expected %v, got %v", 3, priorityQueue.count)
	}

	if priorityQueue.items[0].value != 20 {
		t.Errorf("expected %v, got %v", 20, priorityQueue.items[0].value)
	}
	if priorityQueue.items[0].priority != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.items[0].priority)
	}

	if priorityQueue.items[1].value != 30 {
		t.Errorf("expected %v, got %v", 30, priorityQueue.items[1].value)
	}
	if priorityQueue.items[1].priority != 3 {
		t.Errorf("expected %v, got %v", 3, priorityQueue.items[1].priority)
	}

	if priorityQueue.items[2].value != 10 {
		t.Errorf("expected %v, got %v", 10, priorityQueue.items[2].value)
	}
	if priorityQueue.items[2].priority != 1 {
		t.Errorf("expected %v, got %v", 1, priorityQueue.items[2].priority)
	}
}

func TestDequeue(t *testing.T) {
	// Arrange
	priorityQueue := PriorityQueue{}
	priorityQueue.enqueue(30, 3)
	priorityQueue.enqueue(20, 2)
	priorityQueue.enqueue(10, 1)

	// Act
	output := priorityQueue.dequeue()

	// Assert
	if output.value != 30 {
		t.Errorf("expected %v, got %v", 30, output.value)
	}
	if output.priority != 3 {
		t.Errorf("expected %v, got %v", 3, output.priority)
	}

	if priorityQueue.count != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.count)
	}

	if priorityQueue.items[0].value != 20 {
		t.Errorf("expected %v, got %v", 20, priorityQueue.items[0].value)
	}
	if priorityQueue.items[0].priority != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.items[0].priority)
	}

	if priorityQueue.items[1].value != 10 {
		t.Errorf("expected %v, got %v", 10, priorityQueue.items[1].value)
	}
	if priorityQueue.items[1].priority != 1 {
		t.Errorf("expected %v, got %v", 1, priorityQueue.items[1].priority)
	}
}
