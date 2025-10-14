package benchmarking

import (
	_ "embed"
	"encoding/json"
	"errors"
	"slices"
	"strings"
)

//go:embed movies-1990s.json
var jsonFile []byte

type Film struct {
	NormalisedTitle string
	Title           string
	Extract         string
}

var ErrFilmNotFound = errors.New("film not found")

func NewFilmSlice() (films []Film, err error) {
	err = json.Unmarshal(jsonFile, &films)
	if err != nil {
		return nil, err
	}
	// Pre-normalise titles to reduce work after returning results.
	for i := range films {
		films[i].NormalisedTitle = strings.ToLower(punctuationRemover.Replace(films[i].Title))
	}
	// Pre-sort the slice to enable binary searching.
	slices.SortFunc(films, func(a, b Film) int {
		return strings.Compare(a.Title, b.Title)
	})
	return films, err
}

// Precompute the strings.Replacer.
var punctuationRemover = strings.NewReplacer("-", "", ":", "")

func NewFilmMap() (map[string]Film, error) {
	s, err := NewFilmSlice()
	if err != nil {
		return nil, err
	}
	filmTitleToFilm := make(map[string]Film)
	for _, film := range s {
		filmTitleToFilm[film.Title] = film
	}
	return filmTitleToFilm, nil
}

func SearchFilmMap(films map[string]Film, title string) (Film, error) {
	// Since we pre-computed the normalised titles, we can just do a direct
	// lookup in the map without needing to normalise the title.
	v, ok := films[title]
	if !ok {
		return Film{}, ErrFilmNotFound
	}
	return v, nil
}

func SearchFilmSlice(films []Film, title string) (Film, error) {
	// Since we've pre-sorted the slice, we can use binary search which is much
	// faster than linear search.
	index, ok := slices.BinarySearchFunc(films, title, func(a Film, b string) int {
		return strings.Compare(a.Title, b)
	})
	if !ok {
		return Film{}, ErrFilmNotFound
	}
	return films[index], nil
}
