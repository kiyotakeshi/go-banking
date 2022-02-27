package domain

// CustomerRepositoryStub Mock adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"100", "mike", "tokyo", "001001", "1998-01-01", "1"},
		{"101", "popcorn", "osaka", "001002", "1988-11-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
