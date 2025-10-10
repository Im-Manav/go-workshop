# Importing packages

We've already seen that Go files start with the package name, and have imports at the top.

Here, we import the `fmt` package which contains functions for formatting text, including printing to the console.

```go
package main

import "fmt"

func main() {
}
```

This `fmt` package is part of the Go standard library.

When you type `fmt.` into your editor, assuming you have the Go extension or plugin installed, you should see a list of functions and types available in the `fmt` package.

When you use `fmt.Println`, you're calling the `Println` function from the `fmt` package, which prints text to the console followed by a newline.

The Go extension for your editor should automatically add the import statement to your code when you use a function from a package that isn't already imported, and when you save the file, any unused imports will be automatically removed.

You can see the documentation for the `fmt` package at [https://pkg.go.dev/fmt](https://pkg.go.dev/fmt).

## What's in the standard library?

The standard library is pretty big, it contains:

- `net/http` for building web servers and clients
- `os` for interacting with the operating system - reading and writing files, environment variables, etc.
- `time` for working with dates and times - yes, it's built-in!
- `math` for mathematical functions and constants
- `encoding/json` for working with JSON data
- `io` and `bufio` for input and output, including buffered I/O
- `log` and `slog` for logging (slog is structured JSON logging)
- `context` for managing request-scoped values, cancellation signals, and deadlines
- `sync` for concurrency primitives like mutexes and wait groups

You can find the full list of packages in the standard library at [https://pkg.go.dev/std](https://pkg.go.dev/std).

## Task

The `main.go` file wants to print a name, but with spaces removed from the start and end.

The standard library has a package called `strings` which contains functions for working with strings - import that package and use a function from it to remove the spaces.

## Tips

With Go, you can often find what you need in the standard library, so it's worth getting familiar with it. The standard library is well-documented and maintained, including security updates.

A lot of new Go developers go looking for third-party packages to do things that the standard library can already do, including string manipulation, HTTP servers and frameworks, and JSON handling.

Generally, I avoid third-party packages unless there's a very good reason to use one.

New Go developers often try to find packages that are similar to ones they've used in other languages (e.g. Express, but for Go), but in the long term, you're better off using the standard library wherever you can and adopting the idioms of Go.
