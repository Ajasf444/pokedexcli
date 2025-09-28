package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// TODO: check length of slice
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// TODO: compare each word
		}
	}
}
