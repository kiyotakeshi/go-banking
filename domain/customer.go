package domain

import "banking/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	// 1: active
	FindAll(status string) ([]Customer, *errs.ApplicationError)
	FindById(string) (*Customer, *errs.ApplicationError)
}
