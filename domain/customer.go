package domain

import "banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	// 1: active
	FindAll(status string) ([]Customer, *errs.ApplicationError)
	FindById(string) (*Customer, *errs.ApplicationError)
}
