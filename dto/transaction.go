package dto

import "banking/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"customer_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.ApplicationError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("amount cannot be less than zero")
	}
	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
