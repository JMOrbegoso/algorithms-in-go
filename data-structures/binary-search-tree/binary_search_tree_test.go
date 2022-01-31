package binary_search_tree

import (
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
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 8}
	leaf3 := Node{value: 25}
	parent1 := Node{value: 5, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 20, right: &leaf3}
	root := Node{value: 10, left: &parent1, right: &parent2}

	// Act
	output := root.include(10)

	// Assert
	if output != true {
		t.Errorf("expected %v, got %v", true, output)
	}
}

func TestInsert(t *testing.T) {
	// Arrange
	leaf1 := Node{value: 4}
	leaf2 := Node{value: 8}
	parent1 := Node{value: 5, left: &leaf1, right: &leaf2}
	parent2 := Node{value: 20}
	root := Node{value: 10, left: &parent1, right: &parent2}

	// Act
	root.insert(25)

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
	if parent2.right.value != 25 {
		t.Errorf("expected %v, got %v", 25, parent2.right.value)
	}
}
