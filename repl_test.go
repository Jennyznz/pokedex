package main 

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "I baked an apple pie. .",
			expected: []string{"i", "baked", "an", "apple", "pie.", "."},
		},
		{
			input: "WHat is a CHEESECAKE",
			expected: []string{"what", "is", "a", "cheesecake"},
		},
		{
			input: "I LOST MY PHONE!!!!",
			expected: []string{"i", "lost", "my", "phone!!!!"},
		},
	}

	for _, c := range cases {
		output := cleanInput(c.input)
		if len(output) != len(c.expected) {
			t.Errorf("Output is not the correct length.")
		} else {
			for i := range output {
				word := output[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("Unexpected output")
				}
			}
		}
	}
}

