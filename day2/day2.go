package day2

import "strings"

func part1(input string) int {
	twos := 0
	threes := 0
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		if hasCountOfCharacter(l, 2) {
			twos++
		}
		if hasCountOfCharacter(l, 3) {
			threes++
		}
	}
	return twos * threes
}

func hasCountOfCharacter(s string, count int) bool {
	for _, c := range s {
		if strings.Count(s, string(c)) == count {
			return true
		}
	}
	return false
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	for i, s1 := range lines {
		for j, s2 := range lines {
			if i == j {
				continue
			}
			common := getCommonCharacters(s1, s2)
			if len(common) == len(s1)-1 {
				return common
			}
		}
	}
	return ""
}

func getCommonCharacters(a string, b string) string {
	var chars []string
	for i, c := range a {
		if string(b[i]) == string(c) {
			chars = append(chars, string(c))
		}
	}
	return strings.Join(chars, "")
}
