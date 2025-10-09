# Create a Hello World executable

Time to make a Go project.

## Create a new module

First, start by initialising a new module with `go mod init <module-path>`. This creates a `go.mod` file in the current directory.

The module path is typically the repository location. This allows others to import your code if you publish your module code to a version control system like GitHub.

In my case, I plan to put my code in Github, my Github username is `a-h`, and my repository is `hello`, so I run:

```bash
go mod init github.com/a-h/hello
```

Tip: Go module and package names are always lowercase, e.g. `github.com/a-h/hello`, not `github.com/a-h/Hello`. Package names shouldn't contain underscores or hyphens.

## Add your code

Once you have your module initialised, create a `main.go` file with the following code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Build

Now you can build your code with:

```bash
go build
```

Because the package is `main`, this creates an executable file named `hello` (or `hello.exe` on Windows) in the current directory.

## Run

You can run the executable with:

```bash
./hello
```

Or, on Windows:

```powershell
hello.exe
```

## Advanced

You can combine the build and run steps with:

```bash
go run .
```

This compiles the code to a temporary location and runs it.
