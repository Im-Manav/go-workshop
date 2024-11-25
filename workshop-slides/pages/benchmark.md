---
layout: center
---

# Benchmarking

---

# So what is Benchmarking? 

Simply put, it is a technique for measuring the performance of code to identify bottlenecks and improve efficiency.

---

# But how do we do that in Go?

---

With the standard library testing package of course!

---

# The problem

Two developers are arguing whether it's faster to use a map or a slice to search through a list of movies.

---

# Developer's approach

Using a map 

```go{all|2-5|7-10|12-16}{lines:true}
func SearchFilmMap(films map[string]string, title string) (Film, error) {
	v, ok := films[title]
	if !ok {
		return Film{}, ErrFilmNotFound
	}

	r := strings.NewReplacer("-", "", ":", "")
	title = r.Replace(title)

	title = strings.ToLower(title)

	f := Film{
		Title:   title,
		Extract: v,
	}
	return f, nil
}
```
---

# Developer's approach
Using a slice

```go{all|2-3|4-7|8|11}{lines:true}
func SearchFilmSlice(films []Film, title string) (Film, error) {
	for _, film := range films {
		if film.Title == title {
			r := strings.NewReplacer("-", "", ":", "")
			film.Title = r.Replace(film.Title)

			film.Title = strings.ToLower(film.Title)
			return film, nil
		}
	}
	return Film{}, ErrFilmNotFound
}
```
---

```go{all|3-11|8|13-21|23-24|25-26}{lines:true}
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
```

---
layout: center
---

# Let's get benchmarking!

---
layout: center
---

To run benchmark tests we will have to run the following command
```sh
go test -bench=.
```
---
```sh{all|1|2-5|6-11}
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/a-h/flake-templates/go/benchmarking
cpu: AMD Ryzen 5 PRO 4650U with Radeon Graphics
BenchmarkMapExistingFilm-12               240838              4809 ns/op            6840 B/op          7 allocs/op
BenchmarkMapMissingFilm-12              98418349                12.09 ns/op            0 B/op          0 allocs/op
BenchmarkSliceExistingFilm-12             138393              8689 ns/op            6843 B/op          7 allocs/op
BenchmarkSliceMissingFilm-12              512241              2347 ns/op               3 B/op          0 allocs/op
PASS
ok      github.com/a-h/flake-templates/go/benchmarking  7.560s
```
