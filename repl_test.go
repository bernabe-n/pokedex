package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Go  is awesome  ",
			expected: []string{"go", "is", "awesome"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// First check if lengths match
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %v (length %d), expected %v (length %d)",
				c.input, actual, len(actual), c.expected, len(c.expected))
			continue
		}

		// Compare each element individually
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q)[%d] == %q, expected %q",
					c.input, i, actual[i], c.expected[i])
			}
		}
	}
}
