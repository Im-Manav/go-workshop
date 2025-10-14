package benchmarking

import (
	_ "embed"
	"encoding/json"
	"errors"
	"strings"
)

//go:embed movies-1990s.json
var jsonFile []byte

type Film struct {
	Title   string
	Extract string
}

var ErrFilmNotFound = errors.New("film not found")

func NewFilmSlice() (films []Film, err error) {
	err = json.Unmarshal(jsonFile, &films)
	return films, err
}

func NewFilmMap() (map[string]string, error) {
	s, err := NewFilmSlice()
	if err != nil {
		return nil, err
	}
	films := make(map[string]string)
	for _, film := range s {
		films[film.Title] = film.Extract
	}
	return films, nil
}

// Precompute the strings.Replacer to avoid doing it on every search.
var punctuationRemover = strings.NewReplacer("-", "", ":", "")

func normaliseTitle(title string) string {
	return strings.ToLower(punctuationRemover.Replace(title))
}

func SearchFilmMap(films map[string]string, title string) (Film, error) {
	v, ok := films[title]
	if !ok {
		return Film{}, ErrFilmNotFound
	}
	f := Film{
		Title:   normaliseTitle(title),
		Extract: v,
	}
	return f, nil
}

func SearchFilmSlice(films []Film, title string) (Film, error) {
	for _, film := range films {
		if film.Title == title {
			film.Title = normaliseTitle(film.Title)
			return film, nil
		}
	}
	return Film{}, ErrFilmNotFound
}
