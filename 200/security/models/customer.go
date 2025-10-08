package models

import "errors"

type Customer struct {
	ID      int    `json:"customerId"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Company string `json:"company"`
}

func (c Customer) Validate() error {
	var errs []error
	if c.Name == "" {
		errs = append(errs, errors.New("Name is a required field"))
	}
	if c.Surname == "" {
		errs = append(errs, errors.New("Surname is a required field"))
	}
	return errors.Join(errs...)
}
