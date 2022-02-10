package array

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestArrayNotInitialized(t *testing.T) {
	// Arrange

	// Act
	array := Array{}

	// Assert
	if array.count != 0 {
		t.Errorf("expected %v, got %v", 0, array.count)
	}
	if array.size != 0 {
		t.Errorf("expected %v, got %v", 0, array.size)
	}
	if array.items != nil {
		t.Errorf("expected %v, got %v", nil, array.items)
	}
}

func TestInit(t *testing.T) {
	// Arrange
	array := Array{}

	// Act
	array.init(3)

	// Assert
	if array.count != 0 {
		t.Errorf("expected %v, got %v", 0, array.count)
	}
	if array.size != 3 {
		t.Errorf("expected %v, got %v", 3, array.size)
	}
	if array.items == nil {
		t.Errorf("expected not nil, got %v", array.items)
	}
}

func TestGetIndexNotFound(t *testing.T) {
	// Arrange
	array := Array{}
	array.items = []int{3}

	// Act
	output := array.getIndex(5)

	// Assert
	if output != -1 {
		t.Errorf("expected %v, got %v", -1, output)
	}
}

func TestGetIndex(t *testing.T) {
	// Arrange
	array := Array{}
	array.items = []int{3, 4, 5}

	// Act
	output := array.getIndex(5)

	// Assert
	if output != 2 {
		t.Errorf("expected %v, got %v", 2, output)
	}
}

func TestExpandEmptyArray(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(3)

	// Act
	array.expand()

	// Assert
	if array.count != 0 {
		t.Errorf("expected %v, got %v", 0, array.count)
	}
	if array.size != 6 {
		t.Errorf("expected %v, got %v", 6, array.size)
	}
	if array.items == nil {
		t.Errorf("expected not nil, got %v", array.items)
	}
}

func TestExpand(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(3)
	array.items = []int{3, 4, 5}
	array.count = 3

	// Act
	array.expand()

	// Assert
	if array.count != 3 {
		t.Errorf("expected %v, got %v", 3, array.count)
	}
	if array.size != 6 {
		t.Errorf("expected %v, got %v", 6, array.size)
	}
	if array.items[0] != 3 {
		t.Errorf("expected %v, got %v", 3, array.items[0])
	}
	if array.items[1] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[1])
	}
	if array.items[2] != 5 {
		t.Errorf("expected %v, got %v", 5, array.items[2])
	}
}

func TestInsertWithoutExpansion(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(3)

	// Act
	array.insert(3)
	array.insert(4)
	array.insert(5)

	// Assert
	if array.count != 3 {
		t.Errorf("expected %v, got %v", 3, array.count)
	}
	if array.size != 3 {
		t.Errorf("expected %v, got %v", 3, array.size)
	}
	if array.items[0] != 3 {
		t.Errorf("expected %v, got %v", 3, array.items[0])
	}
	if array.items[1] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[1])
	}
	if array.items[2] != 5 {
		t.Errorf("expected %v, got %v", 5, array.items[2])
	}
}

func TestInsertWithExpansion(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(3)

	// Act
	array.insert(3)
	array.insert(4)
	array.insert(5)
	array.insert(6)
	array.insert(7)

	// Assert
	if array.count != 5 {
		t.Errorf("expected %v, got %v", 5, array.count)
	}
	if array.size != 6 {
		t.Errorf("expected %v, got %v", 6, array.size)
	}
	if array.items[0] != 3 {
		t.Errorf("expected %v, got %v", 3, array.items[0])
	}
	if array.items[1] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[1])
	}
	if array.items[2] != 5 {
		t.Errorf("expected %v, got %v", 5, array.items[2])
	}
	if array.items[3] != 6 {
		t.Errorf("expected %v, got %v", 6, array.items[3])
	}
	if array.items[4] != 7 {
		t.Errorf("expected %v, got %v", 7, array.items[4])
	}
}

func TestRemoveAt(t *testing.T) {
	// Arrang
	array := Array{}
	array.init(3)
	array.items = []int{5, 4, 8}
	array.count = 3

	// Act
	array.removeAt(1)

	// Assert
	if array.count != 2 {
		t.Errorf("expected %v, got %v", 2, array.count)
	}
	if array.size != 3 {
		t.Errorf("expected %v, got %v", 3, array.size)
	}
	if array.items[0] != 5 {
		t.Errorf("expected %v, got %v", 5, array.items[0])
	}
	if array.items[1] != 8 {
		t.Errorf("expected %v, got %v", 8, array.items[1])
	}
}

