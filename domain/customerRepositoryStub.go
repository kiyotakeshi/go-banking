package domain

import "banking/errs"

// CustomerRepositoryStub Mock adapter
type CustomerRepositoryStub struct {
	customers []Customer
	customer  Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, *errs.ApplicationError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) FindById(id string) (*Customer, *errs.ApplicationError) {
	return &s.customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"100", "mike", "tokyo", "001001", "1998-01-01", "1"},
		{"101", "popcorn", "osaka", "001002", "1988-11-01", "1"},
	}

	customer := Customer{"201", "kanye", "chicago", "002001", "1970-06-09", "1"}
	return CustomerRepositoryStub{customers: customers, customer: customer}
}
