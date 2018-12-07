package day1

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	frequency := 0
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		d, _ := strconv.Atoi(l)
		frequency += d
	}
	return frequency
}

func part2(input string) int {
	frequency := 0
	seen := map[int]interface{}{0: nil}
	lines := strings.Split(input, "\n")
	for {
		for _, l := range lines {
			d, _ := strconv.Atoi(l)
			frequency += d
			_, ok := seen[frequency]
			if ok {
				return frequency
			} else {
				seen[frequency] = nil
			}
		}
	}
}
