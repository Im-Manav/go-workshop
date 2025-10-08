package benchmarking

import "testing"

func benchmarkFilmMap(title string, b *testing.B) {
	m, err := NewFilmMap()
	if err != nil {
		b.Error(err)
	}
	for n := 0; n < b.N; n++ {
		SearchFilmMap(m, title)
	}
}

func benchmarkFilmSlice(title string, b *testing.B) {
	s, err := NewFilmSlice()
	if err != nil {
		b.Error(err)
	}
	for n := 0; n < b.N; n++ {
		SearchFilmSlice(s, title)
	}
}

func BenchmarkMapExistingFilm(b *testing.B)   { benchmarkFilmMap("Star Wars: Episode I – The Phantom Menace", b) }
func BenchmarkMapMissingFilm(b *testing.B)    { benchmarkFilmMap("Spider-Man 3", b) }
func BenchmarkSliceExistingFilm(b *testing.B) { benchmarkFilmSlice("Star Wars: Episode I – The Phantom Menace", b) }
func BenchmarkSliceMissingFilm(b *testing.B)  { benchmarkFilmSlice("Spider-Man 3", b) }
