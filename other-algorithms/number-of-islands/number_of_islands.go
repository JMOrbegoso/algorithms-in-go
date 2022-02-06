package number_of_islands

import (
	"fmt"
	"strconv"
)

func printGrid(grid [][]string) {
	for row := range grid {
		for col := range grid[row] {
			fmt.Print("|")
			fmt.Print(grid[row][col])
		}
		fmt.Println()
	}
}

func createGrid(rows, columns int) [][]string {
	grid := make([][]string, rows)

	for row := 0; row < len(grid); row++ {
		grid[row] = make([]string, columns)

		for col := 0; col < len(grid[row]); col++ {
			grid[row][col] = "W"
		}
	}

	return grid
}

func getNumberOfIslands(grid [][]string) int {
	islandsCount := 0
	var visited []string

	for row := range grid {
		for col := range grid[row] {

			if grid[row][col] != "L" {
				continue
			}

			coordinate := strconv.Itoa(row) + "," + strconv.Itoa(col)

			if coordinateWasVisited(visited, coordinate) {
				continue
			}

			islandsCount++

			exploreAdjacents(grid, row, col, &visited)
		}
	}

	return islandsCount
}

func coordinateWasVisited(visited []string, coordinate string) bool {
	for _, item := range visited {
		if item == coordinate {
			return true
		}
	}
	return false
}

func exploreAdjacents(grid [][]string, row int, col int, visited *[]string) {
	// Validate
	if row < 0 {
		return
	}
	if col < 0 {
		return
	}
	if row >= len(grid) {
		return
	}
	if col >= len(grid[0]) {
		return
	}

	if grid[row][col] != "L" {
		return
	}

	coordinate := strconv.Itoa(row) + "," + strconv.Itoa(col)

	if coordinateWasVisited(*visited, coordinate) {
		return
	}

	*visited = append(*visited, coordinate)

	exploreAdjacents(grid, row-1, col, visited)
	exploreAdjacents(grid, row+1, col, visited)
	exploreAdjacents(grid, row, col-1, visited)
	exploreAdjacents(grid, row, col+1, visited)
}
