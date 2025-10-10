# Working with files

You can open files using the `os` package. Here's an example of how to open a file for reading:

```go
f, err := os.Open("data.jsonl")
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}
defer f.Close()
```

You'll note that the `os.Open` function returns two(!) values - the file handle and an error.

In Go, it's idiomatic to return errors as the last return value from functions.

You should always check for errors when calling functions that return an error, and wrap erorrs with additional context using `fmt.Errorf` and the `%w` verb ("wrapped error").

`try` / `catch` exception handling isn't part of Go. There's `panic` and `recover`, but `panic` is only for unrecoverable, fatal errors, and is not used for regular error handling.

### `defer`

You'll also notice the `defer f.Close()` statement.

This is similar to a `using` statement in C# or a `with` statement in Python, in that it ensures that the file is closed when the surrounding function (`main` in this case) returns, even if it returns early due to an error.

It's important to close files, network connections, HTTP response bodies and other resources when you're done with them to avoid resource leaks - basically, if it has a `Close` method, you should call it when you're done.

Go doesn't do a _great_ job of enforcing this - it's up to you.

### Unmarshalling JSON

You can unmarshal and marshal (deserialize, serialize) JSON data into Go structs using the `encoding/json` package.

A struct is a composite type that groups together variables under a single name.

```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address1  string
	Address2  string
	Address3  string
	Address4  string
	Postcode  string
	Country   string
}
```

You can unmarshal JSON data into a struct like this:

```go
data := []byte(`{"FirstName":"John","LastName":"Doe","Age":30,"Address1":"123 Main St","Address2":"","Address3":"","Address4":"","Postcode":"12345","Country":"USA"}`)

var person Person
err = json.Unmarshal(data, &person)
if err != nil {
    return fmt.Errorf("failed to unmarshal JSON: %w", err)
}
```

The `&person` passes a pointer to the `person` variable, allowing the `Unmarshal` function to modify the original variable.

Unlike Java, C# or Python classes, Go structs are value types, so when you pass a struct to a function, it gets copied, but if you pass a pointer to a struct, the function can modify the original struct.

You can find the documentation for the `encoding/json` package at [https://pkg.go.dev/encoding/json](https://pkg.go.dev/encoding/json).

## Customising the field names

By default, the JSON field names are expected to match the struct field names exactly.

In Go, public fields (those that start with an uppercase letter) are exported, and private fields (those that start with a lowercase letter) are not. JSON unmarshalling only works with exported fields(those that start with an uppercase letter).

There's no keywords like `public`, `private`, `internal` or `friend` - the visibility is determined solely by the case of the first letter! It means less typing, but might catch you out depending on your language background.

You can customise the field names using struct tags:

```go
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	Address3  string `json:"address_3"`
	Address4  string `json:"address_4"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
}
```

This tells the `encoding/json` package to look for the specified JSON field names when unmarshalling.

These aren't annotations, coroutines or attributes like you might find in other languages - they're just string literals associated with the struct fields that can be read using reflection at runtime.

Reflection is slow compared to regular code, so you should avoid using it in _really_ performance-critical code, however, in most cases, the performance impact is negligible.

## Task

The `main.go` file reads a file called `data.jsonl` which contains JSON Lines data - one JSON object per line.

The `main.go` file decodes each line into a `Person` struct.

For some reason... most of the data isn't printing, but we need to see the names of everyone under the age of 30.

Can you fix it?
