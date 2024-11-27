package fuzzing

import "regexp"

// Match the path /search/{entity}/{term}
var r = regexp.MustCompile(`/search/([^/]+)/([^/]+)`)

func Parse(path string) (entity, term string, err error) {
	results := r.FindStringSubmatch(path)
	return results[1], results[2], nil
}
