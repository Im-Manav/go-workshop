package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/a-h/go-workshop/100/102/04-web-testing/models"
	"github.com/a-h/kv"
	"github.com/google/uuid"
)

func New(log *slog.Logger, store kv.Store) *DB {
	return &DB{
		log:   log,
		store: store,
	}
}

type DB struct {
	log   *slog.Logger
	store kv.Store
}

func (db *DB) ListUsers(ctx context.Context) ([]models.User, error) {
	records, err := db.store.GetPrefix(ctx, "user/", 0, -1)
	if err != nil {
		return nil, err
	}

	users, err := kv.ValuesOf[models.User](records)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (db *DB) CreateUser(ctx context.Context, user models.UserFields) error {
	key := fmt.Sprintf("user/%s", uuid.New())
	// We expect the UUID to be unique, so the version is 0.
	var version int
	return db.store.Put(ctx, key, version, user)
}
