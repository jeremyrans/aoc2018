package day5

import "strings"

func part1(input string) int {
	return len(react(input))
}

func react(polymer string) string {
	for i := 0; i < len(polymer)-1; i++ {
		// 32 characters between capital and lowercase letters
		if polymer[i]-polymer[i+1] == 32 || polymer[i+1]-polymer[i] == 32 {
			return react(polymer[0:i] + polymer[i+2:])
		}
	}
	return polymer
}

func part2(input string) int {
	best := 999999
	for i := 65; i <= 90; i++ { // A-Z
		newInput := strings.Replace(input, string(i), "", -1)      // capital
		newInput = strings.Replace(newInput, string(i+32), "", -1) // lowercase
		result := react(newInput)
		if len(result) < best {
			best = len(result)
		}
	}
	return best
}
