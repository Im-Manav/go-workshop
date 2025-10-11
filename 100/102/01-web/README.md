# Go Web server

Go ships with a Web framework and Web server.

```go
package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
  })
  http.ListenAndServe(":8080", nil)
}
```

This is all you need to create a web server that listens on port 8080 and responds with "Hello, World!" to any request.

It couldn't really get much simpler.

Of course, real-world web applications are more complex, but the standard library still works for that too.

Let's make a simple REST API.

## Handlers

The core of Go's web server is the `http.Handler` interface.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

A handler is anything that implements this interface.

If you recall adding methods to a struct, you can create your own handlers by defining a struct and adding a `ServeHTTP` method to it.

This is how dependency injection is typically done in Go web applications, you create a struct that holds the dependencies and then implement the `ServeHTTP` method to handle the request.

You can access the dependencies via the struct fields, e.g. a database connection, a logger, etc. but in this case, the dependency is just a name to say hello to.

```go
func NewHelloHandler(name string) HelloHandler {
  return HelloHandler{Name: name}
}

type HelloHandler struct{
  Name string
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %s!", h.Name)
}
```

Now you can create a new `HelloHandler` and register it with the `http` package.

```go
func main() {
  http.Handle("/hello", NewHelloHandler("World"))
  http.ListenAndServe(":8080", nil)
}
```

## Handler functions

You can also use the `http.HandleFunc` function to register a function as a handler, which is useful for very simple handlers.

```go
func main() {
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
  })
  http.ListenAndServe(":8080", nil)
}
```

In this case, the function is converted to a handler by the `http.HandlerFunc` type, which is a type that implements the `http.Handler` interface. Yes, that's right, you can add methods to function types in Go. It's not something you are likely to do often, but it is possible.

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// [Handler] that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

## Code structure

In a real-world application, you would typically have multiple handlers, each handling a different route.

I usually create a `handlers` package and put each handler in its own subdirectory that matches the HTTP route that it handles. This helps people to find the code that handles a specific route.

```
./handlers
  /people/handler.go
  /products/handler.go
  /orders/handler.go
```

Each handler would typically have its own dependencies, e.g. a database connection, a logger, etc. injected via the constructor function.

```go
package people

//... imports elided.

func NewHandler(db *db.DB, logger *slog.Logger) *Handler {
  return &Handler{
    db: db,
    log: logger,
  }
}
```

Each handler must implement the `http.Handler` interface by adding a `ServeHTTP` method.

Since the `ServeHTTP` method is where the request is handled, I typically use a switch statement to route the request to the appropriate method.

```go
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case http.MethodGet:
      h.Get(w, r)
    case http.MethodPost:
      h.Post(w, r)
    default:
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
  }
}
```

## Routing

The standard library includes a router, but it's got a slightly unusual name of `Mux` (short for "multiplexer").

A `Mux` is a type that implements the `http.Handler` interface and routes requests to the appropriate handler based on the request URL.

```go
func main() {
  mux := http.NewServeMux()
  mux.Handle("/hello", NewHelloHandler("World"))
  http.ListenAndServe(":8080", mux)
}
```

If you don't define your own `ServeMux`, the `http` package uses a default one, which is why you can use `http.Handle` and `http.HandleFunc` without creating a `ServeMux` explicitly.

However, in production code, you're better off creating your own `ServeMux` so you can configure it as needed. In addition, any package could potentially register handlers on the default `ServeMux`, which could lead to unexpected behavior or security vulnerabilities.

## Path variables

The `ServeMux` type has built-in support for path variables and HTTP method restrictions.

```go
mux.Handle("GET /users/{id}", userHandler)
```

And you can retrieve the path variables from the request context:

```go
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  if id == "" {
    http.Error(w, "Missing user ID", http.StatusBadRequest)
    return
  }

  // ... rest of the code.
}
```

The full API is documented at https://pkg.go.dev/net/http#hdr-Patterns-ServeMux

## Middleware

Middleware is a way to wrap handlers to add functionality, e.g. logging, authentication, etc.

You can write your own easily. Let's start with a long form version that logs the incoming request method and URL:

```go
func NewLogger(log *slog.Logger, next http.Handler) http.Handler {
  return &logger{
    log: log,
    next: next,
  }
}

type logger struct {
  log *slog.Logger
  next http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  l.log.Info("Request", slog.String("method", r.Method), slog.String("url", r.URL.String()))
  l.next.ServeHTTP(w, r)
}
```

This middleware takes a logger and the next handler as dependencies, and logs the request before calling the next handler.

You can use it like this:

```go
mux.Handle("/hello", NewLogger(logger, NewHelloHandler("World")))
```

Or, you might choose to wrap your entire `ServeMux` with the middleware, e.g. if your midddleware is an authentication middleware that checks for a valid token on every request, and only lets the request proceed if the token is valid.

```go
authenticatedMux := NewAuthMiddleware(authService, mux)
http.ListenAndServe(":8080", authenticatedMux)
```

Let's compress our logger middleware a bit. It's harder to understand if you don't know Go very well, but it's more concise. Generally, you might want to favor clarity over conciseness, but middleware is a common enough pattern that most Go developers will understand this version.

```go
func NewLogger(log *slog.Logger, next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Info("Request", slog.String("method", r.Method), slog.String("url", r.URL.String()))
    next.ServeHTTP(w, r)
  })
}
```

## Task

This project allows you to create users and list them, but there's no way to `GET /user/{id}` to retrieve a specific user. Can you add it?

To setup:

```bash
go run .
```

Then you can create a user, by passing all of the `UserFields` in the request body as JSON:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"first_name":"John","last_name":"Doe","email":"john.doe@example.com"}' http://localhost:8080/users
```

Then you can list all users:

```bash
curl http://localhost:8080/users
```

Can you add support for `GET /user/{id}` to retrieve a specific user by ID?

Bonus points (there's no prizes, don't get too excited) for also supporting `POST /user/{id}` to update an existing user, and `DELETE /user/{id}` to delete a user.
