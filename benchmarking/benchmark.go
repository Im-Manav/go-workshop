package benchmarking

import (
	"errors"
	"strings"
)

type Film struct {
	title     string
	summary   string
	ageRating string
}

var filmsMap = map[string]Film{
	"theshawshankredemption": {
		title:     "The Shawshank Redemption",
		summary:   "A banker convicted of uxoricide forms a friendship over a quarter century with a hardened convict, while maintaining his innocence and trying to remain hopeful through simple compassion.",
		ageRating: "18+",
	},
	"spiderman2": {
		title:     "Spider-Man 2",
		summary:   "Peter Parker is beset with troubles in his failing personal life as he battles a former brilliant scientist titled Otto Octavius.",
		ageRating: "12+",
	},
	"thedarkknight": {
		title:     "The Dark Knight",
		summary:   "When a menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman, James Gordon and Harvey Dent must work together to put an end to the madness.",
		ageRating: "18+",
	},
	"thelordoftheringsthereturnoftheking": {
		title:     "The Lord of the Rings: The Return of the King",
		summary:   "Gandalf and Aragorn lead the World of Men against Sauron's army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.",
		ageRating: "15+",
	},
	"thegoodthebadandtheugly": {
		title:     "The Good, the Bad and the Ugly",
		summary:   "A bounty hunting scam joins two men in an uneasy alliance against a third in a race to find a fortune in gold buried in a remote cemetery.",
		ageRating: "18+",
	},
}

type FilmSlice struct {
	name      string
	title     string
	summary   string
	ageRating string
}

var filmsSlice = []FilmSlice{
	{
		name:      "theshawshankredemption",
		title:     "The Shawshank Redemption",
		summary:   "A banker convicted of uxoricide forms a friendship over a quarter century with a hardened convict, while maintaining his innocence and trying to remain hopeful through simple compassion.",
		ageRating: "18+",
	},
	{
		name:      "spiderman2",
		title:     "Spider-Man 2",
		summary:   "Peter Parker is beset with troubles in his failing personal life as he battles a former brilliant scientist titled Otto Octavius.",
		ageRating: "12+",
	},
	{
		name:      "thedarkknight",
		title:     "The Dark Knight",
		summary:   "When a menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman, James Gordon and Harvey Dent must work together to put an end to the madness.",
		ageRating: "18+",
	},
	{
		name:      "thelordoftheringsthereturnoftheking",
		title:     "The Lord of the Rings: The Return of the King",
		summary:   "Gandalf and Aragorn lead the World of Men against Sauron's army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.",
		ageRating: "15+",
	},
	{
		name:      "thegoodthebadandtheugly",
		title:     "The Good, the Bad and the Ugly",
		summary:   "A bounty hunting scam joins two men in an uneasy alliance against a third in a race to find a fortune in gold buried in a remote cemetery.",
		ageRating: "18+",
	},
}

func SearchFilmMap(title string) (Film, error) {
	s := strings.TrimSpace(title)
	r := strings.NewReplacer(",", "", ":", "", "'", "", "-", "")
	name := r.Replace(s)

	v, ok := filmsMap[name]
	if !ok {
		return Film{}, errors.New("film doesn't exist")
	}

	return v, nil
}

func SearchFilmSlice(title string) (FilmSlice, error) {
	s := strings.TrimSpace(title)
	r := strings.NewReplacer(",", "", ":", "", "'", "", "-", "")
	name := r.Replace(s)

	for i := 0; i < len(filmsSlice); i++ {
		if filmsSlice[i].name == name {
			return filmsSlice[i], nil
		}
	}
	return FilmSlice{}, errors.New("film doesn't exist")
}
