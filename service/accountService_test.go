package service

import (
	real_domain "banking/domain"
	"banking/dto"
	"banking/errs"
	mock_domain "banking/mocks/domain"
	"github.com/golang/mock/gomock"
	"testing"
)

var mockRepository *mock_domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	controller := gomock.NewController(t)
	mockRepository = mock_domain.NewMockAccountRepository(controller)
	service = NewAccountService(mockRepository)
	return func() {
		defer controller.Finish()
	}
}

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	// arrange
	tearDown := setup(t)
	defer tearDown()

	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}
	service = NewAccountService(nil)

	// act
	_, appErr := service.NewAccount(request)

	// assert
	// if appErr == nil {
	if appErr.Message != "to open a new account, you need to deposit at least 5000" {
		t.Error("failed while testing the new account validation")
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_account_cannot_be_created(t *testing.T) {
	// arrange
	tearDown := setup(t)
	defer tearDown()

	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := real_domain.NewAccount(request.CustomerId, request.AccountType, request.Amount)
	mockRepository.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("unexpected database error"))

	// act
	_, appErr := service.NewAccount(request)

	// assert
	if appErr == nil {
		t.Error("test failed while validating error for new account")
	}
}

func Test_should_return_new_account_response_when_a_new_account_is_saved_successfully(t *testing.T) {
	// arrange
	tearDown := setup(t)
	defer tearDown()

	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := real_domain.NewAccount(request.CustomerId, request.AccountType, request.Amount)
	accountWithId := account
	accountWithId.AccountId = "301"
	mockRepository.EXPECT().Save(account).Return(&accountWithId, nil)

	// act
	newAccount, appErr := service.NewAccount(request)
	if appErr != nil {
		t.Error("test failed while validating error for new account")
	}
	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("failed while matching new account id")
	}
}
