package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func run() error {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS uploads (category TEXT NOT NULL, file BLOB NOT NULL)")
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	store := NewStore(db)
	mux.Handle("POST /upload", postCsvHandler(&store))

	s := http.Server{
		Addr:    ":8005",
		Handler: mux,
	}
	fmt.Println("listening on port :8005")
	s.ListenAndServe()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func postCsvHandler(store Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}
