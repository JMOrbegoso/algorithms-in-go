package number_of_islands

import (
	"testing"
)

func TestCreateGrid(t *testing.T) {
	// Arrange

	// Act
	output := createGrid(2, 3)

	// Assert
	if len(output) != 2 {
		t.Errorf("expected %v, got %v", 2, len(output))
	}
	if len(output[0]) != 3 {
		t.Errorf("expected %v, got %v", 3, len(output[0]))
	}
	if len(output[1]) != 3 {
		t.Errorf("expected %v, got %v", 3, len(output[1]))
	}

	if output[0][0] != "W" {
		t.Errorf("expected %v, got %v", "W", output[0][0])
	}
	if output[0][1] != "W" {
		t.Errorf("expected %v, got %v", "W", output[0][1])
	}
	if output[0][2] != "W" {
		t.Errorf("expected %v, got %v", "W", output[0][2])
	}

	if output[1][0] != "W" {
		t.Errorf("expected %v, got %v", "W", output[1][0])
	}
	if output[1][1] != "W" {
		t.Errorf("expected %v, got %v", "W", output[1][1])
	}
	if output[1][2] != "W" {
		t.Errorf("expected %v, got %v", "W", output[1][2])
	}
}

func TestGetNumberOfIslands(t *testing.T) {
	// Arrange
	grid := createGrid(6, 6)

	grid[0][1] = "L"
	grid[1][0] = "L"
	grid[1][1] = "L"
	grid[2][1] = "L"

	grid[0][4] = "L"
	grid[1][4] = "L"

	grid[4][1] = "L"

	grid[3][3] = "L"
	grid[3][4] = "L"
	grid[4][3] = "L"
	grid[4][4] = "L"

	// Act
	output := getNumberOfIslands(grid)

	// Assert
	if output != 4 {
		t.Errorf("expected %v, got %v", 4, output)
	}
}

func TestCoordinateWasVisited(t *testing.T) {
	// Arrange
	visited := [...]string{"1,2", "2,1"}

	// Act
	output1 := coordinateWasVisited(visited[:], "1,2")
	output2 := coordinateWasVisited(visited[:], "1,3")

	// Assert
	if output1 != true {
		t.Errorf("expected %v, got %v", true, output1)
	}
	if output2 != false {
		t.Errorf("expected %v, got %v", false, output2)
	}
}

func TestExploreAdjacents(t *testing.T) {
	// Arrange
	grid := createGrid(3, 3)
	var visited []string

	grid[0][1] = "L"
	grid[1][0] = "L"
	grid[1][1] = "L"
	grid[2][1] = "L"

	// Act
	exploreAdjacents(grid, 0, 1, &visited)

	// Assert
	if len(visited) != 4 {
		t.Errorf("expected %v, got %v", 4, len(visited))
	}
	if visited[0] != "0,1" {
		t.Errorf("expected %v, got %v", "0,1", visited[0])
	}
	if visited[1] != "1,1" {
		t.Errorf("expected %v, got %v", "1,1", visited[1])
	}
	if visited[2] != "2,1" {
		t.Errorf("expected %v, got %v", "2,1", visited[2])
	}
	if visited[3] != "1,0" {
		t.Errorf("expected %v, got %v", "1,0", visited[3])
	}
}
