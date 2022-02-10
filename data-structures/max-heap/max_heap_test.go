package max_heap

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
	heap := MaxHeap{}
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
	heap := MaxHeap{}
	heap.items = []int{50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7, 63}

	// Act
	heap.heapifyBottomToTop()

	// Assert
	if heap.items[0] != 63 {
		t.Errorf("expected %v, got %v", 63, heap.items[0])
	}
	if heap.items[1] != 16 {
		t.Errorf("expected %v, got %v", 16, heap.items[1])
	}
	if heap.items[2] != 50 {
		t.Errorf("expected %v, got %v", 50, heap.items[2])
	}
	if heap.items[3] != 14 {
		t.Errorf("expected %v, got %v", 14, heap.items[3])
	}
	if heap.items[4] != 8 {
		t.Errorf("expected %v, got %v", 8, heap.items[4])
	}
	if heap.items[5] != 48 {
		t.Errorf("expected %v, got %v", 48, heap.items[5])
	}
	if heap.items[6] != 20 {
		t.Errorf("expected %v, got %v", 20, heap.items[6])
	}
	if heap.items[7] != 9 {
		t.Errorf("expected %v, got %v", 9, heap.items[7])
	}
	if heap.items[8] != 1 {
		t.Errorf("expected %v, got %v", 1, heap.items[8])
	}
	if heap.items[9] != 5 {
		t.Errorf("expected %v, got %v", 5, heap.items[9])
	}
	if heap.items[10] != 7 {
		t.Errorf("expected %v, got %v", 7, heap.items[10])
	}
	if heap.items[11] != 34 {
		t.Errorf("expected %v, got %v", 34, heap.items[11])
	}
}

func TestInsert(t *testing.T) {
	// Arrange
	heap := MaxHeap{}
	heap.items = []int{50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7}

	// Act
	heap.insert(63)

	// Assert
	if heap.items[0] != 63 {
		t.Errorf("expected %v, got %v", 63, heap.items[0])
	}
	if heap.items[1] != 16 {
		t.Errorf("expected %v, got %v", 16, heap.items[1])
	}
	if heap.items[2] != 50 {
		t.Errorf("expected %v, got %v", 50, heap.items[2])
	}
	if heap.items[3] != 14 {
		t.Errorf("expected %v, got %v", 14, heap.items[3])
	}
	if heap.items[4] != 8 {
		t.Errorf("expected %v, got %v", 8, heap.items[4])
	}
	if heap.items[5] != 48 {
		t.Errorf("expected %v, got %v", 48, heap.items[5])
	}
	if heap.items[6] != 20 {
		t.Errorf("expected %v, got %v", 20, heap.items[6])
	}
	if heap.items[7] != 9 {
		t.Errorf("expected %v, got %v", 9, heap.items[7])
	}
	if heap.items[8] != 1 {
		t.Errorf("expected %v, got %v", 1, heap.items[8])
	}
	if heap.items[9] != 5 {
		t.Errorf("expected %v, got %v", 5, heap.items[9])
	}
	if heap.items[10] != 7 {
		t.Errorf("expected %v, got %v", 7, heap.items[10])
	}
	if heap.items[11] != 34 {
		t.Errorf("expected %v, got %v", 34, heap.items[11])
	}
}

func TestHeapifyTopToBottom(t *testing.T) {
	// Arrange
	heap := MaxHeap{}
	heap.items = []int{34, 16, 50, 14, 8, 48, 20, 9, 1, 5, 7}

	// Act
	heap.heapifyTopToBottom()

	// Assert
	if heap.items[0] != 50 {
		t.Errorf("expected %v, got %v", 50, heap.items[0])
	}
	if heap.items[1] != 16 {
		t.Errorf("expected %v, got %v", 16, heap.items[1])
	}
	if heap.items[2] != 48 {
		t.Errorf("expected %v, got %v", 48, heap.items[2])
	}
	if heap.items[3] != 14 {
		t.Errorf("expected %v, got %v", 14, heap.items[3])
	}
	if heap.items[4] != 8 {
		t.Errorf("expected %v, got %v", 8, heap.items[4])
	}
	if heap.items[5] != 34 {
		t.Errorf("expected %v, got %v", 34, heap.items[5])
	}
	if heap.items[6] != 20 {
		t.Errorf("expected %v, got %v", 20, heap.items[6])
	}
	if heap.items[7] != 9 {
		t.Errorf("expected %v, got %v", 9, heap.items[7])
	}
	if heap.items[8] != 1 {
		t.Errorf("expected %v, got %v", 1, heap.items[8])
	}
	if heap.items[9] != 5 {
		t.Errorf("expected %v, got %v", 5, heap.items[9])
	}
	if heap.items[10] != 7 {
		t.Errorf("expected %v, got %v", 7, heap.items[10])
	}
}

func TestExtract(t *testing.T) {
	// Arrange
	heap := MaxHeap{}
	heap.items = []int{63, 16, 50, 14, 8, 48, 20, 9, 1, 5, 7, 34}

	// Act
	output := heap.extract()

	// Assert
	if output != 63 {
		t.Errorf("expected %v, got %v", 63, output)
	}
	if heap.items[0] != 50 {
		t.Errorf("expected %v, got %v", 50, heap.items[0])
	}
	if heap.items[1] != 16 {
		t.Errorf("expected %v, got %v", 16, heap.items[1])
	}
	if heap.items[2] != 48 {
		t.Errorf("expected %v, got %v", 48, heap.items[2])
	}
	if heap.items[3] != 14 {
		t.Errorf("expected %v, got %v", 14, heap.items[3])
	}
	if heap.items[4] != 8 {
		t.Errorf("expected %v, got %v", 8, heap.items[4])
	}
	if heap.items[5] != 34 {
		t.Errorf("expected %v, got %v", 34, heap.items[5])
	}
	if heap.items[6] != 20 {
		t.Errorf("expected %v, got %v", 20, heap.items[6])
	}
	if heap.items[7] != 9 {
		t.Errorf("expected %v, got %v", 9, heap.items[7])
	}
	if heap.items[8] != 1 {
		t.Errorf("expected %v, got %v", 1, heap.items[8])
	}
	if heap.items[9] != 5 {
		t.Errorf("expected %v, got %v", 5, heap.items[9])
	}
	if heap.items[10] != 7 {
		t.Errorf("expected %v, got %v", 7, heap.items[10])
	}
}
