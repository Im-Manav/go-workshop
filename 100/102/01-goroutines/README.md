# Goroutines

All code in Go is executed within a goroutine. A goroutine is a lightweight thread managed by the Go runtime - it's lightweight in that it doesn't require a lot of memory (typically around 2KB) and can be created in large numbers (thousands or even millions).

During compilation, a runtime is included in the binary, which manages the execution of goroutines. The Go scheduler multiplexes multiple goroutines onto a smaller number of OS threads.

The built-in Go runtime scheduler allocates goroutines to available OS threads. If a goroutine performs a blocking operation, such as waiting for I/O, the scheduler can pause that goroutine and run another one on the same OS thread. This is why Go doesn't have `async/await` keywords; the scheduler handles concurrency at the macro level for you, you don't need to state when you yield control with a keyword.

This is one of the things that makes Go code quite simple to read and write, as you can write synchronous-looking code that runs concurrently. For example, each HTTP request that is handled by the `http.Server` runs in its own goroutine - this allows the server to handle many requests concurrently without blocking.

## Ceating your own goroutines

However, there's lots of times when you might want to create your own goroutines:

- Performing background tasks:
  - Run a web server in the background, and carry on with other work in the main goroutine.
  - Start a goroutine to periodically check for updates or perform maintenance tasks in the background.
- Handling multiple tasks concurrently:
  - Processing multiple files.
  - Making multiple network requests.
- To improve performance by parallelising CPU-bound tasks across multiple CPU cores.
  - Performing complex calculations or data processing in parallel.

You can create your own goroutines using the `go` keyword followed by a function call. This starts a new goroutine that runs concurrently with the calling code.

```go
func sayHello() {
    fmt.Println("Hello from sayHello!")
}

func main() {
    go sayHello()
    go func() {
            fmt.Println("Hello from anonymous func!")
    }()

    // Give the goroutines time to complete.
    time.Sleep(1 * time.Second)
    fmt.Println("Main function completed")
}
```

Clearly, it's not ideal to use `time.Sleep` to wait for goroutines to complete, since we might not know how long they will take.

A better way to wait for goroutines to complete is to use a `sync.WaitGroup`. A `WaitGroup` allows you to wait for a collection of goroutines to finish. You can add the number of goroutines to the `WaitGroup`, and then call `Done` when each goroutine completes. Finally, you call `Wait` to block until all goroutines have finished.

It's idiomatic to use the `defer` keyword to ensure that `wg.Done()` is called when the goroutine completes.

The following example shows how to use a `WaitGroup` to wait for two goroutines to complete:

```go
func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()

        time.Sleep(1 * time.Second)
        fmt.Println("Hello from func 1!")
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()

        time.Sleep(2 * time.Second)
        fmt.Println("Hello from func 2!")
    }()

    wg.Wait()
    fmt.Println("All goroutines completed")
}
```

Some languages, using `Task.Run` (C#) or `std::async` (C++) creates a new thread. In Go, using the `go` keyword does not create a new OS thread; it creates a new goroutine that is managed by the Go runtime. The Go scheduler will decide how to map goroutines to OS threads, which allows for efficient use of system resources.

Go will schedule goroutines on all available CPU cores by default. You can control the number of OS threads that can execute user-level Go code simultaneously using the `GOMAXPROCS` setting. By default, `GOMAXPROCS` is set to the number of CPU cores available on the machine or container. You'd typically only change this if you have a specific reason to limit the number of threads, such as running in a constrained environment.

In Go 1.25 and later, a new feature was added to the `WaitGroup` to reduce the likelihood of forgetting to defer the `Done` call. You can now use `wg.Go(func() { ... })` to start a goroutine which will automatically handle the `Add` and `Done` calls for you.

This makes the code a little less verbose and reduces the chance of errors:

```go
func main() {
    var wg sync.WaitGroup

    wg.Go(func() {
        time.Sleep(1 * time.Second)
        fmt.Println("Hello from func 1!")
    })

    wg.Go(func() {
        time.Sleep(2 * time.Second)
        fmt.Println("Hello from func 2!")
    })

    wg.Wait()
    fmt.Println("All goroutines completed")
}
```

## Shared state

In the case that we have multiple goroutines accessing a single variable, we need to ensure that access to that variable is synchronised. This is typically done using a `sync.Mutex`, which provides a way to lock and unlock access to a variable.

Let's imagine we're writing to a file, we probably want to ensure that only one goroutine is writing to the file at a time so that each write is atomic. We can use a `Mutex` to ensure this.

```go
func main() {
    var mu sync.Mutex

    f, err := os.Create("output.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    var wg sync.WaitGroup

    for i := range 20 {
        wg.Go(func() {
            mu.Lock()
            defer mu.Unlock()

            _, err := io.WriteString(f, fmt.Sprintf("Hello from goroutine %d\n", i))
            if err != nil {
                log.Println("Error writing to file:", err)
            }
        })
    }

    wg.Wait()
}
```

If you disable the mutex, you may see garbled output in the file, as multiple goroutines write to the file at the same time.

In this specific piece of example code, you're unlikely to see the issue because the amount of data being written is small and the writes are likely to complete before another goroutine starts writing. However, in a real-world scenario with larger writes or more complex operations, the lack of synchronization could lead to data corruption or unexpected results.

## Returning errors from goroutines

One challenge with goroutines is that they don't return values directly. If you want to get a result or an error from a goroutine, you'll need to store the result in a variable that is accessible to the main goroutine, and use a mutex to synchronise access to that variable.

Go's standard library is very stable. Adding new features is done very carefully to avoid breaking existing code. However, the Go team maintain a set of experimental packages in the `golang.org/x/` namespace. One of these packages is `errgroup`, which provides a way to run a group of goroutines and collect errors from them.

These experimental packages are available for use, but they are not part of the standard library and may change in future releases. You can find the `errgroup` package at `golang.org/x/sync/errgroup`.

Remember that you can use `go get golang.org/x/sync/errgroup` to download and install the package.

The documentation contains a great example, which shows how to fetch multiple URLs concurrently, and return an error if any of the fetches fail.

```go
package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
```

If you have a mixture of goroutines that return errors, and ones that don't, you can simply return `nil` from the goroutines that don't return errors.

## Task

The program startup takes too long. The program loads some configuration from a file, connects to a database, makes an API call, and then starts a web server.

It currently takes 3s to start up, because each of those tasks takes 1s, and they are done one after the other.

The configuration loading, database connection, and API call can all be done concurrently, so modify the code to do that.

You may need to use a "closure" to set the values of the variables declared in `main`.
