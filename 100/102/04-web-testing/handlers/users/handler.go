package users

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/a-h/go-workshop/100/102/04-web-testing/models"
)

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

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("Listing users from database")
	users, err := h.db.ListUsers(r.Context())
	if err != nil {
		h.log.Error("Error listing users", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.log.Debug("Returning users", slog.Int("count", len(users)))
	var resp models.UsersGetResponse
	resp.Users = users
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.log.Error("Error encoding response", slog.Any("error", err))
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("Creating user from request body")
	var req models.UsersPostRequest
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

	h.log.Debug("Creating user", slog.Any("user", req))
	err := h.db.CreateUser(r.Context(), req.UserFields)
	if err != nil {
		h.log.Error("Error storing user", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.log.Debug("User created successfully, returning response")
	var resp models.UsersPostResponse
	resp.OK = true
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.log.Error("Error encoding response", slog.Any("error", err))
	}
}
