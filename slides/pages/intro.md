---
layout: section
---

# Introduction to Go

---

# History

- Created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson
- First public release in 2009
- Version 1.0 released in March 2012
- Go 1.0 code still works with the latest Go version, over a decade later
- Allegedly created in response to frustrations with C++ build times and complexity

---

# Where is Go used?

- Docker ecosystem
  - Docker
  - Kubernetes
- Terraform
- Prometheus
- Grafana
- HashiCorp tools (Consul, Vault, Nomad)
- Cloud-native tools (Istio, Kong, Traefik)
- Many companies use Go for backend services (Google, CloudFlare, Just Eat, Uber, Twitch, Dropbox, Monzo, ITV, Aviva, Thought Machine etc.)
- Over 5M Go developers worldwide

---
layout: two-cols-header
---

# Features

::left::

- Simple and easy to use
- Strongly typed
- Compiled to native code - not interpreted, or JVM/CLR
- Memory safe - no pointer arithmetic, garbage collected
- Built-in concurrency support (goroutines, channels)
- Comprehensive standard library
- Fast compilation times
- Cross-platform (Windows, macOS, Linux, ARM64, x86_64)

::right::

<img src="/Go_Logo_Blue.svg" alt="Go logo" style="width: 200px;">

<br>
<br>

> The driving motivation for Go 1 is stability for its users. People should be able to write Go programs and expect that they will continue to compile and run without change...
>  - Go 1.0 release notes

---

# Hello, HTTP...

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, via HTTP!"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

---

# Code organisation

- TODO: Add diagram

- Go projects are organised into "modules"
- A module is a collection of related Go "packages"
- A package is a directory containing Go source files
- Each package has a unique import path (e.g. `github.com/yourusername/hello`)
- The `go.mod` file at the root of a module defines the module's path and its dependencies
 - The `go.sum` file contains checksums for module dependencies to ensure integrity

---
layout: section
---

# Starting a new Go project

---

# Create a new module

- Run `go mod init <module-path>` to create a new module

```bash
go mod init github.com/yourusername/hello
```

- This creates a `go.mod` file in the current directory

---

# Add your code

- The `main` package is the entry point for a Go program.

```
go.mod
main.go
```

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

---

# Build

- Run `go build` to compile the program

```bash
go build
```

- Because the package is `main`, this creates an executable file named `hello` (or `hello.exe` on Windows)

---

# Run

```bash
./hello
```

<!--

Automatically compiled for your OS/architecture, so on Windows it would be `hello.exe`

-->

---

# Build it for another OS/architecture

```bash
GOOS=linux GOARCH=amd64 go build -o hello-linux
```

- This creates a Linux executable named `hello-linux` even if you're on macOS or Windows

---

# Create a library package

```{3}
go.mod
main.go
greetings/hello.go
```

```go
package greetings

import "fmt"

func Hello(name string) string {
  return fmt.Sprintf("Hello, %s!", name)
}
```

---

# Use the library package

```{2}
go.mod
main.go
greetings/hello.go
```

```go
package main

import (
  "fmt"
  "github.com/yourusername/hello/greetings"
)

func main() {
  fmt.Println(greetings.Hello("World"))
}
```

---

# Evolution

New features since 1.0 include:

- Modules (dependency management)
- Generics (type parameters)
- Improved error handling (errors.Is, errors.As, fmt.Errorf with %w)
- Built-in structured logging (slog package)
- Context package for managing request-scoped values, cancellation, and timeouts
- HTTP/2 and HTTP/3 support in the net/http package
- FIPS 140-2 compliant crypto libraries


---
layout: two-cols-header
---

# Simplicity

::left::

- Minimalist language design
- "Batteries included" standard library
  - Networking, HTTP, cryptography, compression, image processing, database/sql, testing, and more.

::right::

```bash
go run
```

```bash
go build
```

```bash
go test
```

```bash
gofmt
```

```bash
gopls
```
