package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/a-h/go-workshop/100/102/03-web/db"
	"github.com/a-h/go-workshop/100/102/03-web/handlers/users"
	"github.com/a-h/kv/sqlitekv"
	"zombiezen.com/go/sqlite/sqlitex"
)

func main() {
	// Decide whether to enable verbose logging.
	cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	v := cmd.Bool("v", false, "enable verbose logging")
	dsn := cmd.String("dsn", "file:example.db?cache=shared&mode=rwc&_busy_timeout=5000&_txlock=immediate", "SQLite DSN")
	cmd.Parse(os.Args[1:])

	// Create a JSON logger that writes to standard error.
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	if v != nil && *v {
		opts.Level = slog.LevelDebug
	}
	log := slog.New(slog.NewJSONHandler(os.Stderr, opts))

	log.Debug("Logger initialised")

	// When you press Ctrl+C, a signal will be sent to the channel called `stop`.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Context can be cancelled to stop downstream operations.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// If the stop channel receives a signal, cancel the context.
	go func() {
		<-stop
		cancel()
	}()

	if err := serve(ctx, log, *dsn); err != nil {
		log.Error("Error", slog.Any("error", err))
		os.Exit(1)
	}
}

func serve(ctx context.Context, log *slog.Logger, dsn string) error {
	// Create a new SQLite database.
	pool, err := sqlitex.NewPool("file:example.db?cache=shared&mode=rwc&_busy_timeout=5000&_txlock=immediate", sqlitex.PoolOptions{})
	if err != nil {
		return fmt.Errorf("failed to create database pool: %w", err)
	}
	defer pool.Close()

	// Create a store.
	store := sqlitekv.NewStore(pool)

	// Initialize the store (creates tables).
	if err := store.Init(ctx); err != nil {
		return fmt.Errorf("failed to initialize store: %w", err)
	}

	// Create the database access object.
	// This wraps the store with application-specific methods.
	db := db.New(log, store)

	// Create the handlers.
	uh := users.NewHandler(log, db)

	// Map the routes to the handlers.
	mux := http.NewServeMux()
	mux.Handle("/users", uh)

	// Log all requests (for demonstration purposes).
	loggingHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Request received", slog.String("method", r.Method), slog.String("url", r.URL.String()), slog.String("remote", r.RemoteAddr))
		mux.ServeHTTP(w, r)
	})

	// Create the server.
	s := http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           loggingHandler,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Minute * 30,
		IdleTimeout:       time.Hour,
	}
	go func() {
		log.Info("Server starting", slog.String("addr", s.Addr))
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("ListenAndServe error", slog.Any("error", err))
		}
	}()
	<-ctx.Done()

	// Allow 5 seconds for graceful shutdown of server.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Shutdown(ctx)
}
