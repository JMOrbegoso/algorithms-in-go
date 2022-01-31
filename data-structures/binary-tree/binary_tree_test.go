package binary_tree

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCreation(t *testing.T) {
	// Arrange

	// Act
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 3, right: &leaf3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Assert
	if root.left != &parent1 {
		t.Errorf("expected %v, got %v", parent1, root.left)
	}
	if root.right != &parent2 {
		t.Errorf("expected %v, got %v", parent2, root.right)
	}
	if parent1.left != &leaf1 {
		t.Errorf("expected %v, got %v", leaf1, parent1.left)
	}
	if parent1.right != &leaf2 {
		t.Errorf("expected %v, got %v", leaf2, parent1.right)
	}
	if parent2.left != nil {
		t.Errorf("expected %v, got %v", nil, parent2.left)
	}
	if parent2.right != &leaf3 {
		t.Errorf("expected %v, got %v", leaf3, parent2.right)
	}
}

func TestPrint(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 20}
	leaf2 := Node{value: 10}
	root := Node{value: 30, left: &leaf1, right: &leaf2}

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	root.printByLevel()

	w.Close()
	out, _ := ioutil.ReadAll(r)

	// Assert
	if !strings.Contains(string(out), strconv.Itoa(root.value)) {
		t.Errorf("expected %v", root.value)
	}
	if !strings.Contains(string(out), strconv.Itoa(leaf1.value)) {
		t.Errorf("expected %v", leaf1.value)
	}
	if !strings.Contains(string(out), strconv.Itoa(leaf2.value)) {
		t.Errorf("expected %v", leaf2.value)
	}
}

func TestIncludeNotFound(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 20}
	leaf2 := Node{value: 10}
	root := Node{value: 30, left: &leaf1, right: &leaf2}

	// Act
	output := root.include(100)

	// Assert
	if output != false {
		t.Errorf("expected %v, got %v", false, output)
	}
}

func TestInclude(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 20}
	leaf2 := Node{value: 10}
	root := Node{value: 30, left: &leaf1, right: &leaf2}

	// Act
	output := root.include(30)

	// Assert
	if output != true {
		t.Errorf("expected %v, got %v", true, output)
	}
}

func TestTotalSum(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 3, right: &leaf3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Act
	output := root.totalSum()

	// Assert
	if output != 21 {
		t.Errorf("expected %v, got %v", 21, output)
	}
}

func TestMinValue(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 3, right: &leaf3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Act
	output := root.minValue()

	// Assert
	if output != 1 {
		t.Errorf("expected %v, got %v", 1, output)
	}
}

func TestMaxValue(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 3, right: &leaf3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Act
	output := root.maxValue()

	// Assert
	if output != 6 {
		t.Errorf("expected %v, got %v", 6, output)
	}
}

func TestMaxSumPath(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 3, right: &leaf3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Act
	output := root.maxSumPath(0)

	// Assert
	if output != 10 {
		t.Errorf("expected %v, got %v", 10, output)
	}
}

func TestInvertTree(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 3, right: &leaf3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Act
	root.invertTree()

	// Assert
	if root.left != &parent2 {
		t.Errorf("expected %v, got %v", parent2, root.left)
	}
	if root.right != &parent1 {
		t.Errorf("expected %v, got %v", parent1, root.right)
	}
	if parent2.left != &leaf3 {
		t.Errorf("expected %v, got %v", leaf3, parent2.left)
	}
	if parent2.right != nil {
		t.Errorf("expected %v, got %v", nil, parent2.right)
	}
	if parent1.left != &leaf2 {
		t.Errorf("expected %v, got %v", leaf2, parent1.left)
	}
	if parent1.right != &leaf1 {
		t.Errorf("expected %v, got %v", leaf1, parent1.right)
	}
}

func TestGetMaxDepthOfOneNodeTree(t *testing.T) {
	// Arrange
	root := Node{value: 1}

	// Act
	output := root.getMaxDepth()

	// Assert
	if output != 0 {
		t.Errorf("expected %v, got %v", 0, output)
	}
}

func TestGetMaxDepth(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent3 := Node{value: 3, right: &leaf3}
	parent2 := Node{value: 8, right: &parent3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Act
	output := root.getMaxDepth()

	// Assert
	if output != 3 {
		t.Errorf("expected %v, got %v", 3, output)
	}
}

func TestGetMaxWidthOfOneNodeTree(t *testing.T) {
	// Arrange
	root := Node{value: 1}

	// Act
	output := root.getMaxWidth()

	// Assert
	if output != 0 {
		t.Errorf("expected %v, got %v", 0, output)
	}
}

func TestGetMaxWidth(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent3 := Node{value: 3, right: &leaf3}
	parent2 := Node{value: 8, right: &parent3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Act
	output := root.getMaxWidth()

	// Assert
	if output != 3 {
		t.Errorf("expected %v, got %v", 3, output)
	}
}

func TestInsert(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 5}
	leaf3 := Node{value: 6}
	parent1 := Node{value: 2, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 3, right: &leaf3}
	root := Node{value: 1, left: &parent1, right: &parent2}

	// Act
	root.insert(20)

	// Assert
	if root.left != &parent1 {
		t.Errorf("expected %v, got %v", parent1, root.left)
	}
	if root.right != &parent2 {
		t.Errorf("expected %v, got %v", parent2, root.right)
	}
	if parent1.left != &leaf1 {
		t.Errorf("expected %v, got %v", leaf1, parent1.left)
	}
	if parent1.right != &leaf2 {
		t.Errorf("expected %v, got %v", leaf2, parent1.right)
	}
	if parent2.left.value != 20 {
		t.Errorf("expected %v, got %v", 20, parent2.left.value)
	}
	if parent2.right != &leaf3 {
		t.Errorf("expected %v, got %v", leaf3, parent2.right)
	}
}

func TestInsertTwoTimes(t *testing.T) {
	// Arrange
	root := Node{value: 10}

	// Act
	root.insert(20)
	root.insert(30)

	// Assert
	if root.left.value != 20 {
		t.Errorf("expected %v, got %v", 20, root.left.value)
	}
	if root.right.value != 30 {
		t.Errorf("expected %v, got %v", 30, root.right.value)
	}
	if root.left.left != nil {
		t.Errorf("expected %v, got %v", nil, root.left.left)
	}
	if root.left.right != nil {
		t.Errorf("expected %v, got %v", nil, root.left.right)
	}
	if root.right.left != nil {
		t.Errorf("expected %v, got %v", nil, root.right.left)
	}
	if root.right.right != nil {
		t.Errorf("expected %v, got %v", nil, root.right.right)
	}
}
