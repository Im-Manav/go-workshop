package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func newDB() (string, error) {
	time.Sleep(1 * time.Second)
	return "Database connection established", nil
}

func getConfig() (string, error) {
	time.Sleep(1 * time.Second)
	return "Config loaded", nil
}

func connectToService() (string, error) {
	time.Sleep(1 * time.Second)
	return "Connected to external service", nil
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	err := run(context.Background(), log)
	if err != nil {
		log.Error("Application failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func run(_ context.Context, log *slog.Logger) error {
	startTime := time.Now()

	var wg errgroup.Group

	// Create variables to store the results of each goroutine.
	var db, config, service string

	// Note the use of closures to capture the variables, and the use of a named
	// return value to simplify error handling.
	wg.Go(func() (err error) {
		log.Debug("Starting database connection")
		db, err = newDB()
		return err
	})

	wg.Go(func() (err error) {
		log.Debug("Loading configuration")
		config, err = getConfig()
		return err
	})

	wg.Go(func() (err error) {
		log.Debug("Connecting to external service")
		service, err = connectToService()
		return err
	})

	err := wg.Wait()
	if err != nil {
		return fmt.Errorf("failed to initialise application: %w", err)
	}

	log.Info("Application started successfully", slog.Duration("startTimeMs", time.Duration(time.Since(startTime).Milliseconds())))
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DB: %s\nConfig: %s\nService: %s\n", db, config, service)
	})
	return http.ListenAndServe(":8080", nil)
}
