# Testing web applications

We've built a web API, but now we want to test it.

We can use a variety of different strategies to test our web application. Some developers prefer to create integration tests that run against a real database, while others are happy to accept the trade off of using mocks and fakes to speed up their tests.

In this example, we'll use a combination of both strategies. We'll use an in-memory SQLite database, and also show how to use mocks.

## Testing a HTTP handler

To test a HTTP handler, we can use the `net/http/httptest` package.

It provides a way to create a request and a response recorder, which we can use to test our handler.

Note the name of the test function - it starts with `Test` and is followed by the name of the function being tested. This is important, because the `go test` command looks for functions that start with `Test` to run as tests.

Secondly, note the sub-test definition using `t.Run`. This allows us to group related tests together, and also gives us a way to run a specific test using the `-run` flag.

I always name my tests in a way that describes the behaviour being tested as a positive assertion. This makes it easier to understand what the test is supposed to be testing, vs a test name like "TestHandlerNotFound" which is less clear.

```go
func TestNotFoundHandler(t *testing.T) {
  t.Run("returns 404 status", func(t *testing.T) {
    // Arrange.
    r := httptest.NewRequest(http.MethodGet, "/users", nil)
    w := httptest.NewRecorder()

    sut := http.NotFoundHandler()

    // Act.
    sut.ServeHTTP(w, r)

    // Assert.
    if w.Code != http.StatusNotFound {
	t.Fatalf("expected status %d, got %d", http.StatusNotFound, w.Code)
    }
  })
}
```

If you remember from the previous sessions, you can run `go test ./...` to run all the tests in the current module or use the built-in feature of your IDE.

In this example, we create a new request using `httptest.NewRequest`, and a new response recorder using `httptest.NewRecorder`.

We then create our system under test (`sut`) - a simple `http.NotFoundHandler`.

We're then able to call the `ServeHTTP` method on the system under test, passing in the response recorder and the request.

Finally, we assert that the status code of the response is what we expect - `http.StatusNotFound`.

## Passing real test dependencies

In our actual application, we are likely to want to test our database access code. It can be tested in a separate layer of tests, but if we over-rely on mocks, we can end up testing the structure of our code rather than its behavior.

https://kentcdodds.com/blog/write-tests

Our HTTP handler requires a logger and a database connection.

One trick I like to do is to write log output to a buffer, and then only print it if the test fails. This means that I can see the log output when I'm debugging a test failure, but I don't have to wade through it when all the tests are passing.

```go
func TestHandler(t *testing.T) {
  logOutput := bytes.NewBuffer(nil)
  log := slog.New(slog.NewTextHandler(logOutput, &slog.HandlerOptions{Level: slog.LevelDebug}))
  defer func() {
    if t.Failed() {
      t.Logf("log output:\n%s", logOutput.String())
    }
  }()

  // ...
}
```


The handler also requires an instance of `*db.DB` - our database access struct. So, lets pass in a real database connection.

```go
pool, err := sqlitex.NewPool("file::memory:?mode=memory&cache=shared", sqlitex.PoolOptions{})
if err != nil {
  t.Fatalf("failed to create in-memory database pool: %v", err)
}
defer pool.Close()

kv := sqlitekv.NewStore(pool)
if err = kv.Init(t.Context()); err != nil {
  t.Fatalf("failed to initialise store: %v", err)
}

db := db.New(log, kv)
```

Finally, we can create our handler, passing in the logger and the database connection.

```go
// Create handler.
h := NewHandler(log, db)

r := httptest.NewRequest(http.MethodGet, "/users", nil)
w := httptest.NewRecorder()

// Act.
h.ServeHTTP(w, r)

// Assert.
if w.Code != http.StatusOK {
	t.Fatalf("expected status %d, got %d", http.StatusNotFound, w.Code)
}

// Read the body.
body := w.Body.String()
expected := `{"users":[]}` + "\n"
if body != expected {
	t.Fatalf("expected body %q, got %q", expected, body)
}
```

