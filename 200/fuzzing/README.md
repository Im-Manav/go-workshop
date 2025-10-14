# Fuzzing

Fuzzing is a technique where you provide random inputs to your functions, and see if they can handle them without crashing or misbehaving.

Go has had built-in support for fuzzing since Go 1.18 (2022).

Like a lot of things that have made their way into the standard library, it started off as a successful open source effort, later becoming adapted and adopted into the standard library.

Go is a memory safe language, so you won't get memory corruption issues like you might in C or C++, for example, being able to write outside the bounds of an array which can lead to arbitrary code execution.

However, you can still have panics, uncontrolled resource consumption leading to denial of service, or other unexpected behaviour.

Fuzzing can help you find these issues by providing unexpected inputs to your functions.

Good places to use fuzzing include:

- Input validation functions
- Parsing functions
- Functions that handle untrusted data
- Functions that handle complex data structures
- Functions that have a lot of branching logic
- Functions that have a history of bugs or vulnerabilities

## Scenario

One of the developers on your team has created a function that parses URL paths for a web application.

The function is supposed to take a path like `/search/abc/def` and return the `abc` and `def` parts as `entity` and `term` variables.

As part of our security process, we always do threat modelling and code review for functions that handle untrusted input, such as URL paths, and the use of a regular expression to parse the path has raised some concerns.

Regular expressions can be hard to understand, and are a regular source of bugs and vulnerabilities.

The developer has written some unit tests for the function, but we need to be sure that it can handle a wide range of inputs without crashing or misbehaving.

So, let's add a fuzz test to the function.

To create a fuzz test, you create a function that starts with `Fuzz` and takes a `*testing.F` parameter, and use the `F.Add` method to provide some initial seed inputs.

Then, you use the `F.Fuzz` method to define the actual fuzz test, which takes a function that accepts a `*testing.T` parameter and any number of other parameters that represent the inputs to the function being tested.

The inputs to the function being tested need to be simple types, such as strings, integers, or byte slices.

In our case, we can just pass in the path string.

```go
func FuzzParse(f *testing.F) {
	f.Add("/search/abc/def")
	f.Add("/search/abc/def/ghi")
	f.Fuzz(func(t *testing.T, path string) {
		_, _, err := Parse(path)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
```

But if the subject under test takes a struct as an input, you can break it down into its component fields and pass them in as separate parameters.

```go
type Input struct {
    Name string
    Age  int
}

func FuzzProcessInput(f *testing.F) {
    f.Add("Alice", 30)
    f.Add("Bob", 25)
    f.Fuzz(func(t *testing.T, name string, age int) {
        input := Input{Name: name, Age: age}
        ProcessInput(input)
    })
}
```

## Running the fuzz test

To run the fuzz test, you use the `go test` command with the `-fuzz` flag, specifying the name of the fuzz test function or `.` to run all fuzz tests in the package.

```bash
go test -fuzz=FuzzParse
```

The fuzzing process will run until you stop it, or until it finds a failure.

When it finds a failure, it will print out the input that caused the failure, and save it to a file in the `testdata/fuzz` directory.

That way, you can reproduce the failure and investigate it further and the failure will be retained and retested in future runs.

You can also use the `-fuzztime` flag to specify how long to run the fuzzing for if you want to do a light fuzzing run in CI, for example.

## Fixing the error

Well, running the fuzzing test has found a panic in the `Parse` function...

```
--- FAIL: FuzzParse (0.00s)
    --- FAIL: FuzzParse/771e938e4458e983 (0.00s)
panic: runtime error: index out of range [1] with length 0 [recovered, repanicked]

goroutine 12 [running]:
testing.tRunner.func1.2({0x5f6ee0, 0xc00001a378})
        /nix/store/0vvxk4j7dmdpmw3s60fp6i2v5nqdjsjs-go-1.25.2/share/go/src/testing/testing.go:1872 +0x237
testing.tRunner.func1()
        /nix/store/0vvxk4j7dmdpmw3s60fp6i2v5nqdjsjs-go-1.25.2/share/go/src/testing/testing.go:1875 +0x35b
panic({0x5f6ee0?, 0xc00001a378?})
        /nix/store/0vvxk4j7dmdpmw3s60fp6i2v5nqdjsjs-go-1.25.2/share/go/src/runtime/panic.go:783 +0x132
github.com/a-h/go-workshop/200/fuzzing.Parse(...)
        /home/adrian-hesketh/github.com/a-h/go-workshop-102/200/fuzzing/paths.go:10
github.com/a-h/go-workshop/200/fuzzing.FuzzParse.func1(0x0?, {0xc000012371?, 0x0?})
        /home/adrian-hesketh/github.com/a-h/go-workshop-102/200/fuzzing/fuzz_test.go:38 +0x52
reflect.Value.call({0x5cb8a0?, 0x61bfd0?, 0x13?}, {0x60a630, 0x4}, {0xc000014a20, 0x2, 0x2?})
        /nix/store/0vvxk4j7dmdpmw3s60fp6i2v5nqdjsjs-go-1.25.2/share/go/src/reflect/value.go:581 +0xcc6
reflect.Value.Call({0x5cb8a0?, 0x61bfd0?, 0x788a80?}, {0xc000014a20?, 0x609bc0?, 0x765bd0?})
        /nix/store/0vvxk4j7dmdpmw3s60fp6i2v5nqdjsjs-go-1.25.2/share/go/src/reflect/value.go:365 +0xb9
testing.(*F).Fuzz.func1.1(0xc0000ec700?)
        /nix/store/0vvxk4j7dmdpmw3s60fp6i2v5nqdjsjs-go-1.25.2/share/go/src/testing/fuzz.go:341 +0x32a
testing.tRunner(0xc0000ec700, 0xc0000fc240)
        /nix/store/0vvxk4j7dmdpmw3s60fp6i2v5nqdjsjs-go-1.25.2/share/go/src/testing/testing.go:1934 +0xea
created by testing.(*F).Fuzz.func1 in goroutine 9
        /nix/store/0vvxk4j7dmdpmw3s60fp6i2v5nqdjsjs-go-1.25.2/share/go/src/testing/fuzz.go:328 +0x637
exit status 2
FAIL    github.com/a-h/go-workshop/200/fuzzing  0.005s
```

Our task is to fix the implementation of the `Parse` function so that it can handle a wide range of inputs without crashing or misbehaving.
