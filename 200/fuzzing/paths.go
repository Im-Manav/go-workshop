package fuzzing

import "regexp"

// Match the path /search/{entity}/{term}
var r = regexp.MustCompile(`/search/([^/]+)/([^/]+)`)

func Parse(path string) (entity, term string, ok bool) {
	results := r.FindStringSubmatch(path)
	return results[1], results[2], true
}
