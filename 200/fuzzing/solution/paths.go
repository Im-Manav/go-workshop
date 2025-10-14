package fuzzing

import (
	"strings"
)

// Parse the path /search/{entity}/{term}
func Parse(path string) (entity, term string, ok bool) {
	if !strings.HasPrefix(path, "/search/") {
		return "", "", false
	}
	segments := strings.Split(path, "/")
	if len(segments) != 4 || segments[1] != "search" {
		return "", "", false
	}
	// Really, we might want to validate entity and term further.
	// For example, checking length, allowed characters, etc.
	// We may also need to URL decode the entity and term.
	entity = segments[2]
	term = segments[3]
	return entity, term, true
}