It's not a very good test, because it doesn't set up any data in the database, but it shows how to create a real database connection and pass it into the handler.

Also, I don't like relying on string comparisons for JSON output, because it can be brittle. A better approach would be to unmarshal the JSON into a struct and compare the struct.

In projects I've ran, we made a rule that we'd always create a Go HTTP client for our APIs alongside the server, so that we could use the client in our tests to make requests to the server and unmarshal the responses into structs.

That way, anyone using our API would have a client they could use, and we could use the client in our tests to make requests to the server and unmarshal the responses into structs.

## Setting up a test client

A simple HTTP client might look like this:

```go
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-h/go-workshop/100/102/04-web-testing/models"
)

func New(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}
}

type Client struct {
	BaseURL    string
	HTTPClient Doer
}

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func (c *Client) UsersGet() (output models.UsersGetResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+"/users", nil)
	if err != nil {
		return output, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return output, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return output, err
	}
	return output, nil
}

func (c *Client) UsersPost(input models.UsersPostRequest) (err error) {
	body, err := json.Marshal(input)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/users", bytes.NewReader(body))
	if err != nil {
		return err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
```

Note the use of the custom `Doer` interface. This allows us to pass in an alternative HTTP client for testing, such as the `httptest` server.

However, this simple client doesn't take care of authentication, retries and mapping error responses to Go errors. To help me, I wrote a library at https://github.com/a-h/jsonapi/ which takes care of some of the common boilerplate code. You can use it for inspiration in your own projects.

Similarly, you might want to implement Open API specifications for your API. I prefer to write code first, and base my spec off that, rather than using YAML to define my objects, so I wrote https://github.com/a-h/rest to help me with that.

### Updating our unit tests to use the client

Since the test HTTP client uses the same models as the server, we can use it in our tests. But, our tests don't start a HTTP server, they test the http Handler directly in code, so we need a `Doer` implementation that calls the handler directly.

It could be possible to start a HTTP server, but that would slow down the test setup - another example of the trade off between speed and practicality.

One nice thing about Go tests it that the results are cached, and the tests are only ran again if the code, or a dependency of the code changes. So, if you have a slow test that takes a few seconds to run, it won't be ran again unless something changes.

```go
func newTestDoer(h http.Handler) *testDoer {
	return &testDoer{h: h}
}

type testDoer struct {
	h http.Handler
	r *http.Response
}

func (d *testDoer) Do(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	d.h.ServeHTTP(w, req)
	d.r = w.Result()
	return d.r, nil
}
```

Our test code is then able to use the client to make requests to the handler, plus it records the last response so that we can make assertions on it.

```go
// Create handler.
h := NewHandler(log, db)

// Create HTTP client.
c := client.New("http://example.com")
testDoer := newTestDoer(h)
c.HTTPClient = testDoer

// Act.
users, err := c.UsersGet()
if err != nil {
	t.Fatalf("unexpected error: %v", err)
}
if len(users.Users) != 0 {
	t.Fatalf("expected 0 users, got %d", len(users.Users))
}
if testDoer.r.StatusCode != http.StatusOK {
	t.Fatalf("expected status code %d, got %d", http.StatusOK, testDoer.r.StatusCode)
}
```

## Running integration tests with a real database

If you take a look at the database library I've used https://github.com/a-h/kv - you'll notice that it supports Postgres, rqlite and SQLite. The unit tests for the kv library cover all of the databases by executing integration tests.

There's a `compose.yaml` in the root of the repo that you can use to start a Postgres and rqlite database to run integration tests. Thanks to Docker, it's easy to set up a real database for testing. Even proprietary databases like DynamoDB have local versions you can run in Docker.

and a simple example of database migrations - for your SQL migrations, you might want to use https://github.com/golang-migrate/migrate instead, because it's more fully featured.

## Running tests with a mock database

