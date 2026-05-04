package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Test",
			expected: []string{"Test"},
		},
		{
			input: " trim me ",
			expected: []string{"trim", "me"},
		},
		{
			input: "extra        spaces",
			expected: []string{"extra", "spaces"},
		},
		{
			input: " ",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("number of words doesn't match number of expected words")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("the word %s doesn't match expected word %s", word, expectedWord)
			}
		}
	}
}
