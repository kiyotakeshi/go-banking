package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

// go generate ./...
//go:generate mockgen -source=$GOFILE -destination=../mocks/$GOPACKAGE/mock_$GOFILE
type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.ApplicationError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.ApplicationError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.ApplicationError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, appErr := s.repo.FindAll(status)
	if appErr != nil {
		return nil, appErr
	}

	response := make([]dto.CustomerResponse, 0)
	for _, customer := range customers {
		response = append(response, customer.ToDto())
	}
	return response, appErr
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.ApplicationError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
