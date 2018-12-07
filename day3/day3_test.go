package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseClaim(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected claim
	}

	testCases := []testCase{
		{
			name:  "sample claim",
			input: "#1 @ 1,3: 4x4",
			expected: claim{
				id:     1,
				x:      1,
				y:      3,
				width:  4,
				height: 4,
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, parseClaim(c.input))
		})
	}
}

func Test_coordsFromClaim(t *testing.T) {
	type testCase struct {
		name     string
		input    claim
		expected [][]int
	}

	testCases := []testCase{
		{
			name: "sample claim",
			input: claim{
				id:     1,
				y:      3,
				x:      1,
				width:  4,
				height: 4,
			},
			expected: [][]int{
				{3, 1}, {3, 2}, {3, 3}, {3, 4},
				{4, 1}, {4, 2}, {4, 3}, {4, 4},
				{5, 1}, {5, 2}, {5, 3}, {5, 4},
				{6, 1}, {6, 2}, {6, 3}, {6, 4},
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, coordsFromClaim(c.input))
		})
	}
}

func Test_part1(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name:     "sample input 1",
			input:    SAMPLEINPUT1,
			expected: 4,
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: 116140,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, part1(c.input))
		})
	}
}

func Test_part2(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name:     "sample input 1",
			input:    SAMPLEINPUT1,
			expected: 3,
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: 574,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, part2(c.input))
		})
	}
}
