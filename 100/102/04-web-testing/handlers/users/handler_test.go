package users

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-h/go-workshop/100/102/04-web-testing/client"
	"github.com/a-h/go-workshop/100/102/04-web-testing/db"
	"github.com/a-h/go-workshop/100/102/04-web-testing/models"
	"github.com/a-h/kv/sqlitekv"
	"zombiezen.com/go/sqlite/sqlitex"
)

func TestHandler(t *testing.T) {
	// Create logger.
	logOutput := bytes.NewBuffer(nil)
	log := slog.New(slog.NewTextHandler(logOutput, &slog.HandlerOptions{Level: slog.LevelDebug}))
	defer func() {
		if t.Failed() {
			t.Logf("log output:\n%s", logOutput.String())
		}
	}()

	// Create DB.
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

	t.Run("GET", func(t *testing.T) {
		t.Run("returns an empty list if there are no users", func(t *testing.T) {
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
		})
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
	})
}

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
