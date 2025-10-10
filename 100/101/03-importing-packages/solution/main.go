package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "  go is pretty easy  "

	fmt.Println(strings.TrimSpace(name))
	fmt.Println(removeSpaces(name))
}

// Don't do this, use the standard library...
func removeSpaces(s string) string {
	var start int
	for i, char := range s {
		if char != ' ' {
			start = i
			break
		}
	}
	var end int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i + 1
			break
		}
	}
	// Strings are slices of bytes in Go.
	// If we're dealing with multi-byte characters this could break.
	// But for spaces it works fine.
	// You can use [start:end] to get a substring.
	return s[start:end]
}
