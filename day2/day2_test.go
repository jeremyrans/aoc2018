package day2

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
			input:    SAMPLEINPUT1,
			expected: 12,
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
		expected string
	}

	testCases := []testCase{
		{
			name:     "sample input 1",
			input:    SAMPLEINPUT2,
			expected: "fgij",
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: "pazvmqbftrbeosiecxlghkwud",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, part2(c.input))
		})
	}
}
