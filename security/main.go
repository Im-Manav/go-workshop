package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/language"
)

type Customer struct {
	Id      int    `json:"customerId"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Company string `json:"company"`
}

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func run() error {
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

	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS customers (id INTEGER PRIMARY KEY, name TEXT NOT NULL, surname TEXT NOT NULL, company TEXT NOT NULL)")
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO customers (name, surname, company)
VALUES
('Gopher', 'McGopher', 'Gopher LTD'),
('John', 'Doe', 'JDoe LTD');`)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	store := NewStore(db)

	mux.Handle("GET /customer/{id}", getCustomerHandler(store))
	mux.Handle("POST /customer", postCustomerHandler(store))

	s := http.Server{
		Addr:    ":8005",
		Handler: mux,
	}
	fmt.Println("listening on port :8005")
	err = s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func postCustomerHandler(s Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var c Customer
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = s.PutCustomer(c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func getCustomerHandler(s Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
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
