package post

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/a-h/go-workshop-102/security/models"
)

type CustomerPutter interface {
	PutCustomer(c models.Customer) error
}

func New(log *slog.Logger, db CustomerPutter) Handler {
	return Handler{
		log: log,
		db:  db,
	}
}

type Handler struct {
	log *slog.Logger
	db  CustomerPutter
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("received request", slog.Any("method", r.Method), slog.Any("url", r.URL))
	var c models.Customer
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		h.log.Info("bad request received", slog.Any("error", err))
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err = c.Validate(); err != nil {
		h.log.Info("bad request received", slog.Any("error", err))
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.db.PutCustomer(c)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