There are some things that a real database can't easily simulate:

- Timeouts (e.g. to test retry logic)
- Errors on demand (e.g. to test error handling paths)
- Large datasets (e.g. to test performance)

For those scenarios, mocks are a better fit.

Our HTTP handlers depend on a `db.DB` struct, so we need to make one of those to pass in. Unlike JavaScript, we can't pass any old object in - it has to be a `*db.DB`.

Here we have two options:

- Pass a mock `kv.Store` to `db.New()` and do the mocking at the key/value store level.
- Refactor the handler to depend on an interface, and then pass in a mock implementation of that interface instead.

In this example, we'll do the latter, because it allows us to test the handler in isolation, without needing to set up any database state.

First, we need to define an interface that our handler will depend on. It should contain the methods that the handler needs to call.

The constructor looks like this:

```go
func NewHandler(log *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		log: log,
		db:  db,
	}
}

type Handler struct {
	log *slog.Logger
	db  *db.DB
}
```

Looking at the code, we can see that the handler only calls two methods on the `db.DB` struct:

- ListUsers
- CreateUser

So, we can define an interface that contains those two methods. We can define this interface where we need it, right next to the Handler.

```go
type UserStore interface {
	ListUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user models.UserFields) error
}

func NewHandler(log *slog.Logger, db UserStore) *Handler {
	return &Handler{
		log: log,
		db:  db,
	}
}

type Handler struct {
	log *slog.Logger
	db  UserStore
}
```

That's it. Our real code still works - `*db.DB` already implements the `UserStore` interface without us needing to make any changes to it. We're just telling the compiler that we only care about those two methods.

That simplifies our testing code, because we only have to implement those two methods on our mock.

In Go, while there are some mocking libraries, I prefer to write my own mocks. It only takes a few lines of code, and it's easy to understand.

The main idea is to take advantage of Go's first class functions and create variables that store the implementation of the function.

```go
type UserStoreMock struct {
	ListUsersFunc  func(ctx context.Context) ([]models.User, error)
	CreateUserFunc func(ctx context.Context, user models.UserFields) error
}

func (m *UserStoreMock) ListUsers(ctx context.Context) ([]models.User, error) {
	return m.ListUsersFunc(ctx)
}

func (m *UserStoreMock) CreateUser(ctx context.Context, user models.UserFields) error {
	return m.CreateUserFunc(ctx, user)
}
```

Putting this all together, we can write a test that uses the mock.

```go
t.Run("database errors return an error", func(t *testing.T) {
	dbMock := &UserStoreMock{
		ListUsersFunc: func(ctx context.Context) ([]models.User, error) {
			return nil, fmt.Errorf("database error")
		},
	}

	// Create handler.
	h := NewHandler(log, dbMock)

	// Create HTTP client.
	c := client.New("http://example.com")
	testDoer := newTestDoer(h)
	c.HTTPClient = testDoer

	// Act.
	_, err := c.UsersGet()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if testDoer.r.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected status code %d, got %d", http.StatusInternalServerError, testDoer.r.StatusCode)
	}
})
```

Yes, there's a bit of setup, but it's just normal Go code. No special libraries that you have to learn or DSLs to understand.

## Summary

- Use `httptest` to test HTTP handlers.
- Use real dependencies where practical, such as an in-memory database.
- Use mocks to simulate errors and edge cases that are hard to reproduce with real dependencies.
- Consider writing a HTTP client for your API, and use it in your tests to make requests to the server and unmarshal the responses into structs.
- Use interfaces to define the dependencies of your handlers, so that you can easily swap out real implementations for mocks in your tests.
- Write your own mocks using first class functions, rather than relying on mocking libraries.
- Name your tests in a way that describes the behaviour being tested as a positive assertion.
- Use sub-tests to group related tests together, and to run specific tests using the `-run` flag.

## Task

Now that you've seen how to do it, can you write tests for the handler you built in the previous session?
