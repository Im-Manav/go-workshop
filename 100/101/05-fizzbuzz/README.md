# Fizz buzz

This challenge demonstrates a few things.

## Command line arguments

Go has built-in support for command line argument parsing.

The first argument is always the name of the executable. The rest are the arguments passed to it.

Using the built-in `flag` package, we can define flags and parse them into variables.

```go
cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
n := cmd.Int("n", 0, "Number to check whether it's Fizz or Buzz")
cmd.Parse(os.Args[1:])
```

In Go, if we want to create a value that can be nil, we need to use a pointer.

This is often used to represent optional values - in this case, a command line argument that may or may not be provided.

```go
var n *int
n = nil

// This allocates memory for an int and returns a pointer to it, at this point, n is not nil and has the zero value of int, which is 0. 
n = new(int)

// We can dereference the pointer to get or set the value it points to.
*n = 42
m := 43
n = &m // This makes n point to m.
```

If you try to access a nil pointer, you'll get a runtime panic, so it's good practice to check for nil before dereferencing a pointer or using it.

```go
var n *int
fmt.Println(*n)
```

A panic is a runtime error that stops the normal execution of a program.

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x497b76]

goroutine 1 [running]:
main.main()
	/tmp/sandbox64169669/prog.go:9 +0x16
```

Built-in complex types in Go, like slices and maps, are pointers under the hood, so their zero value is `nil`.

```go
var s []int
if s != nil {
  // This will never happen because s is nil.
}

var m map[string]int
if m != nil {
  // This will never happen, because m is nil.
}
```

Either way, once we have run `cmd.Parse(os.Args[1:])`, the variable `n` will be a pointer to an int, which may be nil if the flag was not provided, or a pointer to an int with the value provided by the user.

## Code organisation

You'll notice that there's a `./fizzbuzz` directory with a `fb.go` file in it.

This is a package that is imported into our `main.go` file.

Note that the import path is relative to the module root, which is defined in `go.mod`.

```go
import "github.com/a-h/go-workshop/100/101/05-fizzbuzz/fizzbuzz"
```

Once a package is imported, all of the exported functions, types, and variables are available to use.

Take a look at the `./fizzbuzz/fb.go` file to see how the `fizzbuzz.Check` function is defined.

```go
package fizzbuzz

import "fmt"

func Check(n int) string {
 //TODO: Implementation...
 return ""
}
```

The default name of the package is the name of the directory it is in, so here we use `fizzbuzz`.

What's really cool about Go, is that if you publish this on GitHub, or another public repository, anyone can import your package and use it in their own code. There's no need to go through extra steps, like publishing it to a package registry.

If needed, or it clashes, you can rename the package on import, like this:

```go
import fb "github.com/a-h/go-workshop/100/101/05-fizzbuzz/fizzbuzz"
```

In the example, you can see that we call the `fizzbuzz.Check` function, defined in `fb.go`, passing the value of `n`.

```go
fb := fizzbuzz.Check(*n)
```

It's idiomatic to name packages with short, lowercase names, without underscores or camel case.

The names of functions, types, and variables that are exported (i.e. available to other packages) must start with an uppercase letter.

Don't worry about remembering the full import path - your editor will help you with that - if you type `fizzbuzz.` it will suggest the available functions, types, and variables automatically, and import the package for you.

## Testing - introduction

Go has built-in support for testing, so the `fizzbuzz` package has a `fb_test.go` file that contains tests for the `Check` function.

We can run the tests with `go test ./...` to test all packages, or `go test ./fizzbuzz` to test just the `fizzbuzz` package.

Tests in Go are just functions that start with `Test` and take a `*testing.T` parameter.

```go
func TestTrue(t *testing.T) {
  if true == false {
    t.Error("This should never happen")
  }
}
```

There's no special assertions library - you just use normal Go code to check conditions, and call `t.Error` or `t.Fatal` if something is wrong.

One common pattern is to define a slice (list) of test cases, and loop over them, checking each one.

This uses some features we haven't covered yet - anonymous types, and first-class functions, so let's explain them briefly.

## Anonymous types

You can define a named struct type like this:

```go
type Person struct {
  Name string
  Age int
}
```

This defines a new type called `Person` with two fields, `Name` and `Age`.

But, you can define an anonymous struct by omitting the type name:

```go
john := struct {
  Name string
  Age int
}{
  Name: "John",
  Age: 30,
}
```

If you break it down, it's just like defining a named struct, but without the name, and then immediately creating a value of that type.

## First-class functions / anonymous functions

You can do the same thing with functions.

Instead of naming a function like this:

```go
func add(a, b int) int {
  return a + b
}
```

You can define an anonymous function and assign it to a variable:

```go
add := func(a, b int) int {
  return a + b
}
result := add(2, 3)
fmt.Printf("%d", result) // 5
```

This is similar to how you might use lambdas or arrow functions in other languages, e.g. in TypeScript:

```ts
const people = [
  { name: "John", age: 30 },
  { name: "Jane", age: 25 },
];
const under30 = (p: Person) => p.age < 30;
const youngPeople = people.filter(under30);
```

However, there's no first-class Lambda / LINQ style functionality in Go. It's more idiomatic to use simple loops and conditionals.

## Putting it all together

For defining test inputs, a common idiom in Go is to use a table-driven test - a set of test cases defined in a slice (list), and then loop over them, checking each one.

```go
tests := []struct {
  input    int
  expected string
}{
  {
    input:    0,
    expected: "0",
  },
  {
    input:    1,
    expected: "1",
  },
}
```

In the loop, the `t.Run` function is used. This function takes the name of the sub-test, and a function to run the test.

```go
for _, test := range tests {
  t.Run(fmt.Sprintf("input=%d", test.input), func(t *testing.T) {
    actual := Check(test.input)
    if actual != test.expected {
      t.Errorf("expected %q, got %q", test.expected, actual)
    }
  })
}
```

## Task

The tests aren't actually passing. It looks like the `fizzbuzz.Check` function has a bug in it.

Fix it and make the tests pass.
