package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "howdy abigail you're just the cutest thing that ever lived! ",
			expected: []string{"howdy", "abigail", "you're", "just", "the", "cutest",
				"thing", "that", "ever", "lived!"},
		},
		{
			input:    "This is the best day of my ENTIRE Life!!!",
			expected: []string{"this", "is", "the", "best", "day", "of", "my", "entire", "life!!!"},
		},
	}

	passed := 0
	failed := 0
	for j, c := range testCases {
		passing := true
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("TestCleanInput case# %d failed, actual length does not equal expected length", j+1)
			failed += 1
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				passing = false
				t.Errorf("TestCleanInput case# %d failed, actual length does not equal expected length", j+1)
				break
			}
		}
		if passing {
			passed += 1
		} else {
			failed += 1
		}
	}
}
