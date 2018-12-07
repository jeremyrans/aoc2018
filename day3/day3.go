package day3

import (
	"strconv"
	"strings"
)

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func part1(input string) int {
	grid := make(map[int]map[int]int)
	overlapping := 0
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		coords := coordsFromClaim(parseClaim(l))
		for _, coord := range coords {
			_, ok := grid[coord[0]]
			if !ok {
				grid[coord[0]] = map[int]int{}
			}
			grid[coord[0]][coord[1]]++
			if grid[coord[0]][coord[1]] == 2 {
				overlapping++
			}
		}
	}
	return overlapping
}

// Example input
// #1 @ 1,3: 4x4
func parseClaim(s string) claim {
	c := claim{}
	parts := strings.Split(s, " ")

	c.id, _ = strconv.Atoi(parts[0][1:])

	coords := strings.Split(parts[2], ",")
	c.x, _ = strconv.Atoi(coords[0])
	c.y, _ = strconv.Atoi(strings.TrimSuffix(coords[1], ":"))

	size := strings.Split(parts[3], "x")
	c.width, _ = strconv.Atoi(size[0])
	c.height, _ = strconv.Atoi(size[1])

	return c
}

func coordsFromClaim(c claim) [][]int {
	var result [][]int
	for y := c.y; y < c.y+c.height; y++ {
		for x := c.x; x < c.x+c.width; x++ {
			result = append(result, []int{y, x})
		}
	}
	return result
}

func part2(input string) int {
	grid := make(map[int]map[int]int)
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		c := parseClaim(l)
		coords := coordsFromClaim(c)
		for _, coord := range coords {
			_, ok := grid[coord[0]]
			if !ok {
				grid[coord[0]] = map[int]int{}
			}

			_, ok = grid[coord[0]][coord[1]]
			if !ok {
				// if cell unoccupied, claim it
				grid[coord[0]][coord[1]] = c.id
			} else {
				// if cell occupied, mark it as invalid
				grid[coord[0]][coord[1]] = -1
			}
		}
	}

	// count number of squares solely occupied by each claim
	claimCounts := map[int]int{}
	for _, row := range grid {
		for _, claimID := range row {
			if claimID != -1 {
				claimCounts[claimID]++
			}
		}
	}

	// if size of claim equals number of squares it solely occupies, winner winner
	for _, l := range lines {
		c := parseClaim(l)
		if c.width*c.height == claimCounts[c.id] {
			return c.id
		}
	}

	return -1
}
