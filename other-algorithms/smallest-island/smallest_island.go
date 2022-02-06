package smallest_island

import (
	"fmt"
	"math"
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

func getSmallestIsland(grid [][]string) int {
	smallestIsland := math.MaxInt
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

			islandSize := exploreIsland(grid, row, col, &visited)

			if islandSize < smallestIsland {
				smallestIsland = islandSize
			}
		}
	}

	if len(visited) == 0 {
		return 0
	}

	return smallestIsland
}

func coordinateWasVisited(visited []string, coordinate string) bool {
	for _, item := range visited {
		if item == coordinate {
			return true
		}
	}
	return false
}

func exploreIsland(grid [][]string, row int, col int, visited *[]string) int {
	if row < 0 {
		return 0
	}
	if col < 0 {
		return 0
	}
	if row >= len(grid) {
		return 0
	}
	if col >= len(grid[0]) {
		return 0
	}

	if grid[row][col] != "L" {
		return 0
	}

	coordinate := strconv.Itoa(row) + "," + strconv.Itoa(col)

	if coordinateWasVisited(*visited, coordinate) {
		return 0
	}

	*visited = append(*visited, coordinate)

	return 1 +
		exploreIsland(grid, row-1, col, visited) +
		exploreIsland(grid, row+1, col, visited) +
		exploreIsland(grid, row, col-1, visited) +
		exploreIsland(grid, row, col+1, visited)
}
