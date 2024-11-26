# Testing

## Tests

- Compiled separately.
- Use built-in testing framework.
- Have a `_test.go` suffix.
- Start with `TestXXX`
- Have the signature `TestName(t *testing.T)`

```go
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

## Running tests

```
go test ./...
```

---

## Running a specific test

```
go test ./... -run TestAdd
```

---

## Code coverage

```
go test ./... -cover
```

---

## Subtests

```go
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

## Table-driven tests

```go
func TestAdd(t *testing.T) {
  tests := []struct {
    x, y, expected int
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

## Dependencies

- Return structs, accept interfaces
- Define your own interface

```go
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

## Use an interface to limit what you need

```go
type CustomerPutter interface {
  PutCustomer(c models.Customer) (err error)
}

func NewCustomerHandler(cp CustomerPutter) CustomerHandler {
  return CustomerHandler{cp: cp}
}

type CustomerHandler struct {
  cp CustomerPutter
}

func main() {
  db := NewDB("...")
  ch := NewCustomerHandler(db)
  // ...
}
```

---

## Write your own mocks

```go
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
	  return nil
	},
  }
  ch := NewCustomerHandler(db)
  // ...
}
```

---

## Tips

- There are 3rd party assertion libraries, but not really needed
- Dependency injection framworks aren't a thing in Go
- Google's `cmp` library is handy for deep comparisons
- You can pass around params and functions, but it gets messy
- Using struct fields for dependencies works well
