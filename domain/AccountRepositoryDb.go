package domain

import (
	"banking/errs"
	"banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(account Account) (*Account, *errs.ApplicationError) {
	sqlInsert := "INSERT INTO ACCOUNTS (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
