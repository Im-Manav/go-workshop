# Benchmarking

Your task is to find out whether a map or a slice is faster searching through a list of movies. To run the benchmark test run the command in task, open a terminal cd to `benchmarking/` and run the `bench` task. 

By running the benchmark you will see something like this (this is an example)

```
goos: linux
goarch: amd64
pkg: github.com
cpu: Intel(R) Core(TM) i7-7560U CPU @ 2.40GHz
BenchmarkPrimeNumbers-4            14588             82798 ns/op
PASS
ok      github.com/username/packagename     2.091s
```
`goos` the OS the test is running on
`goarch` the CPU architecture the test is running on
`pkg` the package that the tests are part of
`cpu` the cpu model

`BenchmarkPrimeNumbers-4` the name of the Benchmark test with a `-4` suffix to denote the number of CPUs (cores) used to run the benchmark (specified by GOMAXPROCS)
On the right side they are two values `1488` indicates the total number of times the function ran and `82798 ns/op` is the average amount of time each iteration took to complete expressed in nanoseconds per operation.

## Next Steps

- Which method is faster? Slice or Map?
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

Get an output of the memory stats, you can pass an addition flag `-benchmem`

```bash
go test -bench=. -benchmem
```

### profile

Get cpu and memory profile from your benchmark tests
`-cpuprofile` enable and name of the file created for the cpu profile
`-memprofile` enable and name of the file creeated for the memory profile

```bash
go test -cpuprofile cpu.prof -memprofile mem.prof -bench . -benchmem
```

### cpu-web

Visualise the cpu profile in web at http://localhost:8000

```bash
go tool pprof -http=":8000" cpu.prof
```

### mem-web

Visualise the memory profile in web at http://localhost:8000

```bash
go tool pprof -http=":8001" mem.prof
```
