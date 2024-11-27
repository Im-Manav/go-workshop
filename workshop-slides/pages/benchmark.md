---
layout: center
---

# Benchmarking

---

## So what is Benchmarking? 

A technique for measuring the performance of code to identify bottlenecks and improve efficiency.

---
layout: center
---

## But how do we do that in Go?

---
layout: center
---

With the standard library testing package of course!

---
layout: center
---

In the `go-workshop-102` folder there is a folder called `benchmarking`. We will now go through what files are in there and the exercise we are looking to solve.

```bash
benchmarking
├── README.md
├── benchmarking.go
├── benchmarking_test.go
├── movies-1990s.json
└── solution
    └── benchmarking.go
```

---
layout: full
---

## The problem

Two developers are arguing whether it's faster to use a map or a slice to search through a list of movies.
![Remote Image](https://net-informations.com/faq/general/img/dictionary-vs-list-graph.png)
---

## Developer's approach

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

## Developer's approach
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

## Benchmark - benchmark.go

```go{all|1-2|4-18|12|4-18}
//go:embed movies-1990s.json
var jsonFile []byte

type Film struct {
	Title   string
	Extract string
}

func NewFilmSlice() ([]Film, error) {
	var films []Film

	err := json.Unmarshal(jsonFile, &films)
	if err != nil {
		return nil, err
	}

	return films, nil
}


// search methods
```
---

## Benchmark Test - benchmark_test.go
```go{all|1|3-11|8|12-20|21-22|23-24}{lines:true}
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

## Let's get benchmarking!

---
layout: center
---

To run benchmark tests we will have to run the following command
```sh
go test -bench=.
```
---
layout: center
---

```sh{all|1-4|5-7}
goos: linux
goarch: amd64
pkg: github.com/user/packagename
cpu: Intel(R) Core(TM) i7-7560U CPU @ 2.40GHz
BenchmarkPrimeNumbers-4            14588             82798 ns/op
PASS
ok      github.com/username/packagename     2.091s
```

---

# Visualisations

<img src="/cpu-profile.png" alt="CPU Profile" style="height: 90%;"/>

---
layout: center
---

Open the README.md in `benchamrking/` folder and follow the instructions.

We will spend around 25 minutes.

---

## Tips

- Use benchmarking!
- Use profiling
- Focus on isolated units of code
- Focus on areas with significant performance impact
- Data driven

---
layout: center
---

## Wait there is more!

Profile Guided Optimisation (PGO) starting from Go 1.20

---
layout: center
---

## PGO Cycle

Build and release initial binary (without PGO) 

```go

go build
```
---
layout: center
---

## PGO Cycle

Collect profiles from production using `runtime/pprof` or `net/http/pprof` standard library packages.

---
layout: center
---

## PGO Cycle

On your next release build from the latest source and provide the production profile!

```go
go build -pgo default.pgo
```

---
layout: center
---

## PGO Cycle
1. Build and release an initial binary (without PGO).
2. Collect profiles from production.
3. When it’s time to release an updated binary, build from the latest source and provide the production profile.
4. GOTO 2

More info: https://go.dev/doc/pgo

