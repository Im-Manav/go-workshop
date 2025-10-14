package main

import (
	"database/sql"
	"fmt"

	"github.com/a-h/go-workshop-102/security/models"
)

type clientDB struct {
	dbClient *sql.DB
}

func newDB(db *sql.DB) clientDB {
	return clientDB{
		dbClient: db,
	}
}

func (s clientDB) PutCustomer(c models.Customer) error {
	_, err := s.dbClient.Exec("INSERT INTO customers (name, surname, company) VALUES (?,?,?)", c.Name, c.Surname, c.Company)
	if err != nil {
		return err
	}
	return nil
}

func (s clientDB) GetCustomer(id string) (models.Customer, error) {
	query := fmt.Sprintf("SELECT * FROM customers WHERE id = %s", id)
	var c models.Customer
	row := s.dbClient.QueryRow(query)
	if err := row.Scan(&c.ID, &c.Name, &c.Surname, &c.Company); err != nil {
		if err == sql.ErrNoRows {
			return c, fmt.Errorf("id %s: no such id", id)
		}
		return c, fmt.Errorf("id %s: %v", id, err)
	}
	return c, nil
}
