# Benchmarking

Your task is to find out whether a map or a slice is faster searching through a list of movies.

After running the benchmark you will see something like this.

```
goos: linux
goarch: amd64
pkg: github.com/a-h/go-workshop/200/benchmarking
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkMapExistingFilm-16               863116              1392 ns/op
BenchmarkMapMissingFilm-16              157165226                7.919 ns/op
BenchmarkSliceExistingFilm-16             369318              3259 ns/op
BenchmarkSliceMissingFilm-16              579453              2021 ns/op
PASS
ok      github.com/a-h/go-workshop/200/benchmarking     7.782s
```

- `goos` the OS the test is running on
- `goarch` the CPU architecture the test is running on
- `pkg` the package that the tests are part of
- `cpu` the cpu model

`BenchmarkMapExistingFilm-16` the name of the Benchmark test with a `-16` suffix to denote the number of CPUs (cores) used to run the benchmark (specified by GOMAXPROCS).

On the right side they are two values `863116` indicates the total number of times the function ran and `1392 ns/op` is the average amount of time each iteration took to complete expressed in nanoseconds per operation.

It's possible to improve the performance of this code.

## Learning points

- Which method is faster? Slice or Map? Why is that?
- Run `benchmem` task, what do you notice?
- Run the `profile` task then choose `cpu-web` or `mem-web` to see the profiles. What do you notice? Tip: Click on View at the top left and choose Flamegraph(New) or any other option you prefer!
- How can we optimise even further?

## Tasks

### bench

Passing a `-bench` to run benchmark tests, the flag accepts valid regex. To run all benchmarks use `-bench=.` to run specific tests pass `-bench=BenchmarkMap*` 

```bash
go test -bench=.
```

### benchmem

To get an output of the memory stats, you can pass an addition flag `-benchmem`. This tells you how many bytes were allocated and how many allocations were made per operation.

In modern computers memory operations that go beyond the CPU cache are expensive, so it's important to keep an eye on memory allocations and reduce them, if you need high performance code.

```bash
go test -bench=. -benchmem
```

### profile

Go provides built-in support for profiling CPU and memory usage of your code. You can create output files that can be visualised in different ways.

The `go test` command accepts two flags for profiling:

 - `-cpuprofile` enable and name of the file created for the cpu profile
 - `-memprofile` enable and name of the file creeated for the memory profile

And it will output `cpu.prof` and `mem.prof` files.

```bash
go test -cpuprofile cpu.prof -memprofile mem.prof -bench . -benchmem
```

### cpu-web

One way is to explore the profile in a web browser. The following command will start a web server on port 8000.

The bigger the box is, the more time was spent in that function, so that tells you where to focus your optimisation efforts.

```bash
go tool pprof -http=":8000" cpu.prof
```

### mem-web

The memory profile can also be visualised in a web browser. The following command will start a web server on port 8001.

Maybe it's obvious where the memory allocations are happening now?

```bash
go tool pprof -http=":8001" mem.prof
```
