package day6

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
			input:    SAMPLEINPUT,
			expected: 17,
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: 4186,
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
		name        string
		input       string
		maxDistance int
		expected    int
	}

	testCases := []testCase{
		{
			name:        "sample input 1",
			input:       SAMPLEINPUT,
			maxDistance: 32,
			expected:    16,
		},
		{
			name:        "actual input",
			input:       INPUT,
			maxDistance: 10000,
			expected:    12,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, part2(c.input, c.maxDistance))
		})
	}
}
