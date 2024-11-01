package benchmarking

import "testing"

func benchmarkFilmMap(title string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SearchFilmMap(title)
	}
}

func benchmarkFilmSlice(title string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SearchFilmSlice(title)
	}
}

func BenchmarkMapExistingFilm(b *testing.B) { benchmarkFilmMap("Spider-Man 2", b)}
func BenchmarkMapMissingFilm(b *testing.B) { benchmarkFilmMap("Back to the Future: Part 3", b)}
func BenchmarkSliceExistingFilm(b *testing.B) { benchmarkFilmMap("Spider-Man 2", b)}
func BenchmarkSliceMissingFilm(b *testing.B) { benchmarkFilmMap("Back to the Future: Part 3", b)}
