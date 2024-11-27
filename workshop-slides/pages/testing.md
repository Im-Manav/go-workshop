---
layout: section
---

# Testing

---
layout: two-cols
---

# Tests

- Compiled separately
- Use built-in testing framework
- Have a `_test.go` suffix
- Start with `TestXXX`
- Have the signature `TestName(t *testing.T)`

::right::

<br>
<br>

```go {|8-10|1-6|2|3-5}
func TestAdd(t *testing.T) {
  result := add(1, 2)
  if result != 3 {
    t.Errorf("Expected 3, got %d", result)
  }
}

func add(x, y int) int {
  return x + y
}
```

---

# Running tests

```bash
go test ./...
```

<br>

<v-click>

```
ok      github.com/a-h/go-workshop-102    0.179s
```

</v-click>

---

# Code coverage

```bash
go test ./... -cover
```

<br>

<v-click>

```
PASS
coverage: 100.0% of statements
ok      github.com/a-h/go-workshop-102    0.190s
```

</v-click>

---

<img src="/coverage.png" style="height: 95%">

---

<img src="/test-runner.png" style="height: 95%">

---

# Subtests

```go {|1,14|2-7|8-13}
func TestAdd(t *testing.T) {
  t.Run("1+2=3", func(t *testing.T) {
    result := add(1, 2)
    if result != 3 {
      t.Errorf("Expected 3, got %d", result)
    }
  })
  t.Run("2+3=5", func(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
      t.Errorf("Expected 5, got %d", result)
    }
  })
}

func add(x, y int) int {
  return x + y
}
```

---

# Subtests

```bash
go test -v ./...
```

<v-click>

```
=== RUN   TestAdd
=== RUN   TestAdd/1+2=3
=== RUN   TestAdd/2+3=5
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/1+2=3 (0.00s)
    --- PASS: TestAdd/2+3=5 (0.00s)
PASS
ok      github.com/a-h/testt    0.189s
```

</v-click>

---

# Run a specific test

```bash
go test ./... -run TestAdd
```

---

## Table-driven tests

```go {|2-6|6-9|10,18|11|12-17}
func TestAdd(t *testing.T) {
  tests := []struct {
    x int
    y int
    expected int
  }{
    { x: 1, y: 2, expected: 3 },
    { x: 2, y: 3, expected: 5 },
  }
  for _, test := range tests {
    name := fmt.Sprintf("%d+%d=%d", test.x, test.y, test.expected)
    t.Run(name, func(t *testing.T) {
      actual := add(test.x, test.y)
      if actual != test.expected {
        t.Errorf("Expected %d, got %d", test.expected, actual)
      }
    })
  }
}
```

---
layout: section
---

# So far, so good

---

# Dependencies / mocking

- Return structs, take interfaces
- Consumers define interfaces
- Standard library contains some ready-made interfaces:
  - `io.ReadCloser` (use instead of a file, or network connection)
  - `io.Writer` (use instead of a file, or network connection)
  - `http.Handler`

---

# Database implementation

```go {|5-7|9-11|13-15|1-3}
func NewDB(dsn string) DB {
  // ...
}

type DB struct {
  // ...
}

func (db DB) PutCustomer(c models.Customer) (err error) {
  // ...
}

func (db DB) GetCustomer(id string) (err error) {
  // ...
}
```

---

# Subject under test - before refactoring

```go {|5-7|1-3|8-13}
func NewCustomerHandler(db DB) CustomerHandler {
  return CustomerHandler{db: db}
}

type CustomerHandler struct {
  db DB
}

func main() {
  db := NewDB("...")
  ch := NewCustomerHandler(db)
  // ...
}
```

---

# Subject under test - after refactoring

```go {|1-3|9-11|5-7|13-15}
type CustomerPutter interface {
  PutCustomer(c models.Customer) (err error)
}

func NewCustomerHandler(db CustomerPutter) CustomerHandler {
  return CustomerHandler{db: db}
}

type CustomerHandler struct {
  db CustomerPutter
}

func main() {
  db := NewDB("...")
  ch := NewCustomerHandler(db)
  // ...
}
```


---

# Write your own mocks

```go {|1-3|5-7|10-15|16}
type MockDB struct {
  PutCustomerFn func(c models.Customer) (err error)
}

func (db MockDB) PutCustomer(c models.Customer) (err error) {
  return db.PutCustomerFn(c)
}

func TestCustomerHandler(t *testing.T) {
  // ...
  db := MockDB{
    PutCustomerFn: func(c models.Customer) (err error) {
      return errors.New("simulated database failure")
    },
  }
  ch := NewCustomerHandler(db)
  // ...
}
```

---

# Tips

- There are 3rd party assertion libraries, but not really needed
- Dependency injection framworks aren't a thing in Go
- Google's `cmp` library is handy for deep comparisons
- You can pass around params and functions, but it gets messy
- Using struct fields for dependencies works well

---

# Hands-on - 15 minutes

- The `./security` directory contains a simple API.
- One of the HTTP handlers is tested, with 100% code coverage.
- Refactor the other handler to be testable, and write tests for it.
