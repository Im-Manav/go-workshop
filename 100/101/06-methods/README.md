# Methods and interfaces

Go is an object-oriented language. You can attach methods to types, including types you define yourself.

```go
type Person struct {
  FirstName string
  LastName  string
  Age       int
}
```

To attach a method to a type, you define a function with a "receiver" argument. The receiver argument is like the `this` or `self` keyword in other languages, and it represents the instance of the type that the method is being called on.

```go
func (p Person) FullName() string {
  return p.FirstName + " " + p.LastName
}
```

You can call the method on an instance of the type:

```go
person := Person{FirstName: "John", LastName: "Doe"}
fmt.Println(person.FullName()) // Output: John Doe
```

It's idiomatic to use a single-letter abbreviation of the type name for the receiver argument, so for `Person`, we use `p`.

It's tempting to use `this` or `self` for the receiver argument, especially if you're coming from another language, but it's not idiomatic Go.

You can also define methods with pointer receivers, which allows the method to modify the original instance of the type:

```go
func (p *Person) IncrementAge() {
  p.Age++
}
```

You can call the method on an instance of the type, and Go will automatically take the address of the instance if needed:

```go
person := Person{FirstName: "John", LastName: "Doe", Age: 30}
person.IncrementAge()
fmt.Println(person.Age) // Output: 31
```

You may have noticed that we constructed the `Person` struct using a struct literal. This is a common way to create instances of structs in Go, but in cases where you need to perform some initialization logic, you can define a constructor function.

```go
func NewPerson(firstName, lastName string, age int) *Person {
  return &Person{
    FirstName: firstName,
    LastName:  lastName,
    Age:       age,
  }
}
```

This has the advantage that you can ensure that the struct is always initialised correctly, and you get a compiler warning if you forget to set a required field.

## Interfaces

Interfaces in Go are a way to define a set of methods that a type must implement. They allow you to write code that can work with different types as long as they implement the same interface.

Go interfaces only allow methods (behaviour) to be defined. You can't use them to define fields (properties). That's different to some other languages.

The Go standard library uses interfaces extensively, for example, the `io.Reader` interface defines a single method, `Read`, which reads data into a byte slice:

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

This is really useful, because files `io.File`, HTTP responses `http.Response.Body`, and in-memory buffers `bytes.Buffer` all implement the `io.Reader` interface, so you can write functions that can read from any of these types without caring about the specific type.

For example, to copy data from a HTTP response to a file, you can use the `io.Copy` function, which takes the src and destination as `io.Reader` and `io.Writer` interfaces:

```go
// Make a HTTP GET request.
resp, err := http.Get("https://example.com")
if err != nil {
  log.Fatal(err)
}
// Always remember to close the response body.
defer resp.Body.Close()

// Create a file.
file, err := os.Create("example.html")
if err != nil {
  log.Fatal(err)
}
defer file.Close()

// Copy the response body to the file.
_, err = io.Copy(file, resp.Body)
if err != nil {
  log.Fatal(err)
}
```

## Defining your own interfaces

Unlike Java or C#, where you have to explicitly declare that a class implements an interface, in Go, a type implements an interface simply by having the required methods.

This has implications for program structure. In C#, or Java, you might need to define an interface in a shared library, and then have multiple projects reference that library to implement the interface.

But in Go, you server can export a struct with methods, and a client can define an interface that matches the subset of methods of that struct that they use.

That feels odd at first - what if a type accidentally implements an interface? In practice, it doesn't happen!

Let's create an interface `IsYounger`:

```go
type Person struct {
  FirstName string
  LastName  string
  Age       int
}

func (p Person) IsYoung() bool {
  return p.Age < 30
}

type IsYounger interface {
  IsYoung() bool
}
```

We can now write code, including functions that accept any struct that has a method that matches the `IsYoung() bool` signature.

```go
func getYoungOnes(all []IsYounger) (output []IsYounger) {
  for _, o := range all {
    if !o.IsYoung() {
      continue
    }
    output = append(output, o)
  }
}
```

It is not idiomatic to call your interfaces `IYounger`, or `YoungerInterface` etc.

Large interfaces are also uncommon in Go. Interfaces comonly have a single method or just a few methods.

You might have noticed that instead of:

```go
if o.IsYoung() {
  output = append(output, o)
}
```

The loop is continued. Aligning the "happy path" to the left is a common idiom in Go called "Line of Sight".

There's a good talk by Mat Ryer on this topic of idiomatic Go code at: https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88

You may also have heard of this style as being a "never nester" - a good video on that is https://www.youtube.com/watch?v=CFRhGnuXG-4

## Task

Our code lists all the things that can be young. But... what about cats? They can be young.

Add a `Cat` type, implement the `IsYounger` interface, and add some cats to the `thingsThatCanBeYoung`.
