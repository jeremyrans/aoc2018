package day6

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) int {
	grid := map[int]map[int]int{}
	lines := strings.Split(input, "\n")
	coords := map[int][]int{}
	candidates := map[int]interface{}{}
	width := 0
	height := 0
	for i, l := range lines {
		parts := strings.Split(l, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		coords[i] = []int{x, y}
		_, ok := grid[y]
		if !ok {
			grid[y] = map[int]int{}
		}
		grid[y][x] = i
		candidates[i] = nil

		if x > width {
			width = x
		}
		if y > height {
			height = y
		}
	}

	// because x and y were 0 indexed but height/width aren't
	width++
	height++

	// populate each grid square with the coordinate it's closest to
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			_, ok := grid[i]
			if !ok {
				grid[i] = map[int]int{}
			}
			grid[i][j] = findClosestCoord(j, i, coords)

			// if coord is touching an edge, it's no longer a candidate
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				delete(candidates, grid[i][j])
			}
		}
	}

	// for each candidate, which has the most squares occupied?
	bestCoordCount := 0
	for c := range candidates {
		count := 0
		for _, row := range grid {
			for _, cell := range row {
				if cell == c {
					count++
				}
			}
		}
		if count > bestCoordCount {
			bestCoordCount = count
		}
	}

	return bestCoordCount
}

func findClosestCoord(x, y int, coords map[int][]int) int {
	closestCoord := -1
	closestCoordDistance := math.MaxInt8
	for i, c := range coords {
		d := calculateManhattanDistance(x, y, c[0], c[1])
		if d < closestCoordDistance {
			closestCoordDistance = d
			closestCoord = i
		} else if d == closestCoordDistance {
			closestCoord = -1
		}
	}
	return closestCoord
}

func calculateManhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func part2(input string, maxDistance int) int {
	grid := map[int]map[int]int{}
	lines := strings.Split(input, "\n")
	coords := map[int][]int{}
	candidates := map[int]interface{}{}
	width := 0
	height := 0
	for i, l := range lines {
		parts := strings.Split(l, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		coords[i] = []int{x, y}
		_, ok := grid[y]
		if !ok {
			grid[y] = map[int]int{}
		}
		grid[y][x] = i
		candidates[i] = nil

		if x > width {
			width = x
		}
		if y > height {
			height = y
		}
	}

	// because x and y were 0 indexed but height/width aren't
	width++
	height++

	regionSize := 0
	// consider if each grid square is < maxDistance total manhattan distance from all coords
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			totalDistance := 0
			for _, c := range coords {
				totalDistance += calculateManhattanDistance(j, i, c[0], c[1])
			}
			if totalDistance < maxDistance {
				regionSize++
			}
		}
	}

	return regionSize
}
