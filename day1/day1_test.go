package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part1(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name:     "sample input 1",
			input:    "+1\n-2\n+3\n1",
			expected: 3,
		},
		{
			name:     "sample input 2",
			input:    "+1\n+1\n+1",
			expected: 3,
		},
		{
			name:     "sample input 3",
			input:    "+1\n+1\n-2",
			expected: 0,
		},
		{
			name:     "sample input 4",
			input:    "-1\n-2\n-3",
			expected: -6,
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: 540,
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
			input:    "+1\n-1",
			expected: 0,
		},
		{
			name:     "sample input 2",
			input:    "+3\n+3\n+4\n-2\n-4",
			expected: 10,
		},
		{
			name:     "sample input 3",
			input:    "-6\n+3\n+8\n+5\n-6",
			expected: 5,
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: 73056,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, part2(c.input))
		})
	}
}
