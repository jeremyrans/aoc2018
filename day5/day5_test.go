package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_react(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	testCases := []testCase{
		{
			name:     "single reaction",
			input:    "aAbb",
			expected: "bb",
		},
		{
			name:     "compound reaction",
			input:    "ccaBbA",
			expected: "cc",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, react(c.input))
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
			input:    "dabAcCaCBAcCcaDA",
			expected: 10,
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: 12,
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
			input:    "dabAcCaCBAcCcaDA",
			expected: 4,
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: 6872,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, part2(c.input))
		})
	}
}
