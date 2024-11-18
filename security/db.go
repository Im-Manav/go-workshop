package main

import (
	"database/sql"
	"fmt"
)

type Store interface {
  PutCsv(Csv) error
  GetCsv(string) (Csv, error)
}

type Csv struct {
  category string
  file []byte
}

type store struct {
  dbClient *sql.DB
}

func NewStore(db *sql.DB) store {
   return store{
    dbClient: db, 
  } 
}

func (s *store) PutCsv(c Csv) error {
  _, err := s.dbClient.Exec("INSERT INTO uploads (category, file) VALUES (?,?)", c.category, c.file)
  if err != nil {
    return err
  }
  return nil
}

func (s *store) GetCsv(category string) (Csv, error) {
  var c Csv
  row := s.dbClient.QueryRow("SELECT * FROM uploads WHERE category = ?", category)
  if err := row.Scan(&c.category, &c.file); err != nil {
    if err == sql.ErrNoRows {
      return c, fmt.Errorf("category %s: no such category", category)
    }
    return c, fmt.Errorf("categgory %s: %v", category, err)
  } 
  return c, nil
}
