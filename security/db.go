package main

import (
	"database/sql"
	"fmt"
)

type Store interface {
  PutCustomer(Customer) error
  GetCustomer(string) (Customer, error)
}


type clientDB struct {
  dbClient *sql.DB
}

func NewStore(db *sql.DB) clientDB {
   return clientDB{
    dbClient: db, 
  } 
}

func (s clientDB) PutCustomer(c Customer) error {
  _, err := s.dbClient.Exec("INSERT INTO customers (name, surname, company) VALUES (?,?,?)", c.Name, c.Surname, c.Company)
  if err != nil {
    return err
  }
  return nil
}

func (s clientDB) GetCustomer(id string) (Customer, error) {
  query := fmt.Sprintf("SELECT * FROM customers WHERE id = %s", id)
  var c Customer
  row := s.dbClient.QueryRow(query)
  if err := row.Scan(&c.Id,&c.Name, &c.Surname, &c.Company); err != nil {
    if err == sql.ErrNoRows {
      return c, fmt.Errorf("id %d: no such id", id)
    }
    return c, fmt.Errorf("id %s: %v", id, err)
  } 
  return c, nil
}
