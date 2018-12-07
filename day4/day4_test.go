package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_parseEntry(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected entry
	}

	testCases := []testCase{
		{
			name:  "start",
			input: "[1518-11-01 00:00] Guard #10 begins shift",
			expected: entry{
				action:  START,
				t:       time.Date(1518, 11, 1, 0, 0, 0, 0, time.UTC),
				guardID: 10,
			},
		},
		{
			name:  "sleep",
			input: "[1518-11-01 00:05] falls asleep",
			expected: entry{
				action: SLEEP,
				t:      time.Date(1518, 11, 1, 0, 5, 0, 0, time.UTC),
			},
		},
		{
			name:  "wake",
			input: "[1518-11-01 00:25] wakes up",
			expected: entry{
				action: WAKE,
				t:      time.Date(1518, 11, 1, 0, 25, 0, 0, time.UTC),
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, parseEntry(c.input))
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
			expected: 240,
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
			input:    SAMPLEINPUT1,
			expected: 4455,
		},
		{
			name:     "actual input",
			input:    INPUT,
			expected: 141071,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, part2(c.input))
		})
	}
}