func TestRemoveLastItem(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(3)
	array.items = []int{5, 4, 8}
	array.count = 3

	// Act
	array.removeAt(2)

	// Assert
	if array.count != 2 {
		t.Errorf("expected %v, got %v", 2, array.count)
	}
	if array.size != 3 {
		t.Errorf("expected %v, got %v", 3, array.size)
	}
	if array.items[0] != 5 {
		t.Errorf("expected %v, got %v", 5, array.items[0])
	}
	if array.items[1] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[1])
	}
}

func TestPrint(t *testing.T) {
	// Arrange
	array := Array{}
	array.items = []int{3, 4, 5}

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	array.print()

	w.Close()
	out, _ := ioutil.ReadAll(r)

	// Assert
	if !strings.Contains(string(out), strconv.Itoa(3)) {
		t.Errorf("expected %v", 3)
	}
	if !strings.Contains(string(out), strconv.Itoa(4)) {
		t.Errorf("expected %v", 4)
	}
	if !strings.Contains(string(out), strconv.Itoa(5)) {
		t.Errorf("expected %v", 5)
	}
}

func TestSwapBelowZero(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(2)
	array.insert(3)
	array.insert(4)

	// Act
	array.swap(-1, 1)

	// Assert
	if array.count != 2 {
		t.Errorf("expected %v, got %v", 2, array.count)
	}
	if array.size != 2 {
		t.Errorf("expected %v, got %v", 2, array.size)
	}
	if array.items[0] != 3 {
		t.Errorf("expected %v, got %v", 3, array.items[0])
	}
	if array.items[1] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[1])
	}
}

func TestSwapUpperLengh(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(2)
	array.insert(3)
	array.insert(4)

	// Act
	array.swap(1, 2)

	// Assert
	if array.count != 2 {
		t.Errorf("expected %v, got %v", 2, array.count)
	}
	if array.size != 2 {
		t.Errorf("expected %v, got %v", 2, array.size)
	}
	if array.items[0] != 3 {
		t.Errorf("expected %v, got %v", 3, array.items[0])
	}
	if array.items[1] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[1])
	}
}

func TestSwap(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(3)
	array.insert(3)
	array.insert(4)
	array.insert(5)

	// Act
	array.swap(0, 2)

	// Assert
	if array.count != 3 {
		t.Errorf("expected %v, got %v", 3, array.count)
	}
	if array.size != 3 {
		t.Errorf("expected %v, got %v", 3, array.size)
	}
	if array.items[0] != 5 {
		t.Errorf("expected %v, got %v", 5, array.items[0])
	}
	if array.items[1] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[1])
	}
	if array.items[2] != 3 {
		t.Errorf("expected %v, got %v", 3, array.items[2])
	}
}

func TestInvertArrayWithOddLengh(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(3)
	array.insert(3)
	array.insert(4)
	array.insert(5)

	// Act
	array.invert()

	// Assert
	if array.count != 3 {
		t.Errorf("expected %v, got %v", 3, array.count)
	}
	if array.size != 3 {
		t.Errorf("expected %v, got %v", 3, array.size)
	}
	if array.items[0] != 5 {
		t.Errorf("expected %v, got %v", 5, array.items[0])
	}
	if array.items[1] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[1])
	}
	if array.items[2] != 3 {
		t.Errorf("expected %v, got %v", 3, array.items[2])
	}
}

func TestInvertArrayWithEvenLengh(t *testing.T) {
	// Arrange
	array := Array{}
	array.init(4)
	array.insert(3)
	array.insert(4)
	array.insert(5)
	array.insert(6)

	// Act
	array.invert()

	// Assert
	if array.count != 4 {
		t.Errorf("expected %v, got %v", 4, array.count)
	}
	if array.size != 4 {
		t.Errorf("expected %v, got %v", 4, array.size)
	}
	if array.items[0] != 6 {
		t.Errorf("expected %v, got %v", 6, array.items[0])
	}
	if array.items[1] != 5 {
		t.Errorf("expected %v, got %v", 5, array.items[1])
	}
	if array.items[2] != 4 {
		t.Errorf("expected %v, got %v", 4, array.items[2])
	}
	if array.items[3] != 3 {
		t.Errorf("expected %v, got %v", 3, array.items[3])
	}
}
