# Basic types

Go has the usual set of primitive types you'd expect from a programming language:

- int (architecture specific size, either 32 or 64 bits)
- int8, int16, int32, int64 (signed integers of specific sizes)
- uint, uint8, uint16, uint32, uint64 (unsigned integers of specific sizes)
- float32, float64 (floating point numbers)
- byte (alias for uint8)
- string (a sequence of runes - more on that later)
- bool (boolean type, true or false)

Along with some built-in composite types which we'll cover later.

- slice (dynamically-sized array)
- array (fixed-size array)
- map (key-value pairs)
- pointer (a variable that holds the memory address of another variable)

## Defining variables

You can define variables using the `var` keyword. Since Go is a strongly typed language, you need to specify the type of the variable:

```go
var x int
var y float64
var name string
```

The variables `x`, `y`, and `name` are now declared but not initialized, so they hold the zero value for their respective types (0 for int, 0.0 for float64, and "" for string).

It's normal to use this pattern if you're declaring a variable that you want to have the zero value.

You _can_ initialize variables with the `var` keyword, but it's not common.

```go
var x int = 10
var y float64 = 20.5
var name string = "Hello, Go!"
```

If you're setting a value, it's more common to use the shorthand syntax which infers the type automatically:

```go
x := 10
y := 20.5
name := "Hello, Go!"
```

You can cast between types using the type name as a function:

```go
x := int64(10)
y := float32(20.5)
name := string([]byte{'H', 'e', 'l', 'l', 'o', ' ', 'G', 'o', '!'})
```

Be careful though... `string(32)` is not the string "32", it's the string containing the single character with Unicode code point 32 (a space)! You'll want to use `strconv.Itoa(32)` to convert an integer to its string representation, or `fmt.Sprintf("%d", 32)` to format it as part of a string.

You can also declare multiple variables at once, again, not something you see often:

```go
var a, b, c int = 1, 2, 3
x, y := 4, 5
```

## Task

The `main.go` file contains a 
