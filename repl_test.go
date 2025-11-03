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
		{
			input:    "CHARMANDER BuLbAsAuR squirTLE",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:    "  HELLO      charMANDER   ",
			expected: []string{"hello", "charmander"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Number of words did not match expected output")
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("Word mismatch")
				}
			}
		}
	}
}
