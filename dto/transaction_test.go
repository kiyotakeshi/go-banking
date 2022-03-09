package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	// arrange
	request := TransactionRequest{
		TransactionType: "invalid transaction type",
	}

	// act
	applicationError := request.Validate()

	// assert
	if applicationError.Message != "transaction type can only be deposit or withdrawal" {
		t.Error("invalid message while testing transaction type")
	}

	if applicationError.Code != http.StatusUnprocessableEntity {
		t.Error("invalid code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	// arrange
	request := TransactionRequest{TransactionType: DEPOSIT, Amount: -1}

	// act
	applicationError := request.Validate()

	// assert
	if applicationError.Message != "amount cannot be less than zero" {
		t.Error("invalid message while validating amount")
	}

	if applicationError.Code != http.StatusUnprocessableEntity {
		t.Error("invalid code while validation amount")
	}
}
