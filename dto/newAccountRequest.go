package dto

import "banking/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.ApplicationError {
	if r.Amount < 5000 {
		return errs.NewValidationError("to open a new account, you need to deposit at least 5000")
	}
	if r.AccountType != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("account type should be checking or saving")
	}
	return nil
}
