package min_heap

import (
	"testing"
)

func TestGetParentIndex(t *testing.T) {
	// Arrange
	testsValues := []struct {
		input int
		want  int
	}{
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 3},
		{9, 4},
		{10, 4},
	}

	for _, testValue := range testsValues {
		// Act
		output := getParentIndex(testValue.input)

		// Assert
		if output != testValue.want {
			t.Errorf("expected %v, got %v", testValue.want, output)
		}
	}
}

func TestGetLeftChildIndex(t *testing.T) {
	// Arrange
	testsValues := []struct {
		input int
		want  int
	}{
		{0, 1},
		{1, 3},
		{2, 5},
		{3, 7},
		{4, 9},
		{5, 11},
	}

	for _, testValue := range testsValues {
		// Act
		output := getLeftChildIndex(testValue.input)

		// Assert
		if output != testValue.want {
			t.Errorf("expected %v, got %v", testValue.want, output)
		}
	}
}

func TestGetRightChildIndex(t *testing.T) {
	// Arrange
	testsValues := []struct {
		input int
		want  int
	}{
		{0, 2},
		{1, 4},
		{2, 6},
		{3, 8},
		{4, 10},
		{5, 12},
	}

	for _, testValue := range testsValues {
		// Act
		output := getRightChildIndex(testValue.input)

		// Assert
		if output != testValue.want {
			t.Errorf("expected %v, got %v", testValue.want, output)
		}
	}
}

func TestSwap(t *testing.T) {
	// Arrange
	heap := MinHeap{}
	heap.items = []int{50, 16, 12}

	// Act
	heap.swap(0, 1)

	// Assert
	if heap.items[0] != 16 {
		t.Errorf("expected %v, got %v", 16, heap.items[0])
	}
	if heap.items[1] != 50 {
		t.Errorf("expected %v, got %v", 50, heap.items[1])
	}
	if heap.items[2] != 12 {
		t.Errorf("expected %v, got %v", 12, heap.items[2])
	}
}

func TestHeapifyBottomToTop(t *testing.T) {
	// Arrange
	heap := MinHeap{}
	heap.items = []int{5, 8, 19, 10, 15, 2}

	// Act
	heap.heapifyBottomToTop()

	// Assert
	if heap.items[0] != 2 {
		t.Errorf("expected %v, got %v", 2, heap.items[0])
	}
	if heap.items[1] != 8 {
		t.Errorf("expected %v, got %v", 8, heap.items[1])
	}
	if heap.items[2] != 5 {
		t.Errorf("expected %v, got %v", 5, heap.items[2])
	}
	if heap.items[3] != 10 {
		t.Errorf("expected %v, got %v", 10, heap.items[3])
	}
	if heap.items[4] != 15 {
		t.Errorf("expected %v, got %v", 15, heap.items[4])
	}
	if heap.items[5] != 19 {
		t.Errorf("expected %v, got %v", 19, heap.items[5])
	}
}

func TestInsert(t *testing.T) {
	// Arrange
	heap := MinHeap{}
	heap.items = []int{5, 8, 19, 10, 15}

	// Act
	heap.insert(2)

	// Assert
	if heap.items[0] != 2 {
		t.Errorf("expected %v, got %v", 2, heap.items[0])
	}
	if heap.items[1] != 8 {
		t.Errorf("expected %v, got %v", 8, heap.items[1])
	}
	if heap.items[2] != 5 {
		t.Errorf("expected %v, got %v", 5, heap.items[2])
	}
	if heap.items[3] != 10 {
		t.Errorf("expected %v, got %v", 10, heap.items[3])
	}
	if heap.items[4] != 15 {
		t.Errorf("expected %v, got %v", 15, heap.items[4])
	}
	if heap.items[5] != 19 {
		t.Errorf("expected %v, got %v", 19, heap.items[5])
	}
}

func TestHeapifyTopToBottom(t *testing.T) {
	// Arrange
	heap := MinHeap{}
	heap.items = []int{30, 8, 19, 10, 15, 21}

	// Act
	heap.heapifyTopToBottom()

	// Assert
	if heap.items[0] != 8 {
		t.Errorf("expected %v, got %v", 8, heap.items[0])
	}
	if heap.items[1] != 10 {
		t.Errorf("expected %v, got %v", 10, heap.items[1])
	}
	if heap.items[2] != 19 {
		t.Errorf("expected %v, got %v", 19, heap.items[2])
	}
	if heap.items[3] != 30 {
		t.Errorf("expected %v, got %v", 30, heap.items[3])
	}
	if heap.items[4] != 15 {
		t.Errorf("expected %v, got %v", 15, heap.items[4])
	}
	if heap.items[5] != 21 {
		t.Errorf("expected %v, got %v", 21, heap.items[5])
	}
}

func TestExtract(t *testing.T) {
	// Arrange
	heap := MinHeap{}
	heap.items = []int{2, 8, 19, 10, 15, 21, 30}

	// Act
	output := heap.extract()

	// Assert
	if output != 2 {
		t.Errorf("expected %v, got %v", 2, output)
	}
	if heap.items[0] != 8 {
		t.Errorf("expected %v, got %v", 8, heap.items[0])
	}
	if heap.items[1] != 10 {
		t.Errorf("expected %v, got %v", 10, heap.items[1])
	}
	if heap.items[2] != 19 {
		t.Errorf("expected %v, got %v", 19, heap.items[2])
	}
	if heap.items[3] != 30 {
		t.Errorf("expected %v, got %v", 30, heap.items[3])
	}
	if heap.items[4] != 15 {
		t.Errorf("expected %v, got %v", 15, heap.items[4])
	}
	if heap.items[5] != 21 {
		t.Errorf("expected %v, got %v", 21, heap.items[5])
	}
}
