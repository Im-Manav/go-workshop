package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{
			input:    0,
			expected: "0",
		},
		{
			input:    1,
			expected: "1",
		},
		{
			input:    2,
			expected: "2",
		},
		{
			input:    3,
			expected: "Fizz",
		},
		{
			input:    4,
			expected: "4",
		},
		{
			input:    5,
			expected: "Buzz",
		},
		{
			input:    6,
			expected: "Fizz",
		},
		{
			input:    7,
			expected: "7",
		},
		{
			input:    8,
			expected: "8",
		},
		// It's possible to populate struct fields by omitting the name.
		// In this case, the fields are populated with the values in the order they appear.
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input=%d", test.input), func(t *testing.T) {
			actual := Check(test.input)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
