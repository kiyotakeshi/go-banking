package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (account Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: account.AccountId}
}

func (account Account) CanWithdraw(amount float64) bool {
	if account.Amount < amount {
		return false
	}
	return true
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.ApplicationError)
	FindById(accountId string) (*Account, *errs.ApplicationError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.ApplicationError)
}
