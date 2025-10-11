# Solution

## Setup a new route

```go
// Create the handlers.
usersHandler := users.NewHandler(log, db)
userHandler := user.NewHandler(log, db)

// Map the routes to the handlers.
mux := http.NewServeMux()
mux.Handle("/users", usersHandler)
mux.Handle("/user/{id}", userHandler)
```

## Create a new handler

Add to `./handlers/user/handler.go`:

```go
package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/a-h/go-workshop/100/102/03-web/db"
	"github.com/a-h/go-workshop/100/102/03-web/models"
)

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

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		h.Delete(w, r)
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPost:
		h.Post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	user, ok, err := h.db.GetUser(r.Context(), id)
	if err != nil {
		h.log.Error("Error getting user", slog.Any("error", err), slog.String("id", id))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var resp models.UserGetResponse
	resp.User = user
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.log.Error("Error encoding response", slog.Any("error", err))
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var req models.UserPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Error decoding request body", slog.Any("error", err))
		// Never expose internal error details to the client.
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if msgs, ok := req.Validate(); !ok {
		h.log.Debug("Validation failed", slog.Any("messages", msgs))
		http.Error(w, "Validation failed: "+strings.Join(msgs, ", "), http.StatusBadRequest)
		return
	}

	err := h.db.UpdateUser(r.Context(), id, req.UserFields)
	if err != nil {
		h.log.Error("Error storing user", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var resp models.UserPostResponse
	resp.OK = true
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.log.Error("Error encoding response", slog.Any("error", err))
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	err := h.db.DeleteUser(r.Context(), id)
	if err != nil {
		h.log.Error("Error deleting user", slog.Any("error", err), slog.String("id", id))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var resp models.UserDeleteResponse
	resp.OK = true
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.log.Error("Error encoding response", slog.Any("error", err))
	}
}
```

## Create request and response models

Add to `./models/user.go`:

```go
// GET /user/{id}
type UserGetResponse struct {
	User User `json:"user"`
}

// POST /user/{id}
type UserPostRequest struct {
	UserFields
}

type UserPostResponse struct {
	OK bool `json:"ok"`
}

// DELETE /user/{id}
type UserDeleteResponse struct {
	OK bool `json:"ok"`
}
```

### Create database methods

Add to `./db/db.go`:

```go
func (db *DB) GetUser(ctx context.Context, id string) (user models.User, ok bool, err error) {
	key := fmt.Sprintf("user/%s", id)
	_, ok, err = db.store.Get(ctx, key, &user)
	return user, ok, err
}

func (db *DB) UpdateUser(ctx context.Context, id string, user models.UserFields) error {
	key := fmt.Sprintf("user/%s", id)
	// Overwrite any existing user data.
	version := -1
	return db.store.Put(ctx, key, version, user)
}

func (db *DB) DeleteUser(ctx context.Context, id string) error {
	key := fmt.Sprintf("user/%s", id)
	_, err := db.store.Delete(ctx, key)
	return err
}
```
