package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	posthandler "github.com/a-h/go-workshop-102/security/handlers/customer/post"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/language"
)

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func run(ctx context.Context, log *slog.Logger) error {
	for _, arg := range os.Args[1:] {
		tag, err := language.Parse(arg)
		if err != nil {
			fmt.Printf("%s: error: %v\n", arg, err)
		} else if tag == language.Und {
			fmt.Printf("%s: undefined\n", arg)
		} else {
			fmt.Printf("%s: tag %s\n", arg, tag)
		}
	}

	log.Info("Creating database")
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS customers (id INTEGER PRIMARY KEY, name TEXT NOT NULL, surname TEXT NOT NULL, company TEXT NOT NULL)")
	if err != nil {
		return err
	}

	log.Info("Populating database with test data")
	_, err = db.Exec(`INSERT INTO customers (name, surname, company)
VALUES
('Gopher', 'McGopher', 'Gopher LTD'),
('John', 'Doe', 'JDoe LTD');`)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	store := newDB(db)

	mux.Handle("GET /customer/{id}", getCustomerHandler(log, store))
	mux.Handle("POST /customer", posthandler.New(log, store))

	s := http.Server{
		Addr:    ":8005",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		log.Info("shutting down server")
		if err := s.Shutdown(ctx); err != nil {
			log.Error("error shutting down server", slog.Any("error", err))
		}
	}()

	log.Info("listening", slog.Any("addr", s.Addr))
	err = s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{}))

	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		log.Info("received signal, shutting down")
		cancel()
	}()

	var exitCode int
	err := run(ctx, log)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Error("fatal error", slog.Any("error", err))
		exitCode = 1
	}
	os.Exit(exitCode)
}

func getCustomerHandler(log *slog.Logger, s clientDB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		log.Info("received request", slog.Any("method", r.Method), slog.Any("url", r.URL), slog.Any("id", id))
		customer, err := s.GetCustomer(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(&customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
